.PHONY: deps clean build

clean: 

	rm -rf ./release/*
	
build:
	GOOS=linux GOARCH=amd64 go build -o release/getHandler ./app/get/main.go
	GOOS=linux GOARCH=amd64 go build -o release/delHandler ./app/del/main.go
	GOOS=linux GOARCH=amd64 go build -o release/postHandler ./app/post/main.go
	GOOS=linux GOARCH=amd64 go build -o release/cuseventHandler ./app/cusEvent/main.go
	GOOS=linux GOARCH=amd64 go build -o release/s3eventHandler ./app/s3Event/main.go