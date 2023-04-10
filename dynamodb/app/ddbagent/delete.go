package DDBAgent

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (agent *DDBAgent) Delete(obj interface{}) (*dynamodb.DeleteItemOutput, error) {
	av, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}
	return agent.Agent.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(agent.Table),
		Key:       av,
	})
}
