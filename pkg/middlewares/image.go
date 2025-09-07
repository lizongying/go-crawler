package middlewares

import (
	"bytes"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/media"
	"github.com/lizongying/go-crawler/pkg/utils"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ImageMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	store  pkg.Store
}

func (m *ImageMiddleware) ProcessResponse(c pkg.Context, response pkg.Response) (err error) {
	task := c.GetTask()
	if len(response.BodyBytes()) == 0 {
		m.logger.Debug("BodyBytes empty")
		return
	}

	isImage := response.IsImage()
	if isImage {
		options := response.ImageOptions()
		img, ext, e := image.Decode(bytes.NewReader(response.BodyBytes()))
		if e != nil {
			err = e
			m.logger.Error(err)
			return
		}

		rect := img.Bounds()

		i := new(media.Image)
		if options.Url {
			i.SetUrl(response.Url())
		}
		name := utils.StrMd5(response.Url())
		if options.Name {
			i.SetName(name)
		}
		if options.Ext {
			i.SetExt(ext)
		}
		if options.Width {
			i.SetWidth(rect.Dx())
		}
		if options.Height {
			i.SetHeight(rect.Dy())
		}

		key := fmt.Sprintf("%s.%s", name, ext)
		storePath := ""
		storePath, err = m.store.Save("", key, response.BodyBytes())
		if err != nil {
			m.logger.Error(err)
			return
		}
		i.SetStorePath(storePath)

		response.SetImages(append(response.Images(), i))
		stats, ok := task.GetStats().(pkg.StatsWithImage)
		if ok {
			stats.IncImageTotal()
		}
	}

	return
}

func (m *ImageMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(ImageMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	crawler := spider.GetCrawler()
	m.logger = spider.GetLogger()
	m.store, _ = crawler.GetStore(crawler.GetConfig().GetStorage())

	return m
}
