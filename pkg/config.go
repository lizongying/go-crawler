package pkg

import (
	"net/url"
	"time"
)

type Config interface {
	GetEnv() string
	GetBotName() string
	GetProxy() *url.URL
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
	KafkaUri() string

	GetLimitType() LimitType
}
