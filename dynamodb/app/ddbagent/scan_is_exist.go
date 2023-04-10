package DDBAgent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// 屬性類型 : https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#AttributeValue
func (ddb *DDBAgent) ScanIsExist(key, value string) bool {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(ddb.Table),
		Select:    aws.String("COUNT"),
		ExpressionAttributeNames: map[string]*string{
			"#id": aws.String(key),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":id": {
				S: aws.String(value),
			},
		},
		FilterExpression: aws.String("#id = :id"),
	}

	scanOutput, err := ddb.Agent.Scan(scanInput)
	if err != nil {
		panic(err)
	}
	if *scanOutput.Count > 0 {
		return true
	}
	return false
}
