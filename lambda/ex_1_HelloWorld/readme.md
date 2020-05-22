# 目的
- 在本地端建立 hello Lambda, 並且在本地端測試

# flow
1. `docker-compose up -d`
2. develop app
    1. `docker exec -it app /bin/bash`
    2. update code
    3. compile to lambda : `go build -o /root/lambda main.go`
3. test lambda
    - curl : `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
    - aws cli : 
        - `aws --endpoint-url http://localhost:9001 --region us-east-1 lambda invoke --function-name main --no-sign-request --payload '{}' /dev/stdout`

# ps: 使用 docker run 測試 lambda func
- `docker run --rm -v $(pwd)/lambda/task:/var/task lambci/lambda:go1.x main`