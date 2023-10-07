package db

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"go.uber.org/fx"
	"os"
	"testing"
)

// go test -v ./pkg/db/*.go -run TestNewStore
func TestNewStore(t *testing.T) {
	_ = os.Setenv("CRAWLER_CONFIG_FILE", "/Users/lizongying/IdeaProjects/go-crawler/dev.yml")
	fx.New(
		fx.Provide(
			cli.NewCli,
			fx.Annotate(
				loggers.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			fx.Annotate(
				NewStore,
				fx.As(new(pkg.Store)),
			),
			config.NewConfig,
		),
		fx.Invoke(func(logger pkg.Logger, store pkg.Store, shutdowner fx.Shutdowner) {
			var err error
			ctx := context.Background()
			buckets, err := store.S3Client().ListBuckets(ctx, nil)
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
