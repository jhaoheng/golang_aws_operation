.PHONY: clean build

clean: 
	rm -rf ./release/*
	

build:
	@if [ "$$Lambda" ]; then\
		echo "var is 123";\
		handler=$(Lambda)Handler;\
		echo $$handler;\
		GOOS=linux GOARCH=amd64 go build -o release/$$handler ./app/$$Lambda/main.go;\
	else\
		GOOS=linux GOARCH=amd64 go build -o release/getHandler ./app/get/main.go;\
		GOOS=linux GOARCH=amd64 go build -o release/delHandler ./app/del/main.go;\
		GOOS=linux GOARCH=amd64 go build -o release/postHandler ./app/post/main.go;\
		GOOS=linux GOARCH=amd64 go build -o release/cuseventHandler ./app/cusEvent/main.go;\
		GOOS=linux GOARCH=amd64 go build -o release/s3eventHandler ./app/s3Event/main.go;\
	fi

