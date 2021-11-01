package ddbservice

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
- https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.BatchGetItem
- 最多 16MB, 100 個物件, 每個物件不得超過 300 KB
*/
func (ddb *DDBSERVICE) BatchGetItem(items []ITEM) {
	var keys = []map[string]*dynamodb.AttributeValue{}
	for _, item := range items {
		key := map[string]*dynamodb.AttributeValue{
			"did_module_id": {
				S: aws.String(item.DidModuleID),
			},
			"created_ts": {
				S: aws.String(item.CreatedTs),
			},
		}
		keys = append(keys, key)
	}

	batchGetItemInput := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			ddb.table: {
				Keys: keys,
				// ConsistentRead: ,
				// AttributesToGet: ,
				// ExpressionAttributeNames: ,
				// ProjectionExpression: ,
			},
		},
	}

	batchGetItemOutput, err := ddb.agent.BatchGetItem(batchGetItemInput)
	if err != nil {
		panic(err)
	}

	fmt.Println(batchGetItemOutput)
}
