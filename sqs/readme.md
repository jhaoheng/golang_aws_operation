# how
- 本地端啟用 docker-compose.yml
- 透過 sqs/main.go, 將 msg 送到本地端的 sqs service
    - 需要注意的是 credential 的設定
- sqs doc : https://docs.aws.amazon.com/sdk-for-go/api/service/sqs/

# send
- sendMessage()
- sendBatchMessage() : 最多十筆資料，同時