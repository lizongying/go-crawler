package pkg

import (
	"context"
	"net/url"
	"time"
)

type Spider interface {
	GetContext() Context
	WithContext(ctx Context) Spider
	Name() string
	SetName(string) Spider
	GetHost() string
	SetHost(string) Spider
	Username() string
	SetUsername(string) Spider
	Password() string
	SetPassword(string) Spider
	GetPlatforms() []Platform
	SetPlatforms(...Platform) Spider
	GetBrowsers() []Browser
	SetBrowsers(...Browser) Spider
	GetSpider() Spider
	SetSpider(spider Spider) Spider
	CallBacks() map[string]CallBack
	CallBack(name string) (callback CallBack)
	SetCallBacks(map[string]CallBack) Spider
	ErrBacks() map[string]ErrBack
	ErrBack(name string) (errBack ErrBack)
	SetErrBacks(map[string]ErrBack) Spider
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	IsAllowedDomain(*url.URL) bool
	RetryMaxTimes() uint8
	SetRetryMaxTimes(uint8) Spider
	RedirectMaxTimes() uint8
	SetRedirectMaxTimes(uint8) Spider
	Timeout() time.Duration
	SetTimeout(time.Duration) Spider
	OkHttpCodes() []int
	SetOkHttpCodes(...int) Spider
	GetFilter() Filter
	SetFilter(Filter) Spider

	Run(context.Context, string, string, JobMode, string, bool) (string, error)
	Start(Context) error
	Stop(ctx Context) error
	FromCrawler(Crawler) Spider

	GetLogger() Logger
	GetConfig() Config
	YieldItem(Context, Item) error
	MustYieldItem(Context, Item)
	Request(Context, Request) (Response, error)
	YieldRequest(Context, Request) error
	NewRequest(Context, ...RequestOption) error
	MustYieldRequest(Context, Request)
	MustNewRequest(Context, ...RequestOption)
	YieldExtra(Context, any) error
	MustYieldExtra(Context, any)
	GetExtra(Context, any) error
	MustGetExtra(Context, any)
	SetRequestRate(slot string, interval time.Duration, concurrency int)
	AddMockServerRoutes(...Route)

	GetCrawler() Crawler

	Options() []SpiderOption
	WithOptions(options ...SpiderOption) Spider

	Interval() time.Duration
	WithInterval(time.Duration) Spider

	RequestSlotLoad(slot string) (value any, ok bool)
	RequestSlotStore(slot string, value any)

	RerunJob(ctx context.Context, jobId string) (err error)
	KillJob(ctx context.Context, jobId string) (err error)
	JobStopped(Context, error)

	GetDownloader() Downloader
	WithDownloader(Downloader) Spider
	GetExporter() Exporter
	WithExporter(Exporter) Spider

	Downloader
	Exporter
}

type NewSpider func(Spider) (Spider, error)

type SpiderOption func(Spider)

