package spider

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type SimpleSpider struct {
	*BaseSpider
}

func (s *SimpleSpider) WithDumpPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithDumpPipeline()
	return s
}
func (s *SimpleSpider) WithFilePipeline() pkg.SimpleSpider {
	s.BaseSpider.WithFilePipeline()
	return s
}
func (s *SimpleSpider) WithImagePipeline() pkg.SimpleSpider {
	s.BaseSpider.WithImagePipeline()
	return s
}
func (s *SimpleSpider) WithFilterPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithFilterPipeline()
	return s
}
func (s *SimpleSpider) WithNonePipeline() pkg.SimpleSpider {
	s.BaseSpider.WithNonePipeline()
	return s
}
func (s *SimpleSpider) WithCsvPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithCsvPipeline()
	return s
}
func (s *SimpleSpider) WithJsonLinesPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithJsonLinesPipeline()
	return s
}
func (s *SimpleSpider) WithMongoPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithMongoPipeline()
	return s
}
func (s *SimpleSpider) WithSqlitePipeline() pkg.SimpleSpider {
	s.BaseSpider.WithSqlitePipeline()
	return s
}
func (s *SimpleSpider) WithMysqlPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithMysqlPipeline()
	return s
}
func (s *SimpleSpider) WithKafkaPipeline() pkg.SimpleSpider {
	s.BaseSpider.WithKafkaPipeline()
	return s
}
func (s *SimpleSpider) WithCustomPipeline(pipeline pkg.Pipeline) pkg.SimpleSpider {
	s.BaseSpider.WithCustomPipeline(pipeline)
	return s
}
func (s *SimpleSpider) SetName(name string) pkg.SimpleSpider {
	s.BaseSpider.SetName(name)
	return s
}
func (s *SimpleSpider) SetHost(host string) pkg.SimpleSpider {
	s.BaseSpider.SetHost(host)
	return s
}
func (s *SimpleSpider) SetUsername(username string) pkg.SimpleSpider {
	s.BaseSpider.SetUsername(username)
	return s
}
func (s *SimpleSpider) SetPassword(password string) pkg.SimpleSpider {
	s.BaseSpider.SetPassword(password)
	return s
}
func (s *SimpleSpider) SetPlatforms(platforms ...pkg.Platform) pkg.SimpleSpider {
	s.BaseSpider.SetPlatforms(platforms...)
	return s
}
func (s *SimpleSpider) SetBrowsers(browsers ...pkg.Browser) pkg.SimpleSpider {
	s.BaseSpider.SetBrowsers(browsers...)
	return s
}
func (s *SimpleSpider) WithMiddleware(middleware pkg.Middleware, order uint8) pkg.SimpleSpider {
	s.BaseSpider.WithMiddleware(middleware, order)
	return s
}
func (s *SimpleSpider) WithStatsMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithStatsMiddleware()
	return s
}
func (s *SimpleSpider) WithDumpMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithDumpMiddleware()
	return s
}
func (s *SimpleSpider) WithProxyMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithProxyMiddleware()
	return s
}
func (s *SimpleSpider) WithRobotsTxtMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithRobotsTxtMiddleware()
	return s
}
func (s *SimpleSpider) WithFilterMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithFilterMiddleware()
	return s
}
func (s *SimpleSpider) WithFileMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithFileMiddleware()
	return s
}
func (s *SimpleSpider) WithImageMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithImageMiddleware()
	return s
}
func (s *SimpleSpider) WithHttpMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithHttpMiddleware()
	return s
}
func (s *SimpleSpider) WithRetryMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithRetryMiddleware()
	return s
}
func (s *SimpleSpider) WithUrlMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithUrlMiddleware()
	return s
}
func (s *SimpleSpider) WithReferrerMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithReferrerMiddleware()
	return s
}
func (s *SimpleSpider) WithCookieMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithCookieMiddleware()
	return s
}
func (s *SimpleSpider) WithRedirectMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithRedirectMiddleware()
	return s
}
func (s *SimpleSpider) WithChromeMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithChromeMiddleware()
	return s
}
func (s *SimpleSpider) WithHttpAuthMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithHttpAuthMiddleware()
	return s
}
func (s *SimpleSpider) WithCompressMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithCompressMiddleware()
	return s
}
func (s *SimpleSpider) WithDecodeMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithDecodeMiddleware()
	return s
}
func (s *SimpleSpider) WithDeviceMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithDeviceMiddleware()
	return s
}
func (s *SimpleSpider) WithRecordErrorMiddleware() pkg.SimpleSpider {
	s.BaseSpider.WithRecordErrorMiddleware()
	return s
}
func (s *SimpleSpider) WithCustomMiddleware(middleware pkg.Middleware) pkg.SimpleSpider {
	s.BaseSpider.WithCustomMiddleware(middleware)
	return s
}

