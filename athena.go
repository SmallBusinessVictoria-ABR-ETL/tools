package tools

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"strings"
	"time"
)

var athenaSession *athena.Athena
var s3Session *s3.S3

func init() {
	athenaSession = athena.New(session.Must(session.NewSession()))
	s3Session = s3.New(session.Must(session.NewSession()))
}

func Query(sql, location string) {
	resp, err := athenaSession.StartQueryExecution(&athena.StartQueryExecutionInput{
		QueryString: aws.String(sql),
		QueryExecutionContext: &athena.QueryExecutionContext{
			Database: aws.String("sbv_abr"),
		},
		ResultConfiguration: &athena.ResultConfiguration{
			OutputLocation: aws.String(location),
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	for true {
		r2, err := athenaSession.GetQueryExecution(&athena.GetQueryExecutionInput{
			QueryExecutionId: resp.QueryExecutionId,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Print("State: ", *r2.QueryExecution.Status.State)
		time.Sleep(time.Second)

		if *r2.QueryExecution.Status.State != "RUNNING" {
			log.Print(*r2.QueryExecution.Status)

			//fmt.Print("\n\n-----\n" + location + "/" + *resp.QueryExecutionId + ".csv\n\n-----\n")

			o, _ := s3Session.GetObjectRequest(ObjectRequestFromS3Url(location + "/" + *resp.QueryExecutionId + ".csv"))

			signed, err := o.Presign(time.Hour * 24)
			if err != nil {
				log.Fatal("Failed to sign url")
			}

			fmt.Print("\n---------------\n" + signed + "\n---------------\n")
			break
		}

	}
}
func ObjectRequestFromS3Url(s string) *s3.GetObjectInput {

	s = strings.TrimPrefix(s, "s3://")
	bucket_key := strings.SplitN(s, "/", 2)

	return &s3.GetObjectInput{
		Bucket: &bucket_key[0],
		Key:    &bucket_key[1],
	}
}
