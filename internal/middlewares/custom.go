package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
)

type CustomMiddleware struct {
	pkg.UnimplementedMiddleware
	logger *logger.Logger

	spider pkg.Spider
}

func (m *CustomMiddleware) GetName() string {
	return "custom"
}

func (m *CustomMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.logger.Debug("start")
	return
}

func (m *CustomMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	if err = c.NextRequest(); err != nil {
		m.logger.Error(err)
	}
	return
}

func (m *CustomMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	if err = c.NextResponse(); err != nil {
		m.logger.Error(err)
	}
	return
}

func (m *CustomMiddleware) ProcessItem(c *pkg.Context) (err error) {
	if err = c.NextItem(); err != nil {
		m.logger.Error(err)
	}
	return
}

func (m *CustomMiddleware) SpiderStop(_ context.Context) (err error) {
	m.logger.Debug("stop")
	return
}

func NewCustomMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &CustomMiddleware{
		logger: logger,
	}
	return
}
