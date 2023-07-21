package app

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/crawler"
	"github.com/lizongying/go-crawler/pkg/db"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/spider"
	"go.uber.org/fx"
)

func NewApp(newSpider pkg.NewSpider, crawlOptions ...pkg.CrawlOption) *fx.App {
	return fx.New(
		fx.Provide(
			cli.NewCli,
			db.NewMongoDb,
			db.NewMysql,
			db.NewKafka,
			db.NewKafkaReader,
			db.NewRedis,
			db.NewS3,
			fx.Annotate(
				logger.NewLogger,
				fx.As(new(pkg.Logger)),
			),
			config.NewConfig,
			//fx.Annotate(
			//	config.NewConfig,
			//	fx.As(new(pkg.Config)),
			//),
			fx.Annotate(
				spider.NewBaseSpider,
				fx.As(new(pkg.Spider)),
				fx.ResultTags(`name:"baseSpider"`),
			),
			devServer.NewHttpServer,
			fx.Annotate(
				newSpider,
				fx.ParamTags(`name:"baseSpider"`),
			),
			crawler.NewCrawler,
		),
		fx.Invoke(func(logger pkg.Logger, spider pkg.Spider, crawler pkg.Crawler, shutdowner fx.Shutdowner) {
			ctx := context.Background()

			for _, v := range crawlOptions {
				v(crawler)
			}

			crawler.SetSpider(spider)
			err := crawler.Start(ctx)
			if err != nil {
				logger.Error(err)
				_ = shutdowner.Shutdown()
				return
			}

			err = crawler.Stop(ctx)
			if errors.Is(err, pkg.DontStopErr) {
				select {}
			}
			if err != nil {
				logger.Error(err)
				return
			}

			err = shutdowner.Shutdown()
			if err != nil {
				logger.Error(err)
				return
			}
			logger.Info("Shutdown success")

			return
		}),
	)
}
