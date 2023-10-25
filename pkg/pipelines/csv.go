package pipelines

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/items"
	"github.com/lizongying/go-crawler/pkg/utils"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
)

type CsvPipeline struct {
	pkg.UnimplementedPipeline
	files  sync.Map
	logger pkg.Logger
}

func (m *CsvPipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	spider := m.GetSpider()
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		spider.IncItemError()
		return
	}

	if itemWithContext.Name() != pkg.ItemCsv {
		m.logger.Warn("item not support", pkg.ItemCsv)
		return
	}

	itemCsv, ok := itemWithContext.GetItem().(*items.ItemCsv)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemCsv)
		return
	}

	if itemCsv.GetFileName() == "" {
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

	refType := reflect.TypeOf(data)
	refValue := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	var lines []string
	var columns []string
	filename := fmt.Sprintf("%s.csv", itemCsv.GetFileName())
	var file *os.File
	fileValue, ok := m.files.Load(itemCsv.GetFileName())
	create := false
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
			create = true
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				spider.IncItemError()
				return
			}
		}
		m.files.Store(itemCsv.GetFileName(), file)
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
		spider.IncItemError()
		return err
	}

	m.logger.Info("item saved:", filename)
	spider.GetCrawler().GetSignal().ItemSaved(itemWithContext)
	spider.IncItemSuccess()
	return
}

func (m *CsvPipeline) Stop(_ pkg.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func (m *CsvPipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(CsvPipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
