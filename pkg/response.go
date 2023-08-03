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
	SetBodyBytes([]byte) Response
	GetBodyBytes() []byte
	SetFiles([]File) Response
	GetFiles() []File
	SetImages([]Image) Response
	GetImages() []Image
	GetHeaders() http.Header
	GetHeader(string) string
	GetStatusCode() int
	GetBody() io.ReadCloser
	GetCookies() []*http.Cookie
	UnmarshalBody(any) error
	Xpath() (*xpath.Selector, error)
	Query() (*query.Selector, error)
	Json() (gjson.Result, error)
	Re() (*re.Selector, error)

	GetUniqueKey() string
	UnmarshalExtra(any) error
	GetUrl() string
	GetURL() *url.URL
	Context() context.Context
	WithContext(context.Context) Request
	GetFile() bool
	GetImage() bool
	GetSkipMiddleware() bool
	SetSpendTime(time.Duration) Request

	AllLink() []*url.URL
	BodyText() string
}
