# 目的
- 透過 lambda 使用外部的 pkg

# flow

1. `docker-compose up -d`
2. build
    - `docker exec app go build -o /root/lambda main.go`
3. test
    - `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`