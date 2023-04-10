package DDBAgent

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func (ddb *DDBAgent) CheckNotExists(obj interface{}) (*dynamodb.ScanOutput, error) {

	av, err := dynamodbattribute.MarshalMap(obj)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	scanInput := &dynamodb.ScanInput{
		TableName:         aws.String(ddb.Table),
		Select:            aws.String("ALL_ATTRIBUTES"),
		ExclusiveStartKey: av,
		ExpressionAttributeNames: map[string]*string{
			"#Status": aws.String("status"),
		},
		FilterExpression: aws.String("attribute_not_exists(#Status)"),
	}

	if err := scanInput.Validate(); err != nil {
		return nil, err
	}

	scanOutput, err := ddb.Agent.Scan(scanInput)
	return scanOutput, err
}
