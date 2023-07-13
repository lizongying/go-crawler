package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Http struct {
	*http.Request
	Url     string         `json:"url,omitempty"`
	Method  string         `json:"method,omitempty"`
	BodyStr string         `json:"body,omitempty"`
	Header  http.Header    `json:"header,omitempty"`
	Cookies []*http.Cookie `json:"cookies,omitempty"`
}

type Request struct {
	Http
	UniqueKey          string
	CallBack           Callback
	ErrBack            Errback
	Referer            string
	Username           string
	Password           string
	Checksum           string
	CreateTime         string
	SpendTime          time.Duration
	Skip               *bool
	SkipFilter         *bool
	CanonicalHeaderKey *bool
	ProxyEnable        *bool
	Proxy              *url.URL
	RetryMaxTimes      *uint8
	RetryTimes         uint8
	RedirectMaxTimes   *uint8
	RedirectTimes      uint8
	OkHttpCodes        []int
	Slot               string
	Concurrency        *uint8
	Interval           time.Duration
	Timeout            time.Duration
	HttpProto          string
	Platform           []Platform
	Browser            []Browser
	File               *bool
	Image              *bool
	Extra              any
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

func (r *Request) SetHeader(key string, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)

	if r.Request != nil {
		r.Request.Header = r.Header
	}

	return
}
func (r *Request) SetUrl(Url string) *Request {
	r.Url = Url
	if r.Request != nil {
		URL, _ := url.Parse(Url)
		r.Request.URL = URL
	}

	return r
}

func MapToStruct(data map[string]interface{}, obj interface{}) error {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr || objValue.IsNil() {
		return fmt.Errorf("obj must be a non-null pointer")
	}

	objValue = objValue.Elem()
	objType := objValue.Type()

	for i := 0; i < objValue.NumField(); i++ {
		field := objValue.Field(i)
		fieldType := objType.Field(i)

		value, ok := data[fieldType.Name]
		if !ok {
			continue
		}

		fieldValue := reflect.ValueOf(value)

		if fieldType.Type.Kind() == reflect.Struct && fieldValue.Type().Kind() == reflect.Map {
			nestedStruct := reflect.New(field.Type()).Interface()
			err := MapToStruct(fieldValue.Interface().(map[string]interface{}), nestedStruct)
			if err != nil {
				return err
			}
			field.Set(reflect.ValueOf(nestedStruct).Elem())
		} else if fieldValue.Type().ConvertibleTo(field.Type()) {
			field.Set(fieldValue.Convert(field.Type()))
		} else {
			return fmt.Errorf("field %s type does not match", fieldType.Name)
		}
	}

	return nil
}
func (r *Request) SetFile(isFile bool) {
	r.File = &isFile
}
func (r *Request) GetFile() bool {
	if r.File == nil {
		return false
	}
	return *r.File
}
func (r *Request) SetImage(isImage bool) {
	r.Image = &isImage
}
func (r *Request) GetImage() bool {
	if r.Image == nil {
		return false
	}
	return *r.Image
}
func (r *Request) SetProxyEnable(proxyEnable *bool) {
	r.ProxyEnable = proxyEnable
}
func (r *Request) GetProxyEnable() bool {
	if r.ProxyEnable == nil {
		return false
	}
	return *r.ProxyEnable
}
func (r *Request) SetSkip(skip *bool) {
	r.Skip = skip
}
func (r *Request) GetSkip() bool {
	if r.Skip == nil {
		return false
	}
	return *r.Skip
}
func (r *Request) SetSkipFilter(skipFilter *bool) {
	r.SkipFilter = skipFilter
}
func (r *Request) GetSkipFilter() bool {
	if r.SkipFilter == nil {
		return false
	}
	return *r.SkipFilter
}
func (r *Request) SetConcurrency(concurrency *uint8) {
	r.Concurrency = concurrency
}
func (r *Request) GetConcurrency() uint8 {
	if r.Concurrency == nil {
		return uint8(1)
	}
	return *r.Concurrency
}
func (r *Request) SetCanonicalHeaderKey(canonicalHeaderKey *bool) {
	r.CanonicalHeaderKey = canonicalHeaderKey
}
func (r *Request) GetCanonicalHeaderKey() bool {
	if r.CanonicalHeaderKey == nil {
		return true
	}
	return *r.CanonicalHeaderKey
}
func (r *Request) GetExtra(obj interface{}) (err error) {
	objValue := reflect.ValueOf(obj)
	if objValue.Kind() != reflect.Ptr || objValue.IsNil() {
		return fmt.Errorf("obj must be a non-null pointer")
	}
	objValue = objValue.Elem()
	objType := objValue.Type()

	if r.Extra == nil {
		return
	}

	extraValue := reflect.ValueOf(r.Extra)
	if extraValue.Kind() == reflect.Ptr {
		extraValue = extraValue.Elem()
	}
	extraType := extraValue.Type()
	if extraValue.Kind() == reflect.Struct {
		if objType.Kind() == reflect.Struct {
			if extraType == objType {
				objValue.Set(extraValue)
				return
			}
			return
		}
		if objType.Kind() == reflect.Interface {
			if extraType.Implements(objType) {
				objValue.Set(extraValue.Convert(objType))
				return
			}
			return
		}
	}
	if extraValue.Kind() == reflect.Map {
		extra, ok := r.Extra.(map[string]interface{})
		if !ok {
			return
		}
		return MapToStruct(extra, obj)
	}

	return
}

