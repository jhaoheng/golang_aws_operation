## 限制是
- 一個 go binary 算一個 lambda = 一個 container

## 執行方法
1. 建立環境
2. 進入 container:app
    - `go get github.com/aws/aws-lambda-go/lambda`
3. `GOOS=linux GOARCH=amd64 go build -o ./release/main ./app/main.go`
4. 查看 lambda log : `docker logs -f lambda`
5. 調用 lambda
    - 使用 curl : `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
    - 使用 awscli : `aws --endpoint-url http://lambda:9001 --region us-east-1 lambda invoke --function-name {handler} --no-sign-request --payload '{}' /dev/stdout`