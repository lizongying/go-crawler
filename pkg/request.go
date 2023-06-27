package pkg

import (
	"context"
	"net/http"
	"net/url"
	"reflect"
	"runtime"
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
	Skip               bool
	SkipFilter         bool
	CanonicalHeaderKey bool
	ProxyEnable        bool
	Proxy              *url.URL
	RetryMaxTimes      *uint8
	RetryTimes         uint8
	RedirectMaxTimes   *uint8
	RedirectTimes      uint8
	OkHttpCodes        []int
	Slot               string
	Concurrency        int
	Interval           time.Duration
	Timeout            time.Duration
	HttpProto          string
	Platform           []Platform
	Browser            []Browser
	Extra              any
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

func (r *Request) Marshal() (requestJson RequestJson, err error) {
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
	var retryMaxTimes uint8
	if r.RetryMaxTimes != nil {
		retryMaxTimes = *r.RetryMaxTimes
	}
	var redirectMaxTimes uint8
	if r.RedirectMaxTimes != nil {
		redirectMaxTimes = *r.RedirectMaxTimes
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
	requestJson = RequestJson{
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
		RetryMaxTimes:    retryMaxTimes,
		RetryTimes:       r.RetryTimes,
		RedirectMaxTimes: redirectMaxTimes,
		RedirectTimes:    r.RedirectTimes,
		OkHttpCodes:      r.OkHttpCodes,
		Slot:             r.Slot,
		Concurrency:      r.Concurrency,
		Interval:         int(r.Interval),
		Timeout:          int(r.Timeout),
		HttpProto:        r.HttpProto,
		Platform:         platform,
		Browser:          browser,
		Extra:            r.Extra,
	}
	return
}

type RequestJson struct {
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
	Skip               bool     `json:"skip,omitempty"`                 // Not in to schedule
	SkipFilter         bool     `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool     `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        bool     `json:"proxy_enable,omitempty"`
	Proxy              string   `json:"proxy,omitempty"`
	RetryMaxTimes      uint8    `json:"retry_max_times,omitempty"`
	RetryTimes         uint8    `json:"retry_times,omitempty"`
	RedirectMaxTimes   uint8    `json:"redirect_max_times,omitempty"`
	RedirectTimes      uint8    `json:"redirect_times,omitempty"`
	OkHttpCodes        []int    `json:"ok_http_codes,omitempty"`
	Slot               string   `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        int      `json:"concurrency,omitempty"`
	Interval           int      `json:"interval,omitempty"`
	Timeout            int      `json:"timeout,omitempty"`    //seconds
	HttpProto          string   `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Platform           []string `json:"platform,omitempty"`
	Browser            []string `json:"browser,omitempty"`
	Extra              any      `json:"extra,omitempty"`
}

func (r *RequestJson) Unmarshal(callbacks map[string]Callback, errbacks map[string]Errback) (request Request, err error) {
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
	request = Request{
		UniqueKey:          r.UniqueKey,
		CallBack:           callbacks[r.CallBack],
		ErrBack:            errbacks[r.ErrBack],
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
		RetryMaxTimes:      &r.RetryMaxTimes,
		RetryTimes:         r.RetryTimes,
		RedirectMaxTimes:   &r.RedirectMaxTimes,
		RedirectTimes:      r.RedirectTimes,
		OkHttpCodes:        r.OkHttpCodes,
		Slot:               r.Slot,
		Concurrency:        r.Concurrency,
		Interval:           time.Duration(r.Interval),
		Timeout:            time.Duration(r.Timeout),
		HttpProto:          r.HttpProto,
		Platform:           platform,
		Browser:            browser,
		Extra:              r.Extra,
	}

	return
}

type Callback func(context.Context, *Response) error
type Errback func(context.Context, *Response, error)
