package sqsagent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func (q *SQSAGENT) GetQueueAttributes(qURL string) (*sqs.GetQueueAttributesOutput, error) {

	getQueueAttributesInput := &sqs.GetQueueAttributesInput{
		AttributeNames: aws.StringSlice([]string{"All"}),
		QueueUrl:       aws.String(qURL),
	}
	return q.Svc.GetQueueAttributes(getQueueAttributesInput)
}
