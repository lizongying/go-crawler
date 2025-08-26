package request

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	context2 "github.com/lizongying/go-crawler/pkg/context"
	"github.com/lizongying/go-crawler/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strings"
	"time"
)

type Request struct {
	Errors             map[string]error `json:"errors,omitempty"`
	Context            pkg.Context      `json:"context,omitempty"`
	*http.Request      `json:"-"`
	Method             string              `json:"method,omitempty"`
	Url                string              `json:"url,omitempty"`
	Header             map[string][]string `json:"header,omitempty"`
	BodyStr            string              `json:"body_str,omitempty"`
	UniqueKey          string              `json:"unique_key,omitempty"` // for filter
	CallBack           string              `json:"call_back,omitempty"`
	ErrBack            string              `json:"err_back,omitempty"`
	Referrer           string              `json:"referrer,omitempty"`
	Username           string              `json:"username,omitempty"`
	Password           string              `json:"password,omitempty"`
	Checksum           string              `json:"checksum,omitempty"`
	CreateTime         string              `json:"create_time,omitempty"`
	SpendTime          utils.DurationNano  `json:"spend_time,omitempty"`
	SkipMiddleware     bool                `json:"skip_middleware,omitempty"`
	SkipFilter         *bool               `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey *bool               `json:"canonical_header_key,omitempty"` // canonical header key
	ProxyEnable        *bool               `json:"proxy_enable,omitempty"`
	Proxy              utils.Url           `json:"proxy,omitempty"`
	RetryMaxTimes      *uint8              `json:"retry_max_times,omitempty"`
	RetryTimes         uint8               `json:"retry_times,omitempty"`
	RedirectMaxTimes   *uint8              `json:"redirect_max_times,omitempty"`
	RedirectTimes      uint8               `json:"redirect_times,omitempty"`
	OkHttpCodes        []int               `json:"ok_http_codes,omitempty"`
	Slot               string              `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        *uint8              `json:"concurrency,omitempty"`
	Interval           utils.DurationNano  `json:"interval,omitempty"`
	Timeout            utils.DurationNano  `json:"timeout,omitempty"`    // seconds
	HttpProto          string              `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Platforms          []pkg.Platform      `json:"platforms,omitempty"`
	Browsers           []pkg.Browser       `json:"browsers,omitempty"`
	File               bool                `json:"is_file,omitempty"`
	FileOptions        *pkg.FileOptions    `json:"file_options,omitempty"`
	Image              bool                `json:"is_image,omitempty"`
	ImageOptions       *pkg.ImageOptions   `json:"image_options,omitempty"`
	Extra              string              `json:"extra,omitempty"`
	ExtraName          string              `json:"extra_name,omitempty"`
	Priority           uint8               `json:"priority,omitempty"`
	Fingerprint        string              `json:"fingerprint,omitempty"`
	Client             pkg.Client          `json:"client,omitempty"`
	Ajax               bool                `json:"ajax,omitempty"`
	Screenshot         string              `json:"screenshot,omitempty"`
}

func (r *Request) GetUniqueKey() string {
	return r.UniqueKey
}
func (r *Request) SetUniqueKey(uniqueKey string) pkg.Request {
	r.UniqueKey = uniqueKey
	return r
}
func (r *Request) GetCallBack() string {
	return r.CallBack
}
func (r *Request) SetCallBack(callBack pkg.CallBack) pkg.Request {
	name := runtime.FuncForPC(reflect.ValueOf(callBack).Pointer()).Name()
	r.CallBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	return r
}
func (r *Request) GetErrBack() string {
	return r.ErrBack
}
func (r *Request) SetErrBack(errBack pkg.ErrBack) pkg.Request {
	name := runtime.FuncForPC(reflect.ValueOf(errBack).Pointer()).Name()
	r.ErrBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	return r
}

func (r *Request) GetReferrer() string {
	return r.Referrer
}
func (r *Request) SetReferrer(referrer string) pkg.Request {
	r.Referrer = referrer
	return r
}
func (r *Request) SetUsername(username string) pkg.Request {
	r.Username = username
	return r
}
func (r *Request) GetUsername() string {
	return r.Username
}
func (r *Request) GetPassword() string {
	return r.Password
}
func (r *Request) SetPassword(password string) pkg.Request {
	r.Password = password
	return r
}
func (r *Request) GetChecksum() string {
	return r.Checksum
}
func (r *Request) SetChecksum(checksum string) pkg.Request {
	r.Checksum = checksum
	return r
}
func (r *Request) GetCreateTime() string {
	return r.CreateTime
}
func (r *Request) SetCreateTime(createTime string) pkg.Request {
	r.CreateTime = createTime
	return r
}
func (r *Request) SetSpendTime(spendTime time.Duration) pkg.Request {
	r.SpendTime = utils.DurationNano{Duration: spendTime}
	return r
}
func (r *Request) GetSpendTime() time.Duration {
	return r.SpendTime.Duration
}
func (r *Request) IsSkipMiddleware() bool {
	return r.SkipMiddleware
}
func (r *Request) SetSkipMiddleware(skipMiddleware bool) pkg.Request {
	r.SkipMiddleware = skipMiddleware
	return r
}
func (r *Request) IsSkipFilter() *bool {
	return r.SkipFilter
}
func (r *Request) SetSkipFilter(skipFilter *bool) pkg.Request {
	r.SkipFilter = skipFilter
	return r
}
func (r *Request) SetCanonicalHeaderKey(canonicalHeaderKey *bool) pkg.Request {
	r.CanonicalHeaderKey = canonicalHeaderKey
	return r
}
func (r *Request) IsCanonicalHeaderKey() *bool {
	return r.CanonicalHeaderKey
}
func (r *Request) SetProxyEnable(proxyEnable bool) pkg.Request {
	r.ProxyEnable = &proxyEnable
	return r
}
func (r *Request) IsProxyEnable() *bool {
	return r.ProxyEnable
}
func (r *Request) SetProxy(proxy string) pkg.Request {
	u, err := url.Parse(proxy)
	if err == nil {
		r.SetProxyEnable(true)
	}
	r.Proxy = utils.Url{URL: u}
	return r
}
func (r *Request) GetProxy() *url.URL {
	return r.Proxy.URL
}
func (r *Request) SetRetryMaxTimes(retryMaxTimes *uint8) pkg.Request {
	r.RetryMaxTimes = retryMaxTimes
	return r
}
func (r *Request) GetRetryMaxTimes() *uint8 {
	return r.RetryMaxTimes
}
func (r *Request) SetRetryTimes(retryTimes uint8) pkg.Request {
	r.RetryTimes = retryTimes
	return r
}
func (r *Request) GetRetryTimes() uint8 {
	return r.RetryTimes
}
func (r *Request) SetRedirectMaxTimes(redirectMaxTimes *uint8) pkg.Request {
	r.RedirectMaxTimes = redirectMaxTimes
	return r
}
func (r *Request) GetRedirectMaxTimes() *uint8 {
	return r.RedirectMaxTimes
}
func (r *Request) SetRedirectTimes(redirectTimes uint8) pkg.Request {
	r.RedirectTimes = redirectTimes
	return r
}
func (r *Request) GetRedirectTimes() uint8 {
	return r.RedirectTimes
}
func (r *Request) SetOkHttpCodes(okHttpCodes []int) pkg.Request {
	r.OkHttpCodes = okHttpCodes
	return r
}
func (r *Request) GetOkHttpCodes() []int {
	return r.OkHttpCodes
}
func (r *Request) SetSlot(slot string) pkg.Request {
	r.Slot = slot
	return r
}
func (r *Request) GetSlot() string {
	return r.Slot
}
func (r *Request) SetConcurrency(concurrency *uint8) pkg.Request {
	r.Concurrency = concurrency
	return r
}
func (r *Request) GetConcurrency() *uint8 {
	return r.Concurrency
}
func (r *Request) GetInterval() time.Duration {
	return r.Interval.Duration
}
func (r *Request) SetInterval(interval time.Duration) pkg.Request {
	r.Interval = utils.DurationNano{Duration: interval}
	return r
}
func (r *Request) GetTimeout() time.Duration {
	return r.Timeout.Duration
}
func (r *Request) SetTimeout(timeout time.Duration) pkg.Request {
	r.Timeout = utils.DurationNano{Duration: timeout}
	return r
}
func (r *Request) GetHttpProto() string {
	return r.HttpProto
}
func (r *Request) SetHttpProto(httpProto string) pkg.Request {
	r.HttpProto = httpProto
	return r
}
func (r *Request) GetPlatforms() []pkg.Platform {
	return r.Platforms
}
func (r *Request) SetPlatforms(platforms ...pkg.Platform) pkg.Request {
	r.Platforms = platforms
	return r
}
func (r *Request) GetBrowsers() []pkg.Browser {
	return r.Browsers
}
func (r *Request) SetBrowsers(browsers ...pkg.Browser) pkg.Request {
	r.Browsers = browsers
	return r
}
func (r *Request) setExtraName(name string) {
	r.ExtraName = name
}
func (r *Request) GetExtraName() string {
	return r.ExtraName
}
func (r *Request) GetPriority() uint8 {
	return r.Priority
}
func (r *Request) SetPriority(priority uint8) pkg.Request {
	r.Priority = priority
	return r
}
func (r *Request) GetFingerprint() string {
	return r.Fingerprint
}
func (r *Request) SetFingerprint(fingerprint string) pkg.Request {
	r.Fingerprint = fingerprint
	return r
}
func (r *Request) GetClient() pkg.Client {
	return r.Client
}
func (r *Request) SetClient(client pkg.Client) pkg.Request {
	r.Client = client
	return r
}
func (r *Request) IsAjax() bool {
	return r.Ajax
}
func (r *Request) SetAjax(ajax bool) pkg.Request {
	r.Ajax = ajax
	return r
}
func (r *Request) GetScreenshot() string {
	return r.Screenshot
}
func (r *Request) SetScreenshot(screenshot string) pkg.Request {
	r.Screenshot = screenshot
	return r
}
func (r *Request) setErr(key string, value error) {
	if r.Errors == nil {
		r.Errors = make(map[string]error)
	}
	r.Errors[key] = value
}
func (r *Request) Err() map[string]error {
	return r.Errors
}
func (r *Request) delErr(key string) {
	delete(r.Errors, key)
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
	if r.BodyStr != "" {
		newValues, err := url.ParseQuery(r.BodyStr)
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
		r.Request.Method = method
	} else {
		r.setErr("Method", errors.New("method error"))
	}
	return r
}
func (r *Request) GetMethod() string {
	return r.Method
}
func (r *Request) GetBodyStr() string {
	return r.BodyStr
}
func (r *Request) SetBodyStr(bodyStr string) pkg.Request {
	r.BodyStr = bodyStr
	r.Body = io.NopCloser(strings.NewReader(bodyStr))
	return r
}
func (r *Request) BodyBytes() []byte {
	return []byte(r.BodyStr)
}
func (r *Request) SetBodyBytes(bodyBytes []byte) pkg.Request {
	r.BodyStr = string(bodyBytes)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	return r
}
func (r *Request) SetBodyJson(bodyJson any) pkg.Request {
	bodyBytes, err := json.Marshal(bodyJson)
	if err != nil {
		r.setErr("body", err)
		return r
	}
	r.BodyStr = string(bodyBytes)
	r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
	r.SetHeader("content-type", "application/json")
	return r
}
func (r *Request) GetHeader(key string) string {
	return r.Request.Header.Get(key)
}
func (r *Request) SetHeader(key string, value string) pkg.Request {
	if r.Request.Header == nil {
		r.Request.Header = make(http.Header)
	}
	r.Request.Header.Set(http.CanonicalHeaderKey(key), value)
	return r
}
func (r *Request) Headers() http.Header {
	return r.Request.Header
}
func (r *Request) SetHeaders(headers map[string]string) pkg.Request {
	if r.Request.Header == nil {
		r.Request.Header = make(http.Header)
	}
	for key, value := range headers {
		r.Request.Header.Set(http.CanonicalHeaderKey(key), value)
	}
	return r
}
func (r *Request) IsFile() bool {
	return r.File
}
func (r *Request) AsFile(file bool) pkg.Request {
	r.File = file
	return r
}
func (r *Request) GetFileOptions() *pkg.FileOptions {
	return r.FileOptions
}
func (r *Request) SetFileOptions(options pkg.FileOptions) pkg.Request {
	r.FileOptions = &options
	return r
}
func (r *Request) IsImage() bool {
	return r.Image
}
func (r *Request) AsImage(image bool) pkg.Request {
	r.Image = image
	return r
}
func (r *Request) GetImageOptions() *pkg.ImageOptions {
	return r.ImageOptions
}
func (r *Request) SetImageOptions(options pkg.ImageOptions) pkg.Request {
	r.ImageOptions = &options
	return r
}
func (r *Request) GetContext() pkg.Context {
	return r.Context
}
func (r *Request) WithContext(ctx pkg.Context) pkg.Request {
	r.Context = ctx
	return r
}
func (r *Request) SetBasicAuth(username string, password string) pkg.Request {
	r.Request.SetBasicAuth(username, password)
	return r
}
func (r *Request) RequestContext() context.Context {
	return r.Request.Context()
}
func (r *Request) WithRequestContext(ctx context.Context) pkg.Request {
	r.Request = r.Request.WithContext(ctx)
	return r
}
func (r *Request) GetRequest() pkg.Request {
	return r
}
func (r *Request) GetHttpRequest() *http.Request {
	return r.Request
}
func (r *Request) Cookies() []*http.Cookie {
	return r.Request.Cookies()
}
func (r *Request) AddCookie(c *http.Cookie) pkg.Request {
	if r.Request.Header == nil {
		r.Request.Header = make(http.Header)
	}

	cookiesMap := make(map[string]*http.Cookie)
	for _, v := range r.Request.Cookies() {
		cookiesMap[v.Name] = v
	}

	// TODO different domain?
	cookiesMap[c.Name] = c

	r.Request.Header.Set("Cookie", "")
	for _, v := range cookiesMap {
		r.Request.AddCookie(v)
	}
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
	r.Extra = string(bs)
	return r
}
func (r *Request) GetExtra() string {
	return r.Extra
}
func (r *Request) UnmarshalExtra(v any) (err error) {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		return fmt.Errorf("v must be a non-null pointer")
	}

	if r.Extra == "" {
		return
	}

	err = json.Unmarshal([]byte(r.Extra), v)
	return
}
func (r *Request) MustUnmarshalExtra(v any) {
	if err := r.UnmarshalExtra(v); err != nil {
		panic(err)
	}
}
func (r *Request) UnsafeExtra(v any) {
	_ = r.UnmarshalExtra(v)
}
func (r *Request) Marshal() ([]byte, error) {
	r.Method = r.Request.Method
	if r.URL != nil {
		r.Url = r.URL.String()
	} else {
		r.Url = ""
	}
	r.Header = r.Request.Header
	return json.Marshal(r)
}
func (r *Request) Unmarshal(bytes []byte) (err error) {
	r.Context = new(context2.Context)
	err = json.Unmarshal(bytes, r)
	r.Request, err = http.NewRequest(r.Method, r.Url, strings.NewReader(r.BodyStr))
	if err != nil {
		return
	}
	r.Request.Header = r.Header
	return err
}
func (r *Request) Yield() (err error) {
	return r.Context.GetSpider().GetSpider().YieldRequest(r.Context, r)
}
func (r *Request) MustYield() {
	if err := r.Context.GetSpider().GetSpider().YieldRequest(r.Context, r); err != nil {
		panic(fmt.Errorf("%w: %v", pkg.ErrYieldRequestFailed, err))
	}
}
func (r *Request) UnsafeYield() {
	_ = r.Context.GetSpider().GetSpider().YieldRequest(r.Context, r)
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
