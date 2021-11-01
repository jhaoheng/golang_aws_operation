# build

GOOS=linux GOARCH=amd64 go build -o ./lambda/func ./app/main.go

# test it with `lambci/lambda:go1.x`

docker run --rm -v $(pwd)/lambda/func:/var/task lambci/lambda:go1.x main

# run it with sam

1. build template



curl -d '{}' http://localhost:3001/2015-03-31/functions/appHandler/invocations

aws --endpoint-url http://localhost:3001 --region us-east-1 lambda invoke --function-name main --no-sign-request --payload '{}' /dev/stdout