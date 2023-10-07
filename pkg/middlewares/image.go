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

func (m *ImageMiddleware) ProcessResponse(ctx pkg.Context, response pkg.Response) (err error) {
	spider := m.GetSpider()
	if len(response.BodyBytes()) == 0 {
		m.logger.Debug("BodyBytes empty")
		return
	}

	isImage := response.Image()
	if isImage {
		img, ext, e := image.Decode(bytes.NewReader(response.BodyBytes()))
		if e != nil {
			err = e
			m.logger.Error(err)
			return
		}

		rect := img.Bounds()

		i := new(media.Image)
		name := utils.StrMd5(response.GetUrl())
		i.SetName(name)
		i.SetExtension(ext)
		i.SetWidth(rect.Dx())
		i.SetHeight(rect.Dy())

		key := fmt.Sprintf("%s.%s", name, ext)
		storePath := ""
		storePath, err = m.store.Save("", key, response.BodyBytes())
		if err != nil {
			m.logger.Error(err)
			return
		}
		i.SetStorePath(storePath)

		response.SetImages(append(response.Images(), i))
		stats, ok := spider.GetStats().(pkg.StatsWithImage)
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
	m.store = crawler.GetStore()

	return m
}
