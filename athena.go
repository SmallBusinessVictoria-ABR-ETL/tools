package tools

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/athena"
	"log"
	"time"
)

var athenaSession *athena.Athena

func init() {
	athenaSession = athena.New(session.Must(session.NewSession()))
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

			fmt.Print("\n\n-----\n" + location + "/" + *resp.QueryExecutionId + ".csv\n\n-----\n")

			break
		}

	}
}
