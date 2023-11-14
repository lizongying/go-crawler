package pipelines

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"reflect"
	"strings"
)

type ImagePipeline struct {
	pkg.UnimplementedPipeline
	logger pkg.Logger
}

func (m *ImagePipeline) ProcessItem(item pkg.Item) (err error) {
	if item == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	images := item.ImagesRequest()
	if len(images) == 0 {
		return
	}

	field, ok := reflect.TypeOf(item.Data()).Elem().FieldByName("Images")
	isUrl := false
	isName := false
	isExt := false
	isWidth := false
	isHeight := false
	if ok {
		tag := field.Tag.Get("field")
		isUrl = strings.Contains(tag, "url")
		isName = strings.Contains(tag, "name")
		isExt = strings.Contains(tag, "ext")
		isWidth = strings.Contains(tag, "width")
		isHeight = strings.Contains(tag, "height")
	}
	imageOptions := pkg.ImageOptions{
		FileOptions: pkg.FileOptions{
			Url:  isUrl,
			Name: isName,
			Ext:  isExt,
		},
		Width:  isWidth,
		Height: isHeight,
	}
	for _, i := range images {
		ctx := &crawlerContext.Context{}
		r, e := m.Spider().Request(ctx, i.SetImageOptions(imageOptions))
		if e != nil {
			m.logger.Error(e)
			continue
		}
		item.SetImages(r.Images())
	}

	return
}

func (m *ImagePipeline) FromSpider(spider pkg.Spider) (err error) {
	if m == nil {
		return new(ImagePipeline).FromSpider(spider)
	}

	if err = m.UnimplementedPipeline.FromSpider(spider); err != nil {
		return
	}
	m.logger = spider.GetLogger()
	return
}
