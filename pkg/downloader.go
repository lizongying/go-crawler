package pkg

import (
	"context"
)

type Downloader interface {
	SetMiddlewares([]Middleware)
	DoRequest(context.Context, *Request) (*Response, error)
}
