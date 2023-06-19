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

	GetEnableStats() bool
	GetEnableDumpMiddleware() bool
	GetEnableFilterMiddleware() bool
	GetEnableImageMiddleware() bool
	GetEnableRetry() bool
	GetEnableUrl() bool
	GetEnableReferer() bool
	GetEnableCookie() bool
	GetEnableRedirect() bool
	GetEnableChrome() bool
	GetEnableHttpAuth() bool
	GetEnableCompress() bool
	GetEnableDecode() bool
	GetEnableDevice() bool

	GetEnableDumpPipeline() bool
	GetEnableFilterPipeline() bool

	GetRequestConcurrency() uint8
	GetRequestInterval() uint
	GetOkHttpCodes() []int
}
