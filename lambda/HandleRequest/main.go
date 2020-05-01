package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	IsTest bool
	name   string
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println("IsTest : ", event)

	if event.IsTest {
		b, _ := json.Marshal(event)
		return string(b), nil
	}

	//
	return "", nil
}

func main() {

	lambda.Start(HandleRequest)
}
