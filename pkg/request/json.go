package request

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type JsonWithContext struct {
	pkg.ContextJson `json:"context,omitempty"`
	pkg.RequestJson `json:"request,omitempty"`
}

type Json struct {
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
	Task               *pkg.Task           `json:"task,omitempty"`
}

func (r *Json) ToRequest() (request pkg.Request, err error) {
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
		task:               r.Task,
	}

	return
}
