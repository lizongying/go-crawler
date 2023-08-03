package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-query/query"
	"github.com/lizongying/go-re/re"
	"github.com/lizongying/go-xpath/xpath"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"time"
)

type Response struct {
	*http.Response
	request   pkg.Request
	bodyBytes []byte
	files     []pkg.File
	images    []pkg.Image
}

func (r *Response) SetResponse(response *http.Response) pkg.Response {
	r.Response = response
	return r
}
func (r *Response) GetResponse() *http.Response {
	return r.Response
}
func (r *Response) SetRequest(request pkg.Request) pkg.Response {
	r.request = request
	return r
}
func (r *Response) GetRequest() pkg.Request {
	return r.request
}
func (r *Response) SetBodyBytes(bodyBytes []byte) pkg.Response {
	r.bodyBytes = bodyBytes
	return r
}
func (r *Response) GetBodyBytes() []byte {
	return r.bodyBytes
}
func (r *Response) SetFiles(files []pkg.File) pkg.Response {
	r.files = files
	return r
}
func (r *Response) GetFiles() []pkg.File {
	return r.files
}
func (r *Response) SetImages(images []pkg.Image) pkg.Response {
	r.images = images
	return r
}
func (r *Response) GetImages() []pkg.Image {
	return r.images
}

func (r *Response) GetHeaders() http.Header {
	return r.Response.Header
}
func (r *Response) GetHeader(key string) string {
	return r.Response.Header.Get(key)
}
func (r *Response) GetStatusCode() int {
	return r.Response.StatusCode
}
func (r *Response) GetBody() io.ReadCloser {
	return r.Response.Body
}
func (r *Response) GetCookies() []*http.Cookie {
	return r.Response.Cookies()
}
func (r *Response) UnmarshalBody(v any) error {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		return fmt.Errorf("v must be a non-null pointer")
	}

	if len(r.bodyBytes) == 0 {
		return fmt.Errorf("response is empty")
	}

	return json.Unmarshal(r.bodyBytes, v)
}

func (r *Response) GetUniqueKey() string {
	return r.request.GetUniqueKey()
}
func (r *Response) UnmarshalExtra(v any) error {
	return r.request.UnmarshalExtra(v)
}
func (r *Response) GetUrl() string {
	return r.request.GetUrl()
}
func (r *Response) GetURL() *url.URL {
	return r.request.GetURL()
}
func (r *Response) Context() context.Context {
	return r.request.Context()
}
func (r *Response) WithContext(ctx context.Context) pkg.Request {
	return r.request.WithContext(ctx)
}
func (r *Response) GetFile() bool {
	return r.request.GetFile()
}
func (r *Response) GetImage() bool {
	return r.request.GetImage()
}
func (r *Response) GetSkipMiddleware() bool {
	return r.request.GetSkipMiddleware()
}
func (r *Response) SetSpendTime(spendTime time.Duration) pkg.Request {
	return r.request.SetSpendTime(spendTime)
}

// Xpath returns a xpath selector
func (r *Response) Xpath() (selector *xpath.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = xpath.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}

	return
}

// Query returns a query selector
func (r *Response) Query() (selector *query.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = query.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}

	return
}

// Json return a gjson
func (r *Response) Json() (result gjson.Result, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	result = gjson.ParseBytes(r.bodyBytes)

	return
}

// Re return a regex
func (r *Response) Re() (selector *re.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = re.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}

	return
}
func areSameDomain(url1, url2 *url.URL) bool {
	// Check if the scheme and host (domain) are the same
	if url1.Scheme != url2.Scheme || url1.Host != url2.Host {
		return false
	}

	// Check if the port (if specified) is the same
	if url1.Port() != url2.Port() {
		return false
	}

	return true
}

func (r *Response) AllLink() (links []*url.URL) {
	if r == nil {
		return
	}

	if len(r.bodyBytes) == 0 {
		return
	}

	selector, err := xpath.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}

	base := r.request.GetRequest().URL
	for _, v := range selector.FindStrMany("//a/@href") {
		relative, e := url.Parse(v)
		if e != nil {
			continue
		}
		relative = base.ResolveReference(relative)
		if areSameDomain(base, relative) && base.String() != relative.String() {
			links = append(links, relative)
		}
	}

	return
}

func (r *Response) BodyText() (body string) {
	if r == nil {
		return
	}

	if len(r.bodyBytes) == 0 {
		return
	}

	sel, err := query.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}
	body = sel.Remove("script").FindStrOne("body")
	body = regexp.MustCompile(` `).ReplaceAllString(body, "")
	body = regexp.MustCompile(`\n+`).ReplaceAllString(body, "\n")
	return
}
