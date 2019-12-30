package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

/*
CreateBucket ...
*/
func (client *S3Client) CreateBucket(BucketName string) (*s3.CreateBucketOutput, error) {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(BucketName),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String(client.Region),
		},
	}
	result, err := client.SVCObject.CreateBucket(input)
	return result, err
}

/*
DeleteBucket ...
*/
func (client *S3Client) DeleteBucket(BucketName string) (*s3.DeleteBucketOutput, error) {
	input := &s3.DeleteBucketInput{
		Bucket: aws.String(BucketName),
	}
	result, err := client.SVCObject.DeleteBucket(input)
	return result, err
}

/*
ListBuckets ...
*/
func (client *S3Client) ListBuckets() (*s3.ListBucketsOutput, error) {
	input := &s3.ListBucketsInput{}
	result, err := client.SVCObject.ListBuckets(input)
	return result, err
}

/*
ExistBucket ...
*/
func (client *S3Client) ExistBucket(BucketName string) (*s3.HeadBucketOutput, error) {

	input := &s3.HeadBucketInput{
		Bucket: aws.String(BucketName),
	}
	result, err := client.SVCObject.HeadBucket(input)
	return result, err
}
