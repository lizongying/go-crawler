package db

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/utils"
	"go.uber.org/fx"
	"os"
	"path/filepath"
	"strings"
)

type Store struct {
	Config *config.Store
	*s3.Client
	logger pkg.Logger
}

func (s *Store) S3Client() *s3.Client {
	return s.Client
}

// Save
// is s3:
//
//	prefix=bucket
//
// is file:
//
//	prefix=dir
func (s *Store) Save(prefix string, key string, body []byte) (storePath string, err error) {
	if s.Client != nil {
		uploadParams := &s3.PutObjectInput{
			Key:  &key,
			Body: bytes.NewReader(body),
		}
		if prefix == "" {
			prefix = s.Config.Bucket
		}
		prefix = strings.Trim(prefix, "/")
		uploadParams.Bucket = &prefix

		// Upload the file
		_, err = s.PutObject(context.Background(), uploadParams)
		if err != nil {
			s.logger.Error(err)
			return
		}
		storePath = fmt.Sprintf("%s://%s/%s", s.Config.Type, prefix, key)
	} else {
		if s.Config.Endpoint != "" {
			prefix = strings.TrimPrefix(s.Config.Endpoint, "file:/")
		}
		if prefix == "" {
			name, _ := os.MkdirTemp("", s.Config.Bucket)
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
	}
	return
}

func NewStore(config *config.Config, logger pkg.Logger, lc fx.Lifecycle) (store *Store, err error) {
	store = new(Store)
	store.logger = logger

	for _, v := range config.Store {
		//if v.Endpoint == "" {
		//	continue
		//}

		if v.Bucket == "" {
			v.Bucket = config.GetBotName()
		}
		store.Config = v

		if utils.InSlice(v.Type, []string{"s3", "oss", "cos", "minio"}) {
			options := s3.Options{
				UsePathStyle: true,
			}
			if v.Region != "" {
				options.Region = v.Region
			}
			if v.Id != "" && v.Key != "" {
				options.Credentials = credentials.StaticCredentialsProvider{
					Value: aws.Credentials{
						AccessKeyID:     v.Id,
						SecretAccessKey: v.Key,
						SessionToken:    "",
					},
				}
			}
			if v.Endpoint != "" {
				options.EndpointResolver = s3.EndpointResolverFromURL(v.Endpoint)
			}
			store.Client = s3.New(options)
		}

		break
	}

	lc.Append(fx.Hook{
		OnStop: func(_ context.Context) (err error) {
			return
		},
	})
	return
}
