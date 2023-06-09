package pkg

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type RequestJson struct {
	Http
	UniqueKey          string `json:"unique_key,omitempty"` // for filter
	CallBack           string `json:"call_back,omitempty"`
	ErrBack            string `json:"err_back,omitempty"`
	Referer            string `json:"referer,omitempty"`
	Username           string `json:"username,omitempty"`
	Password           string `json:"password,omitempty"`
	Checksum           string `json:"checksum,omitempty"`
	CreateTime         string `json:"create_time,omitempty"`          //create time
	Skip               bool   `json:"skip,omitempty"`                 // Not in to schedule
	SkipFilter         bool   `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool   `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        bool   `json:"proxy_enable,omitempty"`
	Proxy              string `json:"proxy,omitempty"`
	RetryMaxTimes      *uint8 `json:"retry_max_times,omitempty"`
	RetryTimes         uint8  `json:"retry_times,omitempty"`
	RedirectMaxTimes   *uint8 `json:"redirect_max_times,omitempty"`
	RedirectTimes      uint8  `json:"redirect_times,omitempty"`
	OkHttpCodes        []int  `json:"ok_http_codes,omitempty"`
	Slot               string `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        int    `json:"concurrency,omitempty"`
	Interval           int    `json:"interval,omitempty"`
	Timeout            int    `json:"timeout,omitempty"`    //seconds
	HttpProto          string `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Extra              any    `json:"extra,omitempty"`
}

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
	CallBack           func(context.Context, *Response) (err error)
	ErrBack            func(context.Context, *Response, error)
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

func (r *Request) Marshal(requestJson RequestJson) {
	requestJson = RequestJson{
		UniqueKey:        r.UniqueKey,
		Checksum:         r.Checksum,
		CreateTime:       r.CreateTime,
		Skip:             r.Skip,
		SkipFilter:       r.SkipFilter,
		ProxyEnable:      r.ProxyEnable,
		Proxy:            r.Proxy.String(),
		RetryMaxTimes:    r.RetryMaxTimes,
		RetryTimes:       r.RetryTimes,
		RedirectMaxTimes: r.RedirectMaxTimes,
		RedirectTimes:    r.RedirectTimes,
		Slot:             r.Slot,
		OkHttpCodes:      r.OkHttpCodes,
		Concurrency:      r.Concurrency,
		Interval:         int(r.Interval / time.Second),
		Timeout:          int(r.Timeout / time.Second),
		HttpProto:        r.HttpProto,
		Extra:            r.Extra,
	}
}
func (r *RequestJson) Unmarshal(request Request) (err error) {
	proxy, err := url.Parse(r.Proxy)
	request = Request{
		UniqueKey:        r.UniqueKey,
		Checksum:         r.Checksum,
		CreateTime:       r.CreateTime,
		Skip:             r.Skip,
		SkipFilter:       r.SkipFilter,
		ProxyEnable:      r.ProxyEnable,
		Proxy:            proxy,
		RetryMaxTimes:    r.RetryMaxTimes,
		RetryTimes:       r.RetryTimes,
		RedirectMaxTimes: r.RedirectMaxTimes,
		RedirectTimes:    r.RedirectTimes,
		Slot:             r.Slot,
		OkHttpCodes:      r.OkHttpCodes,
		Concurrency:      r.Concurrency,
		Interval:         time.Second * time.Duration(r.Interval),
		Timeout:          time.Second * time.Duration(r.Timeout),
		HttpProto:        r.HttpProto,
		Extra:            r.Extra,
	}

	return
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

type Platform string

const (
	Windows Platform = "windows"
	Mac     Platform = "mac"
	Android Platform = "android"
	Iphone  Platform = "iphone"
	Ipad    Platform = "ipad"
	Linux   Platform = "linux"
)

type Browser string

const (
	Chrome  Browser = "chrome"
	Edge    Browser = "edge"
	Safari  Browser = "safari"
	FireFox Browser = "firefox"
)
