# 目的
- 建立 lambda func / layer
- 讓 func 可以執行 layer 的操作

# flow

1. `docker-compose up -d`
2. build : 
    - func : `docker exec app sh -c "cd appFunc && go build -o /root/lambda/func main.go"`
    - layer : `docker exec app sh -c "cd appLayer && go build -buildmode=plugin -o /root/lambda/layer/layer.so helloLayer.go"`
3. test
    - `curl -d '{"func":"Func_1"}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`

## 注意
- build func 與 layer 必須要在同一個環境
- main func 回傳一定要包含 error