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

	// SetName sets the spider's name.
	SetName(name string) Spider
	GetHost() string

	// SetHost sets the target host.
	SetHost(host string) Spider
	Username() string

	// SetUsername sets the authentication username.
	SetUsername(username string) Spider
	Password() string

	// SetPassword sets the authentication password.
	SetPassword(password string) Spider
	GetPlatforms() []Platform

	// SetPlatforms specifies the target platforms.
	SetPlatforms(platforms ...Platform) Spider
	GetBrowsers() []Browser

	// SetBrowsers specifies the supported browsers.
	SetBrowsers(browsers ...Browser) Spider

	GetCrawler() Crawler

	GetSpider() Spider

	SetSpider(spider Spider) Spider

	// CallBackNames returns the list of all registered callback names.
	CallBackNames() []string

	// SetCallBack registers a new callback function under the given name.
	SetCallBack(name string, callBack CallBack)

	// CallBack retrieves the callback function by name.
	// Returns an error if the callback does not exist.
	CallBack(name string) (callback CallBack, err error)

	// SetErrBack registers a new error handler under the given name.
	SetErrBack(name string, errBack ErrBack)

	// ErrBackNames returns the list of all registered error handler names.
	ErrBackNames() []string

	// ErrBack retrieves the error handler by name.
	// Returns an error if the error handler does not exist.
	ErrBack(name string) (errBack ErrBack, err error)

	// StartFuncNames returns the list of all registered start function names.
	StartFuncNames() []string

	// SetStartFunc registers a new start function under the given name.
	SetStartFunc(name string, startFunc StartFunc)

	// StartFunc looks up a registered StartFunc by its name.
	// It returns the StartFunc and a nil error if found.
	// If no StartFunc exists for the given name, it returns
	// ErrStartFuncNotExist.
	StartFunc(name string) (startFunc StartFunc, err error)

	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	IsAllowedDomain(*url.URL) bool

	// RetryMaxTimes returns the maximum number of retry attempts configured for the spider.
	RetryMaxTimes() uint8

	// SetRetryMaxTimes sets the maximum number of retry attempts for the spider.
	// Returns the spider itself for chaining.
	SetRetryMaxTimes(retryMaxTimes uint8) Spider

	// RedirectMaxTimes returns the maximum number of HTTP redirects the spider will follow automatically.
	RedirectMaxTimes() uint8

	// SetRedirectMaxTimes sets the maximum number of HTTP redirects the spider
	// will follow automatically. Returns the spider itself for chaining.
	SetRedirectMaxTimes(redirectMaxTimes uint8) Spider

	// Timeout returns the timeout duration configured for HTTP requests made by the spider.
	Timeout() time.Duration

	// SetTimeout sets the timeout for HTTP requests made by the spider.
	// Returns the spider itself for chaining.
	SetTimeout(timeout time.Duration) Spider

	// Interval returns the duration between consecutive requests made by the spider.
	Interval() time.Duration

	// SetInterval sets the duration between consecutive requests for the spider and returns the spider for chaining.
	SetInterval(interval time.Duration) Spider

	// Concurrency returns the maximum number of concurrent requests
	// that the spider is allowed to execute.
	Concurrency() uint8

	// SetConcurrency sets the maximum number of concurrent requests
	// the spider is allowed to execute and returns the Spider instance.
	// This controls how many requests can be processed in parallel.
	SetConcurrency(concurrency uint8) Spider

	// OkHttpCodes returns the list of HTTP status codes considered successful by the spider.
	OkHttpCodes() []int

	// SetOkHttpCodes sets the HTTP status codes that are considered successful.
	// Requests returning other codes may trigger retries or errors.
	// Returns the spider itself for chaining.
	SetOkHttpCodes(okHttpCodes ...int) Spider

	// SetRequestRate sets a request rate limiter for the given slot.
	//
	// Parameters:
	//   - slot: an identifier for grouping requests. If empty, "*" (default) is used.
	//   - interval: the total time window during which 'concurrency' number of requests are allowed.
	//   - concurrency: the maximum number of requests allowed within the given interval.
	//
	// Behavior:
	//   - If the slot has no existing limiter, a new one is created.
	//   - If a limiter already exists for the slot, its burst (concurrency) and rate limit
	//     are updated to match the new configuration.
	//   - If concurrency < 1, it will be normalized to 1.
	//
	// Example:
	//   SetRequestRate("api", time.Second, 5)
	//   → allows up to 5 requests per second for the "api" slot.
	SetRequestRate(slot string, interval time.Duration, concurrency int) Limiter

	SetRatePerHour(slot string, ratePerHour int, concurrency int)

	// GetFilter returns the filter function currently set for processing or filtering items.
	GetFilter() Filter

	// SetFilter sets a filter function to process or filter items before they are exported.
	// Returns the spider itself for chaining.
	SetFilter(filter Filter) Spider

	// Run executes a job with the given parameters.
	// Parameters:
	//   - ctx: the context for controlling cancellation and timeout.
	//   - jobFunc: the name of the job function to run.
	//   - args: arguments to pass to the job function.
	//   - mode: the execution mode of the job (e.g., immediate, scheduled).
	//   - spec: scheduling specification (like a cron expression) if applicable.
	//   - onlyOneTask: if true, ensures that for scheduled jobs, a new instance
	//     will not start until the previous run has finished.
	// Returns:
	//   - id: a unique identifier for the job run.
	//   - err: any error encountered when starting the job.
	Run(ctx context.Context, jobFunc string, args string, mode JobMode, spec string, onlyOneTask bool) (id string, err error)

	// Start starts the spider with the given context.
	// It initializes and runs all necessary routines for crawling.
	Start(ctx Context) error

	// Stop stops the spider with the given context.
	// It gracefully shuts down all ongoing tasks and releases resources.
	Stop(ctx Context) error

	// FromCrawler initializes a Spider instance from an existing Crawler.
	// It returns the Spider configured based on the Crawler's settings.
	FromCrawler(Crawler) Spider

	// GetLogger returns the logger associated with the spider.
	// This provides access to logging methods for recording information, warnings, and errors.
	GetLogger() Logger

	// Logger returns the logger associated with the spider.
	// This provides access to logging methods for recording information, warnings, and errors.
	Logger() Logger

	// GetConfig returns the current configuration of the Spider.
	// The returned Config contains all settings such as timeouts, headers, and pipelines.
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

	// NewItemNone creates a new Item that does not output any data.
	// Useful when you only want to collect statistics without persisting items.
	NewItemNone(ctx Context) (item Item)

	// NewItemCsv creates a new Item that outputs data to a CSV file with the given filename.
	NewItemCsv(ctx Context, filename string) (item Item)

	// NewItemJsonl creates a new Item that outputs data to a JSON Lines (JSONL) file with the given filename.
	// The Context is used for managing file lifecycle and concurrency.
	NewItemJsonl(ctx Context, filename string) (item Item)

	// NewItemMongo creates a new Item that outputs data to a MongoDB collection.
	// If update is true, existing documents with the same key will be updated.
	NewItemMongo(ctx Context, collection string, update bool) (item Item)

	// NewItemSqlite creates a new Item that outputs data to a SQLite table.
	// If update is true, existing rows with the same key will be updated.
	NewItemSqlite(ctx Context, table string, update bool) (item Item)

	// NewItemMysql creates a new Item that outputs data to a MySQL table.
	// If update is true, existing rows with the same key will be updated.
	NewItemMysql(ctx Context, table string, update bool) (item Item)

	// NewItemKafka creates a new Item that outputs data to a Kafka topic.
	NewItemKafka(ctx Context, topic string) (item Item)

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

	// NewRequest creates a new Request associated with the given Context.
	// Optional RequestOption functions can be used to configure the request, such as setting the URL, callback, or extra data.
	NewRequest(Context, ...RequestOption) Request

	YieldExtra(Context, any) error

	MustYieldExtra(Context, any)

	UnsafeYieldExtra(Context, any)

	GetExtra(Context, any) error

	MustGetExtra(Context, any)

	UnsafeGetExtra(Context, any)

	AddMockServerRoutes(...Route) Crawler

	Options() []SpiderOption

	WithOptions(options ...SpiderOption) Spider

	// Limiter retrieves the rate limiter associated with the given slot.
	// It returns the limiter and a boolean indicating whether the limiter exists.
	Limiter(slot string) (value Limiter, ok bool)

	// RerunJob restarts an existing job identified by jobId.
	// It looks up the job by its unique ID and attempts to run it again.
	// Returns an error if the job does not exist or cannot be rerun.
	RerunJob(ctx context.Context, jobId string) (err error)

	// KillJob forcefully terminates a running job identified by jobId.
	// It attempts to interrupt and stop the job execution.
	// Returns an error if the job does not exist or has already finished.
	KillJob(ctx context.Context, jobId string) (err error)

	JobStopped(ctx Context, err error)

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
		spider.SetInterval(timeout)
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
func WithRatePerHour(slot string, ratePerHour int) SpiderOption {
	return func(spider Spider) {
		spider.SetRatePerHour(slot, ratePerHour, 1)
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
