package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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
	fmt.Println("===send once sqs===")
	q.sendMessage()
	fmt.Println("===send batch sqs===")
	q.sendBatchMessage()
}

func (q *Queue) sendMessage() {

	var msg = Message{
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

func (q *Queue) sendBatchMessage() {

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
		QueueUrl: aws.String(q.URL),
	}
	batchMsgOutput, err := q.Client.SendMessageBatch(batchMsgInput)
	if err != nil {
		panic(err)
	}
	fmt.Println(batchMsgOutput)
}
