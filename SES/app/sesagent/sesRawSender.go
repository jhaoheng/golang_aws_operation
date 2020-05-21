package sesagent

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
)

type SES_RAW_AGENT struct {
	From         string
	ToAddresses  []string
	CcAddresses  []string
	BccAddresses []string
}

func New_SES_RAW_AGENT(from string, toAddresses []string) *SES_RAW_AGENT {
	agent := &SES_RAW_AGENT{
		From:        from,
		ToAddresses: toAddresses,
	}
	return agent
}

func (SESAgent *SES_RAW_AGENT) SetCc(ccAddresses []string) {
	SESAgent.CcAddresses = ccAddresses
}

func (SESAgent *SES_RAW_AGENT) SetBcc(bccAddresses []string) {
	SESAgent.BccAddresses = bccAddresses
}

func (SESAgent *SES_RAW_AGENT) SendRawEmial(email []byte) (*sesv2.SendEmailOutput, error) {

	//
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
	svc := sesv2.New(sess)
	//
	SendEmailInput := &sesv2.SendEmailInput{
		Content: &sesv2.EmailContent{
			Raw: &sesv2.RawMessage{
				Data: email,
			},
		},
		FromEmailAddress: aws.String(SESAgent.From),
	}

	destination := &sesv2.Destination{}
	if len(SESAgent.ToAddresses) != 0 {
		destination = destination.SetToAddresses(aws.StringSlice(SESAgent.ToAddresses))
	}
	if len(SESAgent.CcAddresses) != 0 {
		destination.SetCcAddresses(aws.StringSlice(SESAgent.CcAddresses))
	}
	if len(SESAgent.BccAddresses) != 0 {
		destination.SetBccAddresses(aws.StringSlice(SESAgent.BccAddresses))
	}
	SendEmailInput.SetDestination(destination)

	return svc.SendEmail(SendEmailInput)
}
