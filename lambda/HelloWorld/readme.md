## build binary
- `GOOS=linux go build -o ./release/task/main main.go`

## 本地測試 : `docker run` 
- `docker run --rm -v $(pwd)/release/task:/var/task lambci/lambda:go1.x main`

## 本地測試 : docker-compose
1. 選擇 func
2. 建立 binary : `GOOS=linux GOARCH=amd64 go build -o main main.go`
3. 將 binary 移動到 release folder
4. 查看 lambda log : `docker logs -f lambda`
5. 調用 lambda
    - 使用 curl : `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
    - 使用 awscli : `aws --endpoint-url http://lambda:9001 --region us-east-1 lambda invoke --function-name {handler} --no-sign-request --payload '{}' /dev/stdout`