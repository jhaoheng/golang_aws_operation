package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	defaultS3Region   = "ap-southeast-1"
	defaultS3Endpoint = "http://minio:9000"
)

/*
S3Client ...
*/
type S3Client struct {
	IsLocal   bool
	SVCObject *s3.S3
	Region    string
}

type S3CONFIG struct {
	Credentials struct {
		Id     string
		Secret string
	}
	Endpoint         string // minio
	DisableSSL       bool   // minio
	S3ForcePathStyle bool   // ninio
}

/*
NewSVC ...
Location : https://docs.aws.amazon.com/zh_cn/general/latest/gr/rande.html
*/
func NewS3Service(isLocal bool, region string) (client S3Client) {
	var awsS3Config = &aws.Config{}
	if isLocal {
		awsS3Config = &aws.Config{
			Region:           aws.String("us-east-1"),
			Credentials:      credentials.NewStaticCredentials("minio", "miniominio", ""),
			Endpoint:         aws.String("http://minio:9000"),
			DisableSSL:       aws.Bool(false),
			S3ForcePathStyle: aws.Bool(true),
		}
	} else {
		awsS3Config = &aws.Config{
			Region: aws.String(region),
		}
	}

	newSession := session.New(awsS3Config)
	svcObject := s3.New(newSession)
	client = S3Client{
		IsLocal:   isLocal,
		SVCObject: svcObject,
		Region:    region,
	}
	return
}

func NewS3ServiceWithKey(Region string, Endpoint string, keyID, keySecret string) (client S3Client) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Endpoint:    aws.String(Endpoint),
		Credentials: credentials.NewStaticCredentials(keyID, keySecret, ""),
	})

	// Create S3 service client
	SVCObject := s3.New(sess)
	client = S3Client{
		SVCObject: SVCObject,
		Region:   Region,
	}
	return
}

func NewS3ServiceWithShared(Region string, Endpoint string, sharedPath string, sharedName string) (client S3Client) {
	creds := credentials.NewSharedCredentials(sharedPath, sharedName)
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Endpoint:    aws.String(Endpoint),
		Credentials: creds,
	})

	// Create S3 service client
	SVCObject := s3.New(sess)
	client = S3Client{
		SVCObject: SVCObject,
		Region:   Region,
	}
	return
}
