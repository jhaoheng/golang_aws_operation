package ddbservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
此 item size 約 128 Bytes
*/
type ITEM struct {
	DidModuleID string `json:"did_module_id"`
	CreatedTs   string `json:"created_ts"`
	Did         string `json:"did"`
	Value       string `json:"value"`
	ExpiredAt   string `json:"expired_at"`
}

type DDBSERVICE struct {
	agent *dynamodb.DynamoDB
	table string
}

var DDBService *DDBSERVICE

func NewDDB() {
	//
	cfg := aws.NewConfig()
	cfg = cfg.WithRegion("us-east-1")
	cfg = cfg.WithEndpoint("http://dynamodb:8000")
	//
	sess := session.Must(session.NewSession())
	dynamodbAgent := dynamodb.New(sess, cfg)

	DDBService = &DDBSERVICE{
		agent: dynamodbAgent,
	}
}
