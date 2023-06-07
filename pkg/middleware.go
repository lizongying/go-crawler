package pkg

import (
	"context"
)

type Middleware interface {
	SpiderStart(context.Context, Spider) error
	ProcessRequest(*Context) error
	ProcessResponse(*Context) error
	ProcessItem(*Context) error
	SpiderStop(context.Context) error
	SetName(string)
	GetName() string
	FromCrawler(Spider) Middleware
}

type UnimplementedMiddleware struct {
	name string
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

func (*UnimplementedMiddleware) ProcessItem(c *Context) (err error) {
	return c.NextItem()
}

func (*UnimplementedMiddleware) SpiderStop(context.Context) (err error) {
	return
}
func (m *UnimplementedMiddleware) SetName(name string) {
	m.name = name
}
func (m *UnimplementedMiddleware) GetName() string {
	return m.name
}

type ProcessFunc func(*Context) error

type Context struct {
	Request              *Request
	Response             *Response
	Item                 Item
	Middlewares          []Middleware
	processRequestIndex  uint8
	processResponseIndex uint8
	processItemIndex     uint8
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

func (m *Context) FirstItem() (err error) {
	m.processItemIndex = 0
	if m.processItemIndex >= uint8(len(m.Middlewares)) {
		return
	}

	err = m.Middlewares[0].ProcessItem(m)
	return
}

func (m *Context) NextItem() (err error) {
	m.processItemIndex++
	if m.processItemIndex >= uint8(len(m.Middlewares)) {
		return
	}

	err = m.Middlewares[m.processItemIndex].ProcessItem(m)
	return
}
