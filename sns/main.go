package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {

	//
	make_method_2_sns_mobile_push()
}

type MYSNS struct {
	snsClient *sns.SNS
}

/*
實踐方法一
*/
func make_method_1_sns_topic_publish() {
	dat, _ := ioutil.ReadFile("./message")

	//
	region := "ap-southeast-1"
	endpoint := "http://sns:9911"
	subject := "test"
	topicArn := "arn:aws:sns:ap-southeast-1:478205036267:atlas-issue"
	mysns := NewSNS(region, endpoint)
	mysns.publish(string(dat), subject, topicArn)
}

func make_method_2_sns_mobile_push() {
	region := "ap-southeast-1"
	endpoint := ""
	mysns := NewSNS(region, endpoint)
	//
	platform_application_arn := "arn:aws:sns:ap-southeast-1:424613967558:app/GCM/fcm"
	device_token := "token"
	input := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: aws.String(platform_application_arn),
		Token:                  aws.String(device_token),
	}
	output, err := mysns.snsClient.CreatePlatformEndpoint(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(*output.EndpointArn)
}

func NewSNS(region string, endpoint string) MYSNS {
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
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
