package internal

import (
	"context"
)

type PipelineOrder struct {
	Pipeline Pipeline
	Order    int
}

type Pipeline interface {
	GetName() string
	SpiderStart(context.Context, Spider) error
	ProcessItem(context.Context, *Item) error
	SpiderStop(context.Context) error
}

type UnimplementedPipeline struct {
}

func (*UnimplementedPipeline) GetName() (name string) {
	return
}

func (*UnimplementedPipeline) SpiderStart(context.Context, Spider) (err error) {
	return
}

func (*UnimplementedPipeline) ProcessItem(context.Context, *Item) (err error) {
	return
}

func (*UnimplementedPipeline) SpiderStop(context.Context) (err error) {
	return
}
