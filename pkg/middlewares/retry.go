package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
)

type RetryMiddleware struct {
	pkg.UnimplementedMiddleware
	logger        pkg.Logger
	spider        pkg.Spider
	okHttpCodes   []int
	retryMaxTimes uint8
}

func (m *RetryMiddleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	m.okHttpCodes = spider.GetOkHttpCodes()
	m.okHttpCodes = append(m.okHttpCodes, http.StatusMovedPermanently, http.StatusFound)
	m.retryMaxTimes = spider.GetInfo().RetryMaxTimes
	return
}

func (m *RetryMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	err = c.NextResponse()

	response := c.Response
	request := c.Request
	m.logger.Debug("after response")

	retryMaxTimes := m.retryMaxTimes
	if request.RetryMaxTimes != nil {
		retryMaxTimes = *request.RetryMaxTimes
	}

	okHttpCodes := m.okHttpCodes
	if len(request.OkHttpCodes) > 0 {
		okHttpCodes = request.OkHttpCodes
	}

	if retryMaxTimes > 0 && (response.Response == nil || !utils.InSlice(response.StatusCode, okHttpCodes)) {
		if request.RetryTimes < retryMaxTimes {
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