func (s *SimpleSpider) Logger() pkg.Logger {
	return s.BaseSpider.Logger()
}

// SetRetryMaxTimes sets the maximum number of retry attempts for the spider.
// Returns the spider itself for chaining.
func (s *SimpleSpider) SetRetryMaxTimes(retryMaxTimes uint8) pkg.SimpleSpider {
	s.BaseSpider.SetRetryMaxTimes(retryMaxTimes)
	return s
}

// SetRedirectMaxTimes sets the maximum number of HTTP redirects the spider
// will follow automatically. Returns the spider itself for chaining.
func (s *SimpleSpider) SetRedirectMaxTimes(redirectMaxTimes uint8) pkg.SimpleSpider {
	s.BaseSpider.SetRedirectMaxTimes(redirectMaxTimes)
	return s
}

// SetTimeout sets the timeout for HTTP requests made by the spider.
// Returns the spider itself for chaining.
func (s *SimpleSpider) SetTimeout(timeout time.Duration) pkg.SimpleSpider {
	s.BaseSpider.SetTimeout(timeout)
	return s
}

// SetInterval sets the duration between consecutive requests for the spider and returns the spider for chaining.
func (s *SimpleSpider) SetInterval(interval time.Duration) pkg.SimpleSpider {
	s.BaseSpider.SetInterval(interval)
	return s
}

// SetOkHttpCodes sets the HTTP status codes that are considered successful.
// Requests returning other codes may trigger retries or errors.
// Returns the spider itself for chaining.
func (s *SimpleSpider) SetOkHttpCodes(httpCodes ...int) pkg.SimpleSpider {
	s.BaseSpider.SetOkHttpCodes(httpCodes...)
	return s
}

// SetFilter sets a filter function to process or filter items before they are exported.
// Returns the spider itself for chaining.
func (s *SimpleSpider) SetFilter(filter pkg.Filter) pkg.SimpleSpider {
	s.BaseSpider.SetFilter(filter)
	return s
}

func (s *SimpleSpider) NewRequest(ctx pkg.Context, options ...pkg.RequestOption) (req pkg.Request) {
	return s.BaseSpider.NewRequest(ctx, options...)
}

// NewItemNone creates a new Item that does not output any data.
// Useful when you only want to collect statistics without persisting items.
func (s *SimpleSpider) NewItemNone(ctx pkg.Context) (item pkg.Item) {
	return s.BaseSpider.NewItemNone(ctx)
}

// NewItemCsv creates a new Item that outputs data to a CSV file with the given filename.
func (s *SimpleSpider) NewItemCsv(ctx pkg.Context, filename string) (item pkg.Item) {
	return s.BaseSpider.NewItemCsv(ctx, filename)
}

// NewItemJsonl creates a new Item that outputs data to a JSON Lines (JSONL) file with the given filename.
// The Context is used for managing file lifecycle and concurrency.
func (s *SimpleSpider) NewItemJsonl(ctx pkg.Context, fileName string) (item pkg.Item) {
	return s.BaseSpider.NewItemJsonl(ctx, fileName)
}

// NewItemMongo creates a new Item that outputs data to a MongoDB collection.
// If update is true, existing documents with the same key will be updated.
func (s *SimpleSpider) NewItemMongo(ctx pkg.Context, collection string, update bool) (item pkg.Item) {
	return s.BaseSpider.NewItemMongo(ctx, collection, update)
}

// NewItemSqlite creates a new Item that outputs data to a SQLite table.
// If update is true, existing rows with the same key will be updated.
func (s *SimpleSpider) NewItemSqlite(ctx pkg.Context, table string, update bool) (item pkg.Item) {
	return s.BaseSpider.NewItemSqlite(ctx, table, update)
}

// NewItemMysql creates a new Item that outputs data to a MySQL table.
// If update is true, existing rows with the same key will be updated.
func (s *SimpleSpider) NewItemMysql(ctx pkg.Context, table string, update bool) (item pkg.Item) {
	return s.BaseSpider.NewItemMysql(ctx, table, update)
}

// NewItemKafka creates a new Item that outputs data to a Kafka topic.
func (s *SimpleSpider) NewItemKafka(ctx pkg.Context, topic string) (item pkg.Item) {
	return s.BaseSpider.NewItemKafka(ctx, topic)
}

func (s *SimpleSpider) YieldExtra(ctx pkg.Context, extra any) (err error) {
	return ctx.GetTask().GetTask().YieldExtra(ctx, extra)
}
func (s *SimpleSpider) MustYieldExtra(ctx pkg.Context, extra any) {
	if err := s.YieldExtra(ctx, extra); err != nil {
		panic(fmt.Errorf("%w: %v", pkg.ErrYieldExtraFailed, err))
	}
}
