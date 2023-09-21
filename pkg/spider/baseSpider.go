package spider

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/filters"
	"github.com/lizongying/go-crawler/pkg/middlewares"
	kafka2 "github.com/lizongying/go-crawler/pkg/scheduler/kafka"
	"github.com/lizongying/go-crawler/pkg/scheduler/memory"
	redis2 "github.com/lizongying/go-crawler/pkg/scheduler/redis"
	"github.com/lizongying/go-crawler/pkg/signals"
	"github.com/lizongying/go-crawler/pkg/stats"
	"github.com/lizongying/go-crawler/pkg/utils"
	"log"
	"reflect"
	"runtime"
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

type BaseSpider struct {
	pkg.Crawler
	pkg.Stats
	pkg.Signal
	pkg.Scheduler
	filter      pkg.Filter
	middlewares pkg.Middlewares

	name                  string
	host                  string
	username              string
	password              string
	platforms             map[pkg.Platform]struct{}
	browsers              map[pkg.Browser]struct{}
	callBacks             map[string]pkg.CallBack
	errBacks              map[string]pkg.ErrBack
	defaultAllowedDomains map[string]struct{}
	allowedDomains        map[string]struct{}
	retryMaxTimes         uint8
	timeout               time.Duration
	okHttpCodes           []int
	config                pkg.Config
	logger                pkg.Logger
	spider                pkg.Spider
	options               []pkg.SpiderOption

	stateRequest *pkg.State
	stateItem    *pkg.State
	stateMethod  *pkg.State
	couldStop    chan struct{}
}

