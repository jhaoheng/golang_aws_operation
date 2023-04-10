package DDBAgent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

/*
- https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.BatchGetItem
- 最多 16MB, 100 個物件, 每個物件不得超過 300 KB
*/

type GetItem struct {
	Key   string
	Value string
}

func (ddb *DDBAgent) BatchGetItem(get_items []GetItem) (*dynamodb.BatchGetItemOutput, error) {

	var avs = []map[string]*dynamodb.AttributeValue{}
	for _, item := range get_items {
		av := map[string]*dynamodb.AttributeValue{
			item.Key: {N: aws.String(item.Value)},
		}
		avs = append(avs, av)
	}

	// var avs = []map[string]*dynamodb.AttributeValue{
	// 	{
	// 		"id": {N: aws.String("1")},
	// 	},
	// 	{
	// 		"id": {N: aws.String("2")},
	// 	},
	// }

	batchGetItemInput := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			ddb.Table: {
				Keys: avs,
			},
		},
	}

	if err := batchGetItemInput.Validate(); err != nil {
		return nil, err
	}

	return ddb.Agent.BatchGetItem(batchGetItemInput)
}

/*
- https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#DynamoDB.BatchGetItem
- 最多 16MB, 100 個物件, 每個物件不得超過 300 KB
*/
// func (ddb *DDBAgent) BatchGetItem(items []ITEM) {
// var keys = []map[string]*dynamodb.AttributeValue{}
// for _, item := range items {
// 	key := map[string]*dynamodb.AttributeValue{
// 		"did_module_id": {
// 			S: aws.String(item.DidModuleID),
// 		},
// 		"created_ts": {
// 			S: aws.String(item.CreatedTs),
// 		},
// 	}
// 	keys = append(keys, key)
// }

// batchGetItemInput := &dynamodb.BatchGetItemInput{
// 	RequestItems: map[string]*dynamodb.KeysAndAttributes{
// 		ddb.Table: {
// 			Keys: keys,
// 			// ConsistentRead: ,
// 			// AttributesToGet: ,
// 			// ExpressionAttributeNames: ,
// 			// ProjectionExpression: ,
// 		},
// 	},
// }

// batchGetItemOutput, err := ddb.Agent.BatchGetItem(batchGetItemInput)
// if err != nil {
// 	panic(err)
// }

// fmt.Println(batchGetItemOutput)
// }
