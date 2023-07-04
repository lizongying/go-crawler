package crawler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/filter"
	"github.com/lizongying/go-crawler/pkg/scheduler/memory"
	redis2 "github.com/lizongying/go-crawler/pkg/scheduler/redis"
	"github.com/lizongying/go-crawler/pkg/signals"
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"time"
)

type Crawler struct {
	mode string

	startFunc string
	args      string

	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}

	retryMaxTimes uint8
	timeout       time.Duration
	username      string
	password      string

	filter    pkg.Filter
	spider    pkg.Spider
	devServer pkg.DevServer

	okHttpCodes []int
	platforms   map[pkg.Platform]struct{}
	browsers    map[pkg.Browser]struct{}

	config pkg.Config

	MongoDb *mongo.Database
	Mysql   *sql.DB
	Redis   *redis.Client
	Kafka   *kafka.Writer
	S3      *s3.Client
	logger  pkg.Logger

	pkg.Scheduler
	pkg.Stats
	pkg.Signal
}

func (c *Crawler) GetMode() string {
	return c.mode
}

func (c *Crawler) SetMode(mode string) {
	c.mode = mode
}

func (c *Crawler) GetPlatforms() (platforms []pkg.Platform) {
	for k := range c.platforms {
		platforms = append(platforms, k)
	}
	return
}

func (c *Crawler) SetPlatforms(platforms ...pkg.Platform) {
	for _, platform := range platforms {
		if platform == "" {
			err := errors.New("platform error")
			c.logger.Warn(err)
			continue
		}
		c.platforms[platform] = struct{}{}
	}
}

func (c *Crawler) GetBrowsers() (browsers []pkg.Browser) {
	for k := range c.browsers {
		browsers = append(browsers, k)
	}
	return
}

func (c *Crawler) SetBrowsers(browsers ...pkg.Browser) {
	for _, browser := range browsers {
		if browser == "" {
			err := errors.New("browser error")
			c.logger.Warn(err)
			continue
		}
		c.browsers[browser] = struct{}{}
	}
}

func (c *Crawler) GetLogger() pkg.Logger {
	return c.logger
}

func (c *Crawler) SetLogger(logger pkg.Logger) {
	c.logger = logger
}
func (c *Crawler) GetSpider() pkg.Spider {
	return c.spider
}
func (c *Crawler) SetSpider(spider pkg.Spider) {
	c.spider = spider
	c.Signal.SetSpider(spider)
}

func (c *Crawler) GetUsername() string {
	return c.username
}
func (c *Crawler) SetUsername(username string) {
	c.username = username
}
func (c *Crawler) GetPassword() string {
	return c.password
}
func (c *Crawler) SetPassword(password string) {
	c.password = password
}
func (c *Crawler) GetStats() pkg.Stats {
	return c.Stats
}
func (c *Crawler) RunDevServer() (err error) {
	err = c.devServer.Run()
	if err != nil {
		c.logger.Error(err)
		return
	}

	return
}

func (c *Crawler) AddDevServerRoutes(routes ...pkg.Route) {
	c.devServer.AddRoutes(routes...)
}

func (c *Crawler) GetDevServerHost() (host string) {
	host = c.devServer.GetHost()
	return
}

func (c *Crawler) GetOkHttpCodes() (httpCodes []int) {
	httpCodes = c.okHttpCodes
	return
}

func (c *Crawler) SetOkHttpCodes(httpCodes ...int) {
	for _, v := range httpCodes {
		if utils.InSlice(v, c.okHttpCodes) {
			continue
		}
		c.okHttpCodes = append(c.okHttpCodes, v)
	}
}

func (c *Crawler) GetCallbacks() map[string]pkg.Callback {
	return c.spider.GetCallbacks()
}
func (c *Crawler) GetErrbacks() map[string]pkg.Errback {
	return c.spider.GetErrbacks()
}

func (c *Crawler) GetConfig() pkg.Config {
	return c.config
}

