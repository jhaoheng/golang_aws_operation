## build

`GOOS=linux go build -o ./release/task/main main.go`

## run localhost test

`docker run --rm -v $(pwd)/release/task:/var/task lambci/lambda:go1.x main`