package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jhaoheng/gobuildpkgdemo/pkg"
)

func Handler() (string, error) {
	pkg.Printhello()
	return "hello, what are you doing!", nil
}

func main() {
	lambda.Start(Handler)
}
