package bucket

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Bucket struct {
	s3   *s3.S3
	name string
	url  string
}

func NewS3(s3 *s3.S3, name, url string) Bucket {
	return &S3Bucket{s3, name, url}
}

func (bk *S3Bucket) Get(key string) (io.ReadCloser, error) {
	o, err := bk.s3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bk.name),
		Key:    aws.String(key),
	})

	if err != nil {
		return nil, err
	}

	return o.Body, nil
}

func (bk *S3Bucket) Put(key string, r io.ReadSeeker) error {
	_, err := bk.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bk.name),
		Key:    aws.String(key),
		Body:   r,
	})

	if err != nil {
		return err
	}

	return nil
}

func (bk *S3Bucket) URL(key string) string {
	return bk.url + key
}
