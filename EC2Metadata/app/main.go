package main

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
)

// https://docs.aws.amazon.com/sdk-for-go/api/aws/ec2metadata/#EC2InstanceIdentityDocument
func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		// Region: aws.String(""),
		// Endpoint: aws.String(""),
	}))
	EC2Metadata := ec2metadata.New(sess)

	value, _ := EC2Metadata.Config.Credentials.Get()

	fmt.Println("IsEC2 ? ", IsEc2(&value))

	if IsEc2(&value) {
		InstanceID := GetInstanceID(EC2Metadata)
		fmt.Println("InstanceID => ", InstanceID)
	}

}

func IsEc2(value *credentials.Value) bool {
	if strings.HasPrefix(value.AccessKeyID, "local") && strings.HasPrefix(value.SecretAccessKey, "local") {
		return false
	}
	return true
}

func GetInstanceID(ec2metadata *ec2metadata.EC2Metadata) string {
	Doc, _ := ec2metadata.GetInstanceIdentityDocument()
	return Doc.InstanceID
}
