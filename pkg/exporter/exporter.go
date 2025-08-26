package exporter

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/pipelines"
	"reflect"
	"sort"
	"sync"
)

type Exporter struct {
	pipelines      []pkg.Pipeline
	processItemFns []func(ctx pkg.Item) error
	spider         pkg.Spider
	logger         pkg.Logger
	locker         sync.Mutex
}

func (e *Exporter) Export(item pkg.Item) (err error) {
	for _, v := range e.pipelines {
		if err = v.ProcessItem(item); err != nil {
			e.logger.Error(err)
		}
	}
	return
}

func (e *Exporter) Names() (pipelines map[uint8]string) {
	e.locker.Lock()
	defer e.locker.Unlock()

	pipelines = make(map[uint8]string)
	for _, v := range e.pipelines {
		pipelines[v.Order()] = v.Name()
	}

	return
}

func (e *Exporter) Pipelines() []pkg.Pipeline {
	return e.pipelines
}

func (e *Exporter) SetPipeline(pipeline pkg.Pipeline, order uint8) {
	e.locker.Lock()
	defer e.locker.Unlock()

	if err := pipeline.FromSpider(e.spider); err != nil {
		e.logger.Error(err)
		return
	}

	name := reflect.TypeOf(pipeline).Elem().String()
	pipeline.SetName(name)
	pipeline.SetOrder(order)
	for k, v := range e.pipelines {
		if v.Name() == name {
			if v.Order() == order {
				return
			} else {
				e.Remove(k)
				break
			}
		}
	}

	e.pipelines = append(e.pipelines, pipeline)

	sort.Slice(e.pipelines, func(i, j int) bool {
		return e.pipelines[i].Order() < e.pipelines[j].Order()
	})

	var processItemFns []func(pkg.Item) error
	for _, v := range e.pipelines {
		processItemFns = append(processItemFns, v.ProcessItem)
	}
	e.processItemFns = processItemFns
}

func (e *Exporter) Remove(index int) {
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

func (e *Exporter) Clean() {
	e.locker.Lock()
	defer e.locker.Unlock()

	e.pipelines = make([]pkg.Pipeline, 0)
}
func (e *Exporter) WithFilePipeline() {
	e.SetPipeline(new(pipelines.FilePipeline), 10)
}
func (e *Exporter) WithImagePipeline() {
	e.SetPipeline(new(pipelines.ImagePipeline), 20)
}
func (e *Exporter) WithDumpPipeline() {
	e.SetPipeline(new(pipelines.DumpPipeline), 30)
}
func (e *Exporter) WithFilterPipeline() {
	e.SetPipeline(new(pipelines.FilterPipeline), 200)
}
func (e *Exporter) WithNonePipeline() {
	e.SetPipeline(new(pipelines.NonePipeline), 101)
}
func (e *Exporter) WithCsvPipeline() {
	e.SetPipeline(new(pipelines.CsvPipeline), 102)
}
func (e *Exporter) WithJsonLinesPipeline() {
	e.SetPipeline(new(pipelines.JsonLinesPipeline), 103)
}
func (e *Exporter) WithMongoPipeline() {
	e.SetPipeline(new(pipelines.MongoPipeline), 104)
}
func (e *Exporter) WithSqlitePipeline() {
	e.SetPipeline(new(pipelines.SqlitePipeline), 105)
}
func (e *Exporter) WithMysqlPipeline() {
	e.SetPipeline(new(pipelines.MysqlPipeline), 106)
}
func (e *Exporter) WithKafkaPipeline() {
	e.SetPipeline(new(pipelines.KafkaPipeline), 107)
}
func (e *Exporter) WithCustomPipeline(pipeline pkg.Pipeline) {
	e.SetPipeline(pipeline, 110)
}
func (e *Exporter) FromSpider(spider pkg.Spider) pkg.Exporter {
	if e == nil {
		return new(Exporter).FromSpider(spider)
	}

	crawler := spider.GetCrawler()
	e.spider = spider
	e.logger = crawler.GetLogger()
	config := crawler.GetConfig()

	// set pipelines
	if config.GetEnableDumpPipeline() {
		e.WithDumpPipeline()
	}
	if config.GetEnableFilePipeline() {
		e.WithFilePipeline()
	}
	if config.GetEnableImagePipeline() {
		e.WithImagePipeline()
	}
	if config.GetEnableFilterPipeline() {
		e.WithFilterPipeline()
	}
	if config.GetEnableNonePipeline() {
		e.WithNonePipeline()
	}
	if config.GetEnableCsvPipeline() {
		e.WithCsvPipeline()
	}
	if config.GetEnableJsonLinesPipeline() {
		e.WithJsonLinesPipeline()
	}
	if config.GetEnableMongoPipeline() {
		e.WithMongoPipeline()
	}
	if config.GetEnableMysqlPipeline() {
		e.WithMysqlPipeline()
	}
	if config.GetEnableKafkaPipeline() {
		e.WithKafkaPipeline()
	}
	return e
}
