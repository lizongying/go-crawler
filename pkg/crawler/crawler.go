package crawler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lizongying/cron"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/api"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/signals"
	"github.com/lizongying/go-crawler/pkg/statistics"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Crawler struct {
	context     pkg.Context
	spiders     []pkg.Spider
	spiderName  string
	startFunc   string
	args        string
	mode        string
	spec        string
	config      pkg.Config
	logger      pkg.Logger
	MongoDb     *mongo.Database
	Mysql       *sql.DB
	Redis       *redis.Client
	Kafka       *kafka.Writer
	KafkaReader *kafka.Reader
	Sqlite      pkg.Sqlite
	Store       pkg.Store
	mockServer  pkg.MockServer
	api         *api.Api
	statistics  pkg.Statistics
	pkg.Signal

	Status    pkg.CrawlerStatus `json:"status,omitempty"`
	StartTime utils.Timestamp   `json:"start_time,omitempty"`
	StopTime  utils.Timestamp   `json:"stop_time,omitempty"`
}

func (c *Crawler) GetContext() pkg.Context {
	return c.context
}
func (c *Crawler) WithContext(ctx pkg.Context) pkg.Crawler {
	c.context = ctx
	return c
}
func (c *Crawler) GetStatistics() pkg.Statistics {
	return c.statistics
}
func (c *Crawler) SetStatistics(statistics pkg.Statistics) {
	c.statistics = statistics
}
func (c *Crawler) GetMode() string {
	return c.mode
}
func (c *Crawler) SetMode(mode string) {
	c.mode = mode
}
func (c *Crawler) GetLogger() pkg.Logger {
	return c.logger
}
func (c *Crawler) SetLogger(logger pkg.Logger) {
	c.logger = logger
}
func (c *Crawler) GetSpiders() []pkg.Spider {
	return c.spiders
}
func (c *Crawler) SetSpiders(spiders []pkg.Spider) {
	c.spiders = spiders
}
func (c *Crawler) AddSpider(spider pkg.Spider) {
	c.spiders = append(c.spiders, spider)
}
func (c *Crawler) RunMockServer() (err error) {
	err = c.mockServer.Run()
	if err != nil {
		c.logger.Error(err)
		return
	}

	return
}
func (c *Crawler) AddMockServerRoutes(routes ...pkg.Route) {
	c.mockServer.AddRoutes(routes...)
}
func (c *Crawler) GetConfig() pkg.Config {
	return c.config
}
func (c *Crawler) GetKafka() *kafka.Writer {
	return c.Kafka
}
func (c *Crawler) GetKafkaReader() *kafka.Reader {
	return c.KafkaReader
}
func (c *Crawler) GetMongoDb() *mongo.Database {
	return c.MongoDb
}
func (c *Crawler) GetMysql() *sql.DB {
	return c.Mysql
}
func (c *Crawler) GetRedis() *redis.Client {
	return c.Redis
}
func (c *Crawler) GetSqlite() pkg.Sqlite {
	return c.Sqlite
}
func (c *Crawler) GetStore() pkg.Store {
	return c.Store
}
func (c *Crawler) GetSignal() pkg.Signal {
	return c.Signal
}
func (c *Crawler) SetSignal(signal pkg.Signal) {
	c.Signal = signal
}
func (c *Crawler) SpiderStart(ctx pkg.Context) (err error) {
	var spider pkg.Spider
	for _, v := range c.spiders {
		if v.Name() == ctx.GetSpiderName() {
			spider = v
			break
		}
	}

	if spider == nil {
		err = errors.New("nil spider")
		c.logger.Error(err)
		return
	}
	c.logger.Info("crawlerId", ctx.GetCrawlerId())
	c.logger.Info("name", ctx.GetSpiderName())
	c.logger.Info("func", ctx.GetStartFunc())
	c.logger.Info("args", ctx.GetArgs())
	c.logger.Info("mode", ctx.GetMode())
	c.logger.Info("allowedDomains", spider.GetAllowedDomains())
	c.logger.Info("okHttpCodes", spider.OkHttpCodes())
	c.logger.Info("platforms", spider.GetPlatforms())
	c.logger.Info("browsers", spider.GetBrowsers())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.logger.Info("redirectMaxTimes", c.config.GetRedirectMaxTimes())
	c.logger.Info("retryMaxTimes", c.config.GetRetryMaxTimes())
	c.logger.Info("filter", c.config.GetFilter())

	switch ctx.GetMode() {
	case "once":
		ctx.WithScheduleId(utils.UUIDV1WithoutHyphens())
		ctx.WithScheduleStatus(pkg.ScheduleStatusStarted)
		ctx.WithScheduleEnable(true)
		ctx.WithScheduleStartTime(time.Now())
		c.Signal.ScheduleStarted(ctx)

		if ctx.GetTaskId() == "" {
			ctx.WithTaskId(utils.UUIDV1WithoutHyphens())
		}
		if err = spider.Start(ctx); err != nil {
			c.logger.Error(err)
			return
		}
	case "loop":
		ctx.WithScheduleId(utils.UUIDV1WithoutHyphens())
		ctx.WithScheduleStatus(pkg.ScheduleStatusStarted)
		ctx.WithScheduleEnable(true)
		ctx.WithScheduleStartTime(time.Now())
		c.Signal.ScheduleStarted(ctx)

		for {
			ctx.WithTaskId(utils.UUIDV1WithoutHyphens())
			if err = spider.Start(ctx); err != nil {
				c.logger.Error(err)
				return
			}
		}
	case "cron":
		ctx.WithScheduleId(utils.UUIDV1WithoutHyphens())
		ctx.WithScheduleStatus(pkg.ScheduleStatusStarted)
		ctx.WithScheduleEnable(true)
		ctx.WithScheduleStartTime(time.Now())
		c.Signal.ScheduleStarted(ctx)

		cr := cron.New(cron.WithLogger(c.logger))
		cr.MustStart()
		job := new(cron.Job).
			MustEverySpec(c.spec).
			Callback(func() {
				ctx.WithTaskId(utils.UUIDV1WithoutHyphens())
				if err = spider.Start(ctx); err != nil {
					c.logger.Error(err)
					return
				}
			})
		cr.MustAddJob(job)
		select {}
	default:
		c.logger.Warn("mode", ctx.GetMode())
	}
	return
}
func (c *Crawler) SpiderStop(ctx pkg.Context) (err error) {
	taskId := ctx.GetTaskId()
	c.logger.Info(taskId)
	return
}
func (c *Crawler) Start(ctx context.Context) (err error) {
	if err = c.api.Run(); err != nil {
		c.logger.Error(err)
		return
	}

	c.context = crawlerContext.NewContext().
		WithGlobalContext(ctx).
		WithCrawlerId(utils.UUIDV1WithoutHyphens()).
		WithCrawlerStatus(pkg.CrawlerStatusOnline).
		WithCrawlerStartTime(time.Now())
	c.Signal.CrawlerStarted(c.context)

	if c.spiderName != "" {
		if err = c.SpiderStart(c.context.
			WithSpiderName(c.spiderName).
			WithStartFunc(c.startFunc).
			WithArgs(c.args).
			WithMode(c.mode)); err != nil {
			c.logger.Error(err)
			return
		}
	} else {
		select {}
	}
	return
}

