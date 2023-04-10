package DDBAgent

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (ddb *DDBAgent) BatchWriteItem(items []map[string]interface{}) (*dynamodb.BatchWriteItemOutput, error) {

	var writes = []*dynamodb.WriteRequest{}

	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
		}
		// fmt.Println(av)
		writes = append(writes, &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: av,
			},
		})
	}

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			ddb.Table: writes,
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	return ddb.Agent.BatchWriteItem(batchWriteItemInput)
}
