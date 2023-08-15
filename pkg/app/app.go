package app

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/crawler"
	"github.com/lizongying/go-crawler/pkg/db"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/mockServer"
	"github.com/lizongying/go-crawler/pkg/spider"
	"go.uber.org/fx"
)

type App struct {
	newSpiders []pkg.NewSpider
}

func NewApp(newSpiders ...pkg.NewSpider) *App {
	return &App{
		newSpiders: newSpiders,
	}
}

func (a *App) Run(crawlOptions ...pkg.CrawlOption) {
	constructors := []any{cli.NewCli,
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
			fx.ResultTags(`name:"baseSpider"`),
		),
		mockServer.NewHttpServer,
		fx.Annotate(
			crawler.NewCrawler,
			fx.ParamTags(`group:"spiders"`),
		),
	}

	for _, v := range a.newSpiders {
		constructors = append(constructors, fx.Annotate(
			v,
			fx.ParamTags(`name:"baseSpider"`),
			fx.ResultTags(`group:"spiders"`),
		))
	}

	fx.New(
		fx.Provide(constructors...),
		fx.Invoke(func(logger pkg.Logger, crawler pkg.Crawler, shutdowner fx.Shutdowner) {
			var err error
			for _, v := range crawlOptions {
				v(crawler)
			}

			ctx := context.Background()

			err = crawler.Start(ctx)
			if err != nil {
				logger.Error(err)
				err = shutdowner.Shutdown()
				if err != nil {
					logger.Error(err)
				}
				return
			}

			err = crawler.Stop(ctx)
			if errors.Is(err, pkg.DontStopErr) {
				select {}
			}
			if err != nil {
				logger.Error(err)
				err = shutdowner.Shutdown()
				if err != nil {
					logger.Error(err)
				}
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
	).Run()
}
