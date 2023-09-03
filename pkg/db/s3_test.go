package db

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"go.uber.org/fx"
	"os"
	"testing"
)

// go test -v ./pkg/db/*.go -run TestNewS3
func TestNewS3(t *testing.T) {
	_ = os.Setenv("CRAWLER_CONFIG_FILE", "/Users/lizongying/IdeaProjects/go-crawler/dev.yml")
	fx.New(
		fx.Provide(
			cli.NewCli,
			NewS3,
			fx.Annotate(
				loggers.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			config.NewConfig,
		),
		fx.Invoke(func(logger pkg.Logger, s3 *s3.Client, shutdowner fx.Shutdowner) {
			var err error
			ctx := context.Background()
			buckets, err := s3.ListBuckets(ctx, nil)
			for _, v := range buckets.Buckets {
				logger.Info(*v.Name, *v.CreationDate)
			}

			err = shutdowner.Shutdown()
			if err != nil {
				logger.Error(err)
				return
			}

			return
		}),
	).Run()
}
