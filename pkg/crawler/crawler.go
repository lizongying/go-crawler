package crawler

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/api"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/loggers"
	"github.com/lizongying/go-crawler/pkg/signals"
	"github.com/lizongying/go-crawler/pkg/statistics"
	"github.com/lizongying/go-crawler/pkg/utils"
	uidv1 "github.com/lizongying/go-uid/v1"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"runtime"
	"time"
)

var (
	buildBranch string
	buildCommit string
	buildTime   string
)

func init() {
	if buildTime != "" {
		str2Int64, err := utils.Str2Int64(buildTime)
		if err == nil {
			buildTime = time.Unix(str2Int64, 0).Format(time.DateTime)
		}
	}
	info := fmt.Sprintf("Branch: %s, Commit: %s, Time: %s, GOVersion: %s, OS: %s, ARCH: %s", buildBranch, buildCommit, buildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	log.Println(info)
}

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
	ug                  *uidv1.Uid

	stream pkg.Stream

	cdp bool
}

func (c *Crawler) GetCDP() bool {
	return c.cdp
}
func (c *Crawler) WithCDP(cdp bool) pkg.Crawler {
	c.cdp = cdp
	return c
}

func (c *Crawler) GetStream() pkg.Stream {
	return c.stream
}
func (c *Crawler) GenUid() uint64 {
	return c.ug.Gen()
}
func (c *Crawler) NextId() string {
	return fmt.Sprintf("%d", c.ug.Gen())
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
func (c *Crawler) AddMockServerRoutes(routes ...pkg.Route) pkg.Crawler {
	if !c.GetConfig().MockServerEnable() {
		c.GetConfig().SetMockServerEnable(true)
	}

	c.mockServer.AddRoutes(routes...)
	return c
}

func (c *Crawler) AddDefaultMocks() pkg.Crawler {
	if !c.GetConfig().MockServerEnable() {
		c.GetConfig().SetMockServerEnable(true)
	}

	c.mockServer.AddDefaultRoutes()
	return c
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
func (c *Crawler) StartFromCLI() bool {
	return c.spiderName != ""
}

func (c *Crawler) RunJob(ctx context.Context, spiderName string, startFunc string, args string, mode pkg.JobMode, spec string) (id string, err error) {
	var spider pkg.Spider
	for _, v := range c.spiders {
		if v.Name() == spiderName {
			spider = v
			break
		}
	}

	if spider == nil {
		err = fmt.Errorf("%w: %s", pkg.ErrSpiderNotFound, spiderName)
		return
	}

	if spider.GetContext().GetSpider().GetStatus() == pkg.SpiderStatusReady {
		if err = spider.Start(c.context); err != nil {
			c.logger.Error(err)
			return
		}
		c.spider.In()
	}

	c.logger.Info("name", spiderName)
	c.logger.Info("func", startFunc)
	c.logger.Info("args", args)
	c.logger.Info("mode", mode.String())
	if mode == pkg.JobModeCron {
		c.logger.Info("spec", spec)
	}

	id, err = spider.Run(ctx, startFunc, args, mode, spec, true)
	if err != nil {
		c.logger.Error(err)
	}

	return
}
func (c *Crawler) RerunJob(ctx context.Context, spiderName string, jobId string) (err error) {
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

	if err = spider.RerunJob(ctx, jobId); err != nil {
		c.logger.Error(err)
	}
	return
}
func (c *Crawler) KillJob(ctx context.Context, spiderName string, jobId string) (err error) {
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

	if err = spider.KillJob(ctx, jobId); err != nil {
		c.logger.Error(err)
	}
	return
}
func (c *Crawler) Start(ctx context.Context) (err error) {
	// load id
	c.WithContext(new(crawlerContext.Context).
		WithCrawler(new(crawlerContext.Crawler).
			WithId(c.NextId())))
	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusReady)
	c.Signal.CrawlerChanged(c.context)

	c.context.GetCrawler().WithContext(ctx)

	if c.GetConfig().MockServerEnable() {
		if err = c.RunMockServer(); err != nil {
			c.logger.Error(err)
			return
		}
	}

	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusStarting)
	c.Signal.CrawlerChanged(c.context)

	for _, v := range c.spiders {
		v.SetSpider(v)
		v.FromCrawler(c)
		c.logger.Info("spider", v.Name(), "loaded")
	}

	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusRunning)
	c.Signal.CrawlerChanged(c.context)

	c.logger.Info("crawler is running. id:", c.context.GetCrawler().GetId())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())

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
		if id, err = c.RunJob(ctx, c.spiderName,
			c.startFunc,
			c.args,
			c.mode,
			c.spec,
		); err != nil {
			c.logger.Error(err)
			return
		}
		c.logger.Info("job id", id)
	}

	<-c.stop
	return
}

func (c *Crawler) Stop(ctx pkg.Context) (err error) {
	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusIdle)
	c.Signal.CrawlerChanged(ctx)
	c.logger.Debug("crawler has idle")

	if !c.StartFromCLI() {
		c.logger.Info("crawler don't need to stop")
		return
	}

	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusStopping)
	c.Signal.CrawlerChanged(ctx)
	c.logger.Debug("crawler wait for stop")

	for _, v := range c.spiders {
		if !utils.InSlice(v.GetContext().GetSpider().GetStatus(), []pkg.SpiderStatus{
			pkg.SpiderStatusReady,
			pkg.SpiderStatusStopped,
		}) {
			c.logger.Warn("crawler has a spider that need to stop", v.GetContext().GetSpider().GetStatus().String())
			return
		}
	}

	c.context.GetCrawler().WithStatus(pkg.CrawlerStatusStopped)
	c.Signal.CrawlerChanged(ctx)
	c.logger.Info("crawler finished. spend time:", ctx.GetCrawler().GetStopTime().Sub(ctx.GetCrawler().GetStartTime()), ctx.GetCrawler().GetId())
	c.stop <- struct{}{}
	return
}
func (c *Crawler) SpiderStopped(_ pkg.Context, _ error) {
	c.spider.Out()
}

func NewCrawler(spiders []pkg.Spider, cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, kafkaReader *kafka.Reader, sqlite pkg.Sqlite, store pkg.Store, mockServer pkg.MockServer, httpApi *api.Api, stream *loggers.Stream) (crawler pkg.Crawler, err error) {
	spider := pkg.NewState("spider")
	spider.RegisterIsReadyAndIsZero(func() {
		_ = crawler.Stop(crawler.GetContext())
	})

	ug := uidv1.NewUid(1, nil)

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
		ug:          ug,
		spiders:     spiders,
		stream:      stream,
	}

	httpApi.AddRoutes(new(api.RouteHome).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteHello).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpider).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobRun).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobRerun).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobStop).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteCrawlers).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteSpiders).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteJobs).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteTasks).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteRequests).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteItems).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteUser).FromCrawler(crawler))
	httpApi.AddRoutes(new(api.RouteLog).FromCrawler(crawler))

	crawler.SetSignal(new(signals.Signal).FromCrawler(crawler))
	crawler.SetStatistics(new(statistics.Statistics).FromCrawler(crawler))

	return crawler, nil
}
