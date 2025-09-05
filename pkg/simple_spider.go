package pkg

import "time"

// SimpleSpider defines a simplified spider interface
// with chainable configuration methods.
type SimpleSpider interface {

	// SetName sets the spider's name.
	SetName(name string) SimpleSpider

	// SetHost sets the target host.
	SetHost(host string) SimpleSpider

	// SetUsername sets the authentication username.
	SetUsername(username string) SimpleSpider

	// SetPassword sets the authentication password.
	SetPassword(password string) SimpleSpider

	// SetPlatforms specifies the target platforms.
	SetPlatforms(platforms ...Platform) SimpleSpider

	// SetBrowsers specifies the supported browsers.
	SetBrowsers(browsers ...Browser) SimpleSpider

	// Logger returns the logger associated with the spider.
	// This provides access to logging methods for recording information, warnings, and errors.
	Logger() Logger

	// SetRetryMaxTimes sets the maximum number of retry attempts for the spider.
	// Returns the spider itself for chaining.
	SetRetryMaxTimes(retryMaxTimes uint8) SimpleSpider

	// SetRedirectMaxTimes sets the maximum number of HTTP redirects the spider
	// will follow automatically. Returns the spider itself for chaining.
	SetRedirectMaxTimes(redirectMaxTimes uint8) SimpleSpider

	// SetTimeout sets the timeout for HTTP requests made by the spider.
	// Returns the spider itself for chaining.
	SetTimeout(timeout time.Duration) SimpleSpider

	// SetOkHttpCodes sets the HTTP status codes that are considered successful.
	// Requests returning other codes may trigger retries or errors.
	// Returns the spider itself for chaining.
	SetOkHttpCodes(okHttpCodes ...int) SimpleSpider

	// SetRatePerHour configures the request rate limit and concurrency for a given slot.
	//
	// Parameters:
	//   - slot:        Identifier to distinguish different request queues or task groups
	//                  (e.g., a domain name or resource category).
	//   - ratePerHour: Maximum allowed requests per hour (rate limit), internally converted
	//                  to an average rate per second.
	//   - concurrency: Maximum number of concurrent requests allowed to run in parallel.
	//
	// Returns:
	//   - The current SimpleSpider instance, enabling method chaining.
	//
	// Example:
	//   spider.SetRatePerHour("example.com", 3600, 1) // 3600 requests/hour (~1 per second), concurrency = 1
	SetRatePerHour(slot string, ratePerHour int, concurrency int) SimpleSpider

	// SetFilter sets a filter function to process or filter items before they are exported.
	// Returns the spider itself for chaining.
	SetFilter(filter Filter) SimpleSpider

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

	// NewRequest creates a new Request associated with the given Context.
	// Optional RequestOption functions can be used to configure the request, such as setting the URL, callback, or extra data.
	NewRequest(Context, ...RequestOption) Request

	// WithMiddleware registers a custom middleware with the given execution order.
	WithMiddleware(middleware Middleware, order uint8) SimpleSpider

	// WithStatsMiddleware configures the spider to use the stats middleware
	// for collecting and reporting runtime statistics.
	WithStatsMiddleware() SimpleSpider

	// WithDumpMiddleware configures the spider to use the dump middleware
	// for logging or debugging request/response data.
	WithDumpMiddleware() SimpleSpider

	// WithProxyMiddleware configures the spider to use the proxy middleware
	// for handling proxy assignment and rotation.
	WithProxyMiddleware() SimpleSpider

	// WithRobotsTxtMiddleware configures the spider to use the robots.txt middleware
	// to respect site crawling rules defined in robots.txt.
	WithRobotsTxtMiddleware() SimpleSpider

	// WithFilterMiddleware configures the spider to use the filter middleware
	// for filtering out duplicate or undesired requests.
	WithFilterMiddleware() SimpleSpider

	// WithFileMiddleware configures the spider to use the file middleware
	// for handling file download requests.
	WithFileMiddleware() SimpleSpider

	// WithImageMiddleware configures the spider to use the image middleware
	// for handling image download requests.
	WithImageMiddleware() SimpleSpider

	// WithHttpMiddleware configures the spider to use the HTTP middleware
	// for customizing or extending HTTP request/response behavior.
	WithHttpMiddleware() SimpleSpider

	// WithRetryMiddleware configures the spider to use the retry middleware
	// for retrying failed requests according to policy.
	WithRetryMiddleware() SimpleSpider

	// WithUrlMiddleware configures the spider to use the URL middleware
	// for normalizing, rewriting, or validating URLs.
	WithUrlMiddleware() SimpleSpider

	// WithReferrerMiddleware configures the spider to use the referrer middleware
	// for automatically setting or propagating referrer headers.
	WithReferrerMiddleware() SimpleSpider

	// WithCookieMiddleware configures the spider to use the cookie middleware
	// for managing cookies across requests.
	WithCookieMiddleware() SimpleSpider

	// WithRedirectMiddleware configures the spider to use the redirect middleware
	// for handling HTTP redirects.
	WithRedirectMiddleware() SimpleSpider

	// WithChromeMiddleware configures the spider to use the Chrome middleware
	// for headless browser-based crawling and rendering.
	WithChromeMiddleware() SimpleSpider

	// WithHttpAuthMiddleware configures the spider to use the HTTP authentication middleware
	// for handling basic or custom authentication.
	WithHttpAuthMiddleware() SimpleSpider

	// WithCompressMiddleware configures the spider to use the compression middleware
	// for handling compressed HTTP responses (e.g., gzip).
	WithCompressMiddleware() SimpleSpider

	// WithDecodeMiddleware configures the spider to use the decode middleware
	// for decoding response bodies (e.g., charset conversions).
	WithDecodeMiddleware() SimpleSpider

	// WithDeviceMiddleware configures the spider to use the device middleware
	// for simulating different devices or user-agents.
	WithDeviceMiddleware() SimpleSpider

	// WithRecordErrorMiddleware configures the spider to use the record-error middleware
	// for capturing and recording errors during crawling.
	WithRecordErrorMiddleware() SimpleSpider

	// WithCustomMiddleware registers a user-defined custom middleware.
	WithCustomMiddleware(middleware Middleware) SimpleSpider

	// WithDumpPipeline configures the spider to use the Dump pipeline.
	WithDumpPipeline() SimpleSpider

	// WithFilePipeline configures the spider to handle file pipelines.
	WithFilePipeline() SimpleSpider

	// WithImagePipeline configures the spider to handle image pipelines.
	WithImagePipeline() SimpleSpider

	// WithFilterPipeline configures the spider to use a filter pipeline.
	//
	// Behavior:
	//   - Filters items according to predefined rules.
	//   - Filtered items are discarded and not sent to other output pipelines.
	WithFilterPipeline() SimpleSpider

	// WithNonePipeline configures the spider to use no output pipeline.
	//
	// Behavior:
	//   - No items will be output or persisted.
	//   - Items are still processed internally, so statistics or counters can still be collected.
	WithNonePipeline() SimpleSpider

	// WithCsvPipeline configures the spider to output CSV data.
	WithCsvPipeline() SimpleSpider

	// WithJsonLinesPipeline configures the spider to output JSON Lines data.
	WithJsonLinesPipeline() SimpleSpider

	// WithMongoPipeline configures the spider to output data to MongoDB.
	WithMongoPipeline() SimpleSpider

	// WithSqlitePipeline configures the spider to output data to SQLite.
	WithSqlitePipeline() SimpleSpider

	// WithMysqlPipeline configures the spider to output data to MySQL.
	WithMysqlPipeline() SimpleSpider

	// WithKafkaPipeline configures the spider to send data to Kafka.
	WithKafkaPipeline() SimpleSpider

	// WithCustomPipeline configures the spider with a custom pipeline.
	// Parameters:
	//   - pipeline: the custom pipeline to add to the spider.
	WithCustomPipeline(pipeline Pipeline) SimpleSpider
}

type NewSimpleSpider func(SimpleSpider) (SimpleSpider, error)
