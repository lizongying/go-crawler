package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"go.uber.org/fx"
)

func NewS3(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (s3Client *s3.Client, err error) {
	if !config.S3Enable {
		logger.Debug("Redis Disable")
		return
	}

	id := config.S3.Example.Id
	if id == "" {
		logger.Warn("id is empty")
		return
	}
	key := config.S3.Example.Key
	if key == "" {
		logger.Warn("key is empty")
		return
	}

	options := s3.Options{
		Region: "",
		Credentials: credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     id,
				SecretAccessKey: key,
				SessionToken:    "",
			},
		},
		UsePathStyle: true,
	}
	endpoint := config.S3.Example.Endpoint
	if endpoint != "" {
		options.EndpointResolver = s3.EndpointResolverFromURL(config.S3.Example.Endpoint)
	}
	s3Client = s3.New(options)

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			if s3Client == nil {
				return
			}
			return
		},
	})
	return
}
