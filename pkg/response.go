package pkg

import (
	"errors"
	"github.com/lizongying/go-query/query"
	"github.com/lizongying/go-re/re"
	"github.com/lizongying/go-xpath/xpath"
	"github.com/tidwall/gjson"
	"net/http"
)

type Response struct {
	*http.Response
	Request   *Request
	BodyBytes []byte
	Files     []File
	Images    []Image
}

// Xpath returns a xpath selector
func (r *Response) Xpath() (selector *xpath.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.BodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = xpath.NewSelectorFromBytes(r.BodyBytes)
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

	if len(r.BodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = query.NewSelectorFromBytes(r.BodyBytes)
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

	if len(r.BodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	result = gjson.ParseBytes(r.BodyBytes)

	return
}

// Re return a regex
func (r *Response) Re() (selector *re.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.BodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	selector, err = re.NewSelectorFromBytes(r.BodyBytes)
	if err != nil {
		return
	}

	return
}
