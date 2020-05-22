package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	IsTest bool   `json:"istest"`
	Name   string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	fmt.Println(event)

	b, _ := json.Marshal(event)
	return string(b), nil
}

func main() {

	lambda.Start(HandleRequest)
}
