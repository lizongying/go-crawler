package pkg

import (
	"context"
	"database/sql"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
	"time"
)

type Crawler interface {
	GetMode() string
	SetMode(string)
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	IsAllowedDomain(*url.URL) bool
	GetUsername() string
	SetUsername(string)
	GetPassword() string
	SetPassword(string)
	GetSpider() Spider
	SetSpider(Spider)
	Start(context.Context) error
	Stop(context.Context) error
	RunDevServer() error
	GetDevServerHost() string
	AddDevServerRoutes(...Route)
	GetLogger() Logger
	SetLogger(Logger)
	GetPlatforms() []Platform
	SetPlatforms(...Platform)
	GetBrowsers() []Browser
	SetBrowsers(...Browser)
	GetConfig() Config
	GetKafka() *kafka.Writer
	GetKafkaReader() *kafka.Reader
	GetRedis() *redis.Client
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetS3() *s3.Client
	GetFilter() Filter
	SetFilter(Filter)
	GetRetryMaxTimes() uint8
	SetRetryMaxTimes(uint8)
	GetTimeout() time.Duration
	SetTimeout(time.Duration)
	GetOkHttpCodes() []int
	SetOkHttpCodes(...int)
	GetScheduler() Scheduler
	SetScheduler(Scheduler)
	YieldItem(context.Context, Item) error
	Request(context.Context, *Request) (*Response, error)
	YieldRequest(context.Context, *Request) error
	GetStats() Stats
	GetCallbacks() map[string]Callback
	GetErrbacks() map[string]Errback
	GetSignal() Signal
	SetSignal(Signal)
}

type CrawlOption func(Crawler)

func WithMode(mode string) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetMode(mode)
	}
}
func WithPlatforms(platforms ...Platform) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetPlatforms(platforms...)
	}
}
func WithBrowsers(browsers ...Browser) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetBrowsers(browsers...)
	}
}
func WithLogger(logger Logger) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetLogger(logger)
	}
}
func WithFilter(filter Filter) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetFilter(filter)
	}
}
func WithDownloader(downloader Downloader) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().SetDownloader(downloader)
	}
}
func WithExporter(exporter Exporter) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().SetExporter(exporter)
	}
}
func WithMiddleware(middleware Middleware, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().SetMiddleware(middleware, order)
	}
}
func WithStatsMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithStatsMiddleware()
	}
}
func WithDumpMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithDumpMiddleware()
	}
}
func WithProxyMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithProxyMiddleware()
	}
}
func WithRobotsTxtMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithRobotsTxtMiddleware()
	}
}
func WithFilterMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithFilterMiddleware()
	}
}
func WithFileMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithFileMiddleware()
	}
}
func WithImageMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithImageMiddleware()
	}
}
func WithHttpMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithHttpMiddleware()
	}
}
func WithRetryMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithRetryMiddleware()
	}
}
func WithUrlMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithUrlMiddleware()
	}
}
func WithRefererMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithRefererMiddleware()
	}
}
func WithCookieMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithCookieMiddleware()
	}
}
func WithRedirectMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithRedirectMiddleware()
	}
}
func WithChromeMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithChromeMiddleware()
	}
}
func WithHttpAuthMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithHttpAuthMiddleware()
	}
}
func WithCompressMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithCompressMiddleware()
	}
}
func WithDecodeMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithDecodeMiddleware()
	}
}
func WithDeviceMiddleware() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithDeviceMiddleware()
	}
}
func WithCustomMiddleware(middleware Middleware) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetDownloader().WithCustomMiddleware(middleware)
	}
}
func WithPipeline(pipeline Pipeline, order uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().SetPipeline(pipeline, order)
	}
}
func WithDumpPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithDumpPipeline()
	}
}
func WithFilePipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithFilePipeline()
	}
}
func WithImagePipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithImagePipeline()
	}
}
func WithFilterPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithFilterPipeline()
	}
}
func WithCsvPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithCsvPipeline()
	}
}
func WithJsonLinesPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithJsonLinesPipeline()
	}
}
func WithMongoPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithMongoPipeline()
	}
}
func WithMysqlPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithMysqlPipeline()
	}
}
func WithKafkaPipeline() CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithKafkaPipeline()
	}
}
func WithCustomPipeline(pipeline Pipeline) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().GetExporter().WithCustomPipeline(pipeline)
	}
}
func WithRetryMaxTimes(retryMaxTimes uint8) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetRetryMaxTimes(retryMaxTimes)
	}
}
func WithTimeout(timeout time.Duration) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetTimeout(timeout)
	}
}
func WithInterval(timeout time.Duration) CrawlOption {
	return func(crawler Crawler) {
		crawler.GetScheduler().SetInterval(timeout)
	}
}
func WithOkHttpCodes(httpCodes ...int) CrawlOption {
	return func(crawler Crawler) {
		crawler.SetOkHttpCodes(httpCodes...)
	}
}
