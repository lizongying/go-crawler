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
	files  map[string]*map[string]*os.File
	mutex  sync.Mutex
	logger pkg.Logger
}

func (m *JsonLinesPipeline) ProcessItem(item pkg.Item) (err error) {
	task := item.GetContext().GetTask()
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		task.IncItemError()
		return
	}
	if item.Name() != pkg.ItemJsonl {
		m.logger.Warn("item not support", pkg.ItemKafka)
		return
	}
	itemJsonl, ok := item.GetItem().(*items.ItemJsonl)
	if !ok {
		m.logger.Warn("item parsing failed with", pkg.ItemKafka)
		return
	}

	if itemJsonl.GetFileName() == "" {
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

	filename := fmt.Sprintf("%s.jsonl", itemJsonl.GetFileName())

	m.mutex.Lock()

	files, ok := m.files[item.GetContext().GetTask().GetId()]
	if !ok {
		fs := make(map[string]*os.File)
		files = &fs
		m.files[item.GetContext().GetTask().GetId()] = files
	}

	file, ok := (*files)[filename]
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

	_, err = file.WriteString(fmt.Sprintf("%s\n", utils.UnsafeJSON(data)))
	if err != nil {
		m.logger.Error(err)
		task.IncItemError()
		return err
	}

	m.logger.Info("item saved:", filename)
	task.IncItemSuccess()
	return
}

func (m *JsonLinesPipeline) taskStopped(task pkg.Task) (err error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	files, ok := m.files[task.GetContext().GetTask().GetId()]
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

func (m *JsonLinesPipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(JsonLinesPipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	m.files = make(map[string]*map[string]*os.File)
	spider.GetCrawler().GetSignal().RegisterTaskChanged(m.taskStopped)
	return
}
