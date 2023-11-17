package config

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
	"strings"
	"time"
)

const defaultKafkaUri = "localhost:9092"
const defaultEnv = "dev"
const defaultBotName = "crawler"
const defaultHttpProto = "2.0"
const defaultApiEnable = true
const defaultApiHttps = true
const defaultApiPort = 8090
const defaultApiAccessKey = ""
const defaultMockServerEnable = false
const defaultMockServerHost = "https://localhost:8081"
const defaultMockServerClientAuth = uint8(0)
const defaultCloseReasonQueueTimeout = uint8(10) // second [0-255], 0 no limit
const defaultEnableJa3 = false
const defaultEnablePriorityQueue = true
const defaultUrlLengthLimit = 2083
const defaultEnableCookieMiddleware = true
const defaultEnableUrlMiddleware = true
const defaultEnableRetryMiddleware = true
const defaultEnableStatsMiddleware = true
const defaultEnableReferrerMiddleware = true
const defaultEnableHttpAuthMiddleware = false
const defaultEnableCompressMiddleware = true
const defaultEnableDecodeMiddleware = true
const defaultEnableRedirectMiddleware = true
const defaultRedirectMaxTimes = uint8(1)
const defaultRetryMaxTimes = uint8(10)
const defaultEnableChromeMiddleware = true
const defaultEnableDeviceMiddleware = false
const defaultEnableDumpMiddleware = true
const defaultEnableFilterMiddleware = true
const defaultEnableFileMiddleware = true
const defaultEnableImageMiddleware = true
const defaultEnableHttpMiddleware = true
const defaultEnableProxyMiddleware = true
const defaultEnableRobotsTxtMiddleware = false
const defaultEnableRecordErrorMiddleware = false
const defaultEnableDumpPipeline = true
const defaultEnableFilePipeline = true
const defaultEnableImagePipeline = true
const defaultEnableFilterPipeline = true
const defaultEnableNonePipeline = false
const defaultEnableCsvPipeline = false
const defaultEnableJsonLinesPipeline = false
const defaultEnableMongoPipeline = false
const defaultEnableMysqlPipeline = false
const defaultEnableKafkaPipeline = false
const defaultRequestConcurrency = uint8(1) // should bigger than 1
const defaultRequestInterval = uint(1000)  // millisecond
const defaultRequestTimeout = uint(60)     //second
const defaultFilterType = pkg.FilterMemory
const defaultSchedulerType = pkg.SchedulerMemory
const defaultLogLongFile = true

type Store struct {
	Name     string `yaml:"name" json:"-"`
	Type     string `yaml:"type" json:"-"`
	Endpoint string `yaml:"endpoint" json:"-"`
	Region   string `yaml:"region" json:"-"`
	Id       string `yaml:"id" json:"-"`
	Key      string `yaml:"key" json:"-"`
	Bucket   string `yaml:"bucket" json:"-"`
}

type Sqlite struct {
	Name string `yaml:"name" json:"-"`
	Path string `yaml:"path" json:"-"`
}

