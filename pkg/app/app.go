package app

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/db"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/spider"
	"go.uber.org/fx"
)

type App struct {
	*fx.App
}

func NewApp(f func(*spider.BaseSpider, *logger.Logger) (pkg.Spider, error)) (app *App) {
	app = &App{
		App: fx.New(
			fx.Provide(
				cli.NewCli,
				config.NewConfig,
				db.NewMongoDb,
				db.NewMysql,
				logger.NewLogger,
				httpClient.NewHttpClient,
				spider.NewBaseSpider,
				devServer.NewHttpServer,
				f,
			),
			fx.Invoke(func(logger *logger.Logger, cli *cli.Cli, spider pkg.Spider, shutdowner fx.Shutdowner) {
				ctx := context.Background()

				if cli.Mode == "dev" {
					err := spider.RunDevServer()
					if err != nil {
						logger.Error(err)
						_ = shutdowner.Shutdown()
						return
					}
				}

				spider.SetSpider(spider)
				err := spider.Start(ctx)
				if err != nil {
					logger.Error(err)
					_ = shutdowner.Shutdown()
					return
				}

				err = spider.Stop(ctx)
				if errors.Is(err, pkg.DontStopErr) {
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
