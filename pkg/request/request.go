package request

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Request struct {
	*http.Request
	bodyStr            string
	uniqueKey          string
	callBack           pkg.CallBack
	errBack            pkg.ErrBack
	referer            string
	username           string
	password           string
	checksum           string
	createTime         string
	spendTime          time.Duration
	skipMiddleware     bool
	skipFilter         *bool
	canonicalHeaderKey *bool
	proxyEnable        *bool
	proxy              *url.URL
	retryMaxTimes      *uint8
	retryTimes         uint8
	redirectMaxTimes   *uint8
	redirectTimes      uint8
	okHttpCodes        []int
	slot               string
	concurrency        *uint8
	interval           time.Duration
	timeout            time.Duration
	httpProto          string
	platform           []pkg.Platform
	browser            []pkg.Browser
	file               bool
	image              bool
	extra              string
	extraName          string
	errors             map[string]error
}

func (r *Request) SetUniqueKey(uniqueKey string) pkg.Request {
	r.uniqueKey = uniqueKey
	return r
}
func (r *Request) GetUniqueKey() string {
	return r.uniqueKey
}
func (r *Request) SetCallBack(callback pkg.CallBack) pkg.Request {
	r.callBack = callback
	return r
}
func (r *Request) GetCallBack() pkg.CallBack {
	return r.callBack
}
func (r *Request) SetErrBack(errback pkg.ErrBack) pkg.Request {
	r.errBack = errback
	return r
}
func (r *Request) GetErrBack() pkg.ErrBack {
	return r.errBack
}
func (r *Request) SetReferer(referer string) pkg.Request {
	r.referer = referer
	return r
}
func (r *Request) GetReferer() string {
	return r.referer
}
func (r *Request) SetUsername(username string) pkg.Request {
	r.username = username
	return r
}
func (r *Request) GetUsername() string {
	return r.username
}
func (r *Request) SetPassword(password string) pkg.Request {
	r.password = password
	return r
}
func (r *Request) GetPassword() string {
	return r.password
}
func (r *Request) SetChecksum(checksum string) pkg.Request {
	r.checksum = checksum
	return r
}
func (r *Request) GetChecksum() string {
	return r.checksum
}
func (r *Request) SetCreateTime(createTime string) pkg.Request {
	r.createTime = createTime
	return r
}
func (r *Request) GetCreateTime() string {
	return r.createTime
}
func (r *Request) SetSpendTime(spendTime time.Duration) pkg.Request {
	r.spendTime = spendTime
	return r
}
func (r *Request) GetSpendTime() time.Duration {
	return r.spendTime
}
func (r *Request) SetSkipMiddleware(skipMiddleware bool) pkg.Request {
	r.skipMiddleware = skipMiddleware
	return r
}
func (r *Request) GetSkipMiddleware() bool {
	return r.skipMiddleware
}
func (r *Request) SetSkipFilter(skipFilter *bool) pkg.Request {
	r.skipFilter = skipFilter
	return r
}
func (r *Request) GetSkipFilter() *bool {
	return r.skipFilter
}
func (r *Request) SetCanonicalHeaderKey(canonicalHeaderKey *bool) pkg.Request {
	r.canonicalHeaderKey = canonicalHeaderKey
	return r
}
func (r *Request) GetCanonicalHeaderKey() *bool {
	return r.canonicalHeaderKey
}
func (r *Request) SetProxyEnable(proxyEnable *bool) pkg.Request {
	r.proxyEnable = proxyEnable
	return r
}
func (r *Request) GetProxyEnable() *bool {
	return r.proxyEnable
}
func (r *Request) SetProxy(proxy *url.URL) pkg.Request {
	r.proxy = proxy
	return r
}
func (r *Request) GetProxy() *url.URL {
	return r.proxy
}
func (r *Request) SetRetryMaxTimes(retryMaxTimes *uint8) pkg.Request {
	r.retryMaxTimes = retryMaxTimes
	return r
}
func (r *Request) GetRetryMaxTimes() *uint8 {
	return r.retryMaxTimes
}
func (r *Request) SetRetryTimes(retryTimes uint8) pkg.Request {
	r.retryTimes = retryTimes
	return r
}
func (r *Request) GetRetryTimes() uint8 {
	return r.retryTimes
}
func (r *Request) SetRedirectMaxTimes(redirectMaxTimes *uint8) pkg.Request {
	r.redirectMaxTimes = redirectMaxTimes
	return r
}
func (r *Request) GetRedirectMaxTimes() *uint8 {
	return r.redirectMaxTimes
}
func (r *Request) SetRedirectTimes(redirectTimes uint8) pkg.Request {
	r.redirectTimes = redirectTimes
	return r
}
func (r *Request) GetRedirectTimes() uint8 {
	return r.redirectTimes
}
func (r *Request) SetOkHttpCodes(okHttpCodes []int) pkg.Request {
	r.okHttpCodes = okHttpCodes
	return r
}
func (r *Request) GetOkHttpCodes() []int {
	return r.okHttpCodes
}
func (r *Request) SetSlot(slot string) pkg.Request {
	r.slot = slot
	return r
}
func (r *Request) GetSlot() string {
	return r.slot
}
func (r *Request) SetConcurrency(concurrency *uint8) pkg.Request {
	r.concurrency = concurrency
	return r
}
func (r *Request) GetConcurrency() *uint8 {
	return r.concurrency
}
func (r *Request) SetInterval(interval time.Duration) pkg.Request {
	r.interval = interval
	return r
}
func (r *Request) GetInterval() time.Duration {
	return r.interval
}
func (r *Request) SetTimeout(timeout time.Duration) pkg.Request {
	r.timeout = timeout
	return r
}
func (r *Request) GetTimeout() time.Duration {
	return r.timeout
}
func (r *Request) SetHttpProto(httpProto string) pkg.Request {
	r.httpProto = httpProto
	return r
}
func (r *Request) GetHttpProto() string {
	return r.httpProto
}
func (r *Request) SetPlatform(platform []pkg.Platform) pkg.Request {
	r.platform = platform
	return r
}
func (r *Request) GetPlatform() []pkg.Platform {
	return r.platform
}
func (r *Request) SetBrowser(browser []pkg.Browser) pkg.Request {
	r.browser = browser
	return r
}
func (r *Request) GetBrowser() []pkg.Browser {
	return r.browser
}
func (r *Request) setExtraName(name string) {
	r.extraName = name
}
func (r *Request) GetExtraName() string {
	return r.extraName
}
func (r *Request) setErr(key string, value error) {
	if r.errors == nil {
		r.errors = make(map[string]error)
	}
	r.errors[key] = value
}
func (r *Request) GetErr() map[string]error {
	return r.errors
}
func (r *Request) delErr(key string) {
	delete(r.errors, key)
}
func (r *Request) SetUrl(Url string) pkg.Request {
	URL, err := url.Parse(Url)
	if err != nil {
		r.setErr("Url", err)
		return r
	}
	r.URL = URL
	return r
}
func (r *Request) GetUrl() string {
	if r.URL == nil {
		return ""
	}
	return r.URL.String()
}
func (r *Request) GetURL() *url.URL {
	return r.URL
}
func (r *Request) AddQuery(key string, value string) pkg.Request {
	r.URL.Query().Add(key, value)
	return r
}
func (r *Request) SetQuery(key string, value string) pkg.Request {
	r.URL.Query().Set(key, value)
	return r
}
func (r *Request) GetQuery(key string) pkg.Request {
	r.URL.Query().Get(key)
	return r
}
func (r *Request) DelQuery(key string) pkg.Request {
	r.URL.Query().Del(key)
	return r
}
func (r *Request) HasQuery(key string) pkg.Request {
	r.URL.Query().Has(key)
	return r
}
func (r *Request) SetForm(key string, value string) pkg.Request {
	if r.Form == nil {
		r.Form = make(url.Values)
	}
	r.Form.Add(key, value)
	err := r.ParseForm()
	if err != nil {
		r.setErr("Form", err)
		return r
	}
	return r
}
func (r *Request) GetForm() url.Values {
	return r.Form
}
func (r *Request) SetPostForm(key string, value string) pkg.Request {
	if r.PostForm == nil {
		r.PostForm = make(url.Values)
	}
	r.PostForm.Add(key, value)
	err := r.ParseForm()
	if err != nil {
		r.setErr("PostForm", err)
		return r
	}
	return r
}
func (r *Request) GetPostForm() url.Values {
	return r.PostForm
}
func (r *Request) SetMethod(method string) pkg.Request {
	method = strings.ToUpper(method)
	ok := false
	for _, v := range []string{"OPTIONS", "GET", "HEAD", "POST", "PUT", "DELETE", "TRACE"} {
		if v == method {
			ok = true
			break
		}
	}
	if ok {
		r.Method = method
	} else {
		r.setErr("Method", errors.New("method error"))
	}
	return r
}
func (r *Request) GetMethod() string {
	return r.Method
}
func (r *Request) SetBody(bodyStr string) pkg.Request {
	r.bodyStr = bodyStr
	r.Body = io.NopCloser(strings.NewReader(bodyStr))
	return r
}
func (r *Request) GetBody() string {
	return r.bodyStr
}
func (r *Request) SetHeader(key string, value string) pkg.Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)

	return r
}
func (r *Request) GetHeaders() http.Header {
	return r.Header
}
func (r *Request) GetHeader(key string) string {
	return r.Header.Get(key)
}
func (r *Request) SetFile(file bool) pkg.Request {
	r.file = file
	return r
}
func (r *Request) GetFile() bool {
	return r.file
}
func (r *Request) SetImage(image bool) pkg.Request {
	r.image = image
	return r
}
func (r *Request) GetImage() bool {
	return r.image
}
func (r *Request) SetBasicAuth(username string, password string) pkg.Request {
	r.Request.SetBasicAuth(username, password)
	return r
}
func (r *Request) Context() context.Context {
	return r.Request.Context()
}
func (r *Request) WithContext(ctx context.Context) pkg.Request {
	r.Request = r.Request.WithContext(ctx)
	return r
}
func (r *Request) GetRequest() *http.Request {
	return r.Request
}
func (r *Request) GetCookies() []*http.Cookie {
	return r.Request.Cookies()
}
func (r *Request) AddCookie(c *http.Cookie) pkg.Request {
	r.Request.AddCookie(c)
	return r
}
func (r *Request) SetExtra(extra any) pkg.Request {
	extraValue := reflect.ValueOf(extra)
	if extraValue.Kind() != reflect.Ptr || extraValue.IsNil() {
		r.setErr("Extra", errors.New("extra must be a non-null pointer"))
		return r
	}
	r.setExtraName(extraValue.Elem().Type().Name())
	bs, err := json.Marshal(extra)
	if err != nil {
		r.setErr("Extra", err)
		return r
	}
	r.extra = string(bs)
	return r
}
func (r *Request) GetExtra() string {
	return r.extra
}
func (r *Request) UnmarshalExtra(v any) (err error) {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		return fmt.Errorf("v must be a non-null pointer")
	}

	if r.extra == "" {
		return
	}

	err = json.Unmarshal([]byte(r.extra), v)
	return
}
func (r *Request) ToRequestJson() (request pkg.RequestJson, err error) {
	var Url string
	if r.URL != nil {
		Url = r.URL.String()
	}
	var proxy string
	if r.proxy != nil {
		proxy = r.proxy.String()
	}
	var callBack string
	if r.GetCallBack() != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.GetCallBack()).Pointer()).Name()
		callBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	}
	var errBack string
	if r.GetErrBack() != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.GetErrBack()).Pointer()).Name()
		errBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	}
	var platform []string
	if len(r.GetPlatform()) > 0 {
		for _, v := range r.GetPlatform() {
			platform = append(platform, string(v))
		}
	}
	var browser []string
	if len(r.GetBrowser()) > 0 {
		for _, v := range r.GetBrowser() {
			browser = append(browser, string(v))
		}
	}

	request = &RequestJson{
		Url:              Url,
		Method:           r.Method,
		BodyStr:          r.GetBody(),
		Header:           r.Header,
		UniqueKey:        r.GetUniqueKey(),
		CallBack:         callBack,
		ErrBack:          errBack,
		Referer:          r.GetReferer(),
		Username:         r.GetUsername(),
		Password:         r.GetPassword(),
		Checksum:         r.GetChecksum(),
		CreateTime:       r.GetCreateTime(),
		SpendTime:        uint(r.GetSpendTime()),
		SkipMiddleware:   r.GetSkipMiddleware(),
		SkipFilter:       r.GetSkipFilter(),
		ProxyEnable:      r.GetProxyEnable(),
		Proxy:            proxy,
		RetryMaxTimes:    r.GetRetryMaxTimes(),
		RetryTimes:       r.GetRetryTimes(),
		RedirectMaxTimes: r.GetRedirectMaxTimes(),
		RedirectTimes:    r.GetRedirectTimes(),
		OkHttpCodes:      r.GetOkHttpCodes(),
		Slot:             r.GetSlot(),
		Concurrency:      r.GetConcurrency(),
		Interval:         int(r.GetInterval()),
		Timeout:          int(r.GetTimeout()),
		HttpProto:        r.GetHttpProto(),
		Platform:         platform,
		Browser:          browser,
		Image:            r.GetImage(),
		File:             r.GetFile(),
		Extra:            r.GetExtra(),
		ExtraName:        r.GetExtraName(),
	}
	return
}