type Config struct {
	Env         string `yaml:"env" json:"-"`
	BotName     string `yaml:"bot_name" json:"-"`
	MongoEnable bool   `yaml:"mongo_enable" json:"-"`
	Mongo       struct {
		Example struct {
			Uri      string `yaml:"uri" json:"-"`
			Database string `yaml:"database" json:"-"`
		} `yaml:"example" json:"-"`
	} `yaml:"mongo" json:"-"`
	MysqlEnable bool `yaml:"mysql_enable" json:"-"`
	Mysql       struct {
		Example struct {
			Uri      string `yaml:"uri" json:"-"`
			Database string `yaml:"database" json:"-"`
		} `yaml:"example" json:"-"`
	} `yaml:"mysql" json:"-"`
	RedisEnable bool `yaml:"redis_enable" json:"-"`
	Redis       struct {
		Example struct {
			Addr     string `yaml:"addr" json:"-"`
			Password string `yaml:"password" json:"-"`
			Db       int    `yaml:"db" json:"-"`
		} `yaml:"example" json:"-"`
	} `yaml:"redis" json:"-"`
	Sqlite      []*Sqlite `yaml:"sqlite" json:"-"`
	Store       []*Store  `yaml:"store" json:"-"`
	KafkaEnable bool      `yaml:"kafka_enable" json:"-"`
	Kafka       struct {
		Example struct {
			Uri string `yaml:"uri" json:"-"`
		} `yaml:"example" json:"-"`
	} `yaml:"kafka" json:"-"`
	Log struct {
		Filename string  `yaml:"filename" json:"-"`
		LongFile *bool   `yaml:"long_file" json:"-"`
		Level    *string `yaml:"level" json:"-"`
	} `yaml:"log" json:"-"`
	Proxy struct {
		Example string `yaml:"example" json:"-"`
	} `yaml:"proxy" json:"-"`
	Request struct {
		Concurrency   *uint8 `yaml:"concurrency" json:"-"`
		Interval      *uint  `yaml:"interval" json:"-"`
		Timeout       *uint  `yaml:"timeout" json:"-"`
		OkHttpCodes   []int  `yaml:"ok_http_codes" json:"-"`
		RetryMaxTimes *uint8 `yaml:"retry_max_times" json:"-"`
		HttpProto     string `yaml:"http_proto" json:"-"`
	} `yaml:"request" json:"-"`
	Api struct {
		Enable    *bool  `yaml:"enable,omitempty" json:"enable"`
		Https     *bool  `yaml:"https,omitempty" json:"https"`
		Port      uint16 `yaml:"port,omitempty" json:"port"`
		AccessKey string `yaml:"access_key,omitempty" json:"access_key"`
	} `yaml:"api" json:"api"`
	MockServer struct {
		Enable     *bool  `yaml:"enable,omitempty" json:"enable"`
		Host       string `yaml:"host,omitempty" json:"host"`
		ClientAuth *uint8 `yaml:"client_auth,omitempty" json:"client_auth"`
	} `yaml:"mock_server" json:"mock_server"`
	CloseReason struct {
		QueueTimeout *uint8 `yaml:"client_auth,omitempty" json:"client_auth"`
	} `yaml:"close_reason" json:"close_reason"`
	EnableJa3                   *bool   `yaml:"enable_ja3,omitempty" json:"enable_ja3"`
	EnablePriorityQueue         *bool   `yaml:"enable_priority_queue,omitempty" json:"enable_priority_queue"`
	EnableReferrerMiddleware    *bool   `yaml:"enable_referrer_middleware,omitempty" json:"enable_referrer_middleware"`
	ReferrerPolicy              *string `yaml:"referrer_policy_middleware,omitempty" json:"referrer_policy_middleware"`
	EnableHttpAuthMiddleware    *bool   `yaml:"enable_http_auth_middleware,omitempty" json:"enable_http_auth_middleware"`
	EnableCookieMiddleware      *bool   `yaml:"enable_cookie_middleware,omitempty" json:"enable_cookie_middleware"`
	EnableStatsMiddleware       *bool   `yaml:"enable_stats_middleware,omitempty" json:"enable_stats_middleware"`
	EnableDumpMiddleware        *bool   `yaml:"enable_dump_middleware,omitempty" json:"enable_dump_middleware"`
	Scheduler                   *string `yaml:"scheduler,omitempty" json:"scheduler"`
	Filter                      *string `yaml:"filter,omitempty" json:"filter"`
	EnableFilterMiddleware      *bool   `yaml:"enable_filter_middleware,omitempty" json:"enable_filter_middleware"`
	EnableFileMiddleware        *bool   `yaml:"enable_file_middleware,omitempty" json:"enable_file_middleware"`
	EnableImageMiddleware       *bool   `yaml:"enable_image_middleware,omitempty" json:"enable_image_middleware"`
	EnableHttpMiddleware        *bool   `yaml:"enable_http_middleware,omitempty" json:"enable_http_middleware"`
	EnableRetryMiddleware       *bool   `yaml:"enable_retry_middleware,omitempty" json:"enable_retry_middleware"`
	EnableUrlMiddleware         *bool   `yaml:"enable_url_middleware,omitempty" json:"enable_url_middleware"`
	UrlLengthLimit              *int    `yaml:"url_length_limit,omitempty" json:"url_length_limit"`
	EnableCompressMiddleware    *bool   `yaml:"enable_compress_middleware,omitempty" json:"enable_compress_middleware"`
	EnableDecodeMiddleware      *bool   `yaml:"enable_decode_middleware,omitempty" json:"enable_decode_middleware"`
	EnableRedirectMiddleware    *bool   `yaml:"enable_redirect_middleware,omitempty" json:"enable_redirect_middleware"`
	RedirectMaxTimes            *uint8  `yaml:"redirect_max_times,omitempty" json:"redirect_max_times"`
	EnableChromeMiddleware      *bool   `yaml:"enable_chrome_middleware,omitempty" json:"enable_chrome_middleware"`
	EnableDeviceMiddleware      *bool   `yaml:"enable_device_middleware,omitempty" json:"enable_device_middleware"`
	EnableProxyMiddleware       *bool   `yaml:"enable_proxy_middleware,omitempty" json:"enable_proxy_middleware"`
	EnableRobotsTxtMiddleware   *bool   `yaml:"enable_robots_txt_middleware,omitempty" json:"enable_robots_txt_middleware"`
	EnableRecordErrorMiddleware *bool   `yaml:"enable_record_error_middleware,omitempty" json:"enable_record_error_middleware"`
	EnableDumpPipeline          *bool   `yaml:"enable_dump_pipeline,omitempty" json:"enable_dump_pipeline"`
	EnableFilePipeline          *bool   `yaml:"enable_file_pipeline,omitempty" json:"enable_file_pipeline"`
	EnableImagePipeline         *bool   `yaml:"enable_image_pipeline,omitempty" json:"enable_image_pipeline"`
	EnableFilterPipeline        *bool   `yaml:"enable_filter_pipeline,omitempty" json:"enable_filter_pipeline"`
	EnableNonePipeline          *bool   `yaml:"enable_none_pipeline,omitempty" json:"enable_none_pipeline"`
	EnableCsvPipeline           *bool   `yaml:"enable_csv_pipeline,omitempty" json:"enable_csv_pipeline"`
	EnableJsonLinesPipeline     *bool   `yaml:"enable_json_lines_pipeline,omitempty" json:"enable_json_lines_pipeline"`
	EnableMongoPipeline         *bool   `yaml:"enable_mongo_pipeline,omitempty" json:"enable_mongo_pipeline"`
	EnableMysqlPipeline         *bool   `yaml:"enable_mysql_pipeline,omitempty" json:"enable_mysql_pipeline"`
	EnableKafkaPipeline         *bool   `yaml:"enable_kafka_pipeline,omitempty" json:"enable_kafka_pipeline"`
}

