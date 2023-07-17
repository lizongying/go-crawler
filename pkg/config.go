package pkg

import (
	"net/url"
	"time"
)

type Config interface {
	GetProxy() *url.URL
	GetHttpProto() string
	GetRequestTimeout() time.Duration
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

	GetEnableDumpPipeline() bool
	GetEnableFilePipeline() bool
	GetEnableImagePipeline() bool
	GetEnableFilterPipeline() bool
	GetEnableCsvPipeline() bool
	GetEnableJsonLinesPipeline() bool
	GetEnableMongoPipeline() bool
	GetEnableMysqlPipeline() bool
	GetEnableKafkaPipeline() bool

	GetRequestConcurrency() uint8
	GetRequestInterval() uint
	GetOkHttpCodes() []int
	GetFilter() FilterType
	GetDevServer() (URL *url.URL, err error)
}
