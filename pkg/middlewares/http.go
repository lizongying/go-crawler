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

func (m *HttpMiddleware) ProcessRequest(ctx context.Context, request *pkg.Request) (err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	if request.GetMethod() == "" {
		request.SetMethod("GET")
	}
	if request.GetUrl() == "" {
		err = errors.New("url is empty")
		return
	}
	request.CreateTime = utils.NowStr()
	request.Checksum = utils.StrMd5(request.GetMethod(), request.GetUrl(), request.GetBody())
	if request.GetCanonicalHeaderKey() {
		headers := make(map[string][]string)
		for k, v := range request.Header {
			headers[http.CanonicalHeaderKey(k)] = v
		}
		request.Header = headers
	}

	request.Request = *request.WithContext(ctx)
	if err != nil {
		m.logger.Error(err)
		m.stats.IncRequestError()
		return
	}

	if request.Header == nil {
		request.Header = make(http.Header)
	}
	request.Request.Header = request.Header

	ok := m.crawler.IsAllowedDomain(request.URL)
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