func (c *Crawler) GetKafka() *kafka.Writer {
	return c.Kafka
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
func (c *Crawler) GetS3() *s3.Client {
	return c.S3
}
func (c *Crawler) GetFilter() pkg.Filter {
	return c.filter
}
func (c *Crawler) SetFilter(filter pkg.Filter) {
	c.filter = filter
}
func (c *Crawler) GetRetryMaxTimes() uint8 {
	return c.retryMaxTimes
}
func (c *Crawler) SetRetryMaxTimes(retryMaxTimes uint8) {
	c.retryMaxTimes = retryMaxTimes
}
func (c *Crawler) GetTimeout() time.Duration {
	return c.timeout
}
func (c *Crawler) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}
func (c *Crawler) GetScheduler() pkg.Scheduler {
	return c.Scheduler
}
func (c *Crawler) SetScheduler(scheduler pkg.Scheduler) {
	c.Scheduler = scheduler
}
func (c *Crawler) GetSignal() pkg.Signal {
	return c.Signal
}
func (c *Crawler) SetSignal(signal pkg.Signal) {
	c.Signal = signal
}
func (c *Crawler) registerParser() {
	callbacks := make(map[string]pkg.Callback)
	errbacks := make(map[string]pkg.Errback)
	rt := reflect.TypeOf(c.spider)
	rv := reflect.ValueOf(c.spider)
	l := rt.NumMethod()
	for i := 0; i < l; i++ {
		name := rt.Method(i).Name
		callback, ok := rv.Method(i).Interface().(func(context.Context, *pkg.Response) error)
		if ok {
			callbacks[name] = callback
		}
		errback, ok := rv.Method(i).Interface().(func(context.Context, *pkg.Response, error))
		if ok {
			errbacks[name] = errback
		}
	}
	c.spider.SetCallbacks(callbacks)
	c.spider.SetErrbacks(errbacks)
}
func (c *Crawler) Start(ctx context.Context) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if c.spider == nil {
		err = errors.New("nil spider")
		c.logger.Error(err)
		return
	}

	name := c.spider.GetName()
	if c.spider.GetName() == "" {
		err = errors.New("spider name is empty")
		return
	}

	c.logger.Info("name", name)
	c.logger.Info("start func", c.startFunc)
	c.logger.Info("args", c.args)
	c.logger.Info("mode", c.mode)
	c.logger.Info("allowedDomains", c.GetAllowedDomains())
	c.logger.Info("okHttpCodes", c.okHttpCodes)
	c.logger.Info("platforms", c.GetPlatforms())
	c.logger.Info("browsers", c.GetBrowsers())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.logger.Info("redirectMaxTimes", c.config.GetRedirectMaxTimes())
	c.logger.Info("retryMaxTimes", c.config.GetRetryMaxTimes())
	c.logger.Info("filter", c.config.GetFilter())

	// must start before scheduler
	err = c.spider.Start(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}
	c.registerParser()

	c.Signal.SpiderOpened()

	err = c.Scheduler.Start(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}

	params := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(c.args),
	}
	caller := reflect.ValueOf(c.spider).MethodByName(c.startFunc)
	if !caller.IsValid() {
		err = errors.New("start func is invalid")
		c.logger.Error(err)
		return
	}

	// TODO handle result and do something
	r := caller.Call(params)[0].Interface()
	if r != nil {
		err = r.(error)
		c.logger.Error(err)
		return
	}

	return
}

func (c *Crawler) Stop(ctx context.Context) (err error) {
	c.logger.Debug("Wait for stop")
	defer func() {
		c.logger.Info("Crawler Stopped")
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	err = c.Scheduler.Stop(ctx)
	if err != nil {
		c.logger.Error(err)
		return
	}

	return c.spider.Stop(ctx)
}

func NewCrawler(cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, redis *redis.Client, kafka *kafka.Writer, s3 *s3.Client, devServer pkg.DevServer) (crawler pkg.Crawler, err error) {
	defaultAllowedDomains := map[string]struct{}{"*": {}}

	crawler = &Crawler{
		retryMaxTimes: config.GetRetryMaxTimes(),
		timeout:       config.GetRequestTimeout(),
		okHttpCodes:   config.GetOkHttpCodes(),
		startFunc:     cli.StartFunc,
		args:          cli.Args,

		devServer: devServer,

		platforms: make(map[pkg.Platform]struct{}, 6),
		browsers:  make(map[pkg.Browser]struct{}, 4),
		config:    config,

		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
		MongoDb:               mongoDb,
		Mysql:                 mysql,
		Kafka:                 kafka,
		Redis:                 redis,
		S3:                    s3,
		Stats:                 &stats.Stats{},
	}

	crawler.SetMode(cli.Mode)
	crawler.SetLogger(logger)
	crawler.SetSignal(new(signals.Signal).FromCrawler(crawler))

	if cli.Mode == "dev" {
		err = crawler.RunDevServer()
		if err != nil {
			logger.Error(err)
			return
		}
	}

	switch config.GetFilter() {
	case pkg.FilterMemory:
		crawler.SetFilter(new(filter.MemoryFilter).FromCrawler(crawler))
	case pkg.FilterRedis:
		crawler.SetFilter(new(filter.RedisFilter).FromCrawler(crawler))
	default:
	}

	switch config.GetScheduler() {
	case pkg.SchedulerMemory:
		crawler.SetScheduler(new(memory.Scheduler).FromCrawler(crawler))
	case pkg.SchedulerRedis:
		crawler.SetScheduler(new(redis2.Scheduler).FromCrawler(crawler))
	default:
	}

	return crawler, nil
}
