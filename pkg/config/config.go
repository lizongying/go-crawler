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
const defaultEnableCookie = true
const defaultEnableUrl = true
const defaultEnableRetry = true
const defaultEnableStats = true
const defaultEnableReferer = true
const defaultEnableHttpAuth = false
const defaultEnableCompress = true
const defaultEnableDecode = true
const defaultEnableRedirect = true
const defaultRedirectMaxTimes = uint8(1)
const defaultRetryMaxTimes = uint8(10)
const defaultEnableChrome = true
const defaultEnableDevice = false
const defaultEnableDumpMiddleware = true
const defaultEnableFilterMiddleware = true
const defaultEnableImageMiddleware = false
const defaultEnableDumpPipeline = true
const defaultEnableFilterPipeline = true
const defaultRequestConcurrency = uint8(1) // should bigger than 1
const defaultRequestInterval = uint(1000)  // millisecond
const defaultRequestTimeout = uint(60)     //second

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
	DevServer              string  `yaml:"dev_server" json:"-"`
	EnableJa3              *bool   `yaml:"enable_ja3,omitempty" json:"enable_ja3"`
	EnableReferer          *bool   `yaml:"enable_referer,omitempty" json:"enable_referer"`
	ReferrerPolicy         *string `yaml:"referrer_policy,omitempty" json:"referrer_policy"`
	EnableHttpAuth         *bool   `yaml:"enable_http_auth,omitempty" json:"enable_http_auth"`
	EnableCookie           *bool   `yaml:"enable_cookie,omitempty" json:"enable_cookie"`
	EnableStats            *bool   `yaml:"enable_stats,omitempty" json:"enable_stats"`
	EnableUrl              *bool   `yaml:"enable_url,omitempty" json:"enable_url"`
	UrlLengthLimit         *int    `yaml:"url_length_limit,omitempty" json:"url_length_limit"`
	EnableRetry            *bool   `yaml:"enable_retry,omitempty" json:"enable_retry"`
	EnableCompress         *bool   `yaml:"enable_compress,omitempty" json:"enable_compress"`
	EnableDecode           *bool   `yaml:"enable_decode,omitempty" json:"enable_decode"`
	EnableRedirect         *bool   `yaml:"enable_redirect,omitempty" json:"enable_redirect"`
	RedirectMaxTimes       *uint8  `yaml:"redirect_max_times" json:"-"`
	EnableChrome           *bool   `yaml:"enable_chrome,omitempty" json:"enable_chrome"`
	EnableDevice           *bool   `yaml:"enable_device,omitempty" json:"enable_device"`
	EnableDumpMiddleware   *bool   `yaml:"enable_dump_middleware,omitempty" json:"enable_dump_middleware"`
	EnableFilterMiddleware *bool   `yaml:"enable_filter_middleware,omitempty" json:"enable_filter_middleware"`
	EnableImageMiddleware  *bool   `yaml:"enable_image_middleware,omitempty" json:"enable_image_middleware"`
	EnableDumpPipeline     *bool   `yaml:"enable_dump_pipeline,omitempty" json:"enable_dump_pipeline"`
	EnableFilterPipeline   *bool   `yaml:"enable_filter_pipeline,omitempty" json:"enable_filter_pipeline"`
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

func (c *Config) GetEnableCookie() bool {
	if c.EnableCookie == nil {
		enableCookie := defaultEnableCookie
		c.EnableCookie = &enableCookie
	}

	return *c.EnableCookie
}

func (c *Config) GetEnableHttpAuth() bool {
	if c.EnableHttpAuth == nil {
		enableHttpAuth := defaultEnableHttpAuth
		c.EnableHttpAuth = &enableHttpAuth
	}

	return *c.EnableHttpAuth
}

func (c *Config) GetEnableReferer() bool {
	if c.EnableReferer == nil {
		enableReferer := defaultEnableReferer
		c.EnableReferer = &enableReferer
	}

	return *c.EnableReferer
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

func (c *Config) GetEnableStats() bool {
	if c.EnableStats == nil {
		enableStats := defaultEnableStats
		c.EnableStats = &enableStats
	}

	return *c.EnableStats
}

func (c *Config) GetEnableUrl() bool {
	if c.EnableUrl == nil {
		enableUrl := defaultEnableUrl
		c.EnableUrl = &enableUrl
	}

	return *c.EnableUrl
}

func (c *Config) GetUrlLengthLimit() int {
	if c.UrlLengthLimit == nil {
		urlLengthLimit := defaultUrlLengthLimit
		c.UrlLengthLimit = &urlLengthLimit
	}

	return *c.UrlLengthLimit
}

func (c *Config) GetEnableRetry() bool {
	if c.EnableRetry == nil {
		enableRetry := defaultEnableRetry
		c.EnableRetry = &enableRetry
	}

	return *c.EnableRetry
}

func (c *Config) GetEnableCompress() bool {
	if c.EnableCompress == nil {
		enableCompress := defaultEnableCompress
		c.EnableCompress = &enableCompress
	}

	return *c.EnableCompress
}

func (c *Config) GetEnableDecode() bool {
	if c.EnableDecode == nil {
		enableDecode := defaultEnableDecode
		c.EnableDecode = &enableDecode
	}

	return *c.EnableDecode
}

func (c *Config) GetEnableRedirect() bool {
	if c.EnableRedirect == nil {
		enableRedirect := defaultEnableRedirect
		c.EnableRedirect = &enableRedirect
	}

	return *c.EnableRedirect
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

func (c *Config) GetEnableChrome() bool {
	if c.EnableChrome == nil {
		enableChrome := defaultEnableChrome
		c.EnableChrome = &enableChrome
	}

	return *c.EnableChrome
}

func (c *Config) GetEnableDevice() bool {
	if c.EnableDevice == nil {
		enableDevice := defaultEnableDevice
		c.EnableDevice = &enableDevice
	}

	return *c.EnableDevice
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
