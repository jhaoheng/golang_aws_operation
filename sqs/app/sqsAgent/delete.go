package sqsagent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (q *SQSAGENT) DeleteMsgs(qURL string, messages []*sqs.Message) (*sqs.DeleteMessageBatchOutput, error) {
	entries := []*sqs.DeleteMessageBatchRequestEntry{}
	for _, v := range messages {
		entry := sqs.DeleteMessageBatchRequestEntry{
			Id:            aws.String(*v.MessageId),
			ReceiptHandle: aws.String(*v.ReceiptHandle),
		}
		entries = append(entries, &entry)
	}

	deleteMsgBatchInput := &sqs.DeleteMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(qURL),
	}
	return q.Svc.DeleteMessageBatch(deleteMsgBatchInput)
}
