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

func (m *RetryMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	m.okHttpCodes = append(m.okHttpCodes, spider.GetOkHttpCodes()...)
	m.retryMaxTimes = spider.GetRetryMaxTimes()
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
	if retryMaxTimes > 0 && (response.GetResponse() == nil || !utils.InSlice(response.GetStatusCode(), okHttpCodes)) {
		if request.GetRetryTimes() < retryMaxTimes {
			request.SetRetryTimes(request.GetRetryTimes() + 1)
			m.logger.Info(request.GetUniqueKey(), "retry times:", request.GetRetryTimes(), "SpendTime:", request.GetSpendTime())
			err = pkg.ErrNeedRetry
			return
		}

		err = errors.New("retry max times")
		m.logger.Error(request.GetUniqueKey(), err, request.GetRetryTimes(), retryMaxTimes)
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
