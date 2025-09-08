package pkg

import (
	"net/url"
	"time"
)

type Config interface {
	GetEnv() string
	GetBotName() string
	GetHttpProto() string
	GetRequestTimeout() time.Duration
	GetEnableJa3() bool
	GetEnablePriorityQueue() bool
	GetReferrerPolicy() ReferrerPolicy
	GetUrlLengthLimit() int
	GetRedirectMaxTimes() uint8
	GetRetryMaxTimes() uint8

	GetEnableStatsMiddleware() bool
	GetEnableDumpMiddleware() bool
	GetEnableProxyMiddleware() bool
	GetEnableRobotsTxtMiddleware() bool
	GetEnableFilterMiddleware() bool
	GetEnableFileMiddleware() bool
	GetEnableImageMiddleware() bool
	GetEnableHttpMiddleware() bool
	GetEnableRetryMiddleware() bool
	GetEnableUrlMiddleware() bool
	GetEnableReferrerMiddleware() bool
	GetEnableCookieMiddleware() bool
	GetEnableRedirectMiddleware() bool
	GetEnableChromeMiddleware() bool
	GetEnableHttpAuthMiddleware() bool
	GetEnableCompressMiddleware() bool
	GetEnableDecodeMiddleware() bool
	GetEnableDeviceMiddleware() bool
	GetEnableRecordErrorMiddleware() bool

	GetEnableDumpPipeline() bool
	GetEnableFilePipeline() bool
	GetEnableImagePipeline() bool
	GetEnableFilterPipeline() bool
	GetEnableNonePipeline() bool
	GetEnableCsvPipeline() bool
	GetEnableJsonLinesPipeline() bool
	GetEnableMongoPipeline() bool
	GetEnableMysqlPipeline() bool
	GetEnableKafkaPipeline() bool

	GetRequestConcurrency() uint8
	GetRequestInterval() uint
	GetRequestRatePerHour() uint
	GetOkHttpCodes() []int
	GetFilter() FilterType
	GetScheduler() SchedulerType
	ApiAccessKey() string
	SetApiAccessKey(accessKey string)
	MockServerEnable() bool
	SetMockServerEnable(enable bool)
	MockServerHost() *url.URL
	CloseReasonQueueTimeout() uint8

	GetLimitType() LimitType

	GetSqliteList() []Sqlite
	GetSqlite() string

	GetRedisList() []Redis
	GetRedis() string

	GetMysqlList() []Mysql
	GetMysql() string

	GetMongoList() []Mongo
	GetMongo() string

	GetKafkaList() []Kafka
	GetKafka() string

	GetStorageList() []Storage
	GetStorage() string

	GetProxyList() []Proxy
	GetProxy() string
}

type Sqlite struct {
	Name string `yaml:"name" json:"-"`
	Path string `yaml:"path" json:"-"`
}

type Redis struct {
	Name     string `yaml:"name" json:"-"`
	Addr     string `yaml:"addr" json:"-"`
	Password string `yaml:"password" json:"-"`
	Db       int    `yaml:"db" json:"-"`
}

type Mysql struct {
	Name     string `yaml:"name" json:"-"`
	Uri      string `yaml:"uri" json:"-"`
	Database string `yaml:"database" json:"-"`
}

type Mongo struct {
	Name     string `yaml:"name" json:"-"`
	Uri      string `yaml:"uri" json:"-"`
	Database string `yaml:"database" json:"-"`
}

type Kafka struct {
	Name string `yaml:"name" json:"-"`
	Uri  string `yaml:"uri" json:"-"`
}

type Storage struct {
	Name     string `yaml:"name" json:"-"`
	Type     string `yaml:"type" json:"-"`
	Endpoint string `yaml:"endpoint" json:"-"`
	Region   string `yaml:"region" json:"-"`
	Id       string `yaml:"id" json:"-"`
	Key      string `yaml:"key" json:"-"`
	Bucket   string `yaml:"bucket" json:"-"`
}
type Proxy struct {
	Name string `yaml:"name" json:"-"`
	Uri  string `yaml:"uri" json:"-"`
}
