package crawler

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/downloader"
	"github.com/lizongying/go-crawler/pkg/exporter"
	"github.com/lizongying/go-crawler/pkg/filter"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/time/rate"
	"reflect"
	"sync"
	"time"
)

const defaultChanRequestMax = 1000 * 1000
const defaultMaxRequestActive = 1000
const defaultChanItemMax = 1000 * 1000
const defaultRequestConcurrency = 1
const defaultRequestInterval = 1

type Crawler struct {
	Spider pkg.Spider
	Mode   string
	Name   string

	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}

	pkg.Downloader
	pkg.Exporter

	*pkg.SpiderInfo
	spider pkg.Spider
	Stats  pkg.Stats

	startFunc           string
	args                string
	itemConcurrency     int
	itemConcurrencyNew  int
	itemConcurrencyChan chan struct{}
	itemDelay           time.Duration
	itemTimer           *time.Timer
	itemChan            chan pkg.Item
	itemActiveChan      chan struct{}
	requestChan         chan *pkg.Request
	requestActiveChan   chan struct{}
	requestSlots        sync.Map

	devServer *devServer.HttpServer

	okHttpCodes []int
	platforms   map[pkg.Platform]struct{}
	browsers    map[pkg.Browser]struct{}

	config *config.Config

	filter pkg.Filter

	MongoDb *mongo.Database
	Mysql   *sql.DB
	Kafka   *kafka.Writer
	logger  pkg.Logger
}

func (c *Crawler) GetMode() string {
	return c.Mode
}

