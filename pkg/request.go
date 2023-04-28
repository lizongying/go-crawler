package pkg

import (
	"context"
	"net/http"
	"net/url"
	"time"
)

type RequestJson struct {
	Http
	UniqueKey string `json:"unique_key,omitempty"` // for filter
	//Parser             func(context.Context, *Response) (err error) `json:"parser,omitempty"`
	Parser             func(context.Context, *Response) (err error) `json:"-"`
	ParserTimeout      int                                          `json:"parser_timeout,omitempty"` // Millisecond
	ErrBack            string                                       `json:"err_back,omitempty"`
	Timeout            int                                          `json:"timeout,omitempty"` // Millisecond
	Checksum           string                                       `json:"checksum,omitempty"`
	CreateTime         string                                       `json:"create_time,omitempty"`          //create time
	Skip               bool                                         `json:"skip,omitempty"`                 // Not in to schedule
	SkipFilter         bool                                         `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool                                         `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        bool                                         `json:"proxy_enable,omitempty"`
	Proxy              string                                       `json:"proxy,omitempty"`
	AllowRefererEmpty  bool                                         `json:"allow_referer_empty,omitempty"`
	RetryEnable        bool                                         `json:"retry_enable,omitempty"`
	RetryMaxTimes      int                                          `json:"retry_max_times,omitempty"`
	RetryTimes         int                                          `json:"retry_times,omitempty"`
	OkHttpCodes        []int                                        `json:"ok_http_codes,omitempty"`
	Slot               string                                       `json:"slot,omitempty"` // same slot same concurrency & delay
	Concurrency        int                                          `json:"concurrency,omitempty"`
	Interval           int                                          `json:"interval,omitempty"`
	HttpProto          string                                       `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Extra              any                                          `json:"extra,omitempty"`
}

type Http struct {
	*http.Request
	Url     string      `json:"url,omitempty"`
	Method  string      `json:"method,omitempty"`
	BodyStr string      `json:"body,omitempty"`
	Header  http.Header `json:"header,omitempty"`
}

type Request struct {
	Http
	UniqueKey          string                                       `json:"unique_key,omitempty"` // for filter
	CallBack           func(context.Context, *Response) (err error) `json:"-"`
	ErrBack            func(context.Context, *Response, error)      `json:"-"`
	TimeoutAll         time.Duration                                `json:"timeout_all,omitempty"`
	Timeout            time.Duration                                `json:"timeout,omitempty"`
	Checksum           string                                       `json:"checksum,omitempty"`
	CreateTime         string                                       `json:"create_time,omitempty"`          //create time
	Skip               bool                                         `json:"skip,omitempty"`                 // don't schedule
	SkipFilter         bool                                         `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool                                         `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        bool
	Proxy              *url.URL
	RetryEnable        bool
	RetryMaxTimes      int
	RetryTimes         int
	OkHttpCodes        []int
	Slot               string
	Concurrency        int
	Interval           time.Duration
	HttpProto          string `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Extra              any    `json:"extra,omitempty"`
}

func (r *Request) Marshal(requestJson RequestJson) {
	requestJson = RequestJson{
		ProxyEnable:   r.ProxyEnable,
		Proxy:         r.Proxy.String(),
		RetryEnable:   r.RetryEnable,
		RetryMaxTimes: r.RetryMaxTimes,
		RetryTimes:    r.RetryTimes,
		OkHttpCodes:   r.OkHttpCodes,
	}
}
func (r *RequestJson) Unmarshal(request Request) (err error) {
	proxy, err := url.Parse(r.Proxy)
	request = Request{
		ProxyEnable:   r.ProxyEnable,
		Proxy:         proxy,
		RetryEnable:   r.RetryEnable,
		RetryMaxTimes: r.RetryMaxTimes,
		RetryTimes:    r.RetryTimes,
		OkHttpCodes:   r.OkHttpCodes,
	}

	return
}

func (r *Request) SetHeader(key string, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)

	return
}
