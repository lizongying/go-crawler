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
	okHttpCodes   []int
	retryMaxTimes uint8
}

func (m *RetryMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	request := response.Request

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
			m.logger.Info(request.GetUniqueKey(), "retry times:", request.RetryTimes, "SpendTime:", request.SpendTime)
			err = pkg.ErrNeedRetry
			return
		}

		err = errors.New("retry max times")
		m.logger.Error(request.GetUniqueKey(), err, request.RetryTimes, request.RetryMaxTimes)
		return
	}

	return
}

func (m *RetryMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(RetryMiddleware).FromCrawler(crawler)
	}
	m.logger = crawler.GetLogger()
	m.okHttpCodes = crawler.GetOkHttpCodes()
	m.okHttpCodes = append(m.okHttpCodes, http.StatusMovedPermanently, http.StatusFound)
	m.retryMaxTimes = crawler.GetRetryMaxTimes()
	return m
}