func (c *Crawler) Stop(ctx context.Context) (err error) {
	c.logger.Debug("Crawler wait for stop")
	defer func() {
		c.context.WithGlobalContext(ctx)
		c.context.WithCrawlerStatus(pkg.CrawlerStatusOffline)
		c.context.WithCrawlerStopTime(time.Now())
		c.Signal.CrawlerStopped(c.context)
		c.logger.Info("Crawler Stopped")
	}()

	for _, v := range c.spiders {
		if err = v.Stop(v.GetContext()); err != nil {
			c.logger.Error(err)
		}
	}

	return
}

func NewCrawler(spiders []pkg.Spider, cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, kafkaReader *kafka.Reader, sqlite pkg.Sqlite, store pkg.Store, mockServer pkg.MockServer, httpApi *api.Api) (crawler pkg.Crawler, err error) {
	crawler = &Crawler{
		spiderName:  cli.SpiderName,
		startFunc:   cli.StartFunc,
		args:        cli.Args,
		mode:        cli.Mode,
		spec:        cli.Spec,
		config:      config,
		logger:      logger,
		MongoDb:     mongoDb,
		Mysql:       mysql,
		Redis:       redis,
		Kafka:       kafka,
		KafkaReader: kafkaReader,
		Sqlite:      sqlite,
		Store:       store,
		mockServer:  mockServer,
		api:         httpApi,
	}

	httpApi.AddRoutes(new(api.RouteHome).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteHello).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpider).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpiderRun).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteNodes).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpiders).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSchedules).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteTasks).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteRecords).FromCrawler(crawler))

	crawler.SetSignal(new(signals.Signal).FromCrawler(crawler))
	crawler.SetStatistics(new(statistics.Statistics).FromCrawler(crawler))

	for _, v := range spiders {
		v.SetSpider(v)
		v.FromCrawler(crawler)
		for _, option := range v.Options() {
			option(v)
		}
		logger.Info("spider", v.Name(), "loaded")
		crawler.AddSpider(v)
	}

	if config.MockServerEnable() {
		err = crawler.RunMockServer()
		if err != nil {
			logger.Error(err)
			return
		}
	}

	return crawler, nil
}
