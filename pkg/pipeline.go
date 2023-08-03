package pkg

import "context"

type Pipeline interface {
	Start(context.Context, Spider) error
	ProcessItem(context.Context, Item) error
	Stop(context.Context) error
	SetName(string)
	GetName() string
	SetOrder(uint8)
	GetOrder() uint8
	FromSpider(Spider) Pipeline
	WithContext(context context.Context) Pipeline
}

type UnimplementedPipeline struct {
	name    string
	order   uint8
	context context.Context
	spider  Spider
}

func (p *UnimplementedPipeline) GetSpider() Spider {
	return p.spider
}
func (p *UnimplementedPipeline) WithContext(context context.Context) Pipeline {
	p.context = context
	return p
}
func (p *UnimplementedPipeline) Start(ctx context.Context, spider Spider) error {
	p.WithContext(ctx)
	p.spider = spider
	return nil
}
func (*UnimplementedPipeline) ProcessItem(context.Context, Item) error {
	return nil
}
func (*UnimplementedPipeline) Stop(context.Context) error {
	return nil
}
func (p *UnimplementedPipeline) GetName() string {
	return p.name
}
func (p *UnimplementedPipeline) SetName(name string) {
	p.name = name
}
func (p *UnimplementedPipeline) GetOrder() uint8 {
	return p.order
}
func (p *UnimplementedPipeline) SetOrder(order uint8) {
	p.order = order
}
func (p *UnimplementedPipeline) FromSpider(spider Spider) Pipeline {
	if p == nil {
		return new(UnimplementedPipeline).FromSpider(spider)
	}

	p.spider = spider
	return p
}
