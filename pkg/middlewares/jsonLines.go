package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"sync"
)

type JsonLinesMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	spider pkg.Spider
	files  sync.Map
	stats  pkg.Stats
}

func (m *JsonLinesMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.stats = spider.GetStats()
	return
}

func (m *JsonLinesMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item, ok := c.Item.(*pkg.ItemJsonl)
	if !ok {
		m.logger.Warning("item not support jsonl")
		err = c.NextItem()
		return
	}

	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	if item.FileName == "" {
		err = errors.New("fileName is empty")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	data := item.GetData()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return
	}

	filename := fmt.Sprintf("%s.jsonl", item.FileName)
	var file *os.File
	fileValue, ok := m.files.Load(item.FileName)
	if !ok {
		if !utils.ExistsDir(filename) {
			err = os.MkdirAll(filepath.Dir(filename), 0744)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				err = c.NextItem()
				return
			}
		}
		if !utils.ExistsFile(filename) {
			file, err = os.Create(filename)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				err = c.NextItem()
				return
			}
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				err = c.NextItem()
				return
			}
		}
		m.files.Store(item.FileName, file)
	} else {
		file = fileValue.(*os.File)
	}

	_, err = file.WriteString(fmt.Sprintf("%s\n", utils.JsonStr(data)))
	if err != nil {
		m.logger.Error(err)
		m.stats.IncItemError()
		err = c.NextItem()
		return err
	}

	m.stats.IncItemSuccess()
	err = c.NextItem()
	return
}

func (m *JsonLinesMiddleware) SpiderStop(_ context.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func NewJsonLinesMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &JsonLinesMiddleware{
		logger: logger,
	}
	return
}
