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
	GetUniqueKey() string
	SetUniqueKey(string) Request
	SetCallBack(CallBack) Request
	GetCallBack() string
	SetErrBack(ErrBack) Request
	GetErrBack() string
	GetReferrer() string
	SetReferrer(string) Request
	GetUsername() string
	SetUsername(string) Request
	GetPassword() string
	SetPassword(string) Request
	GetChecksum() string
	SetChecksum(string) Request
	GetCreateTime() string
	SetCreateTime(string) Request
	GetSpendTime() time.Duration
	SetSpendTime(time.Duration) Request
	IsSkipMiddleware() bool
	SetSkipMiddleware(bool) Request
	IsSkipFilter() *bool
	SetSkipFilter(*bool) Request
	IsCanonicalHeaderKey() *bool
	SetCanonicalHeaderKey(*bool) Request
	IsProxyEnable() *bool
	SetProxyEnable(bool) Request
	GetProxy() *url.URL
	SetProxy(string) Request
	GetRetryMaxTimes() *uint8
	SetRetryMaxTimes(*uint8) Request
	GetRetryTimes() uint8
	SetRetryTimes(uint8) Request
	GetRedirectMaxTimes() *uint8
	SetRedirectMaxTimes(*uint8) Request
	GetRedirectTimes() uint8
	SetRedirectTimes(uint8) Request
	GetOkHttpCodes() []int
	SetOkHttpCodes([]int) Request
	GetSlot() string
	SetSlot(string) Request
	GetConcurrency() *uint8
	SetConcurrency(*uint8) Request
	GetInterval() time.Duration
	SetInterval(time.Duration) Request
	GetTimeout() time.Duration
	SetTimeout(time.Duration) Request
	GetHttpProto() string
	SetHttpProto(string) Request
	GetPlatforms() []Platform
	SetPlatforms(...Platform) Request
	GetBrowsers() []Browser
	SetBrowsers(...Browser) Request
	GetExtraName() string
	GetPriority() uint8
	SetPriority(uint8) Request
	GetFingerprint() string
	SetFingerprint(string) Request
	GetClient() Client
	SetClient(Client) Request
	IsAjax() bool
	SetAjax(bool) Request
	GetScreenshot() string
	SetScreenshot(string) Request
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
	GetBodyStr() string
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
	GetFileOptions() *FileOptions
	IsImage() bool
	AsImage(bool) Request
	SetImageOptions(options ImageOptions) Request
	GetImageOptions() *ImageOptions
	GetContext() Context
	WithContext(Context) Request
	GetExtra() string
	SetExtra(any) Request
	UnmarshalExtra(any) error
	MustUnmarshalExtra(any)
	UnsafeExtra(any)
	Marshal() ([]byte, error)
	SetBasicAuth(string, string) Request
	RequestContext() context.Context
	WithRequestContext(context.Context) Request
	GetRequest() Request
	GetHttpRequest() *http.Request
	Cookies() []*http.Cookie
	AddCookie(c *http.Cookie) Request
	Yield() (err error)
	MustYield()
	UnsafeYield()
}

type CallBack func(Context, Response) error
type ErrBack func(Context, Response, error)
type StartFunc func(Context, string) error

type RequestStatus uint8

const (
	RequestStatusUnknown = iota
	RequestStatusPending
	RequestStatusRunning
	RequestStatusSuccess
	RequestStatusFailure
)

func (s RequestStatus) String() string {
	switch s {
	case RequestStatusPending:
		return "pending"
	case RequestStatusRunning:
		return "running"
	case RequestStatusSuccess:
		return "success"
	case RequestStatusFailure:
		return "failure"
	default:
		return "unknown"
	}
}

type RequestOption func(Request)

func WithUrl(url string) RequestOption {
	return func(request Request) {
		request.SetUrl(url)
	}
}
func WithMethod(method string) RequestOption {
	return func(request Request) {
		request.SetMethod(method)
	}
}
