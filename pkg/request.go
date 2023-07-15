package pkg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	http.Request
	bodyStr            string
	uniqueKey          string
	callBack           Callback
	errBack            Errback
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
	platform           []Platform
	browser            []Browser
	file               bool
	image              bool
	extra              string
	extraName          string
	errors             map[string]error
}

func (r *Request) String() string {
	t := reflect.TypeOf(r).Elem()
	v := reflect.ValueOf(r).Elem()
	l := t.NumField()
	var out []string
	for i := 0; i < l; i++ {
		var value string
		vv := v.Field(i)
		if vv.Kind() == reflect.Ptr {
			if vv.IsNil() {
				continue
			}
			vv = vv.Elem()
		} else {
			if vv.IsZero() {
				continue
			}
		}
		switch vv.Kind() {
		case reflect.String:
			value = vv.Interface().(string)
		case reflect.Bool:
			value = strconv.FormatBool(vv.Interface().(bool))
		case reflect.Uint8:
			value = fmt.Sprintf("%d", vv.Interface().(uint8))
		default:
		}
		out = append(out, fmt.Sprintf("%s: %s", t.Field(i).Name, value))
	}
	return fmt.Sprintf(`{%s}`, strings.Join(out, ", "))
}
func (r *Request) SetUniqueKey(uniqueKey string) *Request {
	r.uniqueKey = uniqueKey
	return r
}
func (r *Request) GetUniqueKey() string {
	return r.uniqueKey
}
func (r *Request) SetCallback(callback Callback) *Request {
	r.callBack = callback
	return r
}
func (r *Request) GetCallback() Callback {
	return r.callBack
}
func (r *Request) SetErrback(errback Errback) *Request {
	r.errBack = errback
	return r
}
func (r *Request) GetErrback() Errback {
	return r.errBack
}
func (r *Request) SetReferer(referer string) *Request {
	r.referer = referer
	return r
}
func (r *Request) GetReferer() string {
	return r.referer
}
func (r *Request) SetUsername(username string) *Request {
	r.username = username
	return r
}
func (r *Request) GetUsername() string {
	return r.username
}
func (r *Request) SetPassword(password string) *Request {
	r.password = password
	return r
}
func (r *Request) GetPassword() string {
	return r.password
}
func (r *Request) SetChecksum(checksum string) *Request {
	r.checksum = checksum
	return r
}
func (r *Request) GetChecksum() string {
	return r.checksum
}
func (r *Request) SetCreateTime(createTime string) *Request {
	r.createTime = createTime
	return r
}
func (r *Request) GetCreateTime() string {
	return r.createTime
}
func (r *Request) SetSpendTime(spendTime time.Duration) *Request {
	r.spendTime = spendTime
	return r
}
func (r *Request) GetSpendTime() time.Duration {
	return r.spendTime
}
func (r *Request) SetInterval(interval time.Duration) *Request {
	r.interval = interval
	return r
}
func (r *Request) GetInterval() time.Duration {
	return r.interval
}
func (r *Request) SetTimeout(timeout time.Duration) *Request {
	r.timeout = timeout
	return r
}
func (r *Request) GetTimeout() time.Duration {
	return r.timeout
}
func (r *Request) SetSkipMiddleware(skipMiddleware bool) *Request {
	r.skipMiddleware = skipMiddleware
	return r
}
func (r *Request) GetSkipMiddleware() bool {
	return r.skipMiddleware
}
func (r *Request) SetSkipFilter(skipFilter *bool) *Request {
	r.skipFilter = skipFilter
	return r
}
func (r *Request) GetSkipFilter() *bool {
	return r.skipFilter
}
func (r *Request) SetCanonicalHeaderKey(canonicalHeaderKey *bool) *Request {
	r.canonicalHeaderKey = canonicalHeaderKey
	return r
}
func (r *Request) GetCanonicalHeaderKey() *bool {
	return r.canonicalHeaderKey
}
func (r *Request) SetProxyEnable(proxyEnable *bool) *Request {
	r.proxyEnable = proxyEnable
	return r
}
func (r *Request) GetProxyEnable() *bool {
	return r.proxyEnable
}
func (r *Request) SetProxy(proxy *url.URL) *Request {
	r.proxy = proxy
	return r
}
func (r *Request) GetProxy() *url.URL {
	return r.proxy
}
func (r *Request) SetRetryMaxTimes(retryMaxTimes *uint8) *Request {
	r.retryMaxTimes = retryMaxTimes
	return r
}
func (r *Request) GetRetryMaxTimes() *uint8 {
	return r.retryMaxTimes
}
func (r *Request) SetRetryTimes(retryTimes uint8) *Request {
	r.retryTimes = retryTimes
	return r
}
func (r *Request) GetRetryTimes() uint8 {
	return r.retryTimes
}
func (r *Request) SetRedirectMaxTimes(redirectMaxTimes *uint8) *Request {
	r.redirectMaxTimes = redirectMaxTimes
	return r
}
func (r *Request) GetRedirectMaxTimes() *uint8 {
	return r.redirectMaxTimes
}
func (r *Request) SetRedirectTimes(redirectTimes uint8) *Request {
	r.redirectTimes = redirectTimes
	return r
}
func (r *Request) GetRedirectTimes() uint8 {
	return r.redirectTimes
}
func (r *Request) SetOkHttpCodes(okHttpCodes []int) *Request {
	r.okHttpCodes = okHttpCodes
	return r
}
func (r *Request) GetOkHttpCodes() []int {
	return r.okHttpCodes
}
func (r *Request) SetSlot(slot string) *Request {
	r.slot = slot
	return r
}
func (r *Request) GetSlot() string {
	return r.slot
}
func (r *Request) SetHttpProto(httpProto string) *Request {
	r.httpProto = httpProto
	return r
}
func (r *Request) GetHttpProto() string {
	return r.httpProto
}
func (r *Request) SetPlatform(platform []Platform) *Request {
	r.platform = platform
	return r
}
func (r *Request) GetPlatform() []Platform {
	return r.platform
}
func (r *Request) SetBrowser(browser []Browser) *Request {
	r.browser = browser
	return r
}
func (r *Request) GetBrowser() []Browser {
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
func (r *Request) SetUrl(Url string) *Request {
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
func (r *Request) AddQuery(key string, value string) *Request {
	r.URL.Query().Add(key, value)
	return r
}
func (r *Request) SetQuery(key string, value string) *Request {
	r.URL.Query().Set(key, value)
	return r
}
func (r *Request) GetQuery(key string) *Request {
	r.URL.Query().Get(key)
	return r
}
func (r *Request) DelQuery(key string) *Request {
	r.URL.Query().Del(key)
	return r
}
func (r *Request) HasQuery(key string) *Request {
	r.URL.Query().Has(key)
	return r
}
func (r *Request) SetForm(key string, value string) *Request {
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
func (r *Request) SetPostForm(key string, value string) *Request {
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
func (r *Request) SetMethod(method string) *Request {
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
func (r *Request) SetBody(bodyStr string) *Request {
	r.bodyStr = bodyStr
	r.Body = io.NopCloser(strings.NewReader(bodyStr))
	return r
}
func (r *Request) GetBody() string {
	return r.bodyStr
}
func (r *Request) SetHeader(key string, value string) *Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)

	return r
}
func (r *Request) GetHeader() http.Header {
	return r.Header
}

func (r *Request) SetFile(file bool) *Request {
	r.file = file
	return r
}
func (r *Request) GetFile() bool {
	return r.file
}
func (r *Request) SetImage(image bool) *Request {
	r.image = image
	return r
}
func (r *Request) GetImage() bool {
	return r.image
}
func (r *Request) SetConcurrency(concurrency *uint8) {
	r.concurrency = concurrency
}
func (r *Request) GetConcurrency() *uint8 {
	return r.concurrency
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
func (r *Request) SetExtra(extra any) *Request {
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

func (r *Request) ToRequestJson() (request *RequestJson, err error) {
	var Url string
	if r.URL != nil {
		Url = r.URL.String()
	}
	var proxy string
	if r.proxy != nil {
		proxy = r.proxy.String()
	}
	var callBack string
	if r.GetCallback() != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.GetCallback()).Pointer()).Name()
		callBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	}
	var errBack string
	if r.GetErrback() != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.GetErrback()).Pointer()).Name()
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
	callbacks          map[string]Callback
	errbacks           map[string]Errback
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

func (r *RequestJson) SetCallbacks(callbacks map[string]Callback) {
	r.callbacks = callbacks
}
func (r *RequestJson) SetErrbacks(errbacks map[string]Errback) {
	r.errbacks = errbacks
}
func (r *RequestJson) ToRequest() (request *Request, err error) {
	req, err := http.NewRequest(r.Method, r.Url, strings.NewReader(r.BodyStr))
	if err != nil {
		return
	}
	req.Header = r.Header

	proxy, err := url.Parse(r.Proxy)
	if err != nil {
		return
	}

	var platform []Platform
	if len(r.Platform) > 0 {
		for _, v := range r.Platform {
			platform = append(platform, Platform(v))
		}
	}
	var browser []Browser
	if len(r.Browser) > 0 {
		for _, v := range r.Browser {
			browser = append(browser, Browser(v))
		}
	}

	request = &Request{
		Request:            *req,
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

type Callback func(context.Context, *Response) error
type Errback func(context.Context, *Response, error)
