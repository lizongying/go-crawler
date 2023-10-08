package response

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-css/css"
	"github.com/lizongying/go-json/gjson"
	"github.com/lizongying/go-query/query"
	"github.com/lizongying/go-re/re"
	"github.com/lizongying/go-xpath/xpath"
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
func (r *Response) Url() string {
	return r.request.GetUrl()
}
func (r *Response) URL() *url.URL {
	return r.request.GetURL()
}
func (r *Response) Context() context.Context {
	return r.request.Context()
}
func (r *Response) WithContext(ctx context.Context) pkg.Request {
	return r.request.WithContext(ctx)
}
func (r *Response) IsFile() bool {
	return r.request.IsFile()
}
func (r *Response) FileOptions() *pkg.FileOptions {
	return r.request.FileOptions()
}
func (r *Response) IsImage() bool {
	return r.request.IsImage()
}
func (r *Response) ImageOptions() *pkg.ImageOptions {
	return r.request.ImageOptions()
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
func (r *Response) MustXpathOne(path string) (result *xpath.Result) {
	result, _ = r.XpathOne(path)
	return
}
func (r *Response) MustXpathMany(path string) (results []*xpath.Result) {
	results, _ = r.XpathMany(path)
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
	return
}
func (r *Response) XpathOne(path string) (result *xpath.Result, err error) {
	var selector *xpath.Selector
	selector, err = r.Xpath()
	if err != nil {
		return
	}
	result = selector.One(path)
	return
}
func (r *Response) XpathMany(path string) (results []*xpath.Result, err error) {
	var selector *xpath.Selector
	selector, err = r.Xpath()
	if err != nil {
		return
	}
	results = selector.Many(path)
	return
}

func (r *Response) MustCss() (selector *css.Selector) {
	selector, _ = r.Css()
	return
}
func (r *Response) MustCssOne(path string) (result *css.Result) {
	result, _ = r.CssOne(path)
	return
}
func (r *Response) MustCssMany(path string) (results []*css.Result) {
	results, _ = r.CssMany(path)
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
	return
}
func (r *Response) CssOne(path string) (result *css.Result, err error) {
	var selector *css.Selector
	selector, err = r.Css()
	if err != nil {
		return
	}
	result = selector.One(path)
	return
}
func (r *Response) CssMany(path string) (results []*css.Result, err error) {
	var selector *css.Selector
	selector, err = r.Css()
	if err != nil {
		return
	}
	results = selector.Many(path)
	return
}

func (r *Response) MustJson() (selector *gjson.Selector) {
	selector, _ = r.Json()
	return
}
func (r *Response) MustJsonOne(path string) (result *gjson.Result) {
	result, _ = r.JsonOne(path)
	return
}
func (r *Response) MustJsonMany(path string) (results []*gjson.Result) {
	results, _ = r.JsonMany(path)
	return
}

// Json return a gjson
func (r *Response) Json() (selector *gjson.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.bodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = gjson.NewSelectorFromBytes(r.bodyBytes)
	return
}
func (r *Response) JsonOne(path string) (result *gjson.Result, err error) {
	var selector *gjson.Selector
	selector, err = r.Json()
	if err != nil {
		return
	}
	result = selector.One(path)
	return
}
func (r *Response) JsonMany(path string) (results []*gjson.Result, err error) {
	var selector *gjson.Selector
	selector, err = r.Json()
	if err != nil {
		return
	}
	results = selector.Many(path)
	return
}

func (r *Response) MustRe() (selector *re.Selector) {
	selector, _ = r.Re()
	return
}
func (r *Response) MustReOne(path string) (result *re.Result) {
	result, _ = r.ReOne(path)
	return
}
func (r *Response) MustReMany(path string) (results []*re.Result) {
	results, _ = r.ReMany(path)
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
	return
}
func (r *Response) ReOne(path string) (result *re.Result, err error) {
	var selector *re.Selector
	selector, err = r.Re()
	if err != nil {
		return
	}
	result = selector.One(path)
	return
}
func (r *Response) ReMany(path string) (results []*re.Result, err error) {
	var selector *re.Selector
	selector, err = r.Re()
	if err != nil {
		return
	}
	results = selector.Many(path)
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

	rootPath := ""
	rootPath = dataType.Tag.Get("_json")
	var rootJsonArray []*gjson.Selector
	if rootPath != "" {
		rootJsonArray = r.MustJson().ManySelector(rootPath)
		eleCount = len(rootJsonArray)
		if eleCount == 0 {
			return
		}
	}

	rootPath = dataType.Tag.Get("_xpath")
	var rootXpathArray []*xpath.Selector
	if rootPath != "" {
		rootXpathArray = r.MustXpath().ManySelector(rootPath)
		eleCount = len(rootXpathArray)
		if eleCount == 0 {
			return
		}
	}

	rootPath = dataType.Tag.Get("_css")
	var rootCssArray []*css.Selector
	if rootPath != "" {
		rootCssArray = r.MustCss().ManySelector(rootPath)
		eleCount = len(rootCssArray)
		if eleCount == 0 {
			return
		}
	}

	rootPath = dataType.Tag.Get("_re")
	var rootReArray []*re.Selector
	if rootPath != "" {
		rootReArray = r.MustRe().ManySelector(rootPath)
		eleCount = len(rootReArray)
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
				eleField := ele.Field(ii)
				kind := eleField.Kind()
				elePath := ""
				elePath = eleType.Field(ii).Tag.Get("_json")
				if elePath != "" {
					var rootJson *gjson.Selector
					if len(rootJsonArray) > 0 {
						rootJson = rootJsonArray[i]
					}
					if len(rootXpathArray) > 0 {
						rootJson, err = gjson.NewSelectorFromStr(rootXpathArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootCssArray) > 0 {
						rootJson, err = gjson.NewSelectorFromStr(rootCssArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootReArray) > 0 {
						rootJson, err = gjson.NewSelectorFromStr(rootReArray[i].String())
						if err != nil {
							return
						}
					}
					eleJson := rootJson.One(elePath)
					switch kind {
					case reflect.Int:
						eleField.SetInt(eleJson.Int64())
					case reflect.Int8:
						eleField.SetInt(eleJson.Int64())
					case reflect.Int16:
						eleField.SetInt(eleJson.Int64())
					case reflect.Int32:
						eleField.SetInt(eleJson.Int64())
					case reflect.Int64:
						eleField.SetInt(eleJson.Int64())
					case reflect.String:
						eleField.SetString(eleJson.String())
					case reflect.Bool:
						eleField.SetBool(eleJson.Bool())
					case reflect.Uint:
						eleField.SetUint(eleJson.Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleJson.Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleJson.Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleJson.Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleJson.Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleJson.Float64())
					case reflect.Float64:
						eleField.SetFloat(eleJson.Float64())
					}
					continue
				}
				elePath = eleType.Field(ii).Tag.Get("_xpath")
				if elePath != "" {
					var rootXpath *xpath.Selector
					if len(rootJsonArray) > 0 {
						rootXpath, err = xpath.NewSelectorFromStr(rootJsonArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootXpathArray) > 0 {
						rootXpath = rootXpathArray[i]
					}
					if len(rootCssArray) > 0 {
						rootXpath, err = xpath.NewSelectorFromStr(rootCssArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootReArray) > 0 {
						rootXpath, err = xpath.NewSelectorFromStr(rootReArray[i].String())
						if err != nil {
							return
						}
					}
					eleXpath := rootXpath.One(elePath)
					switch kind {
					case reflect.Int:
						eleField.SetInt(eleXpath.Int64())
					case reflect.Int8:
						eleField.SetInt(eleXpath.Int64())
					case reflect.Int16:
						eleField.SetInt(eleXpath.Int64())
					case reflect.Int32:
						eleField.SetInt(eleXpath.Int64())
					case reflect.Int64:
						eleField.SetInt(eleXpath.Int64())
					case reflect.String:
						eleField.SetString(eleXpath.String())
					case reflect.Bool:
						eleField.SetBool(eleXpath.Bool())
					case reflect.Uint:
						eleField.SetUint(eleXpath.Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleXpath.Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleXpath.Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleXpath.Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleXpath.Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleXpath.Float64())
					case reflect.Float64:
						eleField.SetFloat(eleXpath.Float64())
					}
					continue
				}
				elePath = eleType.Field(ii).Tag.Get("_css")
				if elePath != "" {
					var rootCss *css.Selector
					if len(rootJsonArray) > 0 {
						rootCss, err = css.NewSelectorFromStr(rootJsonArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootXpathArray) > 0 {
						rootCss, err = css.NewSelectorFromStr(rootXpathArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootCssArray) > 0 {
						rootCss = rootCssArray[i]
					}
					if len(rootReArray) > 0 {
						rootCss, err = css.NewSelectorFromStr(rootReArray[i].String())
						if err != nil {
							return
						}
					}
					eleCss := rootCss.One(elePath)
					switch kind {
					case reflect.Int:
						eleField.SetInt(eleCss.Int64())
					case reflect.Int8:
						eleField.SetInt(eleCss.Int64())
					case reflect.Int16:
						eleField.SetInt(eleCss.Int64())
					case reflect.Int32:
						eleField.SetInt(eleCss.Int64())
					case reflect.Int64:
						eleField.SetInt(eleCss.Int64())
					case reflect.String:
						eleField.SetString(eleCss.String())
					case reflect.Bool:
						eleField.SetBool(eleCss.Bool())
					case reflect.Uint:
						eleField.SetUint(eleCss.Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleCss.Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleCss.Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleCss.Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleCss.Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleCss.Float64())
					case reflect.Float64:
						eleField.SetFloat(eleCss.Float64())
					}
					continue
				}
				elePath = eleType.Field(ii).Tag.Get("_re")
				if elePath != "" {
					var rootRe *re.Selector
					if len(rootJsonArray) > 0 {
						rootRe, err = re.NewSelectorFromStr(rootJsonArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootXpathArray) > 0 {
						rootRe, err = re.NewSelectorFromStr(rootXpathArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootCssArray) > 0 {
						rootRe, err = re.NewSelectorFromStr(rootCssArray[i].String())
						if err != nil {
							return
						}
					}
					if len(rootReArray) > 0 {
						rootRe = rootReArray[i]
					}
					eleRe := rootRe.One(elePath)
					switch kind {
					case reflect.Int:
						eleField.SetInt(eleRe.Int64())
					case reflect.Int8:
						eleField.SetInt(eleRe.Int64())
					case reflect.Int16:
						eleField.SetInt(eleRe.Int64())
					case reflect.Int32:
						eleField.SetInt(eleRe.Int64())
					case reflect.Int64:
						eleField.SetInt(eleRe.Int64())
					case reflect.String:
						eleField.SetString(eleRe.String())
					case reflect.Bool:
						eleField.SetBool(eleRe.Bool())
					case reflect.Uint:
						eleField.SetUint(eleRe.Uint64())
					case reflect.Uint8:
						eleField.SetUint(eleRe.Uint64())
					case reflect.Uint16:
						eleField.SetUint(eleRe.Uint64())
					case reflect.Uint32:
						eleField.SetUint(eleRe.Uint64())
					case reflect.Uint64:
						eleField.SetUint(eleRe.Uint64())
					case reflect.Float32:
						eleField.SetFloat(eleRe.Float64())
					case reflect.Float64:
						eleField.SetFloat(eleRe.Float64())
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
			eleField := vValue.Field(ii)
			kind := eleField.Kind()
			elePath := ""
			elePath = eleType.Field(ii).Tag.Get("_json")
			if elePath != "" {
				var rootJson *gjson.Selector
				if len(rootJsonArray) > 0 {
					rootJson = rootJsonArray[0]
				}
				if len(rootXpathArray) > 0 {
					rootJson, err = gjson.NewSelectorFromStr(rootXpathArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootCssArray) > 0 {
					rootJson, err = gjson.NewSelectorFromStr(rootCssArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootReArray) > 0 {
					rootJson, err = gjson.NewSelectorFromStr(rootReArray[0].String())
					if err != nil {
						return
					}
				}
				eleJson := rootJson.One(elePath)
				switch kind {
				case reflect.Int:
					eleField.SetInt(eleJson.Int64())
				case reflect.Int8:
					eleField.SetInt(eleJson.Int64())
				case reflect.Int16:
					eleField.SetInt(eleJson.Int64())
				case reflect.Int32:
					eleField.SetInt(eleJson.Int64())
				case reflect.Int64:
					eleField.SetInt(eleJson.Int64())
				case reflect.String:
					eleField.SetString(eleJson.String())
				case reflect.Bool:
					eleField.SetBool(eleJson.Bool())
				case reflect.Uint:
					eleField.SetUint(eleJson.Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleJson.Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleJson.Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleJson.Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleJson.Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleJson.Float64())
				case reflect.Float64:
					eleField.SetFloat(eleJson.Float64())
				}
				continue
			}
			elePath = eleType.Field(ii).Tag.Get("_xpath")
			if elePath != "" {
				var rootXpath *xpath.Selector
				if len(rootJsonArray) > 0 {
					rootXpath, err = xpath.NewSelectorFromStr(rootJsonArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootXpathArray) > 0 {
					rootXpath = rootXpathArray[0]
				}
				if len(rootCssArray) > 0 {
					rootXpath, err = xpath.NewSelectorFromStr(rootCssArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootReArray) > 0 {
					rootXpath, err = xpath.NewSelectorFromStr(rootReArray[0].String())
					if err != nil {
						return
					}
				}
				eleXpath := rootXpath.One(elePath)
				switch kind {
				case reflect.Int:
					eleField.SetInt(eleXpath.Int64())
				case reflect.Int8:
					eleField.SetInt(eleXpath.Int64())
				case reflect.Int16:
					eleField.SetInt(eleXpath.Int64())
				case reflect.Int32:
					eleField.SetInt(eleXpath.Int64())
				case reflect.Int64:
					eleField.SetInt(eleXpath.Int64())
				case reflect.String:
					eleField.SetString(eleXpath.String())
				case reflect.Bool:
					eleField.SetBool(eleXpath.Bool())
				case reflect.Uint:
					eleField.SetUint(eleXpath.Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleXpath.Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleXpath.Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleXpath.Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleXpath.Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleXpath.Float64())
				case reflect.Float64:
					eleField.SetFloat(eleXpath.Float64())
				}
				continue
			}
			elePath = eleType.Field(ii).Tag.Get("_css")
			if elePath != "" {
				var rootCss *css.Selector
				if len(rootJsonArray) > 0 {
					rootCss, err = css.NewSelectorFromStr(rootJsonArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootXpathArray) > 0 {
					rootCss, err = css.NewSelectorFromStr(rootXpathArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootCssArray) > 0 {
					rootCss = rootCssArray[0]
				}
				if len(rootReArray) > 0 {
					rootCss, err = css.NewSelectorFromStr(rootReArray[0].String())
					if err != nil {
						return
					}
				}
				eleCss := rootCss.One(elePath)
				switch kind {
				case reflect.Int:
					eleField.SetInt(eleCss.Int64())
				case reflect.Int8:
					eleField.SetInt(eleCss.Int64())
				case reflect.Int16:
					eleField.SetInt(eleCss.Int64())
				case reflect.Int32:
					eleField.SetInt(eleCss.Int64())
				case reflect.Int64:
					eleField.SetInt(eleCss.Int64())
				case reflect.String:
					eleField.SetString(eleCss.String())
				case reflect.Bool:
					eleField.SetBool(eleCss.Bool())
				case reflect.Uint:
					eleField.SetUint(eleCss.Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleCss.Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleCss.Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleCss.Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleCss.Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleCss.Float64())
				case reflect.Float64:
					eleField.SetFloat(eleCss.Float64())
				}
				continue
			}
			elePath = eleType.Field(ii).Tag.Get("_re")
			if elePath != "" {
				var rootRe *re.Selector
				if len(rootJsonArray) > 0 {
					rootRe, err = re.NewSelectorFromStr(rootJsonArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootXpathArray) > 0 {
					rootRe, err = re.NewSelectorFromStr(rootXpathArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootCssArray) > 0 {
					rootRe, err = re.NewSelectorFromStr(rootCssArray[0].String())
					if err != nil {
						return
					}
				}
				if len(rootReArray) > 0 {
					rootRe = rootReArray[0]
				}
				eleRe := rootRe.One(elePath)
				switch kind {
				case reflect.Int:
					eleField.SetInt(eleRe.Int64())
				case reflect.Int8:
					eleField.SetInt(eleRe.Int64())
				case reflect.Int16:
					eleField.SetInt(eleRe.Int64())
				case reflect.Int32:
					eleField.SetInt(eleRe.Int64())
				case reflect.Int64:
					eleField.SetInt(eleRe.Int64())
				case reflect.String:
					eleField.SetString(eleRe.String())
				case reflect.Bool:
					eleField.SetBool(eleRe.Bool())
				case reflect.Uint:
					eleField.SetUint(eleRe.Uint64())
				case reflect.Uint8:
					eleField.SetUint(eleRe.Uint64())
				case reflect.Uint16:
					eleField.SetUint(eleRe.Uint64())
				case reflect.Uint32:
					eleField.SetUint(eleRe.Uint64())
				case reflect.Uint64:
					eleField.SetUint(eleRe.Uint64())
				case reflect.Float32:
					eleField.SetFloat(eleRe.Float64())
				case reflect.Float64:
					eleField.SetFloat(eleRe.Float64())
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
