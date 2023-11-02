package pkg

import (
	"context"
	"github.com/lizongying/go-css/css"
	"github.com/lizongying/go-json/gjson"
	"github.com/lizongying/go-re/re"
	"github.com/lizongying/go-xpath/xpath"
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
	MustXpath() *xpath.Selector
	MustXpathOne(string) *xpath.Result
	MustXpathMany(string) []*xpath.Result
	Xpath() (*xpath.Selector, error)
	XpathOne(string) (*xpath.Result, error)
	XpathMany(string) ([]*xpath.Result, error)
	MustCss() *css.Selector
	MustCssOne(string) *css.Result
	MustCssMany(string) []*css.Result
	Css() (*css.Selector, error)
	CssOne(string) (*css.Result, error)
	CssMany(string) ([]*css.Result, error)
	MustJson() *gjson.Selector
	MustJsonOne(string) *gjson.Result
	MustJsonMany(string) []*gjson.Result
	Json() (*gjson.Selector, error)
	JsonOne(string) (*gjson.Result, error)
	JsonMany(string) ([]*gjson.Result, error)
	MustRe() *re.Selector
	MustReOne(string) *re.Result
	MustReMany(string) []*re.Result
	Re() (*re.Selector, error)
	ReOne(string) (*re.Result, error)
	ReMany(string) ([]*re.Result, error)

	UniqueKey() string
	UnmarshalExtra(any) error
	MustUnmarshalExtra(any)
	Url() string
	URL() *url.URL
	Context() context.Context
	WithRequestContext(context.Context) Request
	IsFile() bool
	FileOptions() *FileOptions
	IsImage() bool
	ImageOptions() *ImageOptions
	SkipMiddleware() bool
	SetSpendTime(time.Duration) Request

	AllLink() []*url.URL
	BodyText() string

	AbsoluteURL(relativeUrl string) (absoluteURL *url.URL, err error)

	MustUnmarshalData(v any)

	// UnmarshalData
	// Parsing data into `v` based on the parsing rules in the `v` field tag.
	// _json="data.name"
	// _re="name"
	// _xpath="//a[@link]"
	// _css=".class"
	UnmarshalData(v any) error
}
