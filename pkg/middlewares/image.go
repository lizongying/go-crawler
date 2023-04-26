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
}

func (m *ImageMiddleware) GetName() string {
	return "image"
}

func (m *ImageMiddleware) ProcessResponse(_ context.Context, r *pkg.Response) (request *pkg.Request, response *pkg.Response, err error) {
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
	}

	return
}

func NewImageMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &ImageMiddleware{
		logger: logger,
	}
	return
}
