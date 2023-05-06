package middlewares

import (
	"bytes"
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ImageMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	spider pkg.Spider
	stats  pkg.StatsWithImage
}

func (m *ImageMiddleware) GetName() string {
	return "image"
}

func (m *ImageMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.stats, _ = spider.GetStats().(pkg.StatsWithImage)
	return
}

func (m *ImageMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	r := c.Response
	m.logger.Debug("response body len:", len(r.BodyBytes))

	if len(r.BodyBytes) == 0 {
		err = errors.New("BodyBytes empty")
		m.logger.Error(err)
		return
	}

	extra, ok := r.Request.Extra.(pkg.OptionImage)
	if ok {
		img, name, e := image.Decode(bytes.NewReader(r.BodyBytes))
		if e != nil {
			err = e
			m.logger.Error(err)
			return
		}

		rect := img.Bounds()
		extra.SetName(utils.StrMd5(r.Request.URL.String()))
		extra.SetExtension(name)
		extra.SetWidth(rect.Dx())
		extra.SetHeight(rect.Dy())
		if m.stats != nil {
			m.stats.IncImageTotal()
		}
	}

	err = c.NextResponse()
	return
}

func NewImageMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &ImageMiddleware{
		logger: logger,
	}
	return
}
