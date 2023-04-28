package spider

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/httpServer"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/time/rate"
	"log"
	"reflect"
	"runtime"
	"sort"
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
const defaultChanItemMax = 1000 * 1000
const defaultRequestConcurrency = 1
const defaultRequestInterval = 1
const defaultRequestRetryMaxTimes = 3

type BaseSpider struct {
	*pkg.SpiderInfo
	spider pkg.Spider

	MongoDb    *mongo.Database
	Logger     pkg.Logger
	httpClient *httpClient.HttpClient

	startFunc             string
	itemConcurrency       int
	itemConcurrencyNew    int
	itemConcurrencyChan   chan struct{}
	itemDelay             time.Duration
	itemTimer             *time.Timer
	itemChan              chan *pkg.Item
	itemActiveChan        chan struct{}
	requestChan           chan *pkg.Request
	requestActiveChan     chan struct{}
	requestSlots          sync.Map
	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}
	middlewares           map[int]pkg.Middleware
	pipelines             map[int]pkg.Pipeline

	devServer *httpServer.HttpServer
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

func (s *BaseSpider) SortedMiddlewares() (o []pkg.Middleware) {
	keys := make([]int, 0)
	for k := range s.middlewares {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		o = append(o, s.middlewares[key])
	}

	return
}

func (s *BaseSpider) SortedPipelines() (o []pkg.Pipeline) {
	keys := make([]int, 0)
	for k := range s.pipelines {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, key := range keys {
		o = append(o, s.pipelines[key])
	}

	return
}

func (s *BaseSpider) GetDevServer() pkg.DevServer {
	return s.devServer
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
	s.Logger.Info("mode", s.Mode)
	s.Logger.Info("allowedDomains", s.spider.GetAllowedDomains())
	s.Logger.Info("middlewares", s.spider.GetMiddlewares())
	s.Logger.Info("pipelines", s.spider.GetPipelines())
	s.Logger.Info("okHttpCodes", s.spider.GetInfo().OkHttpCodes)
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

	for _, v := range s.pipelines {
		e := v.SpiderStart(ctx, s)
		if errors.Is(e, pkg.BreakErr) {
			break
		}
	}

	defer func() {
		for _, v := range s.pipelines {
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
	s.Logger.Info("Wait for stop")
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

func NewBaseSpider(cli *cli.Cli, config *config.Config, logger *logger.Logger, mongoDb *mongo.Database, httpClient *httpClient.HttpClient, server *httpServer.HttpServer) (spider *BaseSpider, err error) {
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

	if cli.Mode == "dev" {
		server.AddRoutes(httpServer.NewNoLimitHandler(logger))
	}

	spider = &BaseSpider{
		SpiderInfo: &pkg.SpiderInfo{
			Concurrency:   concurrency,
			Interval:      time.Second * time.Duration(interval),
			OkHttpCodes:   okHttpCodes,
			RetryMaxTimes: retryMaxTimes,
			Timeout:       timeout,
		},
		startFunc:  cli.StartFunc,
		MongoDb:    mongoDb,
		Logger:     logger,
		httpClient: httpClient,

		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
		middlewares:           make(map[int]pkg.Middleware),
		pipelines:             make(map[int]pkg.Pipeline),
		requestChan:           make(chan *pkg.Request, defaultChanRequestMax),
		requestActiveChan:     make(chan struct{}, defaultChanRequestMax),
		itemChan:              make(chan *pkg.Item, defaultChanItemMax),
		itemActiveChan:        make(chan struct{}, defaultChanItemMax),

		devServer: server,
	}
	spider.Mode = cli.Mode

	spider.SetMiddleware(middlewares.NewRecorderMiddleware(logger), 100)
	spider.SetMiddleware(middlewares.NewFilterMiddleware(logger), 110)
	spider.SetMiddleware(middlewares.NewHttpMiddleware(logger, httpClient), 120)
	spider.SetMiddleware(middlewares.NewRetryMiddleware(logger), 130)
	spider.SetPipeline(pipelines.NewMongoPipeline(logger, mongoDb), 100)

	return
}
