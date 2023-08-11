package pkg

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

type Spider interface {
	GetName() string
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
	GetCallBacks() map[string]CallBack
	SetCallBacks(map[string]CallBack) Spider
	GetErrBacks() map[string]ErrBack
	SetErrBacks(map[string]ErrBack) Spider
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	IsAllowedDomain(*url.URL) bool
	RetryMaxTimes() uint8
	SetRetryMaxTimes(uint8) Spider
	Timeout() time.Duration
	SetTimeout(time.Duration) Spider
	OkHttpCodes() []int
	SetOkHttpCodes(...int) Spider
	GetStats() Stats
	SetStats(Stats) Spider
	GetFilter() Filter
	SetFilter(Filter) Spider
	GetScheduler() Scheduler
	SetScheduler(Scheduler) Spider
	Start(ctx context.Context, startFunc string, args string) error
	Stop(ctx context.Context) error
	FromCrawler(Crawler) Spider

	GetLogger() Logger
	GetConfig() Config
	YieldItem(Context, Item) error
	Request(Context, Request) (Response, error)
	YieldRequest(Context, Request) error
	YieldExtra(Context, any) error
	SetRequestRate(slot string, interval time.Duration, concurrency int)
	AddDevServerRoutes(...Route)
	GetMode() string

	SetLogger(Logger) Spider

	Stats
	//Signal

	GetCrawler() Crawler

	GetSignal() Signal
	SetSignal(Signal)

	Options() []SpiderOption
	WithOptions(options ...SpiderOption) Spider
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
func WithLogger(logger Logger) SpiderOption {
	return func(spider Spider) {
		spider.SetLogger(logger)
	}
}
func WithFilter(filter Filter) SpiderOption {
	return func(spider Spider) {
		spider.SetFilter(filter)
	}
}
func WithDownloader(downloader Downloader) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().SetDownloader(downloader)
	}
}
func WithExporter(exporter Exporter) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().SetExporter(exporter)
	}
}
func WithMiddleware(middleware Middleware, order uint8) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().SetMiddleware(middleware, order)
	}
}
func WithStatsMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithStatsMiddleware()
	}
}
func WithDumpMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithDumpMiddleware()
	}
}
func WithProxyMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithProxyMiddleware()
	}
}
func WithRobotsTxtMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithRobotsTxtMiddleware()
	}
}
func WithFilterMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithFilterMiddleware()
	}
}
func WithFileMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithFileMiddleware()
	}
}
func WithImageMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithImageMiddleware()
	}
}
func WithHttpMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithHttpMiddleware()
	}
}
func WithRetryMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithRetryMiddleware()
	}
}
func WithUrlMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithUrlMiddleware()
	}
}
func WithReferrerMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithReferrerMiddleware()
	}
}
func WithCookieMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithCookieMiddleware()
	}
}
func WithRedirectMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithRedirectMiddleware()
	}
}
func WithChromeMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithChromeMiddleware()
	}
}
func WithHttpAuthMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithHttpAuthMiddleware()
	}
}
func WithCompressMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithCompressMiddleware()
	}
}
func WithDecodeMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithDecodeMiddleware()
	}
}
func WithDeviceMiddleware() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithDeviceMiddleware()
	}
}
func WithCustomMiddleware(middleware Middleware) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetDownloader().WithCustomMiddleware(middleware)
	}
}
func WithPipeline(pipeline Pipeline, order uint8) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().SetPipeline(pipeline, order)
	}
}
func WithDumpPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithDumpPipeline()
	}
}
func WithFilePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithFilePipeline()
	}
}
func WithImagePipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithImagePipeline()
	}
}
func WithFilterPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithFilterPipeline()
	}
}
func WithCsvPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithCsvPipeline()
	}
}
func WithJsonLinesPipeline() SpiderOption {
	return func(spider Spider) {
		fmt.Println(11111111, spider.GetScheduler())
		spider.GetScheduler().GetExporter().WithJsonLinesPipeline()
	}
}
func WithMongoPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithMongoPipeline()
	}
}
func WithMysqlPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithMysqlPipeline()
	}
}
func WithKafkaPipeline() SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithKafkaPipeline()
	}
}
func WithCustomPipeline(pipeline Pipeline) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().GetExporter().WithCustomPipeline(pipeline)
	}
}
func WithRetryMaxTimes(retryMaxTimes uint8) SpiderOption {
	return func(spider Spider) {
		spider.SetRetryMaxTimes(retryMaxTimes)
	}
}
func WithTimeout(timeout time.Duration) SpiderOption {
	return func(spider Spider) {
		spider.SetTimeout(timeout)
	}
}
func WithInterval(timeout time.Duration) SpiderOption {
	return func(spider Spider) {
		spider.GetScheduler().SetInterval(timeout)
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
func WithStats(stats Stats) SpiderOption {
	return func(spider Spider) {
		spider.SetStats(stats)
	}
}
