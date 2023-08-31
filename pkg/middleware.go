package pkg

import "context"

type Middleware interface {
	Start(context.Context, Spider) error
	ProcessRequest(Context, Request) error
	ProcessResponse(Context, Response) error
	ProcessError(Context, error)
	Stop(context.Context) error
	SetName(string)
	Name() string
	SetOrder(uint8)
	Order() uint8
	FromSpider(Spider) Middleware
}

type UnimplementedMiddleware struct {
	name    string
	order   uint8
	context context.Context
	spider  Spider
}

func (m *UnimplementedMiddleware) GetSpider() Spider {
	return m.spider
}
func (m *UnimplementedMiddleware) SetSpider(spider Spider) {
	m.spider = spider
}
func (m *UnimplementedMiddleware) WithContext(context context.Context) Middleware {
	m.context = context
	return m
}
func (m *UnimplementedMiddleware) Start(ctx context.Context, spider Spider) error {
	m.WithContext(ctx)
	m.spider = spider
	return nil
}
func (*UnimplementedMiddleware) ProcessRequest(Context, Request) error {
	return nil
}
func (*UnimplementedMiddleware) ProcessResponse(Context, Response) error {
	return nil
}
func (*UnimplementedMiddleware) ProcessError(Context, error) {
}
func (*UnimplementedMiddleware) Stop(context.Context) error {
	return nil
}
func (m *UnimplementedMiddleware) Name() string {
	return m.name
}
func (m *UnimplementedMiddleware) SetName(name string) {
	m.name = name
}
func (m *UnimplementedMiddleware) Order() uint8 {
	return m.order
}
func (m *UnimplementedMiddleware) SetOrder(order uint8) {
	m.order = order
}
func (m *UnimplementedMiddleware) FromSpider(spider Spider) Middleware {
	if m == nil {
		return new(UnimplementedMiddleware).FromSpider(spider)
	}

	m.spider = spider
	return m
}
