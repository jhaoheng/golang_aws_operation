package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	dat, _ := ioutil.ReadFile("./message")

	//
	region := "ap-southeast-1"
	subject := "test"
	topicArn := "arn:aws:sns:ap-southeast-1:478205036267:atlas-issue"
	mysns := NewSNS(region)
	mysns.publish(string(dat), subject, topicArn)
}

type MYSNS struct {
	snsClient *sns.SNS
}

func NewSNS(region string) MYSNS {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc := sns.New(sess)

	mysns := MYSNS{
		snsClient: svc,
	}
	return mysns
}

func (mysns *MYSNS) publish(message string, subject, topicArn string) {
	publishInput := sns.PublishInput{
		Message:  aws.String(message),
		Subject:  aws.String(subject),
		TopicArn: aws.String(topicArn),
	}
	publishOutput, err := mysns.snsClient.Publish(&publishInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(publishOutput)
}
