package sesagent

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sesv2"
)

type SES_SIMPLE_AGENT struct {
	From string
	To   []string

	HtmlBody string
	TextBody string
	Subject  string
	CharSet  string
}

func New_SES_SIMPLE_AGENT(from string, to []string) *SES_SIMPLE_AGENT {
	return &SES_SIMPLE_AGENT{
		From: from,
		To:   to,
	}
}

func (agent *SES_SIMPLE_AGENT) SetEmail(subject, htmlBody, textBody, charSet string) {
	agent.HtmlBody = htmlBody
	agent.TextBody = textBody
	agent.Subject = subject
	agent.CharSet = charSet
}

func (agent *SES_SIMPLE_AGENT) SendSimpleEmail() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	// Create a SESV2 client with additional configuration
	svc := sesv2.New(sess)

	SendEmailInput := sesv2.SendEmailInput{
		Content: &sesv2.EmailContent{
			Simple: &sesv2.Message{
				Body: &sesv2.Body{
					Html: &sesv2.Content{
						Charset: aws.String(agent.CharSet),
						Data:    aws.String(agent.HtmlBody),
					},
					Text: &sesv2.Content{
						Charset: aws.String(agent.CharSet),
						Data:    aws.String(agent.TextBody),
					},
				},
				Subject: &sesv2.Content{
					Charset: aws.String(agent.CharSet),
					Data:    aws.String(agent.Subject),
				},
			},
		},
		Destination: &sesv2.Destination{
			ToAddresses: aws.StringSlice(agent.To),
		},
		FromEmailAddress: aws.String(agent.From),
	}

	SendEmailOutput, err := svc.SendEmail(&SendEmailInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(SendEmailOutput)
}
