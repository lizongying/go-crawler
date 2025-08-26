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
	ErrBacks() map[string]ErrBack
	ErrBack(name string) (errBack ErrBack)
	StartFuncs() map[string]StartFunc
	StartFunc(name string) (startFunc StartFunc)
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
	Logger() Logger
	GetConfig() Config

	// YieldItem passes an item to the spider's pipelines for processing.
	//
	// Parameters:
	//   - ctx: the current request context.
	//   - item: the item to be processed.
	//
	// Returns:
	//   - error: non-nil if processing failed.
	//
	// Behavior:
	//   - The item is passed through all configured pipelines in order.
	//   - Errors are returned to the caller for handling.
	YieldItem(Context, Item) error

	// MustYieldItem passes an item to the pipelines and panics if an error occurs.
	//
	// Parameters:
	//   - ctx: the current request context.
	//   - item: the item to be processed.
	//
	// Behavior:
	//   - Similar to YieldItem, but will panic on any error.
	//   - Useful when errors are considered fatal and should stop execution.
	MustYieldItem(Context, Item)

	// UnsafeYieldItem passes an item to the pipelines and ignores any errors.
	//
	// Parameters:
	//   - ctx: the current request context.
	//   - item: the item to be processed.
	//
	// Behavior:
	//   - Similar to YieldItem, but all errors are silently ignored.
	//   - Useful for fire-and-forget scenarios where errors can be safely ignored.
	UnsafeYieldItem(Context, Item)
	Request(Context, Request) (Response, error)

	// YieldRequest tries to submit a request to the scheduler/engine.
	// It returns an error if the request cannot be yielded.
	YieldRequest(Context, Request) error

	// MustYieldRequest is like YieldRequest but panics if any error occurs.
	// Use this when failure to yield a request should immediately stop execution.
	MustYieldRequest(Context, Request)

	// UnsafeYieldRequest submits a request but ignores any error silently.
	// Use this only if errors can be safely discarded (not recommended for critical flows).
	UnsafeYieldRequest(Context, Request)
	NewRequest(Context, ...RequestOption) error
	MustNewRequest(Context, ...RequestOption)
	YieldExtra(Context, any) error
	MustYieldExtra(Context, any)
	UnsafeYieldExtra(Context, any)
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

	// WithMiddleware registers a custom middleware with the given execution order.
	WithMiddleware(middleware Middleware, order uint8) Spider

	// WithStatsMiddleware configures the spider to use the stats middleware
	// for collecting and reporting runtime statistics.
	WithStatsMiddleware() Spider

	// WithDumpMiddleware configures the spider to use the dump middleware
	// for logging or debugging request/response data.
	WithDumpMiddleware() Spider

	// WithProxyMiddleware configures the spider to use the proxy middleware
	// for handling proxy assignment and rotation.
	WithProxyMiddleware() Spider

	// WithRobotsTxtMiddleware configures the spider to use the robots.txt middleware
	// to respect site crawling rules defined in robots.txt.
	WithRobotsTxtMiddleware() Spider

	// WithFilterMiddleware configures the spider to use the filter middleware
	// for filtering out duplicate or undesired requests.
	WithFilterMiddleware() Spider

	// WithFileMiddleware configures the spider to use the file middleware
	// for handling file download requests.
	WithFileMiddleware() Spider

	// WithImageMiddleware configures the spider to use the image middleware
	// for handling image download requests.
	WithImageMiddleware() Spider

	// WithHttpMiddleware configures the spider to use the HTTP middleware
	// for customizing or extending HTTP request/response behavior.
	WithHttpMiddleware() Spider

	// WithRetryMiddleware configures the spider to use the retry middleware
	// for retrying failed requests according to policy.
	WithRetryMiddleware() Spider

	// WithUrlMiddleware configures the spider to use the URL middleware
	// for normalizing, rewriting, or validating URLs.
	WithUrlMiddleware() Spider

	// WithReferrerMiddleware configures the spider to use the referrer middleware
	// for automatically setting or propagating referrer headers.
	WithReferrerMiddleware() Spider

	// WithCookieMiddleware configures the spider to use the cookie middleware
	// for managing cookies across requests.
	WithCookieMiddleware() Spider

	// WithRedirectMiddleware configures the spider to use the redirect middleware
	// for handling HTTP redirects.
	WithRedirectMiddleware() Spider

	// WithChromeMiddleware configures the spider to use the Chrome middleware
	// for headless browser-based crawling and rendering.
	WithChromeMiddleware() Spider

	// WithHttpAuthMiddleware configures the spider to use the HTTP authentication middleware
	// for handling basic or custom authentication.
	WithHttpAuthMiddleware() Spider

	// WithCompressMiddleware configures the spider to use the compression middleware
	// for handling compressed HTTP responses (e.g., gzip).
	WithCompressMiddleware() Spider

	// WithDecodeMiddleware configures the spider to use the decode middleware
	// for decoding response bodies (e.g., charset conversions).
	WithDecodeMiddleware() Spider

	// WithDeviceMiddleware configures the spider to use the device middleware
	// for simulating different devices or user-agents.
	WithDeviceMiddleware() Spider

	// WithRecordErrorMiddleware configures the spider to use the record-error middleware
	// for capturing and recording errors during crawling.
	WithRecordErrorMiddleware() Spider

	// WithCustomMiddleware registers a user-defined custom middleware.
	WithCustomMiddleware(middleware Middleware) Spider

	Export(Item) error

	// PipelineNames returns a map of pipeline orders to their corresponding names.
	//
	// Returns:
	//   - map[uint8]string: key is the pipeline order, value is the pipeline's name.
	//
	// Behavior:
	//   - Useful for debugging or logging the current pipeline configuration.
	//   - The order corresponds to the execution sequence of pipelines.
	PipelineNames() map[uint8]string

	// Pipelines returns a slice of all pipelines in execution order.
	//
	// Returns:
	//   - []Pipeline: all pipelines currently configured for the spider, ordered by execution sequence.
	//
	// Behavior:
	//   - Lower index pipelines are executed before higher index pipelines.
	//   - Useful for iterating or inspecting the pipelines for processing.
	Pipelines() []Pipeline

	// SetPipeline sets a pipeline at a specific order in the spider's pipeline sequence.
	//
	// Parameters:
	//   - pipeline: the pipeline instance to set.
	//   - order: the position/order of the pipeline in the sequence; lower values are executed earlier.
	//
	// Behavior:
	//   - If a pipeline already exists at the given order, it will be replaced.
	//   - Pipelines with lower order values are executed before higher ones.
	SetPipeline(pipeline Pipeline, order uint8)

	// RemovePipeline removes the pipeline at the specified index.
	//
	// Parameters:
	//   - index: the position of the pipeline in the pipeline slice to be removed.
	//
	// Behavior:
	//   - If the index is out of range, no operation is performed.
	//   - After deletion, the remaining pipelines shift to fill the gap.
	RemovePipeline(index int)

	// CleanPipelines removes all pipelines from the Spider.
	//
	// Behavior:
	//   - Clears the slice of pipelines, effectively removing all existing pipelines.
	//   - After calling this method, the spider will have no pipelines.
	CleanPipelines()

	// WithDumpPipeline configures the spider to use the Dump pipeline.
	WithDumpPipeline() Spider

	// WithFilePipeline configures the spider to handle file pipelines.
	WithFilePipeline() Spider

	// WithImagePipeline configures the spider to handle image pipelines.
	WithImagePipeline() Spider

	// WithFilterPipeline configures the spider to use a filter pipeline.
	//
	// Behavior:
	//   - Filters items according to predefined rules.
	//   - Filtered items are discarded and not sent to other output pipelines.
	WithFilterPipeline() Spider

	// WithNonePipeline configures the spider to use no output pipeline.
	//
	// Behavior:
	//   - No items will be output or persisted.
	//   - Items are still processed internally, so statistics or counters can still be collected.
	WithNonePipeline() Spider

	// WithCsvPipeline configures the spider to output CSV data.
	WithCsvPipeline() Spider

	// WithJsonLinesPipeline configures the spider to output JSON Lines data.
	WithJsonLinesPipeline() Spider

	// WithMongoPipeline configures the spider to output data to MongoDB.
	WithMongoPipeline() Spider

	// WithSqlitePipeline configures the spider to output data to SQLite.
	WithSqlitePipeline() Spider

	// WithMysqlPipeline configures the spider to output data to MySQL.
	WithMysqlPipeline() Spider

	// WithKafkaPipeline configures the spider to send data to Kafka.
	WithKafkaPipeline() Spider

	// WithCustomPipeline configures the spider with a custom pipeline.
	// Parameters:
	//   - pipeline: the custom pipeline to add to the spider.
	WithCustomPipeline(pipeline Pipeline) Spider
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
	case SpiderStatusReady:
		return "ready"
	case SpiderStatusStarting:
		return "starting"
	case SpiderStatusRunning:
		return "running"
	case SpiderStatusIdle:
		return "idle"
	case SpiderStatusStopping:
		return "stopping"
	case SpiderStatusStopped:
		return "stopped"
	default:
		return "unknown"
	}
}