func (c *Crawler) SetMode(mode string) {
	c.Mode = mode
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

func (c *Crawler) GetPlatforms() (platforms []pkg.Platform) {
	for k := range c.platforms {
		platforms = append(platforms, k)
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

func (c *Crawler) GetBrowsers() (browsers []pkg.Browser) {
	for k := range c.browsers {
		browsers = append(browsers, k)
	}
	return
}

func (c *Crawler) SetLogger(logger pkg.Logger) {
	c.logger = logger
}

func (c *Crawler) SetSpider(spider pkg.Spider) {
	c.spider = spider
}

func (c *Crawler) GetInfo() *pkg.SpiderInfo {
	return c.SpiderInfo
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

func (c *Crawler) AddOkHttpCodes(httpCodes ...int) {
	for _, v := range httpCodes {
		if utils.InSlice(v, c.okHttpCodes) {
			continue
		}
		c.okHttpCodes = append(c.okHttpCodes, v)
	}
}

func (c *Crawler) GetOkHttpCodes() (httpCodes []int) {
	httpCodes = c.okHttpCodes
	return
}

func (c *Crawler) GetConfig() pkg.Config {
	return c.config
}

func (c *Crawler) GetLogger() pkg.Logger {
	return c.logger
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

func (c *Crawler) GetFilter() pkg.Filter {
	return c.filter
}

func (c *Crawler) SetDownloader(downloader pkg.Downloader) {
	c.Downloader = downloader
}
func (c *Crawler) SetExporter(exporter pkg.Exporter) {
	c.Exporter = exporter
}
func (c *Crawler) GetRetryMaxTimes() uint8 {
	return c.RetryMaxTimes
}
func (c *Crawler) SetRetryMaxTimes(RetryMaxTimes uint8) {
	c.RetryMaxTimes = RetryMaxTimes
}
func (c *Crawler) GetTimeout() time.Duration {
	return c.Timeout
}
func (c *Crawler) SetTimeout(timeout time.Duration) {
	c.Timeout = timeout
}
func (c *Crawler) GetInterval() time.Duration {
	return c.Interval
}
func (c *Crawler) SetInterval(interval time.Duration) {
	c.Interval = interval
}

func (c *Crawler) Start(ctx context.Context) (err error) {
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
	c.logger.Info("mode", c.Mode)
	c.logger.Info("allowedDomains", c.GetAllowedDomains())
	c.logger.Info("middlewares", c.GetMiddlewareNames())
	c.logger.Info("pipelines", c.GetPipelineNames())
	c.logger.Info("okHttpCodes", c.okHttpCodes)
	c.logger.Info("platforms", c.GetPlatforms())
	c.logger.Info("browsers", c.GetBrowsers())
	c.logger.Info("referrerPolicy", c.config.GetReferrerPolicy())
	c.logger.Info("urlLengthLimit", c.config.GetUrlLengthLimit())
	c.logger.Info("redirectMaxTimes", c.config.GetRedirectMaxTimes())
	c.logger.Info("retryMaxTimes", c.config.GetRetryMaxTimes())
	if c.spider == nil {
		err = errors.New("spider is empty")
		c.logger.Error(err)
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	for _, v := range c.GetPipelines() {
		e := v.Start(ctx, c)
		if errors.Is(e, pkg.BreakErr) {
			break
		}
	}
	for _, v := range c.GetMiddlewares() {
		e := v.Start(ctx, c)
		if errors.Is(e, pkg.BreakErr) {
			break
		}
	}

	defer func() {
		for _, v := range c.GetMiddlewares() {
			e := v.Stop(ctx)
			if errors.Is(e, pkg.BreakErr) {
				break
			}
		}
		for _, v := range c.GetPipelines() {
			e := v.Stop(ctx)
			if errors.Is(e, pkg.BreakErr) {
				break
			}
		}
	}()

	c.itemTimer = time.NewTimer(c.itemDelay)
	if c.itemConcurrency < 1 {
		c.itemConcurrency = 1
	}
	c.itemConcurrencyNew = c.itemConcurrency
	c.itemConcurrencyChan = make(chan struct{}, c.itemConcurrency)
	for i := 0; i < c.itemConcurrency; i++ {
		c.itemConcurrencyChan <- struct{}{}
	}

	slot := "*"
	if _, ok := c.requestSlots.Load(slot); !ok {
		requestSlot := rate.NewLimiter(rate.Every(c.Interval/time.Duration(c.Concurrency)), c.Concurrency)
		c.requestSlots.Store(slot, requestSlot)
	}

	go c.handleItem(ctx)

	go c.handleRequest(ctx)

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
		c.logger.Info("Stopped")
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if len(c.requestActiveChan) > 0 {
			c.logger.Debug("request is active")
			continue
		}
		if len(c.itemActiveChan) > 0 {
			c.logger.Debug("item is active")
			continue
		}
		break
	}

	return c.spider.Stop(ctx)
}

func NewCrawler(cli *cli.Cli, config *config.Config, logger pkg.Logger, mongoDb *mongo.Database, mysql *sql.DB, kafka *kafka.Writer, server *devServer.HttpServer) (crawler pkg.Crawler, err error) {
	defaultAllowedDomains := map[string]struct{}{"*": {}}

	concurrency := defaultRequestConcurrency
	if config.Request.Concurrency > 1 {
		concurrency = config.Request.Concurrency
	}
	interval := defaultRequestInterval
	if config.Request.Interval > 0 {
		interval = config.Request.Interval
	}
	if config.Request.Interval < 0 {
		interval = 0
	}
	okHttpCodes := []int{200}
	if len(config.Request.OkHttpCodes) > 0 {
		okHttpCodes = config.Request.OkHttpCodes
	}
	timeout := time.Minute
	if config.Request.Timeout > 0 {
		timeout = time.Second * time.Duration(config.Request.Timeout)
	}

	crawler = &Crawler{
		SpiderInfo: &pkg.SpiderInfo{
			Concurrency:   concurrency,
			Interval:      time.Millisecond * time.Duration(interval),
			RetryMaxTimes: config.GetRetryMaxTimes(),
			Timeout:       timeout,
		},
		Stats:       &stats.Stats{},
		okHttpCodes: okHttpCodes,
		startFunc:   cli.StartFunc,
		args:        cli.Args,

		requestChan:       make(chan *pkg.Request, defaultChanRequestMax),
		requestActiveChan: make(chan struct{}, defaultChanRequestMax),
		itemChan:          make(chan pkg.Item, defaultChanItemMax),
		itemActiveChan:    make(chan struct{}, defaultChanItemMax),

		devServer: server,

		platforms: make(map[pkg.Platform]struct{}, 6),
		browsers:  make(map[pkg.Browser]struct{}, 4),
		config:    config,
		filter:    new(filter.Filter),

		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
		MongoDb:               mongoDb,
		Mysql:                 mysql,
		Kafka:                 kafka,
	}

	crawler.SetLogger(logger)
	crawler.SetMode(cli.Mode)

	if cli.Mode == "dev" {
		err = crawler.RunDevServer()
		if err != nil {
			logger.Error(err)
			return
		}
	}

	crawler.SetDownloader(new(downloader.Downloader).FromCrawler(crawler))
	crawler.SetExporter(new(exporter.Exporter).FromCrawler(crawler))

	if config.GetEnableStats() {
		crawler.SetMiddleware(new(middlewares.StatsMiddleware), 10)
	}
	if config.GetEnableDumpMiddleware() {
		crawler.SetMiddleware(new(middlewares.DumpMiddleware), 20)
	}
	if config.GetEnableFilterMiddleware() {
		crawler.SetMiddleware(new(middlewares.FilterMiddleware), 30)
	}
	crawler.SetMiddleware(new(middlewares.HttpMiddleware), 40)
	if config.GetEnableRetry() {
		crawler.SetMiddleware(new(middlewares.RetryMiddleware), 50)
	}
	if config.GetEnableUrl() {
		crawler.SetMiddleware(new(middlewares.UrlMiddleware), 60)
	}
	if config.GetEnableReferer() {
		crawler.SetMiddleware(new(middlewares.RefererMiddleware), 70)
	}
	if config.GetEnableCookie() {
		crawler.SetMiddleware(new(middlewares.CookieMiddleware), 80)
	}
	if config.GetEnableRedirect() {
		crawler.SetMiddleware(new(middlewares.RedirectMiddleware), 90)
	}
	if config.GetEnableChrome() {
		crawler.SetMiddleware(new(middlewares.ChromeMiddleware), 100)
	}
	if config.GetEnableHttpAuth() {
		crawler.SetMiddleware(new(middlewares.HttpAuthMiddleware), 110)
	}
	if config.GetEnableCompress() {
		crawler.SetMiddleware(new(middlewares.CompressMiddleware), 120)
	}
	if config.GetEnableDecode() {
		crawler.SetMiddleware(new(middlewares.DecodeMiddleware), 130)
	}
	if config.GetEnableDevice() {
		crawler.SetMiddleware(new(middlewares.DeviceMiddleware), 140)
	}

	if config.GetEnableDumpPipeline() {
		crawler.SetPipeline(new(pipelines.DumpPipeline), 10)
	}
	if config.GetEnableFilterPipeline() {
		crawler.SetPipeline(new(pipelines.FilterPipeline), 20)
	}

	return crawler, nil
}
