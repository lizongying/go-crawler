package pkg

import (
	"context"
)

type Middleware interface {
	SpiderStart(context.Context, Spider) error
	ProcessRequest(context.Context, *Request) error
	ProcessResponse(context.Context, *Response) error
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

func (m *UnimplementedMiddleware) SpiderStart(_ context.Context, spider Spider) error {
	_ = m.FromCrawler(spider)
	return nil
}
func (*UnimplementedMiddleware) ProcessRequest(context.Context, *Request) error {
	return nil
}
func (*UnimplementedMiddleware) ProcessResponse(context.Context, *Response) error {
	return nil
}
func (*UnimplementedMiddleware) SpiderStop(context.Context) error {
	return nil
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

func (m *UnimplementedMiddleware) FromCrawler(spider Spider) Middleware {
	if m == nil {
		return new(UnimplementedMiddleware).FromCrawler(spider)
	}

	return m
}
