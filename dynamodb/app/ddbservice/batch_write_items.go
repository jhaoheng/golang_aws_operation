package ddbservice

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (ddb *DDBSERVICE) BatchWriteItem(items []ITEM) (*dynamodb.BatchWriteItemOutput, error) {
	var requestItems = []*dynamodb.WriteRequest{}
	for index, item := range items {
		writeRequest := dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					"did_module_id": {
						S: aws.String(strconv.Itoa(index)),
					},
					"created_ts": {
						N: aws.String(item.CreatedTs),
					},
					"did": {
						S: aws.String(item.Did),
					},
					"value": {
						S: aws.String(item.Value),
					},
					"expired_at": {
						N: aws.String(item.ExpiredAt),
					},
				},
			},
		}
		requestItems = append(requestItems, &writeRequest)
	}

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			ddb.table: requestItems,
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}

	return ddb.agent.BatchWriteItem(batchWriteItemInput)
}
