package spider

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/devServer"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/time/rate"
	"log"
	"reflect"
	"runtime"
	"sync"
	"time"
)

var (
	buildBranch string
	buildCommit string
	buildTime   string
)

func init() {
	info := fmt.Sprintf("Branch: %s, Commit: %s, Time: %s, GOVersion: %s, OS: %s, ARCH: %s", buildBranch, buildCommit, buildTime, runtime.Version(), runtime.GOOS, runtime.GOARCH)
	log.Println(info)
}

const defaultChanRequestMax = 1000 * 1000
const defaultMaxRequestActive = 1000
const defaultChanItemMax = 1000 * 1000
const defaultRequestConcurrency = 1
const defaultRequestInterval = 1
const defaultRequestRetryMaxTimes = 3

type BaseSpider struct {
	*pkg.SpiderInfo
	spider pkg.Spider
	Stats  pkg.Stats

	MongoDb    *mongo.Database
	Mysql      *sql.DB
	Kafka      *kafka.Writer
	Logger     pkg.Logger
	httpClient pkg.HttpClient

	startFunc             string
	args                  string
	itemConcurrency       int
	itemConcurrencyNew    int
	itemConcurrencyChan   chan struct{}
	itemDelay             time.Duration
	itemTimer             *time.Timer
	itemChan              chan pkg.Item
	itemActiveChan        chan struct{}
	requestChan           chan *pkg.Request
	requestActiveChan     chan struct{}
	requestSlots          sync.Map
	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}
	middlewares           map[uint8]pkg.Middleware

	devServer *devServer.HttpServer

	okHttpCodes []int
	platforms   map[pkg.Platform]struct{}
	browsers    map[pkg.Browser]struct{}

	config *config.Config
}

func (s *BaseSpider) SetPlatforms(platforms ...pkg.Platform) pkg.Spider {
	for _, platform := range platforms {
		if platform == "" {
			err := errors.New("platform error")
			s.Logger.Warn(err)
			continue
		}
		s.platforms[platform] = struct{}{}
	}

	return s
}

func (s *BaseSpider) GetPlatforms() (platforms []pkg.Platform) {
	for k := range s.platforms {
		platforms = append(platforms, k)
	}
	return
}

func (s *BaseSpider) SetBrowsers(browsers ...pkg.Browser) pkg.Spider {
	for _, browser := range browsers {
		if browser == "" {
			err := errors.New("browser error")
			s.Logger.Warn(err)
			continue
		}
		s.browsers[browser] = struct{}{}
	}

	return s
}

func (s *BaseSpider) GetBrowsers() (browsers []pkg.Browser) {
	for k := range s.browsers {
		browsers = append(browsers, k)
	}
	return
}

func (s *BaseSpider) SetLogger(logger pkg.Logger) {
	s.Logger = logger
}

func (s *BaseSpider) SetSpider(spider pkg.Spider) {
	s.spider = spider
}

func (s *BaseSpider) GetInfo() *pkg.SpiderInfo {
	return s.SpiderInfo
}

func (s *BaseSpider) GetStats() pkg.Stats {
	return s.Stats
}

func (s *BaseSpider) RunDevServer() (err error) {
	err = s.devServer.Run()
	if err != nil {
		s.Logger.Error(err)
		return
	}

	return
}

func (s *BaseSpider) AddDevServerRoutes(routes ...pkg.Route) pkg.Spider {
	s.devServer.AddRoutes(routes...)
	return s
}

func (s *BaseSpider) GetDevServerHost() (host string) {
	host = s.devServer.GetHost()
	return
}

func (s *BaseSpider) AddOkHttpCodes(httpCodes ...int) pkg.Spider {
	for _, v := range httpCodes {
		if utils.InSlice(v, s.okHttpCodes) {
			continue
		}
		s.okHttpCodes = append(s.okHttpCodes, v)
	}
	return s
}

func (s *BaseSpider) GetOkHttpCodes() (httpCodes []int) {
	httpCodes = s.okHttpCodes
	return
}

func (s *BaseSpider) GetConfig() pkg.Config {
	return s.config
}

func (s *BaseSpider) GetLogger() pkg.Logger {
	return s.Logger
}

func (s *BaseSpider) GetHttpClient() pkg.HttpClient {
	return s.httpClient
}

func (s *BaseSpider) GetKafka() *kafka.Writer {
	return s.Kafka
}

func (s *BaseSpider) GetMongoDb() *mongo.Database {
	return s.MongoDb
}

func (s *BaseSpider) GetMysql() *sql.DB {
	return s.Mysql
}

