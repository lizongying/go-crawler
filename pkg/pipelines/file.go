package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"reflect"
	"strings"
)

type FilePipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *FilePipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	files := item.FilesRequest()
	if len(files) == 0 {
		return
	}

	isUrl := false
	isName := false
	isExt := false
	t := reflect.TypeOf(item.Data()).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if tag, ok := field.Tag.Lookup("file"); ok {
			isUrl = strings.Contains(tag, "url")
			isName = strings.Contains(tag, "name")
			isExt = strings.Contains(tag, "ext")
			break
		}
	}

	fileOptions := pkg.FileOptions{
		Url:  isUrl,
		Name: isName,
		Ext:  isExt,
	}

	ctx := item.GetContext()

	for _, i := range files {
		r, e := m.Spider().Request(ctx, i.SetFileOptions(fileOptions))
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetFiles(r.Files())
	}

	return
}

func (m *FilePipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(FilePipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
