package middlewares

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
)

type UrlMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlLengthLimit int
}

func (m *UrlMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	request := c.Request

	if m.urlLengthLimit < len(request.Url) {
		err = pkg.ErrUrlLengthLimit
		m.logger.Error(err)
		err = errors.Join(err, pkg.ErrIgnoreRequest)
		return
	}

	return
}

func (m *UrlMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	m.urlLengthLimit = spider.GetConfig().GetUrlLengthLimit()
	return m
}

func NewUrlMiddleware() pkg.Middleware {
	return &UrlMiddleware{}
}
