package app

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/api"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/crawler"
	"github.com/lizongying/go-crawler/pkg/db"
	"github.com/lizongying/go-crawler/pkg/limiter"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"github.com/lizongying/go-crawler/pkg/mock_servers"
	"github.com/lizongying/go-crawler/pkg/spider"
	"go.uber.org/fx"
	"reflect"
)

type SimpleApp struct {
	newSpiders []pkg.NewSimpleSpider
}

func NewSimpleApp(newSpiders ...pkg.NewSimpleSpider) *SimpleApp {
	return &SimpleApp{
		newSpiders: newSpiders,
	}
}

func (a *SimpleApp) Run(crawlOptions ...pkg.CrawlOption) {
	constructors := []any{cli.NewCli,
		db.NewMongoFactory,
		db.NewKafkaFactory,
		db.NewRedisFactory,
		db.NewSqliteFactory,
		db.NewMysqlFactory,
		db.NewStorageFactory,
		loggers.NewStream,
		fx.Annotate(
			loggers.NewLogger,
			fx.As(new(pkg.Logger)),
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
		limiter.NewManager,
	}

	for _, v := range a.newSpiders {
		constructors = append(constructors, fx.Annotate(
			func(logger pkg.Logger, lm *limiter.Manager) (baseSpider pkg.Spider, err error) {
				baseSpider, _ = spider.NewBaseSpider(logger, lm)
				simpleSpider := &spider.SimpleSpider{
					BaseSpider: baseSpider.(*spider.BaseSpider),
				}

				ss, err := v(simpleSpider)

				rv := reflect.ValueOf(ss)
				rt := rv.Type()
				l := rt.NumMethod()
				for i := 0; i < l; i++ {
					method := rt.Method(i)
					name := method.Name
					fn := rv.Method(i).Interface()
					callBack, ok := fn.(func(pkg.Context, pkg.Response) error)
					if ok {
						baseSpider.SetCallBack(name, callBack)
					}
					errBack, ok := fn.(func(pkg.Context, pkg.Response, error))
					if ok {
						baseSpider.SetErrBack(name, errBack)
					}
					startFunc, ok := fn.(func(pkg.Context, string) error)
					if ok {
						baseSpider.SetStartFunc(name, startFunc)
					}
				}
				return
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
