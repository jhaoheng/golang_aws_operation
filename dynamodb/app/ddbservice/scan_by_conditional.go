package ddbservice

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type conditional string

var (
	Greater_than conditional = ">="
	Lower_than   conditional = "<="
)

func (s *DDBSERVICE) ScanByCreatedTs(cond conditional, ts int64, limitSize int64, exclusiveStartKey map[string]*dynamodb.AttributeValue) (*dynamodb.ScanOutput, error) {
	scanInput := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":created_ts_bounder": {
				N: aws.String(strconv.Itoa(int(ts))),
			},
		},
		FilterExpression:       aws.String(fmt.Sprintf("created_ts %v :created_ts_bounder", cond)),
		Limit:                  aws.Int64(limitSize),
		TableName:              aws.String(s.table),
		ReturnConsumedCapacity: aws.String("TOTAL"),
		ExclusiveStartKey:      exclusiveStartKey,
	}
	return s.agent.Scan(scanInput)
}
