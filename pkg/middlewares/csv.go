package middlewares

import (
	"context"
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

type CsvMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	spider pkg.Spider
	files  sync.Map
	stats  pkg.Stats
}

func (m *CsvMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.stats = spider.GetStats()
	return
}

func (m *CsvMiddleware) ProcessItem(c *pkg.Context) (err error) {
	item, ok := c.Item.(*pkg.ItemCsv)
	if !ok {
		m.logger.Warn("item not support csv")
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

	refType := reflect.TypeOf(data).Elem()
	refValue := reflect.ValueOf(data).Elem()

	var lines []string
	var columns []string
	filename := fmt.Sprintf("%s.csv", item.FileName)
	var file *os.File
	fileValue, ok := m.files.Load(item.FileName)
	create := false
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
			create = true
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

	for i := 0; i < refType.NumField(); i++ {
		if create {
			column := refType.Field(i).Tag.Get("column")
			if column == "" {
				column = refType.Field(i).Name
			}
			if strings.Contains(column, `"`) {
				column = strings.ReplaceAll(column, `"`, `""`)
			}
			if strings.Contains(column, `"`) || strings.Contains(column, ",") {
				column = fmt.Sprintf(`"%s"`, column)
			}
			columns = append(columns, column)
		}

		line := fmt.Sprint(refValue.Field(i))
		if strings.Contains(line, `"`) {
			line = strings.ReplaceAll(line, `"`, `""`)
		}
		if strings.Contains(line, `"`) || strings.Contains(line, ",") {
			line = fmt.Sprintf(`"%s"`, line)
		}
		lines = append(lines, line)
	}
	if create {
		_, err = file.WriteString(fmt.Sprintf("%s\n", strings.Join(columns, ",")))
	}
	_, err = file.WriteString(fmt.Sprintf("%s\n", strings.Join(lines, ",")))
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

func (m *CsvMiddleware) SpiderStop(_ context.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func (m *CsvMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CsvMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
