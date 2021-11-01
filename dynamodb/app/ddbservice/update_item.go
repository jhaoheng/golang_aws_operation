package ddbservice

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (ddb *DDBSERVICE) UpdateItem() {
	var utc_8, _ = time.ParseDuration("+8h") // utc+8 時區
	updateItemInput := &dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#Updated_at": aws.String("updated_at"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Updated_value": {
				S: aws.String(time.Now().UTC().Add(utc_8).Format("2006-01-02 15:04:05")),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1"),
			},
		},
		ReturnValues:     aws.String("ALL_NEW"),
		TableName:        aws.String(ddb.table),
		UpdateExpression: aws.String("SET #Updated_at = :Updated_value"),
	}

	updateItemOutput, err := ddb.agent.UpdateItem(updateItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(updateItemOutput)
}
