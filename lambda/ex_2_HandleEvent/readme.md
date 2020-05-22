# 目的

# flow
1. `docker-compose up -d`
2. build : `docker exec app go build -o /root/lambda main.go`
3. test
    - curl : `curl -d '{"istest":true, "name":"max"}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
    - awscli : `aws --endpoint-url http://localhost:9001 --region us-east-1 lambda invoke --function-name main --no-sign-request --payload '{"istest":true, "name":"max"}' --cli-binary-format raw-in-base64-out /dev/stdout`