func WithName(name string) SpiderOption {
	return func(spider Spider) {
		spider.SetName(name)
	}
}
func WithHost(host string) SpiderOption {
	return func(spider Spider) {
		spider.SetHost(host)
	}
}
func WithUsername(username string) SpiderOption {
	return func(spider Spider) {
		spider.SetUsername(username)
	}
}
func WithPassword(password string) SpiderOption {
	return func(spider Spider) {
		spider.SetPassword(password)
	}
}
func WithPlatforms(platforms ...Platform) SpiderOption {
	return func(spider Spider) {
		spider.SetPlatforms(platforms...)
	}
}
func WithBrowsers(browsers ...Browser) SpiderOption {
	return func(spider Spider) {
		spider.SetBrowsers(browsers...)
	}
}
func WithFilter(filter Filter) SpiderOption {
	return func(spider Spider) {
		spider.SetFilter(filter)
	}
}
func WithDownloader(downloader Downloader) SpiderOption {
	return func(spider Spider) {
		spider.WithDownloader(downloader)
	}
}
func WithExporter(exporter Exporter) SpiderOption {
	return func(spider Spider) {
		spider.WithExporter(exporter)
	}
}
func WithMiddleware(middleware Middleware, order uint8) SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().SetMiddleware(middleware, order)
	}
}
func WithStatsMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithStatsMiddleware()
	}
}
func WithDumpMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithDumpMiddleware()
	}
}
func WithProxyMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithProxyMiddleware()
	}
}
func WithRobotsTxtMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithRobotsTxtMiddleware()
	}
}
func WithFilterMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithFilterMiddleware()
	}
}
func WithFileMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithFileMiddleware()
	}
}
func WithImageMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithImageMiddleware()
	}
}
func WithHttpMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithHttpMiddleware()
	}
}
func WithRetryMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithRetryMiddleware()
	}
}
func WithUrlMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithUrlMiddleware()
	}
}
func WithReferrerMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithReferrerMiddleware()
	}
}
func WithCookieMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithCookieMiddleware()
	}
}
func WithRedirectMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithRedirectMiddleware()
	}
}
func WithChromeMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithChromeMiddleware()
	}
}
func WithHttpAuthMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithHttpAuthMiddleware()
	}
}
func WithCompressMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithCompressMiddleware()
	}
}
func WithDecodeMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithDecodeMiddleware()
	}
}
func WithDeviceMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithDeviceMiddleware()
	}
}
func WithRecordErrorMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithRecordErrorMiddleware()
	}
}
func WithCustomMiddleware(middleware Middleware) SpiderOption {
	return func(spider Spider) {
		spider.GetMiddlewares().WithCustomMiddleware(middleware)
	}
}
func WithPipeline(pipeline Pipeline, order uint8) SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().SetPipeline(pipeline, order)
	}
}
func WithDumpPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithDumpPipeline()
	}
}
func WithFilePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithFilePipeline()
	}
}
func WithImagePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithImagePipeline()
	}
}
func WithFilterPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithFilterPipeline()
	}
}
func WithNonePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithNonePipeline()
	}
}
func WithCsvPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithCsvPipeline()
	}
}
func WithJsonLinesPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithJsonLinesPipeline()
	}
}
func WithMongoPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithMongoPipeline()
	}
}
func WithSqlitePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithSqlitePipeline()
	}
}
func WithMysqlPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithMysqlPipeline()
	}
}
func WithKafkaPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithKafkaPipeline()
	}
}
func WithCustomPipeline(pipeline Pipeline) SpiderOption {
	return func(spider Spider) {
		spider.GetExporter().WithCustomPipeline(pipeline)
	}
}
func WithRetryMaxTimes(retryMaxTimes uint8) SpiderOption {
	return func(spider Spider) {
		spider.SetRetryMaxTimes(retryMaxTimes)
	}
}
func WithRedirectMaxTimes(redirectMaxTimes uint8) SpiderOption {
	return func(spider Spider) {
		spider.SetRedirectMaxTimes(redirectMaxTimes)
	}
}
func WithTimeout(timeout time.Duration) SpiderOption {
	return func(spider Spider) {
		spider.SetTimeout(timeout)
	}
}
func WithInterval(timeout time.Duration) SpiderOption {
	return func(spider Spider) {
		spider.WithInterval(timeout)
	}
}
func WithOkHttpCodes(httpCodes ...int) SpiderOption {
	return func(spider Spider) {
		spider.SetOkHttpCodes(httpCodes...)
	}
}
func WithRequestRate(slot string, interval time.Duration, concurrency int) SpiderOption {
	return func(spider Spider) {
		spider.SetRequestRate(slot, interval, concurrency)
	}
}

type SpiderStatus uint8

const (
	SpiderStatusUnknown = iota
	SpiderStatusReady
	SpiderStatusStarting
	SpiderStatusRunning
	SpiderStatusIdle
	SpiderStatusStopping
	SpiderStatusStopped
)

func (s SpiderStatus) String() string {
	switch s {
	case 1:
		return "ready"
	case 2:
		return "starting"
	case 3:
		return "running"
	case 4:
		return "idle"
	case 5:
		return "stopping"
	case 6:
		return "stopped"
	default:
		return "unknown"
	}
}
