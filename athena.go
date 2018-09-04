package tools

import (
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
		ResultConfiguration: &athena.ResultConfiguration{
			OutputLocation: aws.String(location),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	for c := 0; c < 500; c++ {
		r2, err := athenaSession.GetQueryExecution(&athena.GetQueryExecutionInput{
			QueryExecutionId: resp.QueryExecutionId,
		})
		if err != nil {
			log.Fatal(err)
		}

		log.Print("QueryId: ", *r2.QueryExecution.Status.State)
		time.Sleep(time.Second)
	}
}
