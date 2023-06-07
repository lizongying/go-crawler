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
	RetryMaxTimes int
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
	GetMiddlewares() map[uint8]string
	ReplaceMiddlewares(map[uint8]Middleware) error
	SetMiddleware(Middleware, uint8) Spider
	DelMiddleware(string)
	CleanMiddlewares()
	SortedMiddlewares() []Middleware
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
	GetHttpClient() HttpClient
	GetKafka() *kafka.Writer
	GetMongoDb() *mongo.Database
	GetMysql() *sql.DB
}
