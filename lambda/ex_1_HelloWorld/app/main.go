package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	time.Sleep(time.Second * 1)
	return fmt.Sprintf("Hello ƛ, %v", time.Now()), nil
}

func main() {
	lambda.Start(hello)
}
