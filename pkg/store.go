package pkg

import "github.com/aws/aws-sdk-go-v2/service/s3"

type Store interface {
	S3Client() *s3.Client
	Save(prefix string, key string, body []byte) (storePath string, err error)
}
