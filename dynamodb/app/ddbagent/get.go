package DDBAgent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (agent *DDBAgent) GetItem() (*dynamodb.GetItemOutput, error) {

	av := map[string]*dynamodb.AttributeValue{
		"id": {N: aws.String("1")},
	}

	return agent.Agent.GetItem(&dynamodb.GetItemInput{
		TableName:              aws.String(agent.Table),
		Key:                    av,
		ReturnConsumedCapacity: aws.String("TOTAL"),
	})

	// return agent.Agent.GetItem(&dynamodb.GetItemInput{
	// 	TableName: aws.String(agent.Table),
	// 	Key: map[string]*dynamodb.AttributeValue{
	// 		"id": {N: aws.String("1")},
	// 	},
	// 	ReturnConsumedCapacity: aws.String("TOTAL"),
	// })
}
