package pipelines

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"sync"
)

type JsonLinesPipeline struct {
	pkg.UnimplementedPipeline
	files  sync.Map
	logger pkg.Logger
}

func (m *JsonLinesPipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	spider := m.GetSpider()
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}
	if itemWithContext.Name() != pkg.ItemJsonl {
		m.logger.Warn("item not support", pkg.ItemKafka)
		return
	}
	itemJsonl, ok := itemWithContext.GetItem().(*items.ItemJsonl)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemKafka)
		return
	}

	if itemJsonl.GetFileName() == "" {
		err = errors.New("fileName is empty")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	data := itemWithContext.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	filename := fmt.Sprintf("%s.jsonl", itemJsonl.GetFileName())
	var file *os.File
	fileValue, ok := m.files.Load(itemJsonl.GetFileName())
	if !ok {
		if !utils.ExistsDir(filename) {
			err = os.MkdirAll(filepath.Dir(filename), 0744)
			if err != nil {
				m.logger.Error(err)
				spider.IncItemError()
				return
			}
		}
		if !utils.ExistsFile(filename) {
			file, err = os.Create(filename)
			if err != nil {
				m.logger.Error(err)
				spider.IncItemError()
				return
			}
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				spider.IncItemError()
				return
			}
		}
		m.files.Store(itemJsonl.GetFileName(), file)
	} else {
		file = fileValue.(*os.File)
	}

	_, err = file.WriteString(fmt.Sprintf("%s\n", utils.JsonStr(data)))
	if err != nil {
		m.logger.Error(err)
		spider.IncItemError()
		return err
	}

	m.logger.Info("item saved:", filename)
	spider.GetCrawler().GetSignal().ItemSaved(itemWithContext)
	spider.IncItemSuccess()
	return
}

func (m *JsonLinesPipeline) SpiderStop(_ pkg.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func (m *JsonLinesPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(JsonLinesPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
