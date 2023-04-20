package middlewares

import (
	"bytes"
	"context"
	"github.com/lizongying/go-crawler/internal"
	"github.com/lizongying/go-crawler/internal/logger"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type ImageMiddleware struct {
	internal.UnimplementedMiddleware
	logger *logger.Logger
}

func (m *ImageMiddleware) GetName() string {
	return "image"
}

func (m *ImageMiddleware) ProcessRequest(_ context.Context, _ *internal.Request) (request *internal.Request, response *internal.Response, err error) {
	extra := request.Extra.(internal.OptionImage)
	extra.SetName("test3")
	return
}

func (m *ImageMiddleware) ProcessResponse(_ context.Context, r *internal.Response) (request *internal.Request, response *internal.Response, err error) {
	img, name, err := image.Decode(bytes.NewReader(r.BodyBytes))
	if err != nil {
		m.logger.Error(err)
		return
	}

	extra, ok := r.Request.Extra.(internal.OptionImage)
	if ok {
		rect := img.Bounds()
		extra.SetExtension(name)
		extra.SetWidth(rect.Dx())
		extra.SetHeight(rect.Dy())
	}

	return
}

func NewImageMiddleware(logger *logger.Logger) (m internal.Middleware) {
	m = &ImageMiddleware{
		logger: logger,
	}
	return
}
