package pkg

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type Request interface {
	SetUniqueKey(string) Request
	GetUniqueKey() string
	SetCallBack(CallBack) Request
	GetCallBack() CallBack
	SetErrBack(ErrBack) Request
	GetErrBack() ErrBack
	SetReferrer(string) Request
	GetReferrer() string
	SetUsername(string) Request
	GetUsername() string
	SetPassword(string) Request
	GetPassword() string
	SetChecksum(string) Request
	GetChecksum() string
	SetCreateTime(string) Request
	GetCreateTime() string
	SetSpendTime(time.Duration) Request
	GetSpendTime() time.Duration
	SetSkipMiddleware(bool) Request
	GetSkipMiddleware() bool
	SetSkipFilter(*bool) Request
	GetSkipFilter() *bool
	SetCanonicalHeaderKey(*bool) Request
	GetCanonicalHeaderKey() *bool
	SetProxyEnable(*bool) Request
	GetProxyEnable() *bool
	SetProxy(*url.URL) Request
	GetProxy() *url.URL
	SetRetryMaxTimes(*uint8) Request
	GetRetryMaxTimes() *uint8
	SetRetryTimes(uint8) Request
	GetRetryTimes() uint8
	SetRedirectMaxTimes(*uint8) Request
	GetRedirectMaxTimes() *uint8
	SetRedirectTimes(uint8) Request
	GetRedirectTimes() uint8
	SetOkHttpCodes([]int) Request
	GetOkHttpCodes() []int
	SetSlot(string) Request
	GetSlot() string
	SetConcurrency(*uint8) Request
	GetConcurrency() *uint8
	SetInterval(time.Duration) Request
	GetInterval() time.Duration
	SetTimeout(time.Duration) Request
	GetTimeout() time.Duration
	SetHttpProto(string) Request
	GetHttpProto() string
	SetPlatform([]Platform) Request
	GetPlatform() []Platform
	SetBrowser([]Browser) Request
	GetBrowser() []Browser
	GetExtraName() string
	SetPriority(uint8) Request
	GetPriority() uint8
	SetFingerprint(string) Request
	GetFingerprint() string
	SetClient(Client) Request
	GetClient() Client
	GetErr() map[string]error
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
	SetBody(string) Request
	GetBody() string
	SetHeader(string, string) Request
	GetHeaders() http.Header
	GetHeader(string) string
	SetFile(bool) Request
	GetFile() bool
	SetImage(bool) Request
	GetImage() bool
	SetExtra(any) Request
	GetExtra() string
	UnmarshalExtra(any) error
	ToRequestJson() (RequestJson, error)
	Marshal() ([]byte, error)
	SetBasicAuth(string, string) Request
	Context() context.Context
	WithContext(context.Context) Request
	GetRequest() *http.Request
	GetCookies() []*http.Cookie
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
