package sqsagent

import (
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (q *SQSAGENT) SendMsgToFIFO(qURL string, groupId, message string) (*sqs.SendMessageOutput, error) {
	sendMsgInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(message),
		QueueUrl:       aws.String(qURL),
		MessageGroupId: aws.String(groupId),
	}
	return q.Svc.SendMessage(sendMsgInput)
}

func (q *SQSAGENT) SendMsg(qURL string, message string) (*sqs.SendMessageOutput, error) {
	sendMsgInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		QueueUrl:    aws.String(qURL),
	}
	return q.Svc.SendMessage(sendMsgInput)
}

func (q *SQSAGENT) SendBatchMessage(qURL string) (*sqs.SendMessageBatchOutput, error) {

	entries := []*sqs.SendMessageBatchRequestEntry{}

	// 最多十筆
	for i := 0; i < 10; i++ {
		entry := sqs.SendMessageBatchRequestEntry{
			Id:          aws.String(strconv.Itoa(i)),
			MessageBody: aws.String(time.Now().String()),
		}
		entries = append(entries, &entry)
	}

	batchMsgInput := &sqs.SendMessageBatchInput{
		Entries:  entries,
		QueueUrl: aws.String(qURL),
	}
	return q.Svc.SendMessageBatch(batchMsgInput)
}
