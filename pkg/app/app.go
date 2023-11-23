package app

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/api"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/crawler"
	"github.com/lizongying/go-crawler/pkg/db"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
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
		loggers.NewStream,
		fx.Annotate(
			loggers.NewLogger,
			fx.As(new(pkg.Logger)),
		),
		fx.Annotate(
			db.NewSqlite,
			fx.As(new(pkg.Sqlite)),
		),
		fx.Annotate(
			db.NewStore,
			fx.As(new(pkg.Store)),
		),
		config.NewConfig,
		//fx.Annotate(
		//	config.NewConfig,
		//	fx.As(new(pkg.Config)),
		//),
		mock_servers.NewHttpServer,
		fx.Annotate(
			crawler.NewCrawler,
			fx.ParamTags(`group:"spiders"`),
		),
		api.NewApi,
	}

	for _, v := range a.newSpiders {
		v := v
		constructors = append(constructors, fx.Annotate(
			func(logger pkg.Logger) (pkg.Spider, error) {
				baseSpider, _ := spider.NewBaseSpider(logger)
				return v(baseSpider)
			},
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
			ctx, cancel := context.WithCancel(ctx)
			defer func() {
				cancel()
				if err = shutdowner.Shutdown(); err != nil {
					logger.Error(err)
					return
				}
			}()

			if err = crawler.Start(ctx); err != nil {
				logger.Error(err)
				err = shutdowner.Shutdown()
				if err != nil {
					logger.Error(err)
				}
				return
			}

			return
		}),
	).Run()
}
