package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

var url string = "http://localhost:9324/queue/default"

type Queue struct {
	Client sqsiface.SQSAPI
	URL    string
}

// Message is a concrete representation of the SQS message
type Message struct {
	Url  string
	Time time.Time
	From string
}

func main() {
	// Create a Session with a custom region

	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("eu-central-1"),
		Endpoint: aws.String("http://sqs:9324"),
	}))

	q := Queue{
		Client: sqs.New(sess),
		URL:    url,
	}
	q.sendMessage()
}

func (q *Queue) sendMessage() {

	var msg = Message{
		Url:  "",
		Time: time.Now(),
		From: "12345",
	}

	b, _ := json.Marshal(msg)
	str := string(b)

	sendMsgInput := &sqs.SendMessageInput{
		MessageBody: aws.String(str),
		QueueUrl:    aws.String(q.URL),
	}
	sendMessageOutput, err := q.Client.SendMessage(sendMsgInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(sendMessageOutput)
}