func (s *BaseSpider) Start(ctx context.Context) (err error) {
	if s.spider == nil {
		err = errors.New("nil spider")
		s.Logger.Error(err)
		return
	}

	if s.Name == "" {
		err = errors.New("name is empty")
		return
	}
	s.Logger.Info("name", s.Name)
	s.Logger.Info("start func", s.startFunc)
	s.Logger.Info("args", s.args)
	s.Logger.Info("mode", s.Mode)
	s.Logger.Info("allowedDomains", s.spider.GetAllowedDomains())
	s.Logger.Info("middlewares", s.spider.GetMiddlewares())
	s.Logger.Info("okHttpCodes", s.okHttpCodes)
	s.Logger.Info("platforms", s.spider.GetPlatforms())
	s.Logger.Info("browsers", s.spider.GetBrowsers())
	s.Logger.Info("referrerPolicy", s.config.GetReferrerPolicy())
	s.Logger.Info("urlLengthLimit", s.config.GetUrlLengthLimit())
	if s.spider == nil {
		err = errors.New("spider is empty")
		s.Logger.Error(err)
		return
	}

	if ctx == nil {
		ctx = context.Background()
	}

	for _, v := range s.middlewares {
		e := v.SpiderStart(ctx, s)
		if errors.Is(e, pkg.BreakErr) {
			break
		}
	}

	defer func() {
		for _, v := range s.middlewares {
			e := v.SpiderStop(ctx)
			if errors.Is(e, pkg.BreakErr) {
				break
			}
		}
	}()

	s.itemTimer = time.NewTimer(s.itemDelay)
	if s.itemConcurrency < 1 {
		s.itemConcurrency = 1
	}
	s.itemConcurrencyNew = s.itemConcurrency
	s.itemConcurrencyChan = make(chan struct{}, s.itemConcurrency)
	for i := 0; i < s.itemConcurrency; i++ {
		s.itemConcurrencyChan <- struct{}{}
	}

	slot := "*"
	if _, ok := s.requestSlots.Load(slot); !ok {
		requestSlot := rate.NewLimiter(rate.Every(s.Interval/time.Duration(s.Concurrency)), s.Concurrency)
		s.requestSlots.Store(slot, requestSlot)
	}

	go s.handleItem(ctx)

	go s.handleRequest(ctx)

	params := []reflect.Value{
		reflect.ValueOf(ctx),
		reflect.ValueOf(s.args),
	}
	caller := reflect.ValueOf(s.spider).MethodByName(s.startFunc)
	if !caller.IsValid() {
		err = errors.New("start func is invalid")
		s.Logger.Error(err)
		return
	}

	// TODO handle result and do something
	r := caller.Call(params)[0].Interface()
	if r != nil {
		err = r.(error)
		s.Logger.Error(err)
		return
	}

	return
}

func (s *BaseSpider) Stop(ctx context.Context) (err error) {
	s.Logger.Debug("Wait for stop")
	defer func() {
		s.Logger.Info("Stopped")
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	ticker := time.NewTicker(time.Second)
	for {
		<-ticker.C
		if len(s.requestActiveChan) > 0 {
			s.Logger.Debug("request is active")
			continue
		}
		if len(s.itemActiveChan) > 0 {
			s.Logger.Debug("item is active")
			continue
		}
		break
	}

	return
}

func NewBaseSpider(cli *cli.Cli, config *config.Config, logger *logger.Logger, mongoDb *mongo.Database, mysql *sql.DB, kafka *kafka.Writer, server *devServer.HttpServer) (spider *BaseSpider, err error) {
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
	retryMaxTimes := defaultRequestRetryMaxTimes
	if config.Request.RetryMaxTimes > 0 {
		retryMaxTimes = config.Request.RetryMaxTimes
	}
	if config.Request.RetryMaxTimes < 0 {
		retryMaxTimes = 0
	}
	timeout := time.Minute
	if config.Request.Timeout > 0 {
		timeout = time.Second * time.Duration(config.Request.Timeout)
	}

	spider = &BaseSpider{
		SpiderInfo: &pkg.SpiderInfo{
			Concurrency:   concurrency,
			Interval:      time.Millisecond * time.Duration(interval),
			RetryMaxTimes: retryMaxTimes,
			Timeout:       timeout,
		},
		Stats:       &stats.Stats{},
		okHttpCodes: okHttpCodes,
		startFunc:   cli.StartFunc,
		args:        cli.Args,
		MongoDb:     mongoDb,
		Mysql:       mysql,
		Kafka:       kafka,
		Logger:      logger,

		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
		middlewares:           make(map[uint8]pkg.Middleware),
		requestChan:           make(chan *pkg.Request, defaultChanRequestMax),
		requestActiveChan:     make(chan struct{}, defaultChanRequestMax),
		itemChan:              make(chan pkg.Item, defaultChanItemMax),
		itemActiveChan:        make(chan struct{}, defaultChanItemMax),

		devServer: server,

		platforms: make(map[pkg.Platform]struct{}, 6),
		browsers:  make(map[pkg.Browser]struct{}, 4),
		config:    config,
	}
	spider.Mode = cli.Mode
	spider.httpClient = new(httpClient.HttpClient).FromCrawler(spider)

	spider.
		SetMiddleware(middlewares.NewStatsMiddleware, 100).
		SetMiddleware(middlewares.NewFilterMiddleware, 110).
		SetMiddleware(middlewares.NewRetryMiddleware, 120).
		SetMiddleware(middlewares.NewUrlMiddleware, 130).
		SetMiddleware(middlewares.NewRefererMiddleware, 140).
		SetMiddleware(middlewares.NewHttpMiddleware, 150).
		SetMiddleware(middlewares.NewDumpMiddleware, 160)

	return
}
