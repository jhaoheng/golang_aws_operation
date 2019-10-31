package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamodbObj struct {
	agent *dynamodb.DynamoDB
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://dynamodb:8000"),
	}))
	dynamodbAgent := dynamodb.New(sess)

	var dynamodbObj = DynamodbObj{
		agent: dynamodbAgent,
	}

	key := "id"
	value := "1"
	if dynamodbObj.ScanIsExist(key, value) {
		fmt.Printf("yes, the key[%s] and its value[%s] is exist\n", key, value)
	}

	dynamodbObj.BatchWriteItem()
}

const TableName = "test"

// 屬性類型 : https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#AttributeValue
func (dynamodbObj *DynamodbObj) ScanIsExist(key, value string) bool {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(TableName),
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

	scanOutput, err := dynamodbObj.agent.Scan(scanInput)
	if err != nil {
		panic(err)
	}
	if *scanOutput.Count > 0 {
		return true
	}
	return false
}

func (dynamodbObj *DynamodbObj) BatchWriteItem() {

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			// "test" is a table name
			"test": {
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"id": {
								S: aws.String("Somewhat Famous"),
							},
							"name": {
								S: aws.String("name"),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"id": {
								S: aws.String("Songs About Life"),
							},
						},
					},
				},
				{
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{
							"id": {
								S: aws.String("Blue Sky Blues"),
							},
						},
					},
				},
			},
		},
	}

	batchWriteItemOutput, err := dynamodbObj.agent.BatchWriteItem(batchWriteItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(batchWriteItemOutput)

}
