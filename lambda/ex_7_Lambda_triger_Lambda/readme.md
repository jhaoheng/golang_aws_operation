# 如何透過 lambda triger lambda?

## 使用 gosdk : lambda
- https://github.com/aws/aws-sdk-go/blob/v1.30.29/service/lambda/api.go#L2467
- 目的在程式中調用外部的 lambda func
- 心得 : 
    - 在 lambda 中調用 lambda 可能造成耦合性可能過高
    - 若第二個 lambda 發生未知的 lambda limitation (ex: timeout), 則會無法得知 ... 必須要透過 cloudwatch 判斷

## build
- A: `docker exec -it app sh -c "cd appFunc_A && go build -o /root/lambda/funcA main.go"`
- B: `docker exec -it app sh -c "cd appFunc_B && go build -o /root/lambda/funcB main.go"`


## test
- A: `curl -d '{}' http://localhost:9001/2015-03-31/functions/myfunction/invocations`
- B: `curl -d '{"name":"max"}' http://localhost:9002/2015-03-31/functions/myfunction/invocations`