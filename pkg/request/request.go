package request

import (
	"bytes"
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
	callBack           string
	errBack            string
	referrer           string
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
	platforms          []pkg.Platform
	browsers           []pkg.Browser
	isFile             bool
	fileOptions        *pkg.FileOptions
	isImage            bool
	imageOptions       *pkg.ImageOptions
	extra              string
	extraName          string
	errors             map[string]error
	priority           uint8
	fingerprint        string
	client             pkg.Client
	ajax               bool
}

func (r *Request) UniqueKey() string {
	return r.uniqueKey
}
func (r *Request) SetUniqueKey(uniqueKey string) pkg.Request {
	r.uniqueKey = uniqueKey
	return r
}
func (r *Request) SetCallBack(callBack pkg.CallBack) pkg.Request {
	name := runtime.FuncForPC(reflect.ValueOf(callBack).Pointer()).Name()
	r.callBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	return r
}
func (r *Request) CallBack() string {
	return r.callBack
}
func (r *Request) SetErrBack(errBack pkg.ErrBack) pkg.Request {
	name := runtime.FuncForPC(reflect.ValueOf(errBack).Pointer()).Name()
	r.errBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	return r
}
func (r *Request) ErrBack() string {
	return r.errBack
}
func (r *Request) Referrer() string {
	return r.referrer
}
func (r *Request) SetReferrer(referrer string) pkg.Request {
	r.referrer = referrer
	return r
}
func (r *Request) SetUsername(username string) pkg.Request {
	r.username = username
	return r
}
func (r *Request) Username() string {
	return r.username
}
func (r *Request) SetPassword(password string) pkg.Request {
	r.password = password
	return r
}
func (r *Request) Password() string {
	return r.password
}
func (r *Request) Checksum() string {
	return r.checksum
}
func (r *Request) SetChecksum(checksum string) pkg.Request {
	r.checksum = checksum
	return r
}
func (r *Request) SetCreateTime(createTime string) pkg.Request {
	r.createTime = createTime
	return r
}
func (r *Request) CreateTime() string {
	return r.createTime
}
func (r *Request) SetSpendTime(spendTime time.Duration) pkg.Request {
	r.spendTime = spendTime
	return r
}
func (r *Request) SpendTime() time.Duration {
	return r.spendTime
}
func (r *Request) SkipMiddleware() bool {
	return r.skipMiddleware
}
func (r *Request) SetSkipMiddleware(skipMiddleware bool) pkg.Request {
	r.skipMiddleware = skipMiddleware
	return r
}
func (r *Request) SkipFilter() *bool {
	return r.skipFilter
}
func (r *Request) SetSkipFilter(skipFilter *bool) pkg.Request {
	r.skipFilter = skipFilter
	return r
}
func (r *Request) SetCanonicalHeaderKey(canonicalHeaderKey *bool) pkg.Request {
	r.canonicalHeaderKey = canonicalHeaderKey
	return r
}
func (r *Request) CanonicalHeaderKey() *bool {
	return r.canonicalHeaderKey
}
func (r *Request) SetProxyEnable(proxyEnable bool) pkg.Request {
	r.proxyEnable = &proxyEnable
	return r
}
func (r *Request) ProxyEnable() *bool {
	return r.proxyEnable
}
func (r *Request) SetProxy(proxy string) pkg.Request {
	var err error
	r.proxy, err = url.Parse(proxy)
	if err == nil {
		r.SetProxyEnable(true)
	}
	return r
}
func (r *Request) Proxy() *url.URL {
	return r.proxy
}
func (r *Request) SetRetryMaxTimes(retryMaxTimes *uint8) pkg.Request {
	r.retryMaxTimes = retryMaxTimes
	return r
}
func (r *Request) RetryMaxTimes() *uint8 {
	return r.retryMaxTimes
}
func (r *Request) SetRetryTimes(retryTimes uint8) pkg.Request {
	r.retryTimes = retryTimes
	return r
}
func (r *Request) RetryTimes() uint8 {
	return r.retryTimes
}
func (r *Request) SetRedirectMaxTimes(redirectMaxTimes *uint8) pkg.Request {
	r.redirectMaxTimes = redirectMaxTimes
	return r
}
func (r *Request) RedirectMaxTimes() *uint8 {
	return r.redirectMaxTimes
}
func (r *Request) SetRedirectTimes(redirectTimes uint8) pkg.Request {
	r.redirectTimes = redirectTimes
	return r
}
func (r *Request) RedirectTimes() uint8 {
	return r.redirectTimes
}
func (r *Request) SetOkHttpCodes(okHttpCodes []int) pkg.Request {
	r.okHttpCodes = okHttpCodes
	return r
}
func (r *Request) OkHttpCodes() []int {
	return r.okHttpCodes
}
func (r *Request) SetSlot(slot string) pkg.Request {
	r.slot = slot
	return r
}
func (r *Request) Slot() string {
	return r.slot
}
func (r *Request) SetConcurrency(concurrency *uint8) pkg.Request {
	r.concurrency = concurrency
	return r
}
func (r *Request) Concurrency() *uint8 {
	return r.concurrency
}
func (r *Request) SetInterval(interval time.Duration) pkg.Request {
	r.interval = interval
	return r
}
func (r *Request) Interval() time.Duration {
	return r.interval
}
func (r *Request) Timeout() time.Duration {
	return r.timeout
}
func (r *Request) SetTimeout(timeout time.Duration) pkg.Request {
	r.timeout = timeout
	return r
}
func (r *Request) HttpProto() string {
	return r.httpProto
}
func (r *Request) SetHttpProto(httpProto string) pkg.Request {
	r.httpProto = httpProto
	return r
}
func (r *Request) Platforms() []pkg.Platform {
	return r.platforms
}
func (r *Request) SetPlatforms(platforms ...pkg.Platform) pkg.Request {
	r.platforms = platforms
	return r
}
func (r *Request) Browsers() []pkg.Browser {
	return r.browsers
}
func (r *Request) SetBrowsers(browsers ...pkg.Browser) pkg.Request {
	r.browsers = browsers
	return r
}
func (r *Request) setExtraName(name string) {
	r.extraName = name
}
func (r *Request) GetExtraName() string {
	return r.extraName
}
func (r *Request) Priority() uint8 {
	return r.priority
}
func (r *Request) SetPriority(priority uint8) pkg.Request {
	r.priority = priority
	return r
}
func (r *Request) Fingerprint() string {
	return r.fingerprint
}
func (r *Request) SetFingerprint(fingerprint string) pkg.Request {
	r.fingerprint = fingerprint
	return r
}
func (r *Request) Client() pkg.Client {
	return r.client
}
func (r *Request) SetClient(client pkg.Client) pkg.Request {
	r.client = client
	return r
}
func (r *Request) Ajax() bool {
	return r.ajax
}
func (r *Request) SetAjax(ajax bool) pkg.Request {
	r.ajax = ajax
	return r
}
func (r *Request) setErr(key string, value error) {
	if r.errors == nil {
		r.errors = make(map[string]error)
	}
	r.errors[key] = value
}
func (r *Request) Err() map[string]error {
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
func (r *Request) Query(key string) string {
	return r.URL.Query().Get(key)
}
func (r *Request) AddQuery(key string, value string) pkg.Request {
	query := r.URL.Query()
	query.Add(key, value)
	r.URL.RawQuery = query.Encode()
	return r
}
func (r *Request) SetQuery(key string, value string) pkg.Request {
	query := r.URL.Query()
	query.Set(key, value)
	r.URL.RawQuery = query.Encode()
	return r
}
func (r *Request) DelQuery(key string) pkg.Request {
	query := r.URL.Query()
	query.Del(key)
	r.URL.RawQuery = query.Encode()
	return r
}
func (r *Request) HasQuery(key string) bool {
	return r.URL.Query().Has(key)
}
func (r *Request) SetForm(key string, value string) pkg.Request {
	if r.Form == nil {
		r.Form = make(url.Values)
	}

	if r.URL != nil {
		newValues, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			r.setErr("Form", err)
			return r
		}
		if newValues != nil {
			for k, v := range newValues {
				r.Form[k] = v
			}
		}
	}

	r.Form.Set(key, value)
	r.Request.URL.RawQuery = r.Form.Encode()
	return r
}
func (r *Request) GetForm() url.Values {
	return r.Form
}
func (r *Request) SetPostForm(key string, value string) pkg.Request {
	if r.PostForm == nil {
		r.PostForm = make(url.Values)
	}
	if r.bodyStr != "" {
		newValues, err := url.ParseQuery(r.bodyStr)
		if err != nil {
			r.setErr("PostForm", err)
			return r
		}
		if newValues != nil {
			for k, v := range newValues {
				r.PostForm[k] = v
			}
		}
	}

	r.PostForm.Set(key, value)
	r.SetBodyStr(r.PostForm.Encode())

	return r
}
func (r *Request) GetPostForm() url.Values {
	return r.PostForm
}
func (r *Request) SetMethod(method string) pkg.Request {
	method = strings.ToUpper(method)
	ok := false
	for _, v := range []string{
		http.MethodGet,
		http.MethodHead,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		//http.MethodConnect,
		http.MethodOptions,
		http.MethodTrace,
	} {
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
func (r *Request) BodyStr() string {
	return r.bodyStr
}
func (r *Request) SetBodyStr(bodyStr string) pkg.Request {
	r.bodyStr = bodyStr
	r.Body = io.NopCloser(strings.NewReader(bodyStr))
	return r
}
func (r *Request) BodyBytes() []byte {
	return []byte(r.bodyStr)
}
func (r *Request) SetBodyBytes(bodyBytes []byte) pkg.Request {
	r.bodyStr = string(bodyBytes)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return r
}
func (r *Request) SetBodyJson(bodyJson any) pkg.Request {
	bodyBytes, err := json.Marshal(bodyJson)
	if err != nil {
		r.setErr("body", err)
		return r
	}
	r.bodyStr = string(bodyBytes)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	r.SetHeader("content-type", "application/json")
	return r
}
func (r *Request) GetHeader(key string) string {
	return r.Header.Get(key)
}
func (r *Request) SetHeader(key string, value string) pkg.Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(http.CanonicalHeaderKey(key), value)
	return r
}
func (r *Request) Headers() http.Header {
	return r.Header
}
func (r *Request) SetHeaders(headers map[string]string) pkg.Request {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	for key, value := range headers {
		r.Header.Set(http.CanonicalHeaderKey(key), value)
	}
	return r
}
func (r *Request) IsFile() bool {
	return r.isFile
}
func (r *Request) AsFile(isFile bool) pkg.Request {
	r.isFile = isFile
	return r
}
func (r *Request) SetFileOptions(options pkg.FileOptions) pkg.Request {
	r.fileOptions = &options
	return r
}
func (r *Request) FileOptions() *pkg.FileOptions {
	return r.fileOptions
}
func (r *Request) IsImage() bool {
	return r.isImage
}
func (r *Request) AsImage(isImage bool) pkg.Request {
	r.isImage = isImage
	return r
}
func (r *Request) SetImageOptions(options pkg.ImageOptions) pkg.Request {
	r.imageOptions = &options
	return r
}
func (r *Request) ImageOptions() *pkg.ImageOptions {
	return r.imageOptions
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
func (r *Request) Cookies() []*http.Cookie {
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
func (r *Request) Extra() string {
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
func (r *Request) MustUnmarshalExtra(v any) {
	if err := r.UnmarshalExtra(v); err != nil {
		panic(err)
	}
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
	var platform []string
	if len(r.platforms) > 0 {
		for _, v := range r.platforms {
			platform = append(platform, string(v))
		}
	}
	var browser []string
	if len(r.browsers) > 0 {
		for _, v := range r.browsers {
			browser = append(browser, string(v))
		}
	}

	request = &RequestJson{
		Url:              Url,
		Method:           r.Method,
		BodyStr:          r.bodyStr,
		Header:           r.Header,
		UniqueKey:        r.uniqueKey,
		CallBack:         r.callBack,
		ErrBack:          r.errBack,
		Referrer:         r.referrer,
		Username:         r.username,
		Password:         r.password,
		Checksum:         r.checksum,
		CreateTime:       r.createTime,
		SpendTime:        uint(r.spendTime),
		SkipMiddleware:   r.skipMiddleware,
		SkipFilter:       r.skipFilter,
		ProxyEnable:      r.proxyEnable,
		Proxy:            proxy,
		RetryMaxTimes:    r.retryMaxTimes,
		RetryTimes:       r.retryTimes,
		RedirectMaxTimes: r.redirectMaxTimes,
		RedirectTimes:    r.redirectTimes,
		OkHttpCodes:      r.okHttpCodes,
		Slot:             r.slot,
		Concurrency:      r.concurrency,
		Interval:         int(r.interval),
		Timeout:          int(r.timeout),
		HttpProto:        r.httpProto,
		Platform:         platform,
		Browser:          browser,
		IsImage:          r.isImage,
		ImageOptions:     r.imageOptions,
		IsFile:           r.isFile,
		FileOptions:      r.fileOptions,
		Extra:            r.extra,
		ExtraName:        r.extraName,
		Priority:         r.priority,
		Fingerprint:      r.fingerprint,
		Client:           string(r.client),
		Ajax:             r.ajax,
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
	Method             string              `json:"method,omitempty"`
	Url                string              `json:"url,omitempty"`
	BodyStr            string              `json:"body,omitempty"`
	Header             map[string][]string `json:"header,omitempty"`
	UniqueKey          string              `json:"unique_key,omitempty"` // for filter
	CallBack           string              `json:"call_back,omitempty"`
	ErrBack            string              `json:"err_back,omitempty"`
	Referrer           string              `json:"referrer,omitempty"`
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
	IsFile             bool                `json:"is_file,omitempty"`
	FileOptions        *pkg.FileOptions    `json:"file_options,omitempty"`
	IsImage            bool                `json:"is_image,omitempty"`
	ImageOptions       *pkg.ImageOptions   `json:"image_options,omitempty"`
	Extra              string              `json:"extra,omitempty"`
	ExtraName          string              `json:"extra_name,omitempty"`
	Priority           uint8               `json:"priority,omitempty"`
	Fingerprint        string              `json:"fingerprint,omitempty"`
	Client             string              `json:"client,omitempty"`
	Ajax               bool                `json:"ajax,omitempty"`
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

	var platforms []pkg.Platform
	if len(r.Platform) > 0 {
		for _, v := range r.Platform {
			platforms = append(platforms, pkg.Platform(v))
		}
	}
	var browsers []pkg.Browser
	if len(r.Browser) > 0 {
		for _, v := range r.Browser {
			browsers = append(browsers, pkg.Browser(v))
		}
	}

	request = &Request{
		Request:            req,
		bodyStr:            r.BodyStr,
		uniqueKey:          r.UniqueKey,
		callBack:           r.CallBack,
		errBack:            r.ErrBack,
		referrer:           r.Referrer,
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
		platforms:          platforms,
		browsers:           browsers,
		isFile:             r.IsFile,
		fileOptions:        r.FileOptions,
		isImage:            r.IsImage,
		imageOptions:       r.ImageOptions,
		extra:              r.Extra,
		extraName:          r.ExtraName,
		priority:           r.Priority,
		fingerprint:        r.Fingerprint,
		client:             pkg.Client(r.Client),
		ajax:               r.Ajax,
	}

	return
}

func NewRequest() pkg.Request {
	request := new(Request)
	request.Request = new(http.Request)
	return request
}
func Get() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodGet)
	return request
}
func Post() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodPost)
	return request
}
func Head() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodHead)
	return request
}
func Delete() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodDelete)
	return request
}
func Put() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodPut)
	return request
}
func Patch() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodPatch)
	return request
}
func Options() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodOptions)
	return request
}
func Trace() pkg.Request {
	request := NewRequest()
	request.SetMethod(http.MethodTrace)
	return request
}
