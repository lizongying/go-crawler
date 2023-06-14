package pkg

import (
	"context"
	"database/sql"
	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"net/url"
	"time"
)

type SpiderInfo struct {
	Mode string
	Name string

	Concurrency   int
	Interval      time.Duration
	RetryMaxTimes uint8
	Timeout       time.Duration
	Username      string
	Password      string
}

type Spider interface {
	GetInfo() *SpiderInfo
	GetStats() Stats
	SetLogger(logger Logger)
	SetSpider(spider Spider)
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	GetAllowedDomains() []string
	ReplaceAllowedDomains([]string) error
	SetAllowedDomain(string)
	DelAllowedDomain(string) error
	CleanAllowedDomains()
	SetItemDelay(time.Duration) Spider
	SetItemConcurrency(int) Spider
	SetRequestRate(string, time.Duration, int) Spider
	IsAllowedDomain(*url.URL) bool
	YieldRequest(context.Context, *Request) error
	YieldItem(context.Context, Item) error
	RunDevServer() error
	GetDevServerHost() string
	AddDevServerRoutes(routes ...Route) Spider
	AddOkHttpCodes(httpCodes ...int) Spider
	GetOkHttpCodes() []int
	GetPlatforms() []Platform
	GetBrowsers() []Browser
	SetPlatforms(...Platform) Spider
	SetBrowsers(...Browser) Spider
	GetConfig() Config
	GetLogger() Logger
	GetKafka() *kafka.Writer
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
	GetFilter() Filter
}
