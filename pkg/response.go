package pkg

import (
	"errors"
	query "github.com/lizongying/go-css/selector"
	xpath "github.com/lizongying/go-xpath/selector"
	"net/http"
)

type Response struct {
	*http.Response
	Request   *Request
	BodyBytes []byte
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
