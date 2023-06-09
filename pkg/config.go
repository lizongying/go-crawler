package pkg

import (
	"net/url"
	"time"
)

type Config interface {
	GetProxy() *url.URL
	GetHttpProto() string
	GetTimeout() time.Duration
	GetReferrerPolicy() ReferrerPolicy
	GetUrlLengthLimit() int
	GetEnableCookie() bool
	GetRedirectMaxTimes() uint8
	GetRetryMaxTimes() uint8
}