func (r *Request) Marshal() ([]byte, error) {
	request, err := r.ToRequestJson()
	if err != nil {
		return nil, err
	}
	return json.Marshal(request)
}

type RequestJson struct {
	callbacks          map[string]pkg.CallBack
	errbacks           map[string]pkg.ErrBack
	Method             string              `json:"method,omitempty"`
	Url                string              `json:"url,omitempty"`
	BodyStr            string              `json:"body,omitempty"`
	Header             map[string][]string `json:"header,omitempty"`
	UniqueKey          string              `json:"unique_key,omitempty"` // for filter
	CallBack           string              `json:"call_back,omitempty"`
	ErrBack            string              `json:"err_back,omitempty"`
	Referer            string              `json:"referer,omitempty"`
	Username           string              `json:"username,omitempty"`
	Password           string              `json:"password,omitempty"`
	Checksum           string              `json:"checksum,omitempty"`
	CreateTime         string              `json:"create_time,omitempty"` //create time
	SpendTime          uint                `json:"spend_time,omitempty"`
	SkipMiddleware     bool                `json:"skipMiddleware,omitempty"`
	SkipFilter         *bool               `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey *bool               `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        *bool               `json:"proxy_enable,omitempty"`
	Proxy              string              `json:"proxy,omitempty"`
	RetryMaxTimes      *uint8              `json:"retry_max_times,omitempty"`
	RetryTimes         uint8               `json:"retry_times,omitempty"`
	RedirectMaxTimes   *uint8              `json:"redirect_max_times,omitempty"`
	RedirectTimes      uint8               `json:"redirect_times,omitempty"`
	OkHttpCodes        []int               `json:"ok_http_codes,omitempty"`
	Slot               string              `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        *uint8              `json:"concurrency,omitempty"`
	Interval           int                 `json:"interval,omitempty"`
	Timeout            int                 `json:"timeout,omitempty"`    //seconds
	HttpProto          string              `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Platform           []string            `json:"platform,omitempty"`
	Browser            []string            `json:"browser,omitempty"`
	File               bool                `json:"file,omitempty"`
	Image              bool                `json:"image,omitempty"`
	Extra              string              `json:"extra,omitempty"`
	ExtraName          string              `json:"extra_name,omitempty"`
}

