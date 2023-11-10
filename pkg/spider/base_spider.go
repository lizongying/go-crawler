package spider

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/downloader"
	"github.com/lizongying/go-crawler/pkg/exporter"
	"github.com/lizongying/go-crawler/pkg/filters"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
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

type BaseSpider struct {
	context pkg.Context
	pkg.Crawler
	filter pkg.Filter

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
	redirectMaxTimes      uint8
	timeout               time.Duration
	okHttpCodes           []int
	config                pkg.Config
	logger                pkg.Logger
	spider                pkg.Spider
	options               []pkg.SpiderOption

	jobs      map[string]*Job
	job       *pkg.State
	jobsMutex sync.Mutex

	concurrency uint8
	interval    time.Duration

	pkg.Downloader
	pkg.Exporter

	requestSlots sync.Map
}

func (s *BaseSpider) GetDownloader() pkg.Downloader {
	return s.Downloader
}
func (s *BaseSpider) WithDownloader(downloader pkg.Downloader) pkg.Spider {
	s.Downloader = downloader
	return s
}
func (s *BaseSpider) GetExporter() pkg.Exporter {
	return s.Exporter
}
func (s *BaseSpider) WithExporter(exporter pkg.Exporter) pkg.Spider {
	s.Exporter = exporter
	return s
}
func (s *BaseSpider) Interval() time.Duration {
	return s.interval
}
func (s *BaseSpider) WithInterval(interval time.Duration) pkg.Spider {
	s.interval = interval
	return s
}

