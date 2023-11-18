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
	files  map[string]*map[string]*os.File
	mutex  sync.Mutex
	logger pkg.Logger
}

func (m *CsvPipeline) ProcessItem(item pkg.Item) (err error) {
	spider := m.Spider()
	task := item.GetContext().GetTask()
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	if item.Name() != pkg.ItemCsv {
		m.logger.Warn("item not support", pkg.ItemCsv)
		return
	}

	itemCsv, ok := item.GetItem().(*items.ItemCsv)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemCsv)
		return
	}

	if itemCsv.GetFileName() == "" {
		err = errors.New("fileName is empty")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	data := item.Data()
	if data == nil {
		err = errors.New("nil data")
		m.logger.Error(err)
		task.IncItemError()
		return
	}

	item.GetContext().GetItem().WithSaved(true)

	refType := reflect.TypeOf(data)
	refValue := reflect.ValueOf(data)
	if reflect.ValueOf(data).Kind() == reflect.Ptr {
		refType = refType.Elem()
		refValue = refValue.Elem()
	}

	var lines []string
	var columns []string

	filename := fmt.Sprintf("%s.csv", itemCsv.GetFileName())

	m.mutex.Lock()

	files, ok := m.files[item.GetContext().GetTask().GetId()]
	if !ok {
		fs := make(map[string]*os.File)
		files = &fs
		m.files[item.GetContext().GetTask().GetId()] = files
	}

	file, ok := (*files)[filename]
	create := false
	if !ok {
		if !utils.ExistsDir(filename) {
			err = os.MkdirAll(filepath.Dir(filename), 0744)
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				m.mutex.Unlock()
				return
			}
		}
		if !utils.ExistsFile(filename) {
			file, err = os.Create(filename)
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				m.mutex.Unlock()
				return
			}
			create = true
		} else {
			file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err != nil {
				m.logger.Error(err)
				task.IncItemError()
				m.mutex.Unlock()
				return
			}
		}

		(*files)[filename] = file
	}
	m.mutex.Unlock()

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
		task.IncItemError()
		return err
	}

	m.logger.Info("item saved:", filename)
	item.GetContext().GetItem().WithStatus(pkg.ItemStatusSuccess)
	spider.GetCrawler().GetSignal().ItemChanged(item)
	task.IncItemSuccess()
	return
}

func (m *CsvPipeline) taskStopped(ctx pkg.Context) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	files, ok := m.files[ctx.GetTask().GetId()]
	if !ok {
		return
	}

	for _, v := range *files {
		if err = v.Close(); err != nil {
			m.logger.Error(err)
		}
	}
	return
}

func (m *CsvPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(CsvPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	m.files = make(map[string]*map[string]*os.File)
	spider.GetCrawler().GetSignal().RegisterTaskChanged(m.taskStopped)
	return
}
