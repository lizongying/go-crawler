package pipelines

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"sync"
)

type JsonLinesPipeline struct {
	pkg.UnimplementedPipeline
	files  sync.Map
	stats  pkg.Stats
	logger pkg.Logger
}

func (m *JsonLinesPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	itemJsonl, ok := item.(*pkg.ItemJsonl)
	if !ok {
		m.logger.Warn("item not support jsonl")
		return
	}

	if itemJsonl.FileName == "" {
		err = errors.New("fileName is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	filename := fmt.Sprintf("%s.jsonl", itemJsonl.FileName)
	var file *os.File
	fileValue, ok := m.files.Load(itemJsonl.FileName)
	if !ok {
		if !utils.ExistsDir(filename) {
			err = os.MkdirAll(filepath.Dir(filename), 0744)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}
		}
		if !utils.ExistsFile(filename) {
			file, err = os.Create(filename)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}
		}
		m.files.Store(itemJsonl.FileName, file)
	} else {
		file = fileValue.(*os.File)
	}

	_, err = file.WriteString(fmt.Sprintf("%s\n", utils.JsonStr(data)))
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		return err
	}

	m.stats.IncItemSuccess()
	return
}

func (m *JsonLinesPipeline) SpiderStop(_ context.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func (m *JsonLinesPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(JsonLinesPipeline).FromCrawler(crawler)
	}

	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	return m
}
