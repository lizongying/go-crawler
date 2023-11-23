package middlewares

import (
	"context"
	"fmt"
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

func (m *RetryMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	m.okHttpCodes = append(m.okHttpCodes, spider.OkHttpCodes()...)
	m.retryMaxTimes = spider.RetryMaxTimes()
	return
}

func (m *RetryMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
	request := response.GetRequest()

	retryMaxTimes := m.retryMaxTimes
	if request.GetRetryMaxTimes() != nil {
		retryMaxTimes = *request.GetRetryMaxTimes()
	}

	okHttpCodes := m.okHttpCodes
	if len(request.GetOkHttpCodes()) > 0 {
		okHttpCodes = request.GetOkHttpCodes()
	}

	if response.GetResponse() == nil {
		if request.GetRetryTimes() < retryMaxTimes {
			request.SetRetryTimes(request.GetRetryTimes() + 1)
			m.logger.Infof("retry times: %d/%d, response nil, SpendTime: %v, UniqueKey: %s", request.GetRetryTimes(), retryMaxTimes, request.GetSpendTime(), request.GetUniqueKey())
			err = pkg.ErrNeedRetry
			return
		}
		err = fmt.Errorf("response nil")
		return
	}

	if !utils.InSlice(response.StatusCode(), okHttpCodes) {
		if request.GetRetryTimes() < retryMaxTimes {
			request.SetRetryTimes(request.GetRetryTimes() + 1)
			m.logger.Infof("retry times: %d/%d, status code: %d, SpendTime: %v, UniqueKey: %s", request.GetRetryTimes(), retryMaxTimes, response.StatusCode(), request.GetSpendTime(), request.GetUniqueKey())
			err = pkg.ErrNeedRetry
			return
		}

		err = fmt.Errorf("status code error: %d", response.StatusCode())
		return
	}

	return
}

func (m *RetryMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RetryMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.okHttpCodes = append(m.okHttpCodes, http.StatusMovedPermanently, http.StatusFound)
	return m
}
