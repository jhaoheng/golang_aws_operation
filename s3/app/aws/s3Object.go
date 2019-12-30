package aws

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"

	"bytes"
	"strings"
	"unsafe"
)

/*
ListObjects ...
*/
func (client *S3Client) ListObjects(BucketName string) (*s3.ListObjectsV2Output, error) {
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String(BucketName),
		MaxKeys: aws.Int64(2),
	}
	result, err := client.SVCObject.ListObjectsV2(input)
	return result, err
}

/*
PutObject ...
*/
func (client *S3Client) PutObject(BucketName, Key string, newObj io.Reader) (*s3.PutObjectOutput, error) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(newObj)
	b := buf.Bytes()
	s := *(*string)(unsafe.Pointer(&b))

	input := &s3.PutObjectInput{
		Body: aws.ReadSeekCloser(strings.NewReader(s)),
		// Body:   aws.ReadSeekCloser(newObj),
		Bucket: aws.String(BucketName),
		Key:    aws.String(Key),
	}
	result, err := client.SVCObject.PutObject(input)
	return result, err
}

/*
GetObject ...
*/
func (client *S3Client) GetObject(BucketName string, key string) (*s3.GetObjectOutput, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(key),
	}
	result, err := client.SVCObject.GetObject(input)
	return result, err
}

/*
GetObjectSignedURL ...
*/
func (client *S3Client) GetObjectSignedURL(BucketName string, key string, TTLSec int) string {
	req, _ := client.SVCObject.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(key),
	})
	urlStr, _ := req.Presign(time.Duration(TTLSec) * time.Second)
	return urlStr
}

/*
DeleteOjects ...
*/
func (client *S3Client) DeleteObjects(BucketName string, keys []string) (*s3.DeleteObjectsOutput, error) {

	var objs = make([]*s3.ObjectIdentifier, len(keys))
	for i, v := range keys {
		// Add objects from command line to array
		objs[i] = &s3.ObjectIdentifier{Key: aws.String(v)}
	}

	input := &s3.DeleteObjectsInput{
		Bucket: aws.String(BucketName),
		Delete: &s3.Delete{
			Objects: objs,
			Quiet:   aws.Bool(false),
		},
	}
	result, err := client.SVCObject.DeleteObjects(input)
	return result, err
}
