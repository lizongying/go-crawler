package middlewares

import (
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
)

type HttpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *HttpMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	spider := m.GetSpider()
	if request.GetMethod() == "" {
		if request.BodyStr() != "" {
			request.SetMethod(http.MethodPost)
		} else {
			request.SetMethod(http.MethodGet)
		}
	}

	if request.GetUrl() == "" {
		err = errors.New("url is empty")
		m.logger.Error(err)
		return
	}
	request.SetCreateTime(utils.NowStr())
	request.SetChecksum(utils.StrMd5(request.GetMethod(), request.GetUrl(), request.BodyStr()))

	canonicalHeaderKey := true
	if request.CanonicalHeaderKey() != nil {
		canonicalHeaderKey = *request.CanonicalHeaderKey()
	}
	if canonicalHeaderKey {
		for k, v := range request.Headers() {
			l := len(v)
			if l < 1 {
				continue
			}
			request.SetHeader(http.CanonicalHeaderKey(k), v[l-1])
		}
	}

	ok := spider.IsAllowedDomain(request.GetURL())
	if !ok {
		err = errors.New("it's not a allowed domain")
		m.logger.Error(err)
		return
	}

	return
}

func (m *HttpMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(HttpMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