func (c *Config) KafkaUri() string {
	if c.Kafka.Example.Uri != "" {
		return c.Kafka.Example.Uri
	}

	return defaultKafkaUri
}
func (c *Config) GetEnv() string {
	if c.Env != "" {
		return c.Env
	}

	return defaultEnv
}
func (c *Config) GetBotName() string {
	if c.BotName != "" {
		return c.BotName
	}

	return defaultBotName
}

func (c *Config) GetProxy() *url.URL {
	if c.Proxy.Example != "" {
		proxy, err := url.Parse(c.Proxy.Example)
		if err != nil {
			log.Panicln(err)
		}
		return proxy
	}

	return nil
}

func (c *Config) GetHttpProto() string {
	if c.Request.HttpProto != "" {
		return c.Request.HttpProto
	}

	return defaultHttpProto
}
func (c *Config) SetApiEnable(enable bool) {
	c.Api.Enable = &enable
}
func (c *Config) ApiEnable() bool {
	if c.Api.Enable == nil {
		apiEnable := defaultApiEnable
		c.Api.Enable = &apiEnable
	}
	return *c.Api.Enable
}
func (c *Config) SetApiHttps(https bool) {
	c.Api.Https = &https
}
func (c *Config) ApiHttps() bool {
	if c.Api.Https == nil {
		apiHttps := defaultApiHttps
		c.Api.Https = &apiHttps
	}
	return *c.Api.Https
}
func (c *Config) ApiPort() uint16 {
	if c.Api.Port == 0 {
		c.Api.Port = defaultApiPort
	}
	return c.Api.Port
}
func (c *Config) ApiAccessKey() string {
	if c.Api.AccessKey == "" {
		c.Api.AccessKey = defaultApiAccessKey
	}
	return c.Api.AccessKey
}
func (c *Config) SetApiAccessKey(accessKey string) {
	c.Api.AccessKey = accessKey
}
func (c *Config) MockServerEnable() bool {
	if c.MockServer.Enable == nil {
		mockServerEnable := defaultMockServerEnable
		c.MockServer.Enable = &mockServerEnable
	}
	return *c.MockServer.Enable
}
func (c *Config) SetMockServerEnable(enable bool) {
	c.MockServer.Enable = &enable
}
func (c *Config) MockServerHost() *url.URL {
	var URL *url.URL
	var err error
	if c.MockServer.Host != "" {
		URL, err = url.Parse(c.MockServer.Host)
		if err != nil {
			return nil
		}
		return URL
	}

	URL, err = url.Parse(defaultMockServerHost)
	if err != nil {
		return nil
	}
	return URL
}
func (c *Config) MockServerClientAuth() uint8 {
	if c.MockServer.ClientAuth == nil {
		mockServerClientAuth := defaultMockServerClientAuth
		c.MockServer.ClientAuth = &mockServerClientAuth
	}
	return *c.MockServer.ClientAuth
}
func (c *Config) SetMockServerClientAuth(clientAuth uint8) {
	c.MockServer.ClientAuth = &clientAuth
}
func (c *Config) CloseReasonQueueTimeout() uint8 {
	if c.CloseReason.QueueTimeout == nil {
		closeReasonQueueTimeout := defaultCloseReasonQueueTimeout
		c.CloseReason.QueueTimeout = &closeReasonQueueTimeout
	}
	return *c.CloseReason.QueueTimeout
}
func (c *Config) GetEnableJa3() bool {
	if c.EnableJa3 == nil {
		enableJa3 := defaultEnableJa3
		c.EnableJa3 = &enableJa3
	}

	return *c.EnableJa3
}
func (c *Config) GetEnablePriorityQueue() bool {
	if c.EnablePriorityQueue == nil {
		enablePriorityQueue := defaultEnablePriorityQueue
		c.EnablePriorityQueue = &enablePriorityQueue
	}

	return *c.EnablePriorityQueue
}
func (c *Config) GetReferrerPolicy() pkg.ReferrerPolicy {
	if c.ReferrerPolicy == nil {
		referrerPolicy := string(pkg.DefaultReferrerPolicy)
		c.ReferrerPolicy = &referrerPolicy
	}
	if *c.ReferrerPolicy != "" {
		switch pkg.ReferrerPolicy(*c.ReferrerPolicy) {
		case pkg.DefaultReferrerPolicy:
			return pkg.DefaultReferrerPolicy
		case pkg.NoReferrerPolicy:
			return pkg.NoReferrerPolicy
		default:
			return pkg.DefaultReferrerPolicy
		}
	}

	return pkg.DefaultReferrerPolicy
}

