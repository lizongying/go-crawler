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
	ProcessRequest(*Context) error
	ProcessResponse(*Context) error
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

func (*UnimplementedMiddleware) ProcessRequest(c *Context) error {
	return c.NextRequest()
}

func (*UnimplementedMiddleware) ProcessResponse(c *Context) error {
	return c.NextResponse()
}

func (*UnimplementedMiddleware) ProcessItem(context.Context, *Item) (err error) {
	return
}

func (*UnimplementedMiddleware) SpiderStop(context.Context) (err error) {
	return
}

type ProcessFunc func(*Context) error

type Context struct {
	Request              *Request
	Response             *Response
	Middlewares          []Middleware
	processRequestIndex  uint8
	processResponseIndex uint8
}

func (m *Context) FirstRequest() (err error) {
	m.processRequestIndex = 0
	if m.processRequestIndex >= uint8(len(m.Middlewares)) {
		return
	}
	err = m.Middlewares[0].ProcessRequest(m)
	return
}

func (m *Context) NextRequest() (err error) {
	m.processRequestIndex++
	if m.processRequestIndex >= uint8(len(m.Middlewares)) {
		return
	}

	err = m.Middlewares[m.processRequestIndex].ProcessRequest(m)
	return
}

func (m *Context) FirstResponse() (err error) {
	m.processResponseIndex = 0
	if m.processResponseIndex >= uint8(len(m.Middlewares)) {
		return
	}

	err = m.Middlewares[0].ProcessResponse(m)
	return
}

func (m *Context) NextResponse() (err error) {
	m.processResponseIndex++
	if m.processResponseIndex >= uint8(len(m.Middlewares)) {
		return
	}

	err = m.Middlewares[m.processResponseIndex].ProcessResponse(m)
	return
}
