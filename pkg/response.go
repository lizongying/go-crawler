package pkg

import (
	"errors"
	xpath "github.com/lizongying/go-xpath/selector"
	"net/http"
)

type Response struct {
	*http.Response
	Request   *Request
	BodyBytes []byte
}

// Xpath returns a xpath selector
func (r *Response) Xpath() (x *xpath.Selector, err error) {
	if r == nil {
		err = errors.New("response is invalid")
		return
	}

	if len(r.BodyBytes) == 0 {
		err = errors.New("response body is empty")
		return
	}

	x, err = xpath.NewSelectorFromBytes(r.BodyBytes)
	if err != nil {
		return
	}

	return
}