func (r *Request) ToRequestJson() (request *RequestJson, err error) {
	var proxy string
	if r.Proxy != nil {
		proxy = r.Proxy.String()
	}
	var callBack string
	if r.CallBack != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.CallBack).Pointer()).Name()
		callBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	}
	var errBack string
	if r.ErrBack != nil {
		name := runtime.FuncForPC(reflect.ValueOf(r.ErrBack).Pointer()).Name()
		errBack = name[strings.LastIndex(name, ".")+1 : strings.LastIndex(name, "-")]
	}
	var platform []string
	if len(r.Platform) > 0 {
		for _, v := range r.Platform {
			platform = append(platform, string(v))
		}
	}
	var browser []string
	if len(r.Browser) > 0 {
		for _, v := range r.Browser {
			browser = append(browser, string(v))
		}
	}
	request = &RequestJson{
		Http:             r.Http,
		UniqueKey:        r.UniqueKey,
		CallBack:         callBack,
		ErrBack:          errBack,
		Referer:          r.Referer,
		Username:         r.Username,
		Password:         r.Password,
		Checksum:         r.Checksum,
		CreateTime:       r.CreateTime,
		SpendTime:        uint(r.SpendTime),
		Skip:             r.Skip,
		SkipFilter:       r.SkipFilter,
		ProxyEnable:      r.ProxyEnable,
		Proxy:            proxy,
		RetryMaxTimes:    r.RetryMaxTimes,
		RetryTimes:       r.RetryTimes,
		RedirectMaxTimes: r.RedirectMaxTimes,
		RedirectTimes:    r.RedirectTimes,
		OkHttpCodes:      r.OkHttpCodes,
		Slot:             r.Slot,
		Concurrency:      r.Concurrency,
		Interval:         int(r.Interval),
		Timeout:          int(r.Timeout),
		HttpProto:        r.HttpProto,
		Platform:         platform,
		Browser:          browser,
		Image:            r.Image,
		Extra:            r.Extra,
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
	callbacks map[string]Callback
	errbacks  map[string]Errback
	Http
	UniqueKey          string   `json:"unique_key,omitempty"` // for filter
	CallBack           string   `json:"call_back,omitempty"`
	ErrBack            string   `json:"err_back,omitempty"`
	Referer            string   `json:"referer,omitempty"`
	Username           string   `json:"username,omitempty"`
	Password           string   `json:"password,omitempty"`
	Checksum           string   `json:"checksum,omitempty"`
	CreateTime         string   `json:"create_time,omitempty"` //create time
	SpendTime          uint     `json:"spend_time,omitempty"`
	Skip               *bool    `json:"skip,omitempty"`                 // Not in to schedule
	SkipFilter         *bool    `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey *bool    `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        *bool    `json:"proxy_enable,omitempty"`
	Proxy              string   `json:"proxy,omitempty"`
	RetryMaxTimes      *uint8   `json:"retry_max_times,omitempty"`
	RetryTimes         uint8    `json:"retry_times,omitempty"`
	RedirectMaxTimes   *uint8   `json:"redirect_max_times,omitempty"`
	RedirectTimes      uint8    `json:"redirect_times,omitempty"`
	OkHttpCodes        []int    `json:"ok_http_codes,omitempty"`
	Slot               string   `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        *uint8   `json:"concurrency,omitempty"`
	Interval           int      `json:"interval,omitempty"`
	Timeout            int      `json:"timeout,omitempty"`    //seconds
	HttpProto          string   `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Platform           []string `json:"platform,omitempty"`
	Browser            []string `json:"browser,omitempty"`
	Image              *bool    `json:"image,omitempty"`
	Extra              any      `json:"extra,omitempty"`
}

func (r *RequestJson) SetCallbacks(callbacks map[string]Callback) {
	r.callbacks = callbacks
}
func (r *RequestJson) SetErrbacks(errbacks map[string]Errback) {
	r.errbacks = errbacks
}
func (r *RequestJson) ToRequest() (request *Request, err error) {
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
		Http:               r.Http,
		UniqueKey:          r.UniqueKey,
		CallBack:           r.callbacks[r.CallBack],
		ErrBack:            r.errbacks[r.ErrBack],
		Referer:            r.Referer,
		Username:           r.Username,
		Password:           r.Password,
		Checksum:           r.Checksum,
		CreateTime:         r.CreateTime,
		SpendTime:          time.Duration(r.SpendTime),
		Skip:               r.Skip,
		SkipFilter:         r.SkipFilter,
		CanonicalHeaderKey: r.CanonicalHeaderKey,
		ProxyEnable:        r.ProxyEnable,
		Proxy:              proxy,
		RetryMaxTimes:      r.RetryMaxTimes,
		RetryTimes:         r.RetryTimes,
		RedirectMaxTimes:   r.RedirectMaxTimes,
		RedirectTimes:      r.RedirectTimes,
		OkHttpCodes:        r.OkHttpCodes,
		Slot:               r.Slot,
		Concurrency:        r.Concurrency,
		Interval:           time.Duration(r.Interval),
		Timeout:            time.Duration(r.Timeout),
		HttpProto:          r.HttpProto,
		Platform:           platform,
		Browser:            browser,
		Image:              r.Image,
		Extra:              r.Extra,
	}

	return
}

type Callback func(context.Context, *Response) error
type Errback func(context.Context, *Response, error)
