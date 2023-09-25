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
	if request.RetryMaxTimes() != nil {
		retryMaxTimes = *request.RetryMaxTimes()
	}

	okHttpCodes := m.okHttpCodes
	if len(request.OkHttpCodes()) > 0 {
		okHttpCodes = request.OkHttpCodes()
	}
	if retryMaxTimes > 0 && response.GetResponse() == nil {
		if request.RetryTimes() < retryMaxTimes {
			request.SetRetryTimes(request.RetryTimes() + 1)
			m.logger.Info(request.UniqueKey(), "retry times:", request.RetryTimes(), "SpendTime:", request.SpendTime())
			err = pkg.ErrNeedRetry
			return
		}
		err = fmt.Errorf("response nil")
		m.logger.Error(request.UniqueKey(), err, request.RetryTimes(), retryMaxTimes)
		return
	}

	if retryMaxTimes > 0 && !utils.InSlice(response.StatusCode(), okHttpCodes) {
		if request.RetryTimes() < retryMaxTimes {
			request.SetRetryTimes(request.RetryTimes() + 1)
			m.logger.Info(request.UniqueKey(), "retry times:", request.RetryTimes(), "SpendTime:", request.SpendTime())
			err = pkg.ErrNeedRetry
			return
		}

		err = fmt.Errorf("status code error: %d", response.StatusCode())
		m.logger.Error(request.UniqueKey(), err, request.RetryTimes(), retryMaxTimes)
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
