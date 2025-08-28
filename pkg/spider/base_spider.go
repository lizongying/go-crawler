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
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"golang.org/x/time/rate"
	"net/url"
	"sync"
	"time"
)

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
	callBacks             sync.Map
	errBacks              sync.Map
	startFuncs            sync.Map
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

func (s *BaseSpider) PipelineNames() map[uint8]string {
	return s.Exporter.Names()
}
func (s *BaseSpider) Pipelines() []pkg.Pipeline {
	return s.Exporter.Pipelines()
}
func (s *BaseSpider) SetPipeline(pipeline pkg.Pipeline, order uint8) {
	s.Exporter.SetPipeline(pipeline, order)
}
func (s *BaseSpider) RemovePipeline(index int) {
	s.Exporter.Remove(index)
}
func (s *BaseSpider) CleanPipelines() {
	s.Exporter.Clean()
}
func (s *BaseSpider) WithDumpPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithDumpPipeline())
	return s
}
func (s *BaseSpider) WithFilePipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithFilePipeline())
	return s
}
func (s *BaseSpider) WithImagePipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithImagePipeline())
	return s
}
func (s *BaseSpider) WithFilterPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithFilterPipeline())
	return s
}
func (s *BaseSpider) WithNonePipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithNonePipeline())
	return s
}
func (s *BaseSpider) WithCsvPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithCsvPipeline())
	return s
}
func (s *BaseSpider) WithJsonLinesPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithJsonLinesPipeline())
	return s
}
func (s *BaseSpider) WithMongoPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithMongoPipeline())
	return s
}
func (s *BaseSpider) WithSqlitePipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithSqlitePipeline())
	return s
}
func (s *BaseSpider) WithMysqlPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithMysqlPipeline())
	return s
}
func (s *BaseSpider) WithKafkaPipeline() pkg.Spider {
	s.options = append(s.options, pkg.WithKafkaPipeline())
	return s
}
func (s *BaseSpider) WithCustomPipeline(pipeline pkg.Pipeline) pkg.Spider {
	s.options = append(s.options, pkg.WithCustomPipeline(pipeline))
	return s
}
func (s *BaseSpider) WithPipelineDump() pkg.Spider {
	s.options = append(s.options, pkg.WithDumpPipeline())
	return s
}
func (s *BaseSpider) WithPipelineFile() pkg.Spider {
	s.options = append(s.options, pkg.WithFilePipeline())
	return s
}
func (s *BaseSpider) WithPipelineImage() pkg.Spider {
	s.options = append(s.options, pkg.WithImagePipeline())
	return s
}
func (s *BaseSpider) WithPipelineFilter() pkg.Spider {
	s.options = append(s.options, pkg.WithFilterPipeline())
	return s
}
func (s *BaseSpider) WithPipelineNone() pkg.Spider {
	s.options = append(s.options, pkg.WithNonePipeline())
	return s
}
func (s *BaseSpider) WithPipelineCsv() pkg.Spider {
	s.options = append(s.options, pkg.WithCsvPipeline())
	return s
}
func (s *BaseSpider) WithPipelineJsonLines() pkg.Spider {
	s.options = append(s.options, pkg.WithJsonLinesPipeline())
	return s
}
func (s *BaseSpider) WithPipelineMongo() pkg.Spider {
	s.options = append(s.options, pkg.WithMongoPipeline())
	return s
}
func (s *BaseSpider) WithPipelineSqlite() pkg.Spider {
	s.options = append(s.options, pkg.WithSqlitePipeline())
	return s
}
func (s *BaseSpider) WithPipelineMysql() pkg.Spider {
	s.options = append(s.options, pkg.WithMysqlPipeline())
	return s
}
func (s *BaseSpider) WithPipelineKafka() pkg.Spider {
	s.options = append(s.options, pkg.WithKafkaPipeline())
	return s
}
func (s *BaseSpider) WithPipelineCustom(pipeline pkg.Pipeline) pkg.Spider {
	s.options = append(s.options, pkg.WithCustomPipeline(pipeline))
	return s
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
func (s *BaseSpider) GetMiddlewares() pkg.Middlewares {
	return s.Downloader.GetMiddlewares()
}
func (s *BaseSpider) WithMiddleware(middleware pkg.Middleware, order uint8) pkg.Spider {
	s.options = append(s.options, pkg.WithMiddleware(middleware, order))
	return s
}
func (s *BaseSpider) WithStatsMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithStatsMiddleware())
	return s
}
func (s *BaseSpider) WithDumpMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithDumpMiddleware())
	return s
}
func (s *BaseSpider) WithProxyMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithProxyMiddleware())
	return s
}
func (s *BaseSpider) WithRobotsTxtMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithRobotsTxtMiddleware())
	return s
}
func (s *BaseSpider) WithFilterMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithFilterMiddleware())
	return s
}
func (s *BaseSpider) WithFileMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithFileMiddleware())
	return s
}
func (s *BaseSpider) WithImageMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithImageMiddleware())
	return s
}
func (s *BaseSpider) WithHttpMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithHttpMiddleware())
	return s
}
func (s *BaseSpider) WithRetryMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithRetryMiddleware())
	return s
}
func (s *BaseSpider) WithUrlMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithUrlMiddleware())
	return s
}
func (s *BaseSpider) WithReferrerMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithReferrerMiddleware())
	return s
}
func (s *BaseSpider) WithCookieMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithCookieMiddleware())
	return s
}
func (s *BaseSpider) WithRedirectMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithRedirectMiddleware())
	return s
}
func (s *BaseSpider) WithChromeMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithChromeMiddleware())
	return s
}
func (s *BaseSpider) WithHttpAuthMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithHttpAuthMiddleware())
	return s
}
func (s *BaseSpider) WithCompressMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithCompressMiddleware())
	return s
}
func (s *BaseSpider) WithDecodeMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithDecodeMiddleware())
	return s
}
func (s *BaseSpider) WithDeviceMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithDeviceMiddleware())
	return s
}
func (s *BaseSpider) WithRecordErrorMiddleware() pkg.Spider {
	s.options = append(s.options, pkg.WithRecordErrorMiddleware())
	return s
}
func (s *BaseSpider) WithCustomMiddleware(middleware pkg.Middleware) pkg.Spider {
	s.options = append(s.options, pkg.WithCustomMiddleware(middleware))
	return s
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

// RetryMaxTimes returns the maximum number of retry attempts configured for the spider.
func (s *BaseSpider) RetryMaxTimes() uint8 {
	return s.retryMaxTimes
}

// SetRetryMaxTimes sets the maximum number of retry attempts for the spider.
// Returns the spider itself for chaining.
func (s *BaseSpider) SetRetryMaxTimes(retryMaxTimes uint8) pkg.Spider {
	s.retryMaxTimes = retryMaxTimes
	return s
}

// RedirectMaxTimes returns the maximum number of HTTP redirects the spider will follow automatically.
func (s *BaseSpider) RedirectMaxTimes() uint8 {
	return s.redirectMaxTimes
}

// SetRedirectMaxTimes sets the maximum number of HTTP redirects the spider
// will follow automatically. Returns the spider itself for chaining.
func (s *BaseSpider) SetRedirectMaxTimes(redirectMaxTimes uint8) pkg.Spider {
	s.redirectMaxTimes = redirectMaxTimes
	return s
}

// Timeout returns the timeout duration configured for HTTP requests made by the spider.
func (s *BaseSpider) Timeout() time.Duration {
	return s.timeout
}

// SetTimeout sets the timeout for HTTP requests made by the spider.
// Returns the spider itself for chaining.
func (s *BaseSpider) SetTimeout(timeout time.Duration) pkg.Spider {
	s.timeout = timeout
	return s
}

// Interval returns the duration between consecutive requests made by the spider.
func (s *BaseSpider) Interval() time.Duration {
	return s.interval
}

// SetInterval sets the duration between consecutive requests for the spider and returns the spider for chaining.
func (s *BaseSpider) SetInterval(interval time.Duration) pkg.Spider {
	s.interval = interval
	return s
}

// OkHttpCodes returns the list of HTTP status codes considered successful by the spider.
func (s *BaseSpider) OkHttpCodes() (httpCodes []int) {
	httpCodes = s.okHttpCodes
	return
}

// SetOkHttpCodes sets the HTTP status codes that are considered successful.
// Requests returning other codes may trigger retries or errors.
// Returns the spider itself for chaining.
func (s *BaseSpider) SetOkHttpCodes(httpCodes ...int) pkg.Spider {
	for _, v := range httpCodes {
		if utils.InSlice(v, s.okHttpCodes) {
			continue
		}
		s.okHttpCodes = append(s.okHttpCodes, v)
	}
	return s
}

// GetFilter returns the filter function currently set for processing or filtering items.
func (s *BaseSpider) GetFilter() pkg.Filter {
	return s.filter
}

// SetFilter sets a filter function to process or filter items before they are exported.
// Returns the spider itself for chaining.
func (s *BaseSpider) SetFilter(filter pkg.Filter) pkg.Spider {
	s.filter = filter
	return s
}

func (s *BaseSpider) GetSpider() pkg.Spider {
	return s.spider
}

func (s *BaseSpider) SetSpider(spider pkg.Spider) pkg.Spider {
	s.spider = spider
	return s
}

// SetCallBack registers a new callback function under the given name.
func (s *BaseSpider) SetCallBack(name string, callBack pkg.CallBack) {
	s.callBacks.Store(name, callBack)
}

// CallBackNames returns the list of all registered callback names.
func (s *BaseSpider) CallBackNames() (names []string) {
	names = make([]string, 0)
	s.callBacks.Range(func(k any, _ any) bool {
		if key, ok := k.(string); ok {
			names = append(names, key)
		}
		return true
	})
	return
}

// CallBack retrieves the callback function by name.
// Returns an error if the callback does not exist.
func (s *BaseSpider) CallBack(name string) (callback pkg.CallBack, err error) {
	if v, ok := s.callBacks.Load(name); ok {
		callback = v.(pkg.CallBack)
	} else {
		callback = s.Parse
	}
	return
}

// SetErrBack registers a new error handler under the given name.
func (s *BaseSpider) SetErrBack(name string, errBack pkg.ErrBack) {
	s.errBacks.Store(name, errBack)
}

// ErrBackNames returns the list of all registered error handler names.
func (s *BaseSpider) ErrBackNames() (names []string) {
	names = make([]string, 0)
	s.errBacks.Range(func(k any, _ any) bool {
		if key, ok := k.(string); ok {
			names = append(names, key)
		}
		return true
	})
	return
}

// ErrBack returns the error handler (ErrBack) by name.
// If the name is not found or empty, it falls back to the default Error handler.
func (s *BaseSpider) ErrBack(name string) (errBack pkg.ErrBack, err error) {
	if v, ok := s.callBacks.Load(name); ok {
		errBack = v.(pkg.ErrBack)
	} else {
		errBack = s.Error
	}
	return
}

// SetStartFunc registers a new start function under the given name.
func (s *BaseSpider) SetStartFunc(name string, startFunc pkg.StartFunc) {
	s.startFuncs.Store(name, startFunc)
}

// StartFuncNames returns the list of all registered start function names.
func (s *BaseSpider) StartFuncNames() (names []string) {
	names = make([]string, 0)
	s.startFuncs.Range(func(k any, _ any) bool {
		if key, ok := k.(string); ok {
			names = append(names, key)
		}
		return true
	})
	return
}

// StartFunc looks up a registered StartFunc by its name.
// It returns the StartFunc and a nil error if found.
// If no StartFunc exists for the given name, it returns
// ErrStartFuncNotExist.
func (s *BaseSpider) StartFunc(name string) (startFunc pkg.StartFunc, err error) {
	if v, ok := s.startFuncs.Load(name); ok {
		startFunc = v.(pkg.StartFunc)
	} else {
		err = pkg.ErrStartFuncNotExist
	}
	return
}
func (s *BaseSpider) GetCrawler() pkg.Crawler {
	return s.Crawler
}

func (s *BaseSpider) GetLogger() pkg.Logger {
	return s.logger
}
func (s *BaseSpider) Logger() pkg.Logger {
	return s.logger
}
func (s *BaseSpider) Options() []pkg.SpiderOption {
	return s.options
}
func (s *BaseSpider) WithOptions(options ...pkg.SpiderOption) pkg.Spider {
	s.options = options
	return s
}

func (s *BaseSpider) Request(ctx pkg.Context, request pkg.Request) (response pkg.Response, err error) {
	req := request.GetHttpRequest()
	if req.URL.Scheme == "" || req.URL.Host == "" {
		u, e := url.Parse(s.GetHost())
		if e == nil {
			if req.URL.Scheme == "" {
				req.URL.Scheme = u.Scheme
			}
			if req.URL.Host == "" {
				req.URL.Host = u.Host
			}
		}
	}

	return ctx.GetTask().GetTask().Request(ctx, request)
}
func (s *BaseSpider) YieldRequest(ctx pkg.Context, request pkg.Request) (err error) {
	req := request.GetHttpRequest()
	if req.URL.Scheme == "" || req.URL.Host == "" {
		u, e := url.Parse(s.GetHost())
		if e == nil {
			if req.URL.Scheme == "" {
				req.URL.Scheme = u.Scheme
			}
			if req.URL.Host == "" {
				req.URL.Host = u.Host
			}
		}
	}

	return ctx.GetTask().GetTask().YieldRequest(ctx, request)
}
func (s *BaseSpider) MustYieldRequest(ctx pkg.Context, request pkg.Request) {
	if err := s.YieldRequest(ctx, request); err != nil {
		panic(fmt.Errorf("%w: %v", pkg.ErrYieldRequestFailed, err))
	}
}
func (s *BaseSpider) UnsafeYieldRequest(ctx pkg.Context, request pkg.Request) {
	if err := s.YieldRequest(ctx, request); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) NewRequest(ctx pkg.Context, options ...pkg.RequestOption) (req pkg.Request) {
	req = request.NewRequest()
	req = req.WithContext(ctx)
	for _, v := range options {
		v(req)
	}
	return
}
func (s *BaseSpider) MustYieldItem(c pkg.Context, item pkg.Item) {
	if err := s.YieldItem(c, item); err != nil {
		panic(fmt.Errorf("%w: %v", pkg.ErrYieldItemFailed, err))
	}
}
func (s *BaseSpider) UnsafeYieldItem(c pkg.Context, item pkg.Item) {
	if err := s.YieldItem(c, item); err != nil {
		s.logger.Error(err)
	}
}

// NewItemNone creates a new Item that does not output any data.
// Useful when you only want to collect statistics without persisting items.
func (s *BaseSpider) NewItemNone(ctx pkg.Context) (item pkg.Item) {
	item = items.NewItemNone()
	item = item.WithContext(ctx)
	return
}

// NewItemCsv creates a new Item that outputs data to a CSV file with the given filename.
func (s *BaseSpider) NewItemCsv(ctx pkg.Context, filename string) (item pkg.Item) {
	item = items.NewItemCsv(filename)
	item = item.WithContext(ctx)
	return
}

// NewItemJsonl creates a new Item that outputs data to a JSON Lines (JSONL) file with the given filename.
// The Context is used for managing file lifecycle and concurrency.
func (s *BaseSpider) NewItemJsonl(ctx pkg.Context, fileName string) (item pkg.Item) {
	item = items.NewItemJsonl(fileName)
	item = item.WithContext(ctx)
	return
}

// NewItemMongo creates a new Item that outputs data to a MongoDB collection.
// If update is true, existing documents with the same key will be updated.
func (s *BaseSpider) NewItemMongo(ctx pkg.Context, collection string, update bool) (item pkg.Item) {
	item = items.NewItemMongo(collection, update)
	item = item.WithContext(ctx)
	return
}

// NewItemSqlite creates a new Item that outputs data to a SQLite table.
// If update is true, existing rows with the same key will be updated.
func (s *BaseSpider) NewItemSqlite(ctx pkg.Context, table string, update bool) (item pkg.Item) {
	item = items.NewItemSqlite(table, update)
	item = item.WithContext(ctx)
	return
}

// NewItemMysql creates a new Item that outputs data to a MySQL table.
// If update is true, existing rows with the same key will be updated.
func (s *BaseSpider) NewItemMysql(ctx pkg.Context, table string, update bool) (item pkg.Item) {
	item = items.NewItemMysql(table, update)
	item = item.WithContext(ctx)
	return
}

// NewItemKafka creates a new Item that outputs data to a Kafka topic.
func (s *BaseSpider) NewItemKafka(ctx pkg.Context, topic string) (item pkg.Item) {
	item = items.NewItemKafka(topic)
	item = item.WithContext(ctx)
	return
}

func (s *BaseSpider) YieldExtra(ctx pkg.Context, extra any) (err error) {
	return ctx.GetTask().GetTask().YieldExtra(ctx, extra)
}
func (s *BaseSpider) MustYieldExtra(ctx pkg.Context, extra any) {
	if err := s.YieldExtra(ctx, extra); err != nil {
		panic(fmt.Errorf("%w: %v", pkg.ErrYieldExtraFailed, err))
	}
}
func (s *BaseSpider) UnsafeYieldExtra(ctx pkg.Context, extra any) {
	if err := s.YieldExtra(ctx, extra); err != nil {
		s.logger.Error(err)
	}
}
func (s *BaseSpider) GetExtra(ctx pkg.Context, extra any) (err error) {
	return ctx.GetTask().GetTask().GetExtra(ctx, extra)
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
	return ctx.GetTask().GetTask().YieldItem(ctx, item)
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

// Start starts the spider with the given context.
// It initializes and runs all necessary routines for crawling.
func (s *BaseSpider) Start(_ pkg.Context) (err error) {
	ctx := context.Background()

	s.context.GetSpider().WithStatus(pkg.SpiderStatusRunning)
	s.Crawler.GetSignal().SpiderChanged(s.GetContext())

	s.logger.Info("spiderName:", s.context.GetSpider().GetName())
	s.logger.Info("allowedDomains:", s.GetAllowedDomains())
	s.logger.Info("okHttpCodes:", s.OkHttpCodes())
	s.logger.Info("platforms:", s.GetPlatforms())
	s.logger.Info("browsers:", s.GetBrowsers())
	s.logger.Info("retryMaxTimes:", s.retryMaxTimes)
	s.logger.Info("redirectMaxTimes:", s.redirectMaxTimes)

	//s.logger.Info("filter", s.GetFilter())

	for _, v := range s.Pipelines() {
		s.logger.Info(v.Name(), "loaded. order:", v.Order())
	}

	for _, v := range s.GetMiddlewares().Middlewares() {
		if err = v.Start(ctx, s.spider); err != nil {
			s.logger.Error(err)
			return
		}
		s.logger.Info(v.Name(), "loaded. order:", v.Order())
	}
	return
}

// Run executes a job with the given parameters.
// Parameters:
//   - ctx: the context for controlling cancellation and timeout.
//   - jobFunc: the name of the job function to run.
//   - args: arguments to pass to the job function.
//   - mode: the execution mode of the job (e.g., immediate, scheduled).
//   - spec: scheduling specification (like a cron expression) if applicable.
//   - onlyOneTask: if true, ensures that for scheduled jobs, a new instance
//     will not start until the previous run has finished.
//
// Returns:
//   - id: a unique identifier for the job run.
//   - err: any error encountered when starting the job.
func (s *BaseSpider) Run(ctx context.Context, jobFunc string, args string, mode pkg.JobMode, spec string, onlyOneTask bool) (id string, err error) {
	if s.GetContext() == nil {
		err = errors.New("spider hasn't started")
		s.logger.Error(err)
		return
	}

	s.context.GetSpider().WithContext(ctx)

	s.jobsMutex.Lock()
	defer s.jobsMutex.Unlock()

	job := new(Job).FromSpider(s.spider)
	if err = job.start(s.context, jobFunc, args, mode, spec, onlyOneTask); err != nil {
		s.logger.Error(err)
		return
	}

	id = job.context.GetJob().GetId()
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
	if !utils.InSlice(job.GetContext().GetJob().GetStatus(), []pkg.JobStatus{
		pkg.JobStatusSuccess,
		pkg.JobStatusFailure,
	}) {
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
	if !utils.InSlice(job.context.GetJob().GetStatus(), []pkg.JobStatus{
		pkg.SpiderStatusReady,
		pkg.JobStatusRunning,
	}) {
		err = errors.New("the job can be killed in the ready or running state")
		return
	}
	err = job.kill(ctx)
	return
}
func (s *BaseSpider) JobStopped(ctx pkg.Context, err error) {
	if err != nil {
		s.logger.Info(s.spider.Name(), "job finished with an error:", err, "spend time:", ctx.GetJob().GetStopTime().Sub(ctx.GetJob().GetStartTime()), ctx.GetJob().GetId())
	} else {
		s.logger.Info(s.spider.Name(), "job finished. spend time:", ctx.GetJob().GetStopTime().Sub(ctx.GetJob().GetStartTime()), ctx.GetJob().GetId())
	}

	s.job.Out()
}
func (s *BaseSpider) Parse(_ pkg.Context, response pkg.Response) (err error) {
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.Text())
	return
}
func (s *BaseSpider) Error(_ pkg.Context, response pkg.Response, err error) {
	if response.GetResponse() == nil {
		s.logger.Error("response nil")
		return
	}
	s.logger.Info("header", response.Headers())
	s.logger.Info("body", response.Text())
	s.logger.Info("error", err)
	return
}

// Stop stops the spider with the given context.
// It gracefully shuts down all ongoing tasks and releases resources.
func (s *BaseSpider) Stop(_ pkg.Context) (err error) {
	if s.context == nil || s.context.GetSpider() == nil {
		s.logger.Warn("spider hasn't started")
		return
	}

	if s.context.GetSpider().GetStatus() == pkg.SpiderStatusStopping {
		s.logger.Debug("spider unimplemented Stop")
		return
	}

	s.context.GetSpider().WithStatus(pkg.SpiderStatusIdle)
	s.Crawler.GetSignal().SpiderChanged(s.context)
	s.logger.Debug("spider has idle")

	if !s.Crawler.StartFromCLI() {
		s.logger.Info("spider don't need to stop")
		return
	}

	s.context.GetSpider().WithStatus(pkg.SpiderStatusStopping)
	s.Crawler.GetSignal().SpiderChanged(s.context)
	s.logger.Debug("spider wait for stop")

	defer func() {
		err = s.spider.Stop(s.context)
		if errors.Is(err, pkg.DontStopErr) {
			s.logger.Error(err)
			select {}
		}

		stopTime := time.Now()
		s.context.GetSpider().WithStatus(pkg.SpiderStatusStopped)
		s.Crawler.GetSignal().SpiderChanged(s.context)

		spendTime := stopTime.Sub(s.context.GetSpider().GetStartTime())
		s.logger.Info(s.spider.Name(), s.context.GetSpider().GetId(), "spider finished. spend time:", spendTime)
		s.Crawler.SpiderStopped(s.context, err)
	}()

	for _, v := range s.GetMiddlewares().Middlewares() {
		e := v.Stop(s.context)
		if errors.Is(e, pkg.BreakErr) {
			s.logger.Debug("middlewares break", v.Name())
			break
		}
	}

	return
}

// FromCrawler initializes a Spider instance from an existing Crawler.
// It returns the Spider configured based on the Crawler's settings.
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
			WithSpider(s.spider).
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

	s.job = pkg.NewState("job")
	s.job.RegisterIsReadyAndIsZero(func() {
		_ = s.Stop(s.GetContext())
	})

	return s, nil
}
