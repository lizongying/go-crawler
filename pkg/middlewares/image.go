package middlewares

import (
	"bytes"
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
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

func (m *ImageMiddleware) ProcessRequest(_ context.Context, _ *pkg.Request) (request *pkg.Request, response *pkg.Response, err error) {
	extra := request.Extra.(pkg.OptionImage)
	extra.SetName("test3")
	return
}

func (m *ImageMiddleware) ProcessResponse(_ context.Context, r *pkg.Response) (request *pkg.Request, response *pkg.Response, err error) {
	img, name, err := image.Decode(bytes.NewReader(r.BodyBytes))
	if err != nil {
		m.logger.Error(err)
		return
	}

	extra, ok := r.Request.Extra.(pkg.OptionImage)
	if ok {
		rect := img.Bounds()
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
