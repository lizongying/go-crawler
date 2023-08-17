package pkg

import (
	"context"
	"github.com/lizongying/go-query/query"
	"github.com/lizongying/go-re/re"
	"github.com/lizongying/go-xpath/xpath"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Response interface {
	SetResponse(*http.Response) Response
	GetResponse() *http.Response
	SetRequest(Request) Response
	GetRequest() Request
	BodyBytes() []byte
	SetBodyBytes([]byte) Response
	BodyStr() string
	SetBodyStr(string) Response
	Files() []File
	SetFiles([]File) Response
	Images() []Image
	SetImages([]Image) Response
	Headers() http.Header
	GetHeader(string) string
	StatusCode() int
	SetStatusCode(statusCode int) Response
	GetBody() io.ReadCloser
	Cookies() []*http.Cookie
	SetCookies(...*http.Cookie) Response
	UnmarshalBody(any) error
	Xpath() (*xpath.Selector, error)
	Query() (*query.Selector, error)
	Json() (gjson.Result, error)
	Re() (*re.Selector, error)

	UniqueKey() string
	UnmarshalExtra(any) error
	MustUnmarshalExtra(any)
	GetUrl() string
	GetURL() *url.URL
	Context() context.Context
	WithContext(context.Context) Request
	File() bool
	Image() bool
	SkipMiddleware() bool
	SetSpendTime(time.Duration) Request

	AllLink() []*url.URL
	BodyText() string
}
