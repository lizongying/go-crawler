package pkg

import (
	"context"
)

type Pipeline interface {
	SpiderStart(context.Context, Spider) error
	ProcessItem(context.Context, Item) error
	SpiderStop(context.Context) error
	SetName(string)
	GetName() string
	SetOrder(uint8)
	GetOrder() uint8
	FromCrawler(Spider) Pipeline
}

type UnimplementedPipeline struct {
	name  string
	order uint8
}

func (p *UnimplementedPipeline) SpiderStart(_ context.Context, spider Spider) error {
	_ = p.FromCrawler(spider)
	return nil
}
func (*UnimplementedPipeline) ProcessItem(context.Context, Item) error {
	return nil
}
func (*UnimplementedPipeline) SpiderStop(context.Context) error {
	return nil
}
func (p *UnimplementedPipeline) SetName(name string) {
	p.name = name
}
func (p *UnimplementedPipeline) GetName() string {
	return p.name
}
func (p *UnimplementedPipeline) SetOrder(order uint8) {
	p.order = order
}
func (p *UnimplementedPipeline) GetOrder() uint8 {
	return p.order
}
func (p *UnimplementedPipeline) FromCrawler(spider Spider) Pipeline {
	if p == nil {
		return new(UnimplementedPipeline).FromCrawler(spider)
	}

	return p
}
