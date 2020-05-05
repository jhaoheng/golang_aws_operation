package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var (
	creds *credentials.Credentials
	sess  *session.Session
)

func main() {

	/*
		根據 credential file 中的設定, 可選擇
		- default
		- custom_profile
	*/
	profileName := "default"

	/*
		選擇從 default profile 取得參數, 建立 session
	*/
	sess = GetSession(profileName)

	//
	value, _ := sess.Config.Credentials.Get()

	fmt.Printf("ProviderName : %v\n", value.ProviderName)
	fmt.Printf("AccessKeyID : %v\n", value.AccessKeyID)
	fmt.Printf("SecretAccessKey : %v\n", value.SecretAccessKey)
	fmt.Printf("SecretAccessKey : %v\n", value.SessionToken)
	fmt.Printf("Region : %v\n", *sess.Config.Region)
	fmt.Printf("LogLevel : %v\n", *sess.Config.LogLevel)
}

func GetSession(profile string) *session.Session {

	config := SetConfig()

	/*
		方法一
	*/
	// sess, _ = session.NewSession(&config)

	/*
		方法二
	*/
	sess, _ = session.NewSessionWithOptions(session.Options{
		// SharedConfigState: session.SharedConfigEnable,
		Config:  config,
		Profile: profile, // 若不設定, 則預設為 default
	})
	return sess

}

/*
透過操作 aws.Config 來覆蓋 session 的預設變數取得方法
*/
func SetConfig() aws.Config {
	// 使用 credentials file
	// creds := credentials.NewSharedCredentials("/root/.aws/credentials", "default")
	// 使用 arg
	// creds = credentials.NewStaticCredentials("localFromStaticVariable", "localFromStaticVariable", "")
	// 使用環境變數
	// creds = credentials.NewEnvCredentials()

	// config 參數 : https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
	return aws.Config{
		Endpoint: aws.String(""),
		// Region:      aws.String(""),
		Credentials: creds,
	}
}
