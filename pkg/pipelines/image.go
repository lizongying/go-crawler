package pipelines

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"reflect"
	"strings"
)

type ImagePipeline struct {
	pkg.UnimplementedPipeline
	scheduler pkg.Scheduler
	logger    pkg.Logger
}

func (m *ImagePipeline) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedPipeline.Start(ctx, spider)
	m.scheduler = spider.GetScheduler()
	return nil
}

func (m *ImagePipeline) ProcessItem(itemWithContext pkg.ItemWithContext) (err error) {
	if itemWithContext == nil {
		err = errors.New("nil item")
		m.logger.Error(err)
		return
	}

	images := itemWithContext.ImagesRequest()
	if len(images) == 0 {
		return
	}

	field, ok := reflect.TypeOf(itemWithContext.Data()).Elem().FieldByName("Images")
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
		r, e := m.scheduler.Request(ctx, i.SetImageOptions(imageOptions))
		if e != nil {
			m.logger.Error(e)
			continue
		}
		itemWithContext.SetImages(r.Images())
	}

	return
}

func (m *ImagePipeline) FromSpider(spider pkg.Spider) pkg.Pipeline {
	if m == nil {
		return new(ImagePipeline).FromSpider(spider)
	}

	m.UnimplementedPipeline.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
