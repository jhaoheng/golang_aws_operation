package sqsagent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSAGENT struct {
	Svc *sqs.SQS
}

func NewSQSAgent() *SQSAGENT {

	keyID := "local"
	keySecret := "local"

	creds := credentials.NewStaticCredentials(keyID, keySecret, "")

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
		Endpoint:    aws.String("http://localhost:9324"),
		Credentials: creds,
	}))

	svc := sqs.New(sess)

	return &SQSAGENT{
		Svc: svc,
	}
}
