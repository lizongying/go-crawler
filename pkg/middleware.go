package pkg

import (
	"context"
)

type Middleware interface {
	SpiderStart(context.Context, Spider) error
	ProcessRequest(context.Context, *Request) error
	ProcessResponse(context.Context, *Response) error
	ProcessItem(*Context) error
	SpiderStop(context.Context) error
	SetName(string)
	GetName() string
	SetOrder(uint8)
	GetOrder() uint8
	FromCrawler(Spider) Middleware
}

type UnimplementedMiddleware struct {
	name  string
	order uint8
}

func (*UnimplementedMiddleware) SpiderStart(context.Context, Spider) (err error) {
	return
}

func (*UnimplementedMiddleware) ProcessRequest(context.Context, *Request) error {
	return nil
}

func (*UnimplementedMiddleware) ProcessResponse(context.Context, *Response) error {
	return nil
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
func (m *UnimplementedMiddleware) SetOrder(order uint8) {
	m.order = order
}
func (m *UnimplementedMiddleware) GetOrder() uint8 {
	return m.order
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

	ctx context.Context
}

func (m *Context) SetContext(ctx context.Context) {
	m.ctx = ctx
}
func (m *Context) GetContext() context.Context {
	return m.ctx
}

func (m *Context) FirstItem() (err error) {
	m.processItemIndex = 0
	if m.processItemIndex >= uint8(len(m.Middlewares)-1) {
		return
	}

	err = m.Middlewares[0].ProcessItem(m)
	return
}

func (m *Context) NextItem() (err error) {
	m.processItemIndex++
	if m.processItemIndex >= uint8(len(m.Middlewares)-1) {
		return
	}

	err = m.Middlewares[m.processItemIndex].ProcessItem(m)
	return
}
