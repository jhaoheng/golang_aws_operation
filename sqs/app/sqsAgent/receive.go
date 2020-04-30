package sqsagent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (q *SQSAGENT) ReceiveMessage(qURL string) (*sqs.ReceiveMessageOutput, error) {
	receiveMsgInput := &sqs.ReceiveMessageInput{
		MaxNumberOfMessages: aws.Int64(10), //1~10
		QueueUrl:            aws.String(qURL),
	}
	return q.Svc.ReceiveMessage(receiveMsgInput)
}
