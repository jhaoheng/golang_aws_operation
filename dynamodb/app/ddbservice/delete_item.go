package ddbservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (s *DDBSERVICE) Delete(item ITEM) (*dynamodb.DeleteItemOutput, error) {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"did_module_id": {
				S: aws.String(item.DidModuleID),
			},
			"created_ts": {
				N: aws.String(item.CreatedTs),
			},
		},
		TableName:              aws.String(s.table),
		ReturnConsumedCapacity: aws.String("TOTAL"),
	}
	return s.agent.DeleteItem(input)
}
