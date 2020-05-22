package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Message string `json:"message"`
}

func handler(ctx context.Context, event MyEvent) (MyEvent, error) {
	fmt.Println("event : ", event)
	return event, nil
}

func main() {
	lambda.Start(handler)
}
