package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

var t float64 = 10 // 程式預設執行時間 second

func hello(ctx context.Context) (string, error) {
	defer func() {
		timeStart := time.Now()
		if ctx.Err() != nil {
			fmt.Println(ctx.Err())
			for {
				fmt.Printf("執行 lambda 產生錯誤後的動作, 模擬處理時間 : %v\n", time.Now())
				time.Sleep(time.Second * 1)
				if time.Now().Sub(timeStart).Seconds() >= 20 {
					break
				}
			}
		}
	}()

	timeStart := time.Now()
	deadline, _ := ctx.Deadline()
	fmt.Println(deadline)
	deadline = deadline.Add(-100 * time.Millisecond)
	timeoutChannel := time.After(time.Until(deadline))
	for {
		select {
		case <-timeoutChannel:
			fmt.Println("觸發 timeout")
			return "", nil // 因為程式還需要結束, 但 lambda 最後會給出 error
		default:
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Sub(timeStart))
			if time.Now().Sub(timeStart).Seconds() >= t {
				goto FINISH
			}
		}
	}

FINISH:
	return "lambda 在指定時間內完成任務", nil
}

func main() {
	lambda.Start(hello)
}
