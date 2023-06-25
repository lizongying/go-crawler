package config

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/cli"
	"gopkg.in/yaml.v3"
	"log"
	"net/url"
	"os"
	"time"
)

const defaultHttpProto = "2.0"
const defaultDevServer = "http://localhost:8081"
const defaultEnableJa3 = false
const defaultUrlLengthLimit = 2083
const defaultEnableCookieMiddleware = true
const defaultEnableUrlMiddleware = true
const defaultEnableRetryMiddleware = true
const defaultEnableStatsMiddleware = true
const defaultEnableRefererMiddleware = true
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
const defaultEnableImageMiddleware = true
const defaultEnableHttpMiddleware = true
const defaultEnableDumpPipeline = true
const defaultEnableFilterPipeline = true
const defaultRequestConcurrency = uint8(1) // should bigger than 1
const defaultRequestInterval = uint(1000)  // millisecond
const defaultRequestTimeout = uint(60)     //second
const defaultFilterType = pkg.FilterMemory

type Config struct {
	MongoEnable bool `yaml:"mongo_enable" json:"-"`
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
	KafkaEnable bool `yaml:"kafka_enable" json:"-"`
	Kafka       struct {
		Example struct {
			Uri string `yaml:"uri" json:"-"`
		} `yaml:"example" json:"-"`
	} `yaml:"kafka" json:"-"`
	Log struct {
		Filename string `yaml:"filename" json:"-"`
		LongFile bool   `yaml:"long_file" json:"-"`
		Level    string `yaml:"level" json:"-"`
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
	DevServer                string  `yaml:"dev_server" json:"-"`
	EnableJa3                *bool   `yaml:"enable_ja3,omitempty" json:"enable_ja3"`
	EnableRefererMiddleware  *bool   `yaml:"enable_referer,omitempty" json:"enable_referer"`
	ReferrerPolicy           *string `yaml:"referrer_policy,omitempty" json:"referrer_policy"`
	EnableHttpAuthMiddleware *bool   `yaml:"enable_http_auth,omitempty" json:"enable_http_auth"`
	EnableCookieMiddleware   *bool   `yaml:"enable_cookie,omitempty" json:"enable_cookie"`
	EnableStatsMiddleware    *bool   `yaml:"enable_stats,omitempty" json:"enable_stats"`
	EnableDumpMiddleware     *bool   `yaml:"enable_dump_middleware,omitempty" json:"enable_dump_middleware"`
	Filter                   *string `yaml:"filter,omitempty" json:"filter"`
	EnableFilterMiddleware   *bool   `yaml:"enable_filter_middleware,omitempty" json:"enable_filter_middleware"`
	EnableImageMiddleware    *bool   `yaml:"enable_image_middleware,omitempty" json:"enable_image_middleware"`
	EnableHttpMiddleware     *bool   `yaml:"enable_http_middleware,omitempty" json:"enable_http_middleware"`
	EnableRetryMiddleware    *bool   `yaml:"enable_retry,omitempty" json:"enable_retry"`
	EnableUrlMiddleware      *bool   `yaml:"enable_url,omitempty" json:"enable_url"`
	UrlLengthLimit           *int    `yaml:"url_length_limit,omitempty" json:"url_length_limit"`
	EnableCompressMiddleware *bool   `yaml:"enable_compress,omitempty" json:"enable_compress"`
	EnableDecodeMiddleware   *bool   `yaml:"enable_decode,omitempty" json:"enable_decode"`
	EnableRedirectMiddleware *bool   `yaml:"enable_redirect,omitempty" json:"enable_redirect"`
	RedirectMaxTimes         *uint8  `yaml:"redirect_max_times" json:"-"`
	EnableChromeMiddleware   *bool   `yaml:"enable_chrome,omitempty" json:"enable_chrome"`
	EnableDeviceMiddleware   *bool   `yaml:"enable_device,omitempty" json:"enable_device"`
	EnableDumpPipeline       *bool   `yaml:"enable_dump_pipeline,omitempty" json:"enable_dump_pipeline"`
	EnableFilterPipeline     *bool   `yaml:"enable_filter_pipeline,omitempty" json:"enable_filter_pipeline"`
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

func (c *Config) GetDevServer() (url *url.URL, err error) {
	if c.DevServer != "" {
		url, err = url.Parse(c.DevServer)
		return
	}

	url, err = url.Parse(defaultDevServer)
	return
}

func (c *Config) GetEnableJa3() bool {
	if c.EnableJa3 == nil {
		enableJa3 := defaultEnableJa3
		c.EnableJa3 = &enableJa3
	}

	return *c.EnableJa3
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

func (c *Config) GetEnableRefererMiddleware() bool {
	if c.EnableRefererMiddleware == nil {
		enableRefererMiddleware := defaultEnableRefererMiddleware
		c.EnableRefererMiddleware = &enableRefererMiddleware
	}

	return *c.EnableRefererMiddleware
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
func (c *Config) GetEnableImageMiddleware() bool {
	if c.EnableImageMiddleware == nil {
		enableImageMiddleware := defaultEnableImageMiddleware
		c.EnableImageMiddleware = &enableImageMiddleware
	}

	return *c.EnableFilterMiddleware
}
func (c *Config) GetEnableDumpPipeline() bool {
	if c.EnableDumpPipeline == nil {
		enableDumpPipeline := defaultEnableDumpPipeline
		c.EnableDumpPipeline = &enableDumpPipeline
	}

	return *c.EnableDumpPipeline
}
func (c *Config) GetEnableFilterPipeline() bool {
	if c.EnableFilterPipeline == nil {
		enableFilterPipeline := defaultEnableFilterPipeline
		c.EnableFilterPipeline = &enableFilterPipeline
	}

	return *c.EnableFilterPipeline
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
