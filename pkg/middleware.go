package pkg

import (
	"context"
)

type Middleware interface {
	Start(context.Context, Crawler) error
	ProcessRequest(context.Context, *Request) error
	ProcessResponse(context.Context, *Response) error
	Stop(context.Context) error
	SetName(string)
	GetName() string
	SetOrder(uint8)
	GetOrder() uint8
	FromCrawler(Crawler) Middleware
}

type UnimplementedMiddleware struct {
	name  string
	order uint8
}

func (m *UnimplementedMiddleware) Start(_ context.Context, crawler Crawler) error {
	_ = m.FromCrawler(crawler)
	return nil
}
func (*UnimplementedMiddleware) ProcessRequest(context.Context, *Request) error {
	return nil
}
func (*UnimplementedMiddleware) ProcessResponse(context.Context, *Response) error {
	return nil
}
func (*UnimplementedMiddleware) Stop(context.Context) error {
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

func (m *UnimplementedMiddleware) FromCrawler(crawler Crawler) Middleware {
	if m == nil {
		return new(UnimplementedMiddleware).FromCrawler(crawler)
	}

	return m
}
