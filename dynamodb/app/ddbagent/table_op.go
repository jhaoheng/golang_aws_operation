package DDBAgent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func (agent *DDBAgent) CreateTable() (*dynamodb.CreateTableOutput, error) {
	input := &dynamodb.CreateTableInput{
		BillingMode: aws.String("PAY_PER_REQUEST"),
		TableName:   aws.String(agent.Table),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			0: {
				AttributeName: aws.String("id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			0: {
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
	}
	if err := input.Validate(); err != nil {
		return nil, err
	}

	return agent.Agent.CreateTable(input)
}

/*
- 可用於查看 table 是否存在
- 若 table 不存在, 則 err = ResourceNotFoundException: Cannot do operations on a non-existent table

	if awsErr, ok := err.(awserr.Error); ok {
			if awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
				fmt.Println("找不到 table")
				t.Fatal(err)
			}
	}
*/
func (agent *DDBAgent) DescribeTable() (*dynamodb.DescribeTableOutput, error) {
	input := &dynamodb.DescribeTableInput{
		TableName: aws.String(agent.Table),
	}
	output, err := agent.Agent.DescribeTable(input)
	if err != nil {
		return nil, err
	}
	return output, err
}
