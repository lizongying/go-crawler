package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-css/css"
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
func (r *Response) BodyBytes() []byte {
	return r.bodyBytes
}
func (r *Response) SetBodyBytes(bodyBytes []byte) pkg.Response {
	r.bodyBytes = bodyBytes
	return r
}
func (r *Response) BodyStr() string {
	return string(r.bodyBytes)
}
func (r *Response) SetBodyStr(bodyStr string) pkg.Response {
	r.bodyBytes = []byte(bodyStr)
	return r
}
func (r *Response) SetFiles(files []pkg.File) pkg.Response {
	r.files = files
	return r
}
func (r *Response) Files() []pkg.File {
	return r.files
}
func (r *Response) SetImages(images []pkg.Image) pkg.Response {
	r.images = images
	return r
}
func (r *Response) Images() []pkg.Image {
	return r.images
}

func (r *Response) Headers() http.Header {
	return r.Response.Header
}
func (r *Response) GetHeader(key string) string {
	return r.Response.Header.Get(key)
}
func (r *Response) StatusCode() int {
	return r.Response.StatusCode
}
func (r *Response) SetStatusCode(statusCode int) pkg.Response {
	r.Response.StatusCode = statusCode
	return r
}
func (r *Response) GetBody() io.ReadCloser {
	return r.Response.Body
}
func (r *Response) Cookies() []*http.Cookie {
	return r.Response.Cookies()
}
func (r *Response) SetCookies(cookies ...*http.Cookie) pkg.Response {
	if r.Response == nil {
		r.Response = new(http.Response)
	}
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	for _, cookie := range cookies {
		if v := cookie.String(); v != "" {
			r.Header.Add("Set-Cookie", v)
		}
	}

	return r
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

func (r *Response) UniqueKey() string {
	return r.request.UniqueKey()
}
func (r *Response) UnmarshalExtra(v any) error {
	return r.request.UnmarshalExtra(v)
}
func (r *Response) MustUnmarshalExtra(v any) {
	r.request.MustUnmarshalExtra(v)
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
func (r *Response) File() bool {
	return r.request.File()
}
func (r *Response) Image() bool {
	return r.request.Image()
}
func (r *Response) SkipMiddleware() bool {
	return r.request.SkipMiddleware()
}
func (r *Response) SetSpendTime(spendTime time.Duration) pkg.Request {
	return r.request.SetSpendTime(spendTime)
}
func (r *Response) MustXpath() (selector *xpath.Selector) {
	selector, _ = r.Xpath()
	return
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

func (r *Response) MustCss() (selector *css.Selector) {
	selector, _ = r.Css()
	return
}

// Css returns a css selector
func (r *Response) Css() (selector *css.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = css.NewSelectorFromBytes(r.bodyBytes)
	if err != nil {
		return
	}

	return
}

func (r *Response) MustJson() (result gjson.Result) {
	result, _ = r.Json()
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

// AbsoluteURL Generating an absolute URL based on a relative URL.
func (r *Response) AbsoluteURL(relativeUrl string) (absoluteURL *url.URL, err error) {
	base := r.request.GetRequest().URL
	var relativeURL *url.URL
	relativeURL, err = url.Parse(relativeUrl)
	if err != nil {
		return
	}
	absoluteURL = base.ResolveReference(relativeURL)
	return
}

func (r *Response) MustUnmarshalData(v any) {
	_ = r.UnmarshalData(v)
}
func (r *Response) UnmarshalData(v any) (err error) {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		return fmt.Errorf("`v` must be a non-null pointer")
	}

	vValue = vValue.Elem()
	if vValue.Kind() != reflect.Struct {
		err = errors.New("`*v` must be a struct")
		return
	}

	dataType, _ := vValue.Type().FieldByName("Data")
	eleCount := 0
	rootPathJson := dataType.Tag.Get("_json")
	var rootJson gjson.Result
	if rootPathJson != "" {
		rootJson = r.MustJson().Get(rootPathJson)
		if rootJson.IsArray() {
			eleCount = len(rootJson.Array())
		}
	}

	rootPathXpath := dataType.Tag.Get("_xpath")
	var rootXpath []*xpath.Selector
	if rootPathXpath != "" {
		rootXpath = r.MustXpath().FindNodeMany(rootPathXpath)
		eleCount = len(rootXpath)
		if eleCount == 0 {
			return
		}
	}

	rootPathCss := dataType.Tag.Get("_css")
	var rootCss []*css.Selector
	if rootPathCss != "" {
		rootCss = r.MustCss().FindNodeMany(rootPathCss)
		eleCount = len(rootCss)
		if eleCount == 0 {
			return
		}
	}

	vValue = vValue.FieldByName("Data")
	if vValue.Kind() == reflect.Slice {
		eleType := vValue.Type().Elem()
		if eleType.Kind() != reflect.Struct {
			err = errors.New("elements of `v` must be a struct")
			return
		}

		root := reflect.MakeSlice(vValue.Type(), 0, 0)
		l := eleType.NumField()
		for i := 0; i < eleCount; i++ {
			ele := reflect.New(eleType).Elem()
			for ii := 0; ii < l; ii++ {
				elePathJson := eleType.Field(ii).Tag.Get("_json")
				if elePathJson != "" {
					eleJson := rootJson.Array()[i].Get(elePathJson)
					eleField := ele.Field(ii)
					switch eleType.Field(ii).Type.Kind() {
					case reflect.Int:
						eleField.SetInt(eleJson.Int())
					case reflect.Int8:
						eleField.SetInt(eleJson.Int())
					case reflect.Int16:
						eleField.SetInt(eleJson.Int())
					case reflect.Int32:
						eleField.SetInt(eleJson.Int())
					case reflect.Int64:
						eleField.SetInt(eleJson.Int())
					case reflect.String:
						eleField.SetString(eleJson.String())
					case reflect.Bool:
						eleField.SetBool(eleJson.Bool())
					case reflect.Uint:
						eleField.SetUint(eleJson.Uint())
					case reflect.Uint8:
						eleField.SetUint(eleJson.Uint())
					case reflect.Uint16:
						eleField.SetUint(eleJson.Uint())
					case reflect.Uint32:
						eleField.SetUint(eleJson.Uint())
					case reflect.Uint64:
						eleField.SetUint(eleJson.Uint())
					case reflect.Float32:
						eleField.SetFloat(eleJson.Float())
					case reflect.Float64:
						eleField.SetFloat(eleJson.Float())
					}
					continue
				}
				elePathXpath := eleType.Field(ii).Tag.Get("_xpath")
				if elePathXpath != "" {
					eleXpath := rootXpath[i]
					eleField := ele.Field(ii)
					switch eleType.Field(ii).Type.Kind() {
					case reflect.Int:
						eleField.SetInt(eleXpath.One(elePathXpath).Int64())
					case reflect.Int8:
						eleField.SetInt(eleXpath.One(elePathXpath).Int64())
					case reflect.Int16:
						eleField.SetInt(eleXpath.One(elePathXpath).Int64())
					case reflect.Int32:
						eleField.SetInt(eleXpath.One(elePathXpath).Int64())
					case reflect.Int64:
						eleField.SetInt(eleXpath.One(elePathXpath).Int64())
					case reflect.String:
						eleField.SetString(eleXpath.One(elePathXpath).String())
					case reflect.Bool:
						eleField.SetBool(eleXpath.One(elePathXpath).Bool())
					case reflect.Uint:
						eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleXpath.One(elePathXpath).Float64())
					case reflect.Float64:
						eleField.SetFloat(eleXpath.One(elePathXpath).Float64())
					}
					continue
				}
				elePathCss := eleType.Field(ii).Tag.Get("_css")
				if elePathCss != "" {
					eleCss := rootCss[i]
					eleField := ele.Field(ii)
					switch eleType.Field(ii).Type.Kind() {
					case reflect.Int:
						eleField.SetInt(eleCss.One(elePathCss).Int64())
					case reflect.Int8:
						eleField.SetInt(eleCss.One(elePathCss).Int64())
					case reflect.Int16:
						eleField.SetInt(eleCss.One(elePathCss).Int64())
					case reflect.Int32:
						eleField.SetInt(eleCss.One(elePathCss).Int64())
					case reflect.Int64:
						eleField.SetInt(eleCss.One(elePathCss).Int64())
					case reflect.String:
						eleField.SetString(eleCss.One(elePathCss).String())
					case reflect.Bool:
						eleField.SetBool(eleCss.One(elePathCss).Bool())
					case reflect.Uint:
						eleField.SetUint(eleCss.One(elePathCss).Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleCss.One(elePathCss).Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleCss.One(elePathCss).Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleCss.One(elePathCss).Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleCss.One(elePathCss).Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleCss.One(elePathCss).Float64())
					case reflect.Float64:
						eleField.SetFloat(eleCss.One(elePathCss).Float64())
					}
					continue
				}
			}
			root = reflect.Append(root, ele)
		}
		vValue.Set(root)
	} else if vValue.Kind() == reflect.Struct {
		eleType := vValue.Type()

		l := eleType.NumField()
		for ii := 0; ii < l; ii++ {
			elePathJson := eleType.Field(ii).Tag.Get("_json")
			if elePathJson != "" {
				eleJson := rootJson.Get(elePathJson)
				eleField := vValue.Field(ii)
				switch eleType.Field(ii).Type.Kind() {
				case reflect.Int:
					eleField.SetInt(eleJson.Int())
				case reflect.Int8:
					eleField.SetInt(eleJson.Int())
				case reflect.Int16:
					eleField.SetInt(eleJson.Int())
				case reflect.Int32:
					eleField.SetInt(eleJson.Int())
				case reflect.Int64:
					eleField.SetInt(eleJson.Int())
				case reflect.String:
					eleField.SetString(eleJson.String())
				case reflect.Bool:
					eleField.SetBool(eleJson.Bool())
				case reflect.Uint:
					eleField.SetUint(eleJson.Uint())
				case reflect.Uint8:
					eleField.SetUint(eleJson.Uint())
				case reflect.Uint16:
					eleField.SetUint(eleJson.Uint())
				case reflect.Uint32:
					eleField.SetUint(eleJson.Uint())
				case reflect.Uint64:
					eleField.SetUint(eleJson.Uint())
				case reflect.Float32:
					eleField.SetFloat(eleJson.Float())
				case reflect.Float64:
					eleField.SetFloat(eleJson.Float())
				}
				continue
			}
			elePathXpath := eleType.Field(ii).Tag.Get("_xpath")
			if elePathXpath != "" {
				eleXpath := rootXpath[0]
				eleField := vValue.Field(ii)
				switch eleType.Field(ii).Type.Kind() {
				case reflect.Int:
					eleField.SetInt(eleXpath.One(elePathXpath).Int64())
				case reflect.Int8:
					eleField.SetInt(eleXpath.One(elePathXpath).Int64())
				case reflect.Int16:
					eleField.SetInt(eleXpath.One(elePathXpath).Int64())
				case reflect.Int32:
					eleField.SetInt(eleXpath.One(elePathXpath).Int64())
				case reflect.Int64:
					eleField.SetInt(eleXpath.One(elePathXpath).Int64())
				case reflect.String:
					eleField.SetString(eleXpath.One(elePathXpath).String())
				case reflect.Bool:
					eleField.SetBool(eleXpath.One(elePathXpath).Bool())
				case reflect.Uint:
					eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleXpath.One(elePathXpath).Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleXpath.One(elePathXpath).Float64())
				case reflect.Float64:
					eleField.SetFloat(eleXpath.One(elePathXpath).Float64())
				}
				continue
			}
			elePathCss := eleType.Field(ii).Tag.Get("_css")
			if elePathCss != "" {
				eleCss := rootCss[0]
				eleField := vValue.Field(ii)
				switch eleType.Field(ii).Type.Kind() {
				case reflect.Int:
					eleField.SetInt(eleCss.One(elePathCss).Int64())
				case reflect.Int8:
					eleField.SetInt(eleCss.One(elePathCss).Int64())
				case reflect.Int16:
					eleField.SetInt(eleCss.One(elePathCss).Int64())
				case reflect.Int32:
					eleField.SetInt(eleCss.One(elePathCss).Int64())
				case reflect.Int64:
					eleField.SetInt(eleCss.One(elePathCss).Int64())
				case reflect.String:
					eleField.SetString(eleCss.One(elePathCss).String())
				case reflect.Bool:
					eleField.SetBool(eleCss.One(elePathCss).Bool())
				case reflect.Uint:
					eleField.SetUint(eleCss.One(elePathCss).Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleCss.One(elePathCss).Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleCss.One(elePathCss).Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleCss.One(elePathCss).Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleCss.One(elePathCss).Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleCss.One(elePathCss).Float64())
				case reflect.Float64:
					eleField.SetFloat(eleCss.One(elePathCss).Float64())
				}
				continue
			}
		}
	} else {
		err = errors.New("`*v.Data` must be a slice or struct")
		return
	}

	return
}
