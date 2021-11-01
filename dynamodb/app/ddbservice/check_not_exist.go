package ddbservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (ddb *DDBSERVICE) CheckNotExists() (*dynamodb.ScanOutput, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(ddb.table),
		Select:    aws.String("ALL_ATTRIBUTES"),
		ExpressionAttributeNames: map[string]*string{
			"#Status": aws.String("status"),
		},
		FilterExpression: aws.String("attribute_not_exists(#Status)"),
	}
	scanOutput, err := ddb.agent.Scan(scanInput)
	return scanOutput, err
}
