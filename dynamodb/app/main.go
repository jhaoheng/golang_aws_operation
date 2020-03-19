package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DynamodbObj struct {
	agent *dynamodb.DynamoDB
	table string
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

	datas := []map[string]string{
		0: {
			"name": "max",
		},
		1: {
			"name": "sunny",
		},
	}
	dynamodbObj.BatchWriteItem(datas)

	dynamodbObj.GetItem()

	dynamodbObj.UpdateItem()

	dynamodbObj.BatchGetItem()
}

// 屬性類型 : https://docs.aws.amazon.com/sdk-for-go/api/service/dynamodb/#AttributeValue
func (dynamodbObj *DynamodbObj) ScanIsExist(key, value string) bool {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(dynamodbObj.table),
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

func (dynamodbObj *DynamodbObj) GetItem() {
	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String("1"),
			},
		},
		TableName: aws.String(dynamodbObj.table),
	}
	getItemOutput, err := dynamodbObj.agent.GetItem(getItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(getItemOutput)
	if len(getItemOutput.Item) == 0 {
		fmt.Println("don't exist")
	} else {
		fmt.Println("exist")
	}
}

func (dynamodbObj *DynamodbObj) BatchWriteItem(datas []map[string]string) {
	var requestItems = []*dynamodb.WriteRequest{}
	for index, value := range datas {
		writeRequest := dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: map[string]*dynamodb.AttributeValue{
					"id": {
						S: aws.String(strconv.Itoa(index)),
					},
					"name": {
						S: aws.String(value["name"]),
					},
				},
			},
		}
		requestItems = append(requestItems, &writeRequest)
	}

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			dynamodbObj.table: requestItems,
		},
	}

	batchWriteItemOutput, err := dynamodbObj.agent.BatchWriteItem(batchWriteItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(batchWriteItemOutput)
}

func (dynamodbObj *DynamodbObj) UpdateItem() {
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
		TableName:        aws.String(dynamodbObj.table),
		UpdateExpression: aws.String("SET #Updated_at = :Updated_value"),
	}

	updateItemOutput, err := dynamodbObj.agent.UpdateItem(updateItemInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(updateItemOutput)
}

/*

 */
func (dynamodbObj *DynamodbObj) BatchGetItem() {
	//
	datas := []string{
		"max",
		"sunny",
	}

	//
	var keys = []map[string]*dynamodb.AttributeValue{}
	for _, value := range datas {
		key := map[string]*dynamodb.AttributeValue{
			"id": &dynamodb.AttributeValue{
				S: aws.String(value),
			},
		}
		keys = append(keys, key)
	}

	batchGetItemInput := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{
			dynamodbObj.table: {
				Keys: keys,
			},
		},
	}

	batchGetItemOutput, err := dynamodbObj.agent.BatchGetItem(batchGetItemInput)
	if err != nil {
		panic(err)
	}

	fmt.Println(batchGetItemOutput)
}

func (dynamodbObj *DynamodbObj) CheckNotExists() (*dynamodb.ScanOutput, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String(dynamodbObj.table),
		Select:    aws.String("ALL_ATTRIBUTES"),
		ExpressionAttributeNames: map[string]*string{
			"#Status": aws.String("status"),
		},
		FilterExpression: aws.String("attribute_not_exists(#Status)"),
	}
	scanOutput, err := dynamodbObj.agent.Scan(scanInput)
	return scanOutput, err
}