func (r *RequestJson) SetCallBacks(callbacks map[string]pkg.CallBack) {
	r.callbacks = callbacks
}
func (r *RequestJson) SetErrBacks(errbacks map[string]pkg.ErrBack) {
	r.errbacks = errbacks
}
func (r *RequestJson) ToRequest() (request pkg.Request, err error) {
	req, err := http.NewRequest(r.Method, r.Url, strings.NewReader(r.BodyStr))
	if err != nil {
		return
	}
	req.Header = r.Header

	proxy, err := url.Parse(r.Proxy)
	if err != nil {
		return
	}

	var platform []pkg.Platform
	if len(r.Platform) > 0 {
		for _, v := range r.Platform {
			platform = append(platform, pkg.Platform(v))
		}
	}
	var browser []pkg.Browser
	if len(r.Browser) > 0 {
		for _, v := range r.Browser {
			browser = append(browser, pkg.Browser(v))
		}
	}

	request = &Request{
		Request:            req,
		bodyStr:            r.BodyStr,
		uniqueKey:          r.UniqueKey,
		callBack:           r.callbacks[r.CallBack],
		errBack:            r.errbacks[r.ErrBack],
		referer:            r.Referer,
		username:           r.Username,
		password:           r.Password,
		checksum:           r.Checksum,
		createTime:         r.CreateTime,
		spendTime:          time.Duration(r.SpendTime),
		skipMiddleware:     r.SkipMiddleware,
		skipFilter:         r.SkipFilter,
		canonicalHeaderKey: r.CanonicalHeaderKey,
		proxyEnable:        r.ProxyEnable,
		proxy:              proxy,
		retryMaxTimes:      r.RetryMaxTimes,
		retryTimes:         r.RetryTimes,
		redirectMaxTimes:   r.RedirectMaxTimes,
		redirectTimes:      r.RedirectTimes,
		okHttpCodes:        r.OkHttpCodes,
		slot:               r.Slot,
		concurrency:        r.Concurrency,
		interval:           time.Duration(r.Interval),
		timeout:            time.Duration(r.Timeout),
		httpProto:          r.HttpProto,
		platform:           platform,
		browser:            browser,
		file:               r.File,
		image:              r.Image,
		extra:              r.Extra,
		extraName:          r.ExtraName,
	}

	return
}

func NewRequest() pkg.Request {
	request := new(Request)
	request.Request = new(http.Request)
	return request
}
