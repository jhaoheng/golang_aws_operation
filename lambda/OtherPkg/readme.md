1. build : `GOOS=linux go build -o ./release/task/handler .`
2. run : `docker run --rm -v $(pwd)/release/task:/var/task lambci/lambda:go1.x handler`