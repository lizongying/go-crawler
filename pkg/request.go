package pkg

import (
	"context"
	"net/http"
	"time"
)

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
	TimeoutRequest     time.Duration                                `json:"timeout_request,omitempty"`
	Checksum           string                                       `json:"checksum,omitempty"`
	CreateTime         string                                       `json:"create_time,omitempty"`          //create time
	Skip               bool                                         `json:"skip,omitempty"`                 // don't schedule
	SkipFilter         bool                                         `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool                                         `json:"canonical_header_key,omitempty"` //canonical header key
	ProxyEnable        bool                                         `json:"proxy_enable,omitempty"`
	RetryEnable        bool                                         `json:"retry_enable,omitempty"`
	RetryTimes         int                                          `json:"retry_times,omitempty"`
	RetryMaxTimes      int                                          `json:"retry_max_times,omitempty"`
	RetryHttpCodes     []int                                        `json:"retry_http_codes,omitempty"`
	Slot               string                                       // same slot same request delay
	Concurrency        int
	Delay              time.Duration
	HttpProto          string `json:"http_proto,omitempty"` // e.g. 1.0/1.1/2.0
	Extra              any    `json:"extra,omitempty"`
}

type RequestFormat struct {
	Http
	UniqueKey string `json:"unique_key,omitempty"` // for filter
	//Parser             func(context.Context, *Response) (err error) `json:"parser,omitempty"`
	Parser             func(context.Context, *Response) (err error) `json:"-"`
	ParserTimeout      int                                          `json:"parser_timeout,omitempty"` // Millisecond
	ErrBack            string                                       `json:"err_back,omitempty"`
	ErrBackTimeout     int                                          `json:"err_back_timeout,omitempty"` // Millisecond
	Checksum           string                                       `json:"checksum,omitempty"`
	CreateTime         string                                       `json:"create_time,omitempty"`          //create time
	Skip               bool                                         `json:"skip,omitempty"`                 // Not in to schedule
	SkipFilter         bool                                         `json:"skip_filter,omitempty"`          // Allow duplicate requests if set "true"
	CanonicalHeaderKey bool                                         `json:"canonical_header_key,omitempty"` //canonical header key
	AllowRefererEmpty  bool                                         `json:"allow_referer_empty,omitempty"`
	RetryEnable        bool                                         `json:"retry_enable,omitempty"`
	RetryTimes         int                                          `json:"retry_times,omitempty"`
	RetryMaxTimes      int                                          `json:"retry_max_times,omitempty"`
	RetryHttpCodes     []int                                        `json:"retry_http_codes,omitempty"`
	DownloadDelay      int                                          `json:"download_delay,omitempty"` // ms
	HttpProto          string                                       `json:"http_proto,omitempty"`     // e.g. 1.0/1.1/2.0
	Extra              any                                          `json:"extra,omitempty"`
}

func (r *Request) SetHeader(key string, value string) {
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	r.Header.Set(key, value)

	return
}

type RequestSlot struct {
	Concurrency     int
	ConcurrencyChan chan struct{}
	Delay           time.Duration
	Timer           *time.Timer
}
