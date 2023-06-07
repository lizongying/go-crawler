package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
)

type RetryMiddleware struct {
	pkg.UnimplementedMiddleware
	logger        pkg.Logger
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

func (m *RetryMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	request := c.Request

	if request.RetryMaxTimes == 0 {
		request.RetryMaxTimes = m.retryMaxTimes
	}
	if len(request.OkHttpCodes) == 0 {
		request.OkHttpCodes = m.okHttpCodes
	}

	err = c.NextRequest()
	return
}

func (m *RetryMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	err = c.NextResponse()

	response := c.Response
	request := c.Request
	m.logger.Debug("after response")

	if request.RetryMaxTimes > 0 && (response.Response == nil || !utils.InSlice(response.StatusCode, request.OkHttpCodes)) {
		if request.RetryTimes < request.RetryMaxTimes {
			request.RetryTimes++
			m.logger.Info(request.UniqueKey, "retry times:", request.RetryTimes, "SpendTime:", request.SpendTime)
			err = c.FirstRequest()
			return
		}

		err = errors.New("retry max times")
		m.logger.Error(request.UniqueKey, err, request.RetryTimes, request.RetryMaxTimes)
		return
	}

	return
}

func (m *RetryMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RetryMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
