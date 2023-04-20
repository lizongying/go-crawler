package app

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/cli"
	"github.com/lizongying/go-crawler/internal/config"
	"github.com/lizongying/go-crawler/internal/httpClient"
	"github.com/lizongying/go-crawler/internal/logger"
	"github.com/lizongying/go-crawler/internal/mongodb"
	"github.com/lizongying/go-crawler/internal/spider"
	"go.uber.org/fx"
)

type App struct {
	*fx.App
}

func NewApp(f func(*spider.BaseSpider, *logger.Logger) (internal.Spider, error)) (app *App) {
	app = &App{
		App: fx.New(
			fx.Provide(
				cli.NewCli,
				config.NewConfig,
				mongodb.NewMongoDb,
				logger.NewLogger,
				httpClient.NewHttpClient,
				spider.NewBaseSpider,
				f,
			),
			fx.Invoke(func(logger *logger.Logger, spider internal.Spider, shutdowner fx.Shutdowner) {
				ctx := context.Background()
				spider.SetSpider(spider)
				err := spider.Start(ctx)
				if err != nil {
					logger.Error(err)
					_ = shutdowner.Shutdown()
					return
				}

				err = spider.Stop(ctx)
				if errors.Is(err, internal.DontStopErr) {
					select {}
				}
				if err != nil {
					return
				}

				err = shutdowner.Shutdown()
				if err != nil {
					return
				}

				return
			}),
		),
	}

	return
}
