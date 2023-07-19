package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"net/http"
)

type HttpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger  pkg.Logger
	crawler pkg.Crawler
	stats   pkg.Stats
}

func (m *HttpMiddleware) ProcessRequest(ctx context.Context, request pkg.Request) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if request.GetMethod() == "" {
		if request.GetBody() != "" {
			request.SetMethod("POST")
		} else {
			request.SetMethod("GET")
		}
	}

	if request.GetUrl() == "" {
		err = errors.New("url is empty")
		m.logger.Error(err)
		m.stats.IncRequestError()
		return
	}
	request.SetCreateTime(utils.NowStr())
	request.SetChecksum(utils.StrMd5(request.GetMethod(), request.GetUrl(), request.GetBody()))

	canonicalHeaderKey := true
	if request.GetCanonicalHeaderKey() != nil {
		canonicalHeaderKey = *request.GetCanonicalHeaderKey()
	}
	if canonicalHeaderKey {
		for k, v := range request.GetHeaders() {
			l := len(v)
			if l < 1 {
				continue
			}
			request.SetHeader(http.CanonicalHeaderKey(k), v[l-1])
		}
	}

	ok := m.crawler.IsAllowedDomain(request.GetURL())
	if !ok {
		err = errors.New("it's not a allowed domain")
		m.logger.Error(err)
		m.stats.IncRequestError()
		return
	}

	return
}

func (m *HttpMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(HttpMiddleware).FromCrawler(crawler)
	}

	m.crawler = crawler
	m.logger = crawler.GetLogger()
	m.stats = crawler.GetStats()
	return m
}
