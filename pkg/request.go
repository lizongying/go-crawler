package pkg

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
	HEAD = "HEAD"
)

type Request interface {
	UniqueKey() string
	SetUniqueKey(string) Request
	SetCallBack(CallBack) Request
	GetCallBack() CallBack
	SetErrBack(ErrBack) Request
	GetErrBack() ErrBack
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
	SetProxyEnable(*bool) Request
	Proxy() *url.URL
	SetProxy(*url.URL) Request
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
	AddQuery(string, string) Request
	SetQuery(string, string) Request
	GetQuery(string) Request
	DelQuery(string) Request
	HasQuery(string) Request
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
	File() bool
	SetFile(bool) Request
	Image() bool
	SetImage(bool) Request
	Extra() string
	SetExtra(any) Request
	UnmarshalExtra(any) error
	MustUnmarshalExtra(any)
	ToRequestJson() (RequestJson, error)
	Marshal() ([]byte, error)
	SetBasicAuth(string, string) Request
	Context() context.Context
	WithContext(context.Context) Request
	GetRequest() *http.Request
	Cookies() []*http.Cookie
	AddCookie(c *http.Cookie) Request
}
type RequestJson interface {
	SetCallBacks(map[string]CallBack)
	SetErrBacks(map[string]ErrBack)
	ToRequest() (Request, error)
}

type CallBack func(Context, Response) error
type ErrBack func(Context, Response, error)

type Client string

const (
	ClientDefault Client = ""
	ClientGo      Client = "go"
	ClientBrowser Client = "browser"
)
