package pkg

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type HTTPMethod uint8

const (
	Unknown HTTPMethod = iota
	GET
	POST
	HEAD
	PUT
	DELETE
	PATCH
	OPTIONS
	CONNECT
	TRACE
)

func (h HTTPMethod) String() string {
	switch h {
	case GET:
		return http.MethodGet
	case POST:
		return http.MethodPost
	case HEAD:
		return http.MethodHead
	case PATCH:
		return http.MethodPatch
	case PUT:
		return http.MethodPut
	case DELETE:
		return http.MethodDelete
	case OPTIONS:
		return http.MethodOptions
	case CONNECT:
		return http.MethodConnect
	case TRACE:
		return http.MethodTrace
	case Unknown:
		return ""
	default:
		return ""
	}
}

type Request interface {
	UniqueKey() string
	SetUniqueKey(string) Request
	SetCallBack(CallBack) Request
	CallBack() string
	SetErrBack(ErrBack) Request
	ErrBack() string
	Referrer() string
	SetReferrer(string) Request
	Username() string
	SetUsername(string) Request
	Password() string
	SetPassword(string) Request
	Checksum() string
	SetChecksum(string) Request
	CreateTime() string
	SetCreateTime(string) Request
	SpendTime() time.Duration
	SetSpendTime(time.Duration) Request
	SkipMiddleware() bool
	SetSkipMiddleware(bool) Request
	SkipFilter() *bool
	SetSkipFilter(*bool) Request
	CanonicalHeaderKey() *bool
	SetCanonicalHeaderKey(*bool) Request
	ProxyEnable() *bool
	SetProxyEnable(bool) Request
	Proxy() *url.URL
	SetProxy(string) Request
	RetryMaxTimes() *uint8
	SetRetryMaxTimes(*uint8) Request
	RetryTimes() uint8
	SetRetryTimes(uint8) Request
	RedirectMaxTimes() *uint8
	SetRedirectMaxTimes(*uint8) Request
	RedirectTimes() uint8
	SetRedirectTimes(uint8) Request
	OkHttpCodes() []int
	SetOkHttpCodes([]int) Request
	Slot() string
	SetSlot(string) Request
	Concurrency() *uint8
	SetConcurrency(*uint8) Request
	Interval() time.Duration
	SetInterval(time.Duration) Request
	Timeout() time.Duration
	SetTimeout(time.Duration) Request
	HttpProto() string
	SetHttpProto(string) Request
	Platforms() []Platform
	SetPlatforms(...Platform) Request
	Browsers() []Browser
	SetBrowsers(...Browser) Request
	GetExtraName() string
	Priority() uint8
	SetPriority(uint8) Request
	Fingerprint() string
	SetFingerprint(string) Request
	Client() Client
	SetClient(Client) Request
	Ajax() bool
	SetAjax(bool) Request
	Err() map[string]error
	SetUrl(string) Request
	GetUrl() string
	GetURL() *url.URL
	Query(string) string
	AddQuery(string, string) Request
	SetQuery(string, string) Request
	DelQuery(string) Request
	HasQuery(string) bool
	SetForm(string, string) Request
	GetForm() url.Values
	SetPostForm(string, string) Request
	GetPostForm() url.Values
	SetMethod(string) Request
	GetMethod() string
	BodyStr() string
	SetBodyStr(string) Request
	BodyBytes() []byte
	SetBodyBytes([]byte) Request
	SetBodyJson(bodyJson any) Request
	GetHeader(string) string
	SetHeader(string, string) Request
	Headers() http.Header
	SetHeaders(map[string]string) Request
	IsFile() bool
	AsFile(bool) Request
	SetFileOptions(options FileOptions) Request
	FileOptions() *FileOptions
	IsImage() bool
	AsImage(bool) Request
	SetImageOptions(options ImageOptions) Request
	ImageOptions() *ImageOptions
	SetTask(task Task) Request
	Task() *Task
	Extra() string
	SetExtra(any) Request
	UnmarshalExtra(any) error
	MustUnmarshalExtra(any)
	ToRequestJson() RequestJson
	Marshal() ([]byte, error)
	SetBasicAuth(string, string) Request
	RequestContext() context.Context
	WithContext(context.Context) Request
	GetRequest() *http.Request
	Cookies() []*http.Cookie
	AddCookie(c *http.Cookie) Request
}

type RequestWithContext interface {
	Context
	Request
}

type RequestJson interface {
	ToRequest() (Request, error)
}

type CallBack func(Context, Response) error
type ErrBack func(Context, Response, error)
