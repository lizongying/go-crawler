package crawler

import (
	"context"
	"database/sql"
	"errors"
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
	mode        pkg.JobMode
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

	spider *pkg.State
	stop   chan struct{}

	// item limit
	itemDelay           time.Duration
	itemConcurrency     uint8
	itemConcurrencyChan chan struct{}
	itemTimer           *time.Timer
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
	if err = c.mockServer.Run(); err != nil {
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
func (c *Crawler) GetItemDelay() time.Duration {
	return c.itemDelay
}
func (c *Crawler) WithItemDelay(itemDelay time.Duration) pkg.Crawler {
	c.itemDelay = itemDelay
	return c
}
func (c *Crawler) GetItemConcurrency() uint8 {
	return c.itemConcurrency
}
func (c *Crawler) WithItemConcurrency(itemConcurrency uint8) pkg.Crawler {
	if itemConcurrency < 1 {
		itemConcurrency = 1
	}

	c.itemConcurrency = itemConcurrency
	return c
}
func (c *Crawler) ItemTimer() *time.Timer {
	return c.itemTimer
}
func (c *Crawler) ItemConcurrencyChan() chan struct{} {
	return c.itemConcurrencyChan
}
func (c *Crawler) Run(ctx context.Context, spiderName string, startFunc string, args string, mode pkg.JobMode, spec string) (id string, err error) {
	var spider pkg.Spider
	for _, v := range c.spiders {
		if v.Name() == spiderName {
			spider = v
			break
		}
	}

	if spider == nil {
		err = errors.New("nil spider")
		c.logger.Error(err)
		return
	}

	if spider.GetContext() == nil {
		if err = spider.Start(c.context); err != nil {
			c.logger.Error(err)
			return
		}
	}

	c.logger.Info("name", spiderName)
	c.logger.Info("func", startFunc)
	c.logger.Info("args", args)
	c.logger.Info("mode", mode)
	c.logger.Info("spec", spec)

	id, err = spider.Run(ctx, startFunc, args, mode, spec, true)
	if err != nil {
		c.logger.Error(err)
	}

	return
}

func (c *Crawler) SpiderStop(ctx pkg.Context) (err error) {
	taskId := ctx.GetTaskId()
	c.logger.Info(taskId)
	return
}
func (c *Crawler) Start(ctx context.Context) (err error) {
	crawler := new(crawlerContext.Crawler).
		WithContext(ctx).
		WithId(utils.UUIDV1WithoutHyphens()).
		WithStatus(pkg.CrawlerStatusOnline).
		WithStartTime(time.Now())
	c.context = new(crawlerContext.Context).WithCrawler(crawler)
	c.logger.Info("crawlerId", c.context.GetCrawlerId())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.Signal.CrawlerStarted(c.context)

	// init item limit
	c.itemTimer = time.NewTimer(c.itemDelay)
	if c.itemConcurrency < 1 {
		c.itemConcurrency = 1
	}
	c.itemConcurrencyChan = make(chan struct{}, c.itemConcurrency)
	for i := 0; i < int(c.itemConcurrency); i++ {
		c.itemConcurrencyChan <- struct{}{}
	}

	if err = c.api.Run(); err != nil {
		c.logger.Error(err)
		return
	}

	if c.spiderName != "" {
		var id string
		if id, err = c.Run(ctx, c.spiderName,
			c.startFunc,
			c.args,
			c.mode,
			c.spec,
		); err != nil {
			c.logger.Error(err)
			return
		}
		c.logger.Info("job id", id)
		c.spider.In()
	} else {
		<-c.stop
	}

	return
}

func (c *Crawler) Stop(ctx context.Context) (err error) {
	c.logger.Debug("Crawler wait for stop")
	defer func() {
		c.context.WithCrawlerContext(ctx)
		c.context.WithCrawlerStatus(pkg.CrawlerStatusOffline)
		c.context.WithCrawlerStopTime(time.Now())
		c.Signal.CrawlerStopped(c.context)
		c.logger.Info("Crawler Stopped")
	}()

	for _, v := range c.spiders {
		if err = v.Stop(c.context); err != nil {
			c.logger.Error(err)
		}
	}

	return
}
func (c *Crawler) StopSpider() {
	defer c.spider.Out()
}

func NewCrawler(spiders []pkg.Spider, cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, kafkaReader *kafka.Reader, sqlite pkg.Sqlite, store pkg.Store, mockServer pkg.MockServer, httpApi *api.Api) (crawler pkg.Crawler, err error) {
	spider := pkg.NewState()
	spider.RegisterIsReadyAndIsZero(func() {
		_ = crawler.Stop(crawler.GetContext().GetCrawlerContext())
	})

	crawler = &Crawler{
		spiderName:  cli.SpiderName,
		startFunc:   cli.StartFunc,
		args:        cli.Args,
		mode:        pkg.JobModeFromString(cli.Mode),
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
		spider:      spider,
		stop:        make(chan struct{}),
	}

	httpApi.AddRoutes(new(api.RouteHome).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteHello).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpider).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobRun).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteNodes).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpiders).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobs).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteTasks).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteRecords).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteUser).FromCrawler(crawler))

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
		if err = crawler.RunMockServer(); err != nil {
			logger.Error(err)
			return
		}
	}

	return crawler, nil
}
