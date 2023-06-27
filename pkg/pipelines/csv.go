package pipelines

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

type CsvPipeline struct {
	pkg.UnimplementedPipeline
	files  sync.Map
	stats  pkg.Stats
	logger pkg.Logger
}

func (m *CsvPipeline) ProcessItem(ctx context.Context, item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		m.stats.IncItemError()
		return
	}

	itemCsv, ok := item.(*pkg.ItemCsv)
	if !ok {
		m.logger.Warn("item not support csv")
		return
	}

	if itemCsv.FileName == "" {
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

	refType := reflect.TypeOf(data)
	refValue := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	var lines []string
	var columns []string
	filename := fmt.Sprintf("%s.csv", itemCsv.FileName)
	var file *os.File
	fileValue, ok := m.files.Load(itemCsv.FileName)
	create := false
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
			create = true
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				m.stats.IncItemError()
				return
			}
		}
		m.files.Store(itemCsv.FileName, file)
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
		return err
	}

	m.stats.IncItemSuccess()
	return
}

func (m *CsvPipeline) Stop(_ context.Context) (err error) {
	m.files.Range(func(key, value any) bool {
		err = value.(*os.File).Close()
		if err != nil {
			m.logger.Error(err)
		}
		return true
	})
	return
}

func (m *CsvPipeline) FromCrawler(crawler pkg.Crawler) pkg.Pipeline {
	if m == nil {
		return new(CsvPipeline).FromCrawler(crawler)
	}

	m.stats = crawler.GetStats()
	m.logger = crawler.GetLogger()
	return m
}