func (c *Config) GetEnableCookieMiddleware() bool {
	if c.EnableCookieMiddleware == nil {
		enableCookieMiddleware := defaultEnableCookieMiddleware
		c.EnableCookieMiddleware = &enableCookieMiddleware
	}

	return *c.EnableCookieMiddleware
}
func (c *Config) GetEnableHttpAuthMiddleware() bool {
	if c.EnableHttpAuthMiddleware == nil {
		enableHttpAuthMiddleware := defaultEnableHttpAuthMiddleware
		c.EnableHttpAuthMiddleware = &enableHttpAuthMiddleware
	}

	return *c.EnableHttpAuthMiddleware
}

func (c *Config) GetEnableReferrerMiddleware() bool {
	if c.EnableReferrerMiddleware == nil {
		enableReferrerMiddleware := defaultEnableReferrerMiddleware
		c.EnableReferrerMiddleware = &enableReferrerMiddleware
	}

	return *c.EnableReferrerMiddleware
}
func (c *Config) GetEnableDumpMiddleware() bool {
	if c.EnableDumpMiddleware == nil {
		enableDumpMiddleware := defaultEnableDumpMiddleware
		c.EnableDumpMiddleware = &enableDumpMiddleware
	}

	return *c.EnableDumpMiddleware
}
func (c *Config) GetEnableFilterMiddleware() bool {
	if c.EnableFilterMiddleware == nil {
		enableFilterMiddleware := defaultEnableFilterMiddleware
		c.EnableFilterMiddleware = &enableFilterMiddleware
	}

	return *c.EnableFilterMiddleware
}
func (c *Config) GetEnableFileMiddleware() bool {
	if c.EnableFileMiddleware == nil {
		enableFileMiddleware := defaultEnableFileMiddleware
		c.EnableFileMiddleware = &enableFileMiddleware
	}

	return *c.EnableFileMiddleware
}
func (c *Config) GetEnableImageMiddleware() bool {
	if c.EnableImageMiddleware == nil {
		enableImageMiddleware := defaultEnableImageMiddleware
		c.EnableImageMiddleware = &enableImageMiddleware
	}

	return *c.EnableImageMiddleware
}
func (c *Config) GetEnableDumpPipeline() bool {
	if c.EnableDumpPipeline == nil {
		enableDumpPipeline := defaultEnableDumpPipeline
		c.EnableDumpPipeline = &enableDumpPipeline
	}

	return *c.EnableDumpPipeline
}
func (c *Config) GetEnableFilePipeline() bool {
	if c.EnableFilePipeline == nil {
		enableFilePipeline := defaultEnableFilePipeline
		c.EnableFilePipeline = &enableFilePipeline
	}

	return *c.EnableFilePipeline
}
func (c *Config) GetEnableImagePipeline() bool {
	if c.EnableImagePipeline == nil {
		enableImagePipeline := defaultEnableImagePipeline
		c.EnableImagePipeline = &enableImagePipeline
	}

	return *c.EnableImagePipeline
}
func (c *Config) GetEnableFilterPipeline() bool {
	if c.EnableFilterPipeline == nil {
		enableFilterPipeline := defaultEnableFilterPipeline
		c.EnableFilterPipeline = &enableFilterPipeline
	}

	return *c.EnableFilterPipeline
}
func (c *Config) GetEnableNonePipeline() bool {
	if c.EnableCsvPipeline == nil {
		enableNonePipeline := defaultEnableNonePipeline
		c.EnableCsvPipeline = &enableNonePipeline
	}

	return *c.EnableNonePipeline
}
func (c *Config) GetEnableCsvPipeline() bool {
	if c.EnableCsvPipeline == nil {
		enableCsvPipeline := defaultEnableCsvPipeline
		c.EnableCsvPipeline = &enableCsvPipeline
	}

	return *c.EnableCsvPipeline
}
func (c *Config) GetEnableJsonLinesPipeline() bool {
	if c.EnableJsonLinesPipeline == nil {
		enableJsonLinesPipeline := defaultEnableJsonLinesPipeline
		c.EnableJsonLinesPipeline = &enableJsonLinesPipeline
	}

	return *c.EnableJsonLinesPipeline
}
func (c *Config) GetEnableMongoPipeline() bool {
	if c.EnableMongoPipeline == nil {
		enableMongoPipeline := defaultEnableMongoPipeline
		c.EnableMongoPipeline = &enableMongoPipeline
	}

	return *c.EnableMongoPipeline
}
func (c *Config) GetEnableMysqlPipeline() bool {
	if c.EnableMysqlPipeline == nil {
		enableMysqlPipeline := defaultEnableMysqlPipeline
		c.EnableMysqlPipeline = &enableMysqlPipeline
	}

	return *c.EnableMysqlPipeline
}
func (c *Config) GetEnableKafkaPipeline() bool {
	if c.EnableKafkaPipeline == nil {
		enableKafkaPipeline := defaultEnableKafkaPipeline
		c.EnableKafkaPipeline = &enableKafkaPipeline
	}

	return *c.EnableKafkaPipeline
}
func (c *Config) GetEnableStatsMiddleware() bool {
	if c.EnableStatsMiddleware == nil {
		enableStatsMiddleware := defaultEnableStatsMiddleware
		c.EnableStatsMiddleware = &enableStatsMiddleware
	}

	return *c.EnableStatsMiddleware
}

