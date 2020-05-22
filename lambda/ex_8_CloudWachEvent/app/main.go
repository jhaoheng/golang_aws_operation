package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context) {

}

func main() {
	lambda.Start(Handler)
}
