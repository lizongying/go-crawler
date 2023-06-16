package exporter

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"sort"
	"sync"
)

type Exporter struct {
	pipelines      []pkg.Pipeline
	processItemFns []func(context.Context, pkg.Item) error
	crawler        pkg.Crawler
	logger         pkg.Logger
	locker         sync.Mutex
}

func (e *Exporter) Export(ctx context.Context, item pkg.Item) (err error) {
	for _, v := range e.pipelines {
		er := v.ProcessItem(ctx, item)
		if er != nil {
			e.logger.Error(err)
			err = errors.Join(err, er)
		}
	}
	return
}

func (e *Exporter) GetPipelineNames() (pipelines map[uint8]string) {
	e.locker.Lock()
	defer e.locker.Unlock()

	pipelines = make(map[uint8]string)
	for _, v := range e.pipelines {
		pipelines[v.GetOrder()] = v.GetName()
	}

	return
}

func (e *Exporter) GetPipelines() []pkg.Pipeline {
	return e.pipelines
}

func (e *Exporter) SetPipeline(pipeline pkg.Pipeline, order uint8) {
	e.locker.Lock()
	defer e.locker.Unlock()

	pipeline = pipeline.FromCrawler(e.crawler)

	name := reflect.TypeOf(pipeline).Elem().String()
	pipeline.SetName(name)
	pipeline.SetOrder(order)
	for k, v := range e.pipelines {
		if v.GetName() == name && v.GetOrder() != order {
			e.DelPipeline(k)
			break
		}
	}

	e.pipelines = append(e.pipelines, pipeline)

	sort.Slice(e.pipelines, func(i, j int) bool {
		return e.pipelines[i].GetOrder() < e.pipelines[j].GetOrder()
	})
	var processItemFns []func(context.Context, pkg.Item) error
	for _, v := range e.pipelines {
		processItemFns = append(processItemFns, v.ProcessItem)
	}
	e.processItemFns = processItemFns
}

func (e *Exporter) DelPipeline(index int) {
	e.locker.Lock()
	defer e.locker.Unlock()

	if index < 0 {
		return
	}
	if index >= len(e.pipelines) {
		return
	}

	e.pipelines = append(e.pipelines[:index], e.pipelines[index+1:]...)
	return
}

func (e *Exporter) CleanPipelines() {
	e.locker.Lock()
	defer e.locker.Unlock()

	e.pipelines = make([]pkg.Pipeline, 0)
}

func (e *Exporter) FromCrawler(crawler pkg.Crawler) pkg.Exporter {
	if e == nil {
		return new(Exporter).FromCrawler(crawler)
	}

	e.crawler = crawler
	e.logger = crawler.GetLogger()
	return e
}