func (s *BaseSpider) GetContext() pkg.Context {
	return s.context
}
func (s *BaseSpider) WithContext(ctx pkg.Context) pkg.Spider {
	s.context = ctx
	return s
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
func (s *BaseSpider) RedirectMaxTimes() uint8 {
	return s.redirectMaxTimes
}
func (s *BaseSpider) SetRedirectMaxTimes(redirectMaxTimes uint8) pkg.Spider {
	s.redirectMaxTimes = redirectMaxTimes
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
func (s *BaseSpider) GetFilter() pkg.Filter {
	return s.filter
}
func (s *BaseSpider) SetFilter(filter pkg.Filter) pkg.Spider {
	s.filter = filter
	return s
}
func (s *BaseSpider) GetLogger() pkg.Logger {
	return s.logger
}
func (s *BaseSpider) Options() []pkg.SpiderOption {
	return s.options
}
func (s *BaseSpider) WithOptions(options ...pkg.SpiderOption) pkg.Spider {
	s.options = options
	return s
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

func (s *BaseSpider) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	return ctx.GetTask().Request(ctx, request)
}
func (s *BaseSpider) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
	return ctx.GetTask().YieldRequest(ctx, request)
}
func (s *BaseSpider) MustYieldRequest(ctx pkg.Context, request pkg.Request) {
	if err := s.YieldRequest(ctx, request); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) NewRequest(ctx pkg.Context, options ...pkg.RequestOption) (err error) {
	req := request.NewRequest()
	for _, v := range options {
		v(req)
	}
	return ctx.GetTask().YieldRequest(ctx, req)
}
func (s *BaseSpider) MustNewRequest(ctx pkg.Context, options ...pkg.RequestOption) {
	if err := s.NewRequest(ctx, options...); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) MustYieldItem(c pkg.Context, item pkg.Item) {
	if err := s.YieldItem(c, item); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) YieldExtra(ctx pkg.Context, extra any) (err error) {
	return ctx.GetTask().YieldExtra(ctx, extra)
}
func (s *BaseSpider) MustYieldExtra(ctx pkg.Context, extra any) {
	if err := s.YieldExtra(ctx, extra); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) GetExtra(ctx pkg.Context, extra any) (err error) {
	return ctx.GetTask().GetExtra(ctx, extra)
}
func (s *BaseSpider) MustGetExtra(ctx pkg.Context, extra any) {
	if err := s.GetExtra(ctx, extra); err != nil {
		s.logger.Error(err)
		if errors.Is(err, pkg.ErrQueueTimeout) {
			panic(pkg.ErrQueueTimeout)
		}
	}
}
func (s *BaseSpider) YieldItem(ctx pkg.Context, item pkg.Item) (err error) {
	return ctx.GetTask().YieldItem(ctx, item)
}
func (s *BaseSpider) RequestSlotLoad(slot string) (value any, ok bool) {
	return s.requestSlots.Load(slot)
}
func (s *BaseSpider) RequestSlotStore(slot string, value any) {
	s.requestSlots.Store(slot, value)
}
func (s *BaseSpider) SetRequestRate(slot string, interval time.Duration, concurrency int) {
	if slot == "" {
		slot = "*"
	}

	if concurrency < 1 {
		concurrency = 1
	}

	slotValue, ok := s.requestSlots.Load(slot)
	if !ok {
		requestSlot := rate.NewLimiter(rate.Every(interval/time.Duration(concurrency)), concurrency)
		s.requestSlots.Store(slot, requestSlot)
		return
	}

	limiter := slotValue.(*rate.Limiter)
	limiter.SetBurst(concurrency)
	limiter.SetLimit(rate.Every(interval / time.Duration(concurrency)))

	return
}
func (s *BaseSpider) Start(c pkg.Context) (err error) {
	ctx := context.Background()

	s.context.WithSpiderStatus(pkg.SpiderStatusRunning)
	s.Crawler.GetSignal().SpiderChanged(s.GetContext())

	s.logger.Info("spiderName", s.context.GetSpiderName())
	s.logger.Info("allowedDomains", s.GetAllowedDomains())
	s.logger.Info("okHttpCodes", s.OkHttpCodes())
	s.logger.Info("platforms", s.GetPlatforms())
	s.logger.Info("browsers", s.GetBrowsers())
	s.logger.Info("retryMaxTimes", s.retryMaxTimes)
	s.logger.Info("redirectMaxTimes", s.redirectMaxTimes)

	//s.logger.Info("filter", s.GetFilter())

	s.logger.Info("pipelines", s.PipelineNames())

	for _, v := range s.Pipelines() {
		e := v.Start(ctx, s.spider)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.Name())
			break
		}
	}

	for _, v := range s.GetMiddlewares().Middlewares() {
		if err = v.Start(ctx, s.spider); err != nil {
			s.logger.Error(err)
			return
		}
		s.logger.Info(v.Name(), "started")
	}
	return
}
func (s *BaseSpider) Run(ctx context.Context, jobFunc string, args string, mode pkg.JobMode, spec string, onlyOneTask bool) (id string, err error) {
	if s.GetContext() == nil {
		err = errors.New("spider hasn't started")
		s.logger.Error(err)
		return
	}

	s.context.WithSpiderContext(ctx)

	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()

	job := new(Job).FromSpider(s.spider)
	if err = job.start(s.context, jobFunc, args, mode, spec, onlyOneTask); err != nil {
		s.logger.Error(err)
		return
	}

	id = job.context.GetJobId()
	s.jobs[id] = job

	if err = job.run(ctx); err != nil {
		s.logger.Error(err)
		return
	}

	s.job.In()
	return
}
func (s *BaseSpider) RerunJob(ctx context.Context, jobId string) (err error) {
	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()

	job, ok := s.jobs[jobId]
	if !ok {
		err = errors.New("job is not exists")
		return
	}
	if job.GetContext().GetJobStatus() != pkg.JobStatusStopped {
		err = errors.New("job is not stopped")
		return
	}
	err = job.run(ctx)
	return
}
func (s *BaseSpider) KillJob(ctx context.Context, jobId string) (err error) {
	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()

	job, ok := s.jobs[jobId]
	if !ok {
		err = errors.New("the job is not exists")
		return
	}
	if !utils.InSlice(job.context.GetJobStatus(), []pkg.JobStatus{
		pkg.JobStatusRunning,
	}) {
		err = errors.New("the job can be killed in the running state")
		return
	}
	err = job.kill(ctx)
	return
}
func (s *BaseSpider) JobStopped(ctx pkg.Context, err error) {
	if err != nil {
		s.logger.Info(s.spider.Name(), ctx.GetJobId(), "job finished with an error:", err, "spend time:", ctx.GetJobStopTime().Sub(ctx.GetJobStartTime()))
	} else {
		s.logger.Info(s.spider.Name(), ctx.GetJobId(), "job finished successfully. spend time:", ctx.GetJobStopTime().Sub(ctx.GetJobStartTime()))
	}

	s.job.Out()
}
func (s *BaseSpider) Parse(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	return
}
func (s *BaseSpider) Error(_ pkg.Context, response pkg.Response, err error) {
	if response.GetResponse() == nil {
		s.logger.Error("response nil")
		return
	}
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.BodyStr())
	s.logger.Info("error", err)
	return
}
func (s *BaseSpider) Stop(_ pkg.Context) (err error) {
	if s.context == nil || s.context.GetSpider() == nil {
		s.logger.Debug("spider hasn't started")
		return
	}

	if s.context.GetSpiderStatus() == pkg.SpiderStatusStopping || s.context.GetSpiderStatus() == pkg.SpiderStatusStopped {
		s.logger.Debug("stopped")
		return
	}

	s.context.WithSpiderStatus(pkg.SpiderStatusStopping)
	s.Crawler.GetSignal().SpiderChanged(s.context)

	s.logger.Debug("BaseSpider wait for stop")
	defer func() {
		err = s.spider.Stop(s.context)
		if errors.Is(err, pkg.DontStopErr) {
			s.logger.Error(err)
			select {}
		}

		stopTime := time.Now()
		s.context.WithSpiderStatus(pkg.SpiderStatusStopped)
		s.Crawler.GetSignal().SpiderChanged(s.context)

		spendTime := stopTime.Sub(s.context.GetSpiderStartTime())
		s.logger.Info(s.spider.Name(), "spider finished. spend time:", spendTime)
		s.Crawler.SpiderStopped(s.context, err)
	}()

	for _, v := range s.GetMiddlewares().Middlewares() {
		e := v.Stop(s.context)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("middlewares break", v.Name())
			break
		}
	}
	for _, v := range s.Pipelines() {
		e := v.Stop(s.context)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("pipeline break", v.Name())
			break
		}
	}

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

	s.concurrency = config.GetRequestConcurrency()
	s.interval = time.Millisecond * time.Duration(int(config.GetRequestInterval()))

	switch config.GetFilter() {
	case pkg.FilterMemory:
		s.SetFilter(new(filters.MemoryFilter).FromSpider(s))
	case pkg.FilterRedis:
		s.SetFilter(new(filters.RedisFilter).FromSpider(s))
	default:
	}

	s.WithDownloader(new(downloader.Downloader).FromSpider(s))
	s.WithExporter(new(exporter.Exporter).FromSpider(s))

	slot := "*"
	if _, ok := s.RequestSlotLoad(slot); !ok {
		requestSlot := rate.NewLimiter(rate.Every(s.interval/time.Duration(s.concurrency)), int(s.concurrency))
		s.RequestSlotStore(slot, requestSlot)
	}

	for _, option := range s.Options() {
		option(s)
	}

	s.WithContext(new(crawlerContext.Context).
		WithCrawler(crawler.GetContext().GetCrawler()).
		WithSpider(new(crawlerContext.Spider).
			WithId(s.Crawler.GenUid()).
			WithName(s.spider.Name()).
			WithStatus(pkg.SpiderStatusReady)))
	s.Crawler.GetSignal().SpiderChanged(s.GetContext())

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
		jobs:                  make(map[string]*Job),
	}

	s.job = pkg.NewState()
	s.job.RegisterIsReadyAndIsZero(func() {
		_ = s.Stop(s.GetContext())
	})

	return s, nil
}
