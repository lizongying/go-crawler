package pkg

import (
	"context"
)

type MiddlewareOrder struct {
	Middleware Middleware
	Order      int
}

type Middleware interface {
	GetName() string
	SpiderStart(context.Context, Spider) error
	ProcessRequest(context.Context, *Request) (*Request, *Response, error)
	ProcessResponse(context.Context, *Response) (*Request, *Response, error)
	ProcessItem(context.Context, *Item) error
	SpiderStop(context.Context) error
}

type UnimplementedMiddleware struct {
}

func (*UnimplementedMiddleware) GetName() (name string) {
	return
}

func (*UnimplementedMiddleware) SpiderStart(context.Context, Spider) (err error) {
	return
}

func (*UnimplementedMiddleware) ProcessRequest(context.Context, *Request) (request *Request, response *Response, err error) {
	return
}

func (*UnimplementedMiddleware) ProcessResponse(context.Context, *Response) (request *Request, response *Response, err error) {
	return
}

func (*UnimplementedMiddleware) ProcessItem(context.Context, *Item) (err error) {
	return
}

func (*UnimplementedMiddleware) SpiderStop(context.Context) (err error) {
	return
}