func (c *Config) GetEnableUrlMiddleware() bool {
	if c.EnableUrlMiddleware == nil {
		enableUrlMiddleware := defaultEnableUrlMiddleware
		c.EnableUrlMiddleware = &enableUrlMiddleware
	}

	return *c.EnableUrlMiddleware
}

func (c *Config) GetUrlLengthLimit() int {
	if c.UrlLengthLimit == nil {
		urlLengthLimit := defaultUrlLengthLimit
		c.UrlLengthLimit = &urlLengthLimit
	}

	return *c.UrlLengthLimit
}

func (c *Config) GetEnableRetryMiddleware() bool {
	if c.EnableRetryMiddleware == nil {
		enableRetryMiddleware := defaultEnableRetryMiddleware
		c.EnableRetryMiddleware = &enableRetryMiddleware
	}

	return *c.EnableRetryMiddleware
}

func (c *Config) GetEnableCompressMiddleware() bool {
	if c.EnableCompressMiddleware == nil {
		enableCompressMiddleware := defaultEnableCompressMiddleware
		c.EnableCompressMiddleware = &enableCompressMiddleware
	}

	return *c.EnableCompressMiddleware
}

func (c *Config) GetEnableDecodeMiddleware() bool {
	if c.EnableDecodeMiddleware == nil {
		enableDecodeMiddleware := defaultEnableDecodeMiddleware
		c.EnableDecodeMiddleware = &enableDecodeMiddleware
	}

	return *c.EnableDecodeMiddleware
}

