## build 
> 注意 build func 與 layer 必須要在同一個環境

1. `docker run --rm -it -v $(pwd):/go/src/app -w /go/src/app golang:1.13.1 /bin/bash`
2. `go build -o ./release/task/main .`
3. `go build -buildmode=plugin -o ./release/opt/layer.so ./layer/helloLayer.go`

# run local test

```
docker run --rm -v $(pwd)/release/task:/var/task:ro,delegated -v $(pwd)/release/opt:/opt:ro,delegated lambci/lambda:go1.x main
```