package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func Handler(ctx context.Context, event MyEvent) (string, error) {
	// time.Sleep(time.Second * 3)
	fmt.Println(event)
	b, err := json.Marshal(event)
	return string(b), err
}

func main() {
	lambda.Start(Handler)
}