func (c *Config) GetEnableRedirectMiddleware() bool {
	if c.EnableRedirectMiddleware == nil {
		enableRedirectMiddleware := defaultEnableRedirectMiddleware
		c.EnableRedirectMiddleware = &enableRedirectMiddleware
	}

	return *c.EnableRedirectMiddleware
}
func (c *Config) GetEnableProxyMiddleware() bool {
	if c.EnableProxyMiddleware == nil {
		enableProxyMiddleware := defaultEnableProxyMiddleware
		c.EnableProxyMiddleware = &enableProxyMiddleware
	}

	return *c.EnableProxyMiddleware
}
func (c *Config) GetEnableRobotsTxtMiddleware() bool {
	if c.EnableRobotsTxtMiddleware == nil {
		enableRobotsTxtMiddleware := defaultEnableRobotsTxtMiddleware
		c.EnableRobotsTxtMiddleware = &enableRobotsTxtMiddleware
	}

	return *c.EnableRobotsTxtMiddleware
}
func (c *Config) GetEnableRecordErrorMiddleware() bool {
	if c.EnableRecordErrorMiddleware == nil {
		enableRecordErrorMiddleware := defaultEnableRecordErrorMiddleware
		c.EnableRecordErrorMiddleware = &enableRecordErrorMiddleware
	}

	return *c.EnableRecordErrorMiddleware
}
func (c *Config) GetRedirectMaxTimes() uint8 {
	if c.RedirectMaxTimes == nil {
		redirectMaxTimes := defaultRedirectMaxTimes
		c.RedirectMaxTimes = &redirectMaxTimes
	}

	return *c.RedirectMaxTimes
}

func (c *Config) GetRetryMaxTimes() uint8 {
	if c.Request.RetryMaxTimes == nil {
		retryMaxTimes := defaultRetryMaxTimes
		c.Request.RetryMaxTimes = &retryMaxTimes
	}

	return *c.Request.RetryMaxTimes
}

func (c *Config) GetEnableChromeMiddleware() bool {
	if c.EnableChromeMiddleware == nil {
		enableChromeMiddleware := defaultEnableChromeMiddleware
		c.EnableChromeMiddleware = &enableChromeMiddleware
	}

	return *c.EnableChromeMiddleware
}

