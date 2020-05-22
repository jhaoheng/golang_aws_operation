package main

import (
	sqsagent "app/sqsAgent"
	"fmt"
)

func main() {

	// base
	var qUrl string = "http://localhost:9324/queue/logging"
	fmt.Println(sqsagent.NewSQSAgent().SendMsg(qUrl, "hello world SendMsg"))

	// fifo
	var qUrl_fifo string = "http://localhost:9324/queue/logging_fifo"
	fmt.Println(sqsagent.NewSQSAgent().SendMsgToFIFO(qUrl_fifo, "groupId", "hello world fifo"))
}
