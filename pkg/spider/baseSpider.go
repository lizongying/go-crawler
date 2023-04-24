package spider

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"github.com/lizongying/go-crawler/pkg/config"
	"github.com/lizongying/go-crawler/pkg/httpClient"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"go.mongodb.org/mongo-driver/mongo"
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
	requestSlots          sync.Map
	requestSlotsCurrent   map[string]pkg.RequestSlot
	requestChan           chan *pkg.Request
	requestActiveChan     chan struct{}
	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}
	middlewares           map[int]pkg.Middleware
	pipelines             map[int]pkg.Pipeline

	TimeoutRequest time.Duration

	locker sync.Mutex
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

func (s *BaseSpider) Start(ctx context.Context) (err error) {
	s.locker.Lock()
	defer s.locker.Unlock()
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

	s.requestSlots.Range(func(key, value any) bool {
		requestSlot := value.(*pkg.RequestSlot)
		if requestSlot.Delay > 0 {
			requestSlot.Timer = time.NewTimer(requestSlot.Delay)
		}
		requestSlot.ConcurrencyChan = make(chan struct{}, requestSlot.Concurrency)
		for i := 0; i < requestSlot.Concurrency; i++ {
			requestSlot.ConcurrencyChan <- struct{}{}
		}
		s.requestSlotsCurrent[key.(string)] = *requestSlot
		return true
	})

	slot := "*"
	if _, ok := s.requestSlots.Load(slot); !ok {
		requestSlot := new(pkg.RequestSlot)
		requestSlot.Concurrency = 1
		requestSlot.ConcurrencyChan = make(chan struct{}, requestSlot.Concurrency)
		for i := 0; i < requestSlot.Concurrency; i++ {
			requestSlot.ConcurrencyChan <- struct{}{}
		}
		s.requestSlots.Store(slot, requestSlot)
		s.requestSlotsCurrent[slot] = *requestSlot
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

func NewBaseSpider(cli *cli.Cli, _ *config.Config, logger *logger.Logger, mongoDb *mongo.Database, httpClient *httpClient.HttpClient) (spider *BaseSpider, err error) {
	defaultAllowedDomains := map[string]struct{}{"*": {}}

	spider = &BaseSpider{
		SpiderInfo: new(pkg.SpiderInfo),
		startFunc:  cli.StartFunc,
		MongoDb:    mongoDb,
		Logger:     logger,
		httpClient: httpClient,

		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
		middlewares:           make(map[int]pkg.Middleware),
		pipelines:             make(map[int]pkg.Pipeline),
		requestSlotsCurrent:   make(map[string]pkg.RequestSlot),
		requestChan:           make(chan *pkg.Request, defaultChanRequestMax),
		requestActiveChan:     make(chan struct{}, defaultChanRequestMax),
		itemChan:              make(chan *pkg.Item, defaultChanItemMax),
		itemActiveChan:        make(chan struct{}, defaultChanItemMax),
	}
	spider.Mode = cli.Mode

	spider.SetMiddleware(middlewares.NewRecorderMiddleware(logger), 100)
	spider.SetMiddleware(middlewares.NewFilterMiddleware(logger), 110)
	spider.SetMiddleware(middlewares.NewHttpMiddleware(logger, httpClient), 120)
	spider.SetPipeline(pipelines.NewMongoPipeline(logger, mongoDb), 100)

	return
}