func (c *Config) GetEnableDeviceMiddleware() bool {
	if c.EnableDeviceMiddleware == nil {
		enableDeviceMiddleware := defaultEnableDeviceMiddleware
		c.EnableDeviceMiddleware = &enableDeviceMiddleware
	}

	return *c.EnableDeviceMiddleware
}

func (c *Config) GetEnableHttpMiddleware() bool {
	if c.EnableHttpMiddleware == nil {
		enableHttpMiddleware := defaultEnableHttpMiddleware
		c.EnableHttpMiddleware = &enableHttpMiddleware
	}

	return *c.EnableHttpMiddleware
}
func (c *Config) GetLogLongFile() bool {
	if c.Log.LongFile == nil {
		logLongFile := defaultLogLongFile
		c.Log.LongFile = &logLongFile
	}

	return *c.Log.LongFile
}
func (c *Config) GetLogLevel() pkg.Level {
	if c.Log.Level == nil {
		return pkg.LevelInfo
	}

	if level, ok := pkg.LevelMap[strings.ToUpper(*c.Log.Level)]; ok {
		return level
	}

	return pkg.LevelInfo
}

func (c *Config) GetRequestConcurrency() uint8 {
	if c.Request.Concurrency == nil || *c.Request.Concurrency == 0 {
		requestConcurrency := defaultRequestConcurrency
		c.Request.Concurrency = &requestConcurrency
	}

	return *c.Request.Concurrency
}

func (c *Config) GetRequestInterval() uint {
	if c.Request.Interval == nil {
		requestInterval := defaultRequestInterval
		c.Request.Interval = &requestInterval
	}

	return *c.Request.Interval
}

func (c *Config) GetRequestTimeout() time.Duration {
	if c.Request.Timeout == nil || *c.Request.Timeout == 0 {
		requestTimeout := defaultRequestTimeout
		c.Request.Timeout = &requestTimeout
	}

	return time.Second * time.Duration(int(*c.Request.Timeout))
}

func (c *Config) GetOkHttpCodes() []int {
	if len(c.Request.OkHttpCodes) == 0 {
		c.Request.OkHttpCodes = []int{200}
	}

	return c.Request.OkHttpCodes
}
func (c *Config) GetScheduler() pkg.SchedulerType {
	if c.Scheduler == nil {
		schedulerType := string(defaultSchedulerType)
		c.Scheduler = &schedulerType
	}
	if *c.Scheduler != "" {
		switch pkg.SchedulerType(*c.Scheduler) {
		case pkg.SchedulerMemory:
			return pkg.SchedulerMemory
		case pkg.SchedulerRedis:
			return pkg.SchedulerRedis
		case pkg.SchedulerKafka:
			return pkg.SchedulerKafka
		default:
			return pkg.SchedulerUnknown
		}
	}

	return pkg.SchedulerUnknown
}
func (c *Config) GetFilter() pkg.FilterType {
	if c.Filter == nil {
		filterType := string(defaultFilterType)
		c.Filter = &filterType
	}
	if *c.Filter != "" {
		switch pkg.FilterType(*c.Filter) {
		case pkg.FilterMemory:
			return pkg.FilterMemory
		case pkg.FilterRedis:
			return pkg.FilterRedis
		default:
			return pkg.FilterUnknown
		}
	}

	return pkg.FilterUnknown
}
func (c *Config) GetSqlite() []*Sqlite {
	return c.Sqlite
}
func (c *Config) GetStore() []*Store {
	return c.Store
}

func (c *Config) LoadConfig(configPath string) (err error) {
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Panicln(err)
	}

	err = yaml.Unmarshal(configData, c)
	if err != nil {
		log.Panicln(err)
	}

	return
}

func NewConfig(cli *cli.Cli) (config *Config, err error) {
	config = &Config{}
	configFile := cli.ConfigFile
	if configFile != "" {
		err = config.LoadConfig(configFile)
		if err != nil {
			log.Panicln(err)
		}
	}

	return
}
