# how
- 本地端啟用 docker-compose.yml
- 透過 sqs/main.go, 將 msg 送到本地端的 sqs service
    - 需要注意的是 credential 的設定
- sqs admin : `http://localhost:9325`
- sqs doc : https://docs.aws.amazon.com/sdk-for-go/api/service/sqs/

# send
- sendMessage()
- sendBatchMessage() : 最多十筆資料，同時
- receiveMessage() : 最多同時接收十筆資料
- deleteMessage() : 最多同時刪除十筆資料  
- getQueueAttributes() : 取得目前 Queue 中的狀態, 可得到等待的數量

# getQueueAttributes()
> Response is below

```
{
  Attributes: {
    VisibilityTimeout: "10",
    ReceiveMessageWaitTimeSeconds: "0",
    ApproximateNumberOfMessages: "1", // 等待
    ApproximateNumberOfMessagesNotVisible: "0",
    ApproximateNumberOfMessagesDelayed: "0",
    LastModifiedTimestamp: "1572356863",
    DelaySeconds: "5",
    CreatedTimestamp: "1572356863",
    QueueArn: "arn:aws:sqs:elasticmq:000000000000:default"
  }
}
```