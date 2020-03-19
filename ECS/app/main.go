package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
)

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		// Region: aws.String(""),
		// Endpoint: aws.String(""),
	}))
	// Create a ECS client from just a session.
	svc := ecs.New(sess)
	fmt.Println(svc)
}
