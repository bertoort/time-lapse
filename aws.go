package main

import (
	"errors"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Image information from bucket list
type Image struct {
	ETag         string
	Key          string
	LastModified time.Time
	Size         int64
	StorageClass string
}

// s3BucketList gets a list of image keys from a bucket
func s3BucketList(bucket string) (resp *s3.ListObjectsV2Output, err error) {
	if bucket == "" {
		return nil, errors.New("No Bucket Found")
	}

	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	awsKey := os.Getenv("AWS_ACCESS_KEY_ID")
	awsSecret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	svc := s3.New(sess, &aws.Config{
		Region:      aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(awsKey, awsSecret, ""),
	})

	params := &s3.ListObjectsV2Input{
		Bucket:       aws.String(bucket),
		EncodingType: aws.String("url"),
	}
	resp, err = svc.ListObjectsV2(params)

	if err != nil {
		return nil, err
	}

	return
}
