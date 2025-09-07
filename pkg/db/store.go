package db

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"strings"
)

type S3Client struct {
	config pkg.Storage
	logger pkg.Logger

	*s3.Client
}

func NewS3Client(config pkg.Storage, logger pkg.Logger) (*S3Client, error) {
	options := s3.Options{
		UsePathStyle: true,
	}
	if config.Region != "" {
		options.Region = config.Region
	}
	if config.Id != "" && config.Key != "" {
		options.Credentials = credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     config.Id,
				SecretAccessKey: config.Key,
				SessionToken:    "",
			},
		}
	}
	if config.Endpoint != "" {
		options.EndpointResolver = s3.EndpointResolverFromURL(config.Endpoint)
	}

	return &S3Client{
		config: config,
		logger: logger,
		Client: s3.New(options),
	}, nil
}

func (s *S3Client) Save(prefix string, key string, body []byte) (storePath string, err error) {
	uploadParams := &s3.PutObjectInput{
		Key:  &key,
		Body: bytes.NewReader(body),
	}
	if prefix == "" {
		prefix = s.config.Bucket
	}
	prefix = strings.Trim(prefix, "/")
	uploadParams.Bucket = &prefix

	// Upload the file
	_, err = s.PutObject(context.Background(), uploadParams)
	if err != nil {
		s.logger.Error(err)
		return
	}
	storePath = fmt.Sprintf("%s://%s/%s", s.config.Type, prefix, key)
	return
}

func (s *S3Client) Close(_ context.Context) (err error) {
	return
}

type FileClient struct {
	config pkg.Storage
	logger pkg.Logger
}

func NewFileClient(config pkg.Storage, logger pkg.Logger) (*FileClient, error) {
	return &FileClient{
		config: config,
		logger: logger,
	}, nil
}

func (s *FileClient) Save(prefix string, key string, body []byte) (storePath string, err error) {
	if s.config.Endpoint != "" {
		prefix = strings.TrimPrefix(s.config.Endpoint, "file:/")
	}
	if prefix == "" {
		name, _ := os.MkdirTemp("", s.config.Bucket)
		prefix = name
	}
	if prefix == "" {
		prefix = os.TempDir()
	}
	prefix = strings.Trim(prefix, "/")

	filename := fmt.Sprintf("/%s/%s", prefix, key)
	if !utils.ExistsDir(filename) {
		err = os.MkdirAll(filepath.Dir(filename), 0744)
		if err != nil {
			s.logger.Error(err)
			return
		}
	}

	var file *os.File
	if !utils.ExistsFile(filename) {
		file, err = os.Create(filename)
		if err != nil {
			s.logger.Error(err)
			return
		}
	} else {
		file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			s.logger.Error(err)
			return
		}
	}
	_, err = file.Write(body)
	if err != nil {
		s.logger.Error(err)
		return
	}
	storePath = fmt.Sprintf("file:/%s", filename)
	return
}

func (s *FileClient) Close(_ context.Context) (err error) {
	return
}
