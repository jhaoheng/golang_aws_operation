package DDBAgent

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (agent *DDBAgent) PutItem(obj interface{}) (*dynamodb.PutItemOutput, error) {
	av, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	return agent.Agent.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(agent.Table),
		Item:      av,
	})
}
