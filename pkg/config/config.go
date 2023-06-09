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
const defaultTimeout = time.Minute
const defaultDevServer = "http://localhost:8081"
const defaultEnableJa3 = false
const defaultUrlLengthLimit = 2083
const defaultEnableCookie = true
const defaultEnableDump = true
const defaultEnableUrl = true
const defaultEnableRetry = true
const defaultEnableStats = true
const defaultEnableFilter = true
const defaultEnableReferer = true
const defaultEnableHttpAuth = false
const defaultEnableCompress = true
const defaultEnableDecode = true
const defaultEnableRedirect = true
const defaultRedirectMaxTimes = uint8(1)
const defaultRetryMaxTimes = uint8(10)

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
		Concurrency   int    `yaml:"concurrency" json:"-"`
		Interval      int    `yaml:"interval" json:"-"`
		Timeout       int    `yaml:"timeout" json:"-"`
		OkHttpCodes   []int  `yaml:"ok_http_codes" json:"-"`
		RetryMaxTimes *uint8 `yaml:"retry_max_times" json:"-"`
		HttpProto     string `yaml:"http_proto" json:"-"`
	} `yaml:"request" json:"-"`
	DevServer        string  `yaml:"dev_server" json:"-"`
	EnableJa3        *bool   `yaml:"enable_ja3,omitempty" json:"enable_ja3"`
	EnableReferer    *bool   `yaml:"enable_referer,omitempty" json:"enable_referer"`
	ReferrerPolicy   *string `yaml:"referrer_policy,omitempty" json:"referrer_policy"`
	EnableHttpAuth   *bool   `yaml:"enable_http_auth,omitempty" json:"enable_http_auth"`
	EnableCookie     *bool   `yaml:"enable_cookie,omitempty" json:"enable_cookie"`
	EnableDump       *bool   `yaml:"enable_dump,omitempty" json:"enable_dump"`
	EnableFilter     *bool   `yaml:"enable_filter,omitempty" json:"enable_filter"`
	EnableStats      *bool   `yaml:"enable_stats,omitempty" json:"enable_stats"`
	EnableUrl        *bool   `yaml:"enable_url,omitempty" json:"enable_url"`
	UrlLengthLimit   *int    `yaml:"url_length_limit,omitempty" json:"url_length_limit"`
	EnableRetry      *bool   `yaml:"enable_retry,omitempty" json:"enable_retry"`
	EnableCompress   *bool   `yaml:"enable_compress,omitempty" json:"enable_compress"`
	EnableDecode     *bool   `yaml:"enable_decode,omitempty" json:"enable_decode"`
	EnableRedirect   *bool   `yaml:"enable_redirect,omitempty" json:"enable_redirect"`
	RedirectMaxTimes *uint8  `yaml:"redirect_max_times" json:"-"`
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

func (c *Config) GetTimeout() time.Duration {
	if c.Request.Timeout > 0 {
		return time.Second * time.Duration(c.Request.Timeout)
	}

	return defaultTimeout
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

func (c *Config) GetEnableDump() bool {
	if c.EnableDump == nil {
		enableDump := defaultEnableDump
		c.EnableDump = &enableDump
	}

	return *c.EnableDump
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

func (c *Config) GetEnableFilter() bool {
	if c.EnableFilter == nil {
		enableFilter := defaultEnableFilter
		c.EnableFilter = &enableFilter
	}

	return *c.EnableFilter
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
