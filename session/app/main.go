package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	// 使用 file
	// creds := credentials.NewSharedCredentials("/root/.aws/credentials", "default")

	// 使用 arg
	// creds := credentials.NewStaticCredentials(keyID, keySecret, "")

	// config 參數 : https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
	config := aws.Config{
		Endpoint: aws.String(""),
		// Credentials: creds,
	}

	sess := session.Must(session.NewSession(&config))

	value, _ := sess.Config.Credentials.Get()
	fmt.Printf("%#v\n", value)
}