func (s *BaseSpider) Name() string {
	return s.name
}
func (s *BaseSpider) SetName(name string) pkg.Spider {
	s.name = name
	return s
}
func (s *BaseSpider) GetHost() string {
	return s.host
}
func (s *BaseSpider) SetHost(host string) pkg.Spider {
	s.host = host
	return s
}
func (s *BaseSpider) Username() string {
	return s.username
}
func (s *BaseSpider) SetUsername(username string) pkg.Spider {
	s.username = username
	return s
}
func (s *BaseSpider) Password() string {
	return s.password
}
func (s *BaseSpider) SetPassword(password string) pkg.Spider {
	s.password = password
	return s
}
func (s *BaseSpider) GetPlatforms() (platforms []pkg.Platform) {
	for k := range s.platforms {
		platforms = append(platforms, k)
	}
	return
}
func (s *BaseSpider) SetPlatforms(platforms ...pkg.Platform) pkg.Spider {
	for _, platform := range platforms {
		if platform == "" {
			err := errors.New("platform error")
			s.logger.Warn(err)
			continue
		}
		s.platforms[platform] = struct{}{}
	}
	return s
}
func (s *BaseSpider) GetBrowsers() (browsers []pkg.Browser) {
	for k := range s.browsers {
		browsers = append(browsers, k)
	}
	return
}
func (s *BaseSpider) SetBrowsers(browsers ...pkg.Browser) pkg.Spider {
	for _, browser := range browsers {
		if browser == "" {
			err := errors.New("browser error")
			s.logger.Warn(err)
			continue
		}
		s.browsers[browser] = struct{}{}
	}
	return s
}
func (s *BaseSpider) RetryMaxTimes() uint8 {
	return s.retryMaxTimes
}
func (s *BaseSpider) SetRetryMaxTimes(retryMaxTimes uint8) pkg.Spider {
	s.retryMaxTimes = retryMaxTimes
	return s
}
func (s *BaseSpider) Timeout() time.Duration {
	return s.timeout
}
func (s *BaseSpider) SetTimeout(timeout time.Duration) pkg.Spider {
	s.timeout = timeout
	return s
}
func (s *BaseSpider) OkHttpCodes() (httpCodes []int) {
	httpCodes = s.okHttpCodes
	return
}
func (s *BaseSpider) SetOkHttpCodes(httpCodes ...int) pkg.Spider {
	for _, v := range httpCodes {
		if utils.InSlice(v, s.okHttpCodes) {
			continue
		}
		s.okHttpCodes = append(s.okHttpCodes, v)
	}
	return s
}
func (s *BaseSpider) GetSpider() pkg.Spider {
	return s.spider
}
func (s *BaseSpider) SetSpider(spider pkg.Spider) pkg.Spider {
	s.spider = spider
	s.registerParser()
	return s
}
func (s *BaseSpider) CallBacks() map[string]pkg.CallBack {
	return s.callBacks
}
func (s *BaseSpider) CallBack(name string) (callback pkg.CallBack) {
	if name != "" {
		callback = s.callBacks[name]
	}
	if callback == nil {
		callback = s.Parse
	}
	return
}
func (s *BaseSpider) SetCallBacks(callBacks map[string]pkg.CallBack) pkg.Spider {
	s.callBacks = callBacks
	return s
}
func (s *BaseSpider) ErrBacks() map[string]pkg.ErrBack {
	return s.errBacks
}
func (s *BaseSpider) ErrBack(name string) (errBack pkg.ErrBack) {
	if name != "" {
		errBack = s.errBacks[name]
	}
	if errBack == nil {
		errBack = s.Error
	}
	return
}
func (s *BaseSpider) SetErrBacks(errBacks map[string]pkg.ErrBack) pkg.Spider {
	s.errBacks = errBacks
	return s
}
func (s *BaseSpider) GetCrawler() pkg.Crawler {
	return s.Crawler
}
func (s *BaseSpider) GetStats() pkg.Stats {
	return s.Stats
}
func (s *BaseSpider) SetStats(stats pkg.Stats) pkg.Spider {
	s.Stats = stats
	return s
}
func (s *BaseSpider) GetSignal() pkg.Signal {
	return s.Signal
}
func (s *BaseSpider) SetSignal(signal pkg.Signal) {
	s.Signal = signal
}
func (s *BaseSpider) GetFilter() pkg.Filter {
	return s.filter
}
func (s *BaseSpider) SetFilter(filter pkg.Filter) pkg.Spider {
	s.filter = filter
	return s
}
func (s *BaseSpider) GetScheduler() pkg.Scheduler {
	return s.Scheduler
}
func (s *BaseSpider) SetScheduler(scheduler pkg.Scheduler) pkg.Spider {
	scheduler.SetScheduler(scheduler)
	scheduler.SetLogger(s.logger)
	s.Scheduler = scheduler
	return s
}
func (s *BaseSpider) GetMiddlewares() pkg.Middlewares {
	return s.middlewares
}
func (s *BaseSpider) SetMiddlewares(middlewares pkg.Middlewares) pkg.Spider {
	s.middlewares = middlewares
	return s
}
func (s *BaseSpider) GetLogger() pkg.Logger {
	return s.logger
}
func (s *BaseSpider) SetLogger(logger pkg.Logger) pkg.Spider {
	s.logger = logger
	return s
}
func (s *BaseSpider) Options() []pkg.SpiderOption {
	return s.options
}
func (s *BaseSpider) WithOptions(options ...pkg.SpiderOption) pkg.Spider {
	s.options = options
	return s
}
func (s *BaseSpider) StateRequest() *pkg.State {
	return s.stateRequest
}
func (s *BaseSpider) StateItem() *pkg.State {
	return s.stateItem
}
func (s *BaseSpider) StateMethod() *pkg.State {
	return s.stateMethod
}
func (s *BaseSpider) registerParser() {
	callBacks := make(map[string]pkg.CallBack)
	errBacks := make(map[string]pkg.ErrBack)
	rv := reflect.ValueOf(s.spider)
	rt := rv.Type()
	l := rt.NumMethod()
	for i := 0; i < l; i++ {
		name := rt.Method(i).Name
		callBack, ok := rv.Method(i).Interface().(func(pkg.Context, pkg.Response) error)
		if ok {
			callBacks[name] = callBack
		}
		errBack, ok := rv.Method(i).Interface().(func(pkg.Context, pkg.Response, error))
		if ok {
			errBacks[name] = errBack
		}
	}
	s.SetCallBacks(callBacks)
	s.SetErrBacks(errBacks)
}
func (s *BaseSpider) Start(ctx context.Context, taskId string, startFunc string, args string) (err error) {
	timeBegin := time.Now()
	defer func() {
		s.logger.Info("spend time", time.Now().Sub(timeBegin))
	}()

	if ctx == nil {
		ctx = context.Background()
	}

	s.Stats = &stats.Stats{}

	resultChan := make(chan struct{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				if e, ok := r.(error); ok {
					if errors.Is(e, pkg.ErrQueueTimeout) {
						s.logger.Error(pkg.ErrQueueTimeout)
						return
					}
				}
				buf := make([]byte, 1<<16)
				runtime.Stack(buf, true)
				err = errors.New(string(buf))
				s.logger.Error(err)
			}
			resultChan <- struct{}{}
		}()

		states := pkg.NewMultiState(s.stateRequest, s.stateItem, s.stateMethod)
		states.RegisterSetAndZeroFn(func() {
			for _, v := range s.middlewares.Middlewares() {
				e := v.Stop(ctx)
				if errors.Is(e, pkg.BreakErr) {
					s.logger.Debug("middlewares break", v.Name())
					break
				}
			}
			for _, v := range s.Pipelines() {
				e := v.Stop(ctx)
				if errors.Is(e, pkg.BreakErr) {
					s.logger.Debug("pipeline break", v.Name())
					break
				}
			}
			s.couldStop <- struct{}{}
		})

		if err = s.StartScheduler(ctx); err != nil {
			s.logger.Error(err)
			return
		}

		s.Signal.SpiderOpened()

		c := pkg.Context{
			Spider: s.spider,
			TaskId: taskId,
		}
		params := []reflect.Value{
			reflect.ValueOf(c),
			reflect.ValueOf(args),
		}
		caller := reflect.ValueOf(s.spider).MethodByName(startFunc)
		if !caller.IsValid() {
			err = errors.New("start func is invalid")
			s.logger.Error(err)
			return
		}

		res := caller.Call(params)
		if len(res) < 1 {
			return
		}

		if !res[0].IsNil() {
			var ok bool
			err, ok = res[0].Interface().(error)
			if !ok {
				err = errors.New("spider returns an unknown type")
				s.logger.Error(err)
				return
			}

			s.logger.Error(err)
		}

		<-s.couldStop
	}()

	select {
	case <-resultChan:
		s.logger.Info("finished")
		return
	case <-ctx.Done():
		close(resultChan)
		err = pkg.ErrSpiderTimeout
		s.logger.Error(err)
		return
	}
}
func (s *BaseSpider) Parse(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	return
}
func (s *BaseSpider) Error(_ pkg.Context, response pkg.Response, err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	s.logger.Info("error", err)
	return
}
func (s *BaseSpider) Stop(ctx context.Context) (err error) {
	s.logger.Debug("BaseSpider wait for stop")
	defer func() {
		if err == nil {
			s.logger.Info("BaseSpider Stopped")
		}
	}()

	if err = s.StopScheduler(ctx); err != nil {
		s.logger.Error(err)
		return
	}

	s.Signal.SpiderClosed()
	return
}
func (s *BaseSpider) FromCrawler(crawler pkg.Crawler) pkg.Spider {
	if s == nil {
		return new(BaseSpider).FromCrawler(crawler)
	}

	s.logger = crawler.GetLogger()

	s.Crawler = crawler

	config := crawler.GetConfig()
	s.retryMaxTimes = config.GetRetryMaxTimes()
	s.timeout = config.GetRequestTimeout()
	s.okHttpCodes = config.GetOkHttpCodes()

	s.SetSignal(new(signals.Signal).FromSpider(s))

	switch config.GetFilter() {
	case pkg.FilterMemory:
		s.SetFilter(new(filters.MemoryFilter).FromSpider(s))
	case pkg.FilterRedis:
		s.SetFilter(new(filters.RedisFilter).FromSpider(s))
	default:
	}

	s.middlewares = new(middlewares.Middlewares).FromSpider(s)

	switch config.GetScheduler() {
	case pkg.SchedulerMemory:
		s.SetScheduler(new(memory.Scheduler).FromSpider(s))
	case pkg.SchedulerRedis:
		s.SetScheduler(new(redis2.Scheduler).FromSpider(s))
	case pkg.SchedulerKafka:
		s.SetScheduler(new(kafka2.Scheduler).FromSpider(s))
	default:
	}

	return s
}
func NewBaseSpider(logger pkg.Logger) (pkg.Spider, error) {
	defaultAllowedDomains := map[string]struct{}{"*": {}}
	s := &BaseSpider{
		logger:                logger,
		platforms:             make(map[pkg.Platform]struct{}, 6),
		browsers:              make(map[pkg.Browser]struct{}, 4),
		defaultAllowedDomains: defaultAllowedDomains,
		allowedDomains:        defaultAllowedDomains,
	}
	s.stateRequest = pkg.NewState()
	s.stateItem = pkg.NewState()
	s.stateMethod = pkg.NewState()
	s.stateItem.Set()
	s.stateMethod.Set()
	s.couldStop = make(chan struct{}, 10) // TODO

	return s, nil
}
