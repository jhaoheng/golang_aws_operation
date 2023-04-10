package DDBAgent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
此 item size 約 128 Bytes
*/
// type ITEM struct {
// 	DidModuleID string `json:"did_module_id"`
// 	CreatedTs   string `json:"created_ts"`
// 	Did         string `json:"did"`
// 	Value       string `json:"value"`
// 	ExpiredAt   string `json:"expired_at"`
// }

type IDDBAgent interface {
	CreateTable() (*dynamodb.CreateTableOutput, error)
	DescribeTable() (*dynamodb.DescribeTableOutput, error)
	PutItem(obj interface{}) (*dynamodb.PutItemOutput, error)
	GetItem() (*dynamodb.GetItemOutput, error)
	BatchWriteItem(items []map[string]interface{}) (*dynamodb.BatchWriteItemOutput, error)
	BatchGetItem(get_items []GetItem) (*dynamodb.BatchGetItemOutput, error)
}

type DDBAgent struct {
	Agent *dynamodb.DynamoDB
	Table string
}

func NewDDBAgent(table_name string) IDDBAgent {
	//
	cfg := aws.NewConfig()
	cfg = cfg.WithRegion("us-east-1")
	cfg = cfg.WithEndpoint("http://localhost:8000")
	//
	sess := session.Must(session.NewSession())
	dynamodbAgent := dynamodb.New(sess, cfg)

	return &DDBAgent{
		Agent: dynamodbAgent,
		Table: table_name,
	}
}
