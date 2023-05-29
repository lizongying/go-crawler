package pkg

import (
	"context"
)

type HttpClient interface {
	BuildRequest(context.Context, *Request) (err error)
	BuildResponse(context.Context, *Request) (response *Response, err error)
}
