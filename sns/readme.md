# doc
https://docs.aws.amazon.com/sdk-for-go/api/service/sns/

# 設定 sns
1. 進入 sns 頁面
2. 設定 topic 
3. 設定 subscription

# 執行測試

## 用 golang 來執行 aws sns api
1. `docker-compose up -d`
2. 在 `.aws/` 中新增 `credentials` 檔案
```
[default]
aws_access_key_id=
aws_secret_access_key=
```
3. 在 main.go 填寫好
    - region
    - topicArn
4. 執行 `go run main.go`

## 使用 awscli 來執行 aws sns api
1. `docker-compose up -d`
2. 在 `.aws/` 中新增 `credentials` 檔案
3. `aws --region=ap-southeast-1 sns publish --topic-arn=arn:aws:sns:ap-southeast-1:478205036267:atlas-issue --message="hello from cli"`