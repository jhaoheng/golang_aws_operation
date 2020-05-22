package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	lambdasdk "github.com/aws/aws-sdk-go/service/lambda"
)

func Handler() (*lambdasdk.InvokeOutput, error) {
	config := &aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String("http://lambdaB:9001"),
	}
	newSession := session.New(config)

	svc := lambdasdk.New(newSession)
	input := &lambdasdk.InvokeInput{
		FunctionName:   aws.String("main"),
		InvocationType: aws.String("Event"),
		Payload:        []byte("{\"name\":\"max\"}"),
	}
	result, err := svc.Invoke(input)
	return result, err
}

func main() {
	lambda.Start(Handler)
}
