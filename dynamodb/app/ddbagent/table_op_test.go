package DDBAgent

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Test_CreateTable(t *testing.T) {
	table_name := "my_table_" + time.Now().Format("20060102150405")
	output, err := NewDDBAgent(table_name).CreateTable()
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(output, "", "	")
	fmt.Println(string(b))
}

func Test_ExistTable(t *testing.T) {
	table_name := "my_table"
	output, err := NewDDBAgent(table_name).DescribeTable()
	if awsErr, ok := err.(awserr.Error); ok {
		if awsErr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			fmt.Println("找不到 table")
			t.Fatal(err)
		}
	}
	b, _ := json.MarshalIndent(output, "", "	")
	fmt.Println(string(b))
}
