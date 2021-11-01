package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

/*
type APIGatewayProxyRequest struct {
    Resource                        string                        `json:"resource"` // The resource path defined in API Gateway
    Path                            string                        `json:"path"`     // The url path for the caller
    HTTPMethod                      string                        `json:"httpMethod"`
    Headers                         map[string]string             `json:"headers"`
    MultiValueHeaders               map[string][]string           `json:"multiValueHeaders"`
    QueryStringParameters           map[string]string             `json:"queryStringParameters"`
    MultiValueQueryStringParameters map[string][]string           `json:"multiValueQueryStringParameters"`
    PathParameters                  map[string]string             `json:"pathParameters"`
    StageVariables                  map[string]string             `json:"stageVariables"`
    RequestContext                  APIGatewayProxyRequestContext `json:"requestContext"`
    Body                            string                        `json:"body"`
    IsBase64Encoded                 bool                          `json:"isBase64Encoded,omitempty"`
}
*/

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("databaseName : %v\n\n", os.Getenv("databaseName"))
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, del"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
