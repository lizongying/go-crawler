package pkg

import (
	"context"
)

type Pipeline interface {
	Start(context.Context, Crawler) error
	ProcessItem(context.Context, Item) error
	Stop(context.Context) error
	SetName(string)
	GetName() string
	SetOrder(uint8)
	GetOrder() uint8
	FromCrawler(Crawler) Pipeline
}

type UnimplementedPipeline struct {
	name  string
	order uint8
}

func (p *UnimplementedPipeline) Start(_ context.Context, crawler Crawler) error {
	_ = p.FromCrawler(crawler)
	return nil
}
func (*UnimplementedPipeline) ProcessItem(context.Context, Item) error {
	return nil
}
func (*UnimplementedPipeline) Stop(context.Context) error {
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
func (p *UnimplementedPipeline) FromCrawler(crawler Crawler) Pipeline {
	if p == nil {
		return new(UnimplementedPipeline).FromCrawler(crawler)
	}

	return p
}
