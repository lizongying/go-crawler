package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/logger"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type RetryMiddleware struct {
	pkg.UnimplementedMiddleware
	logger        *logger.Logger
	spider        pkg.Spider
	okHttpCodes   []int
	retryMaxTimes int
}

func (m *RetryMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.okHttpCodes = spider.GetOkHttpCodes()
	m.retryMaxTimes = spider.GetInfo().RetryMaxTimes
	return
}

func (m *RetryMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	r := c.Response
	request := c.Request
	m.logger.Debug("response len:", len(r.BodyBytes))
	okHttpCodes := m.okHttpCodes
	if len(c.Request.OkHttpCodes) > 0 {
		okHttpCodes = c.Request.OkHttpCodes
	}
	retryMaxTimes := m.retryMaxTimes
	if request.RetryMaxTimes > 0 {
		retryMaxTimes = request.RetryMaxTimes
	}
	if request.RetryMaxTimes < 0 {
		retryMaxTimes = 0
	}
	if !utils.InSlice(r.StatusCode, okHttpCodes) {
		request.RetryTimes++

		if request.RetryTimes > retryMaxTimes {
			err = errors.New("retry max times")
			m.logger.Error(request.UniqueKey, err, retryMaxTimes)
			return
		}
		m.logger.Info(request.UniqueKey, "retry times", request.RetryTimes)
		err = c.FirstRequest()
		return
	}

	err = c.NextResponse()
	return
}

func NewRetryMiddleware(logger *logger.Logger) (m pkg.Middleware) {
	m = &RetryMiddleware{
		logger: logger,
	}
	return
}
