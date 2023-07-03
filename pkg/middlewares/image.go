package middlewares

import (
	"bytes"
	"context"
	"errors"
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

	stats pkg.StatsWithImage
}

func (m *ImageMiddleware) ProcessResponse(ctx context.Context, response *pkg.Response) (err error) {
	if len(response.BodyBytes) == 0 {
		err = errors.New("BodyBytes empty")
		m.logger.Error(err)
		return
	}

	isImage := response.Request.GetImage()
	if isImage {
		img, name, e := image.Decode(bytes.NewReader(response.BodyBytes))
		if e != nil {
			err = e
			m.logger.Error(err)
			return
		}

		rect := img.Bounds()

		i := new(media.Image)
		i.SetName(utils.StrMd5(response.Request.URL.String()))
		i.SetExtension(name)
		i.SetWidth(rect.Dx())
		i.SetHeight(rect.Dy())
		response.Images = append(response.Images, i)
		if m.stats != nil {
			m.stats.IncImageTotal()
		}
	}

	return
}

func (m *ImageMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(ImageMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()

	m.stats, _ = crawler.GetStats().(pkg.StatsWithImage)
	return m
}
