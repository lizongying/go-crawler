package middlewares

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"io"
	"net/http"
	"net/url"
	"strings"
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

	request.CreateTime = utils.NowStr()
	request.Checksum = utils.StrMd5(request.Method, request.Url, request.BodyStr)
	if request.CanonicalHeaderKey {
		headers := make(map[string][]string)
		for k, v := range request.Header {
			headers[http.CanonicalHeaderKey(k)] = v
		}
		request.Header = headers
	}

	if request.Request == nil {
		Url, e := url.Parse(request.Url)
		if e != nil {
			err = e
			m.logger.Error(err)
			m.stats.IncRequestError()
			return
		}

		var body io.Reader
		if request.BodyStr != "" {
			body = strings.NewReader(request.BodyStr)
		}

		request.Request, err = http.NewRequestWithContext(ctx, request.Method, Url.String(), body)
		if err != nil {
			m.logger.Error(err)
			m.stats.IncRequestError()
			return
		}
	}
	if request.Header == nil {
		request.Header = make(http.Header)
	}
	request.Request.Header = request.Header
	if len(request.Cookies) > 0 {
		for _, cookie := range request.Cookies {
			request.AddCookie(cookie)
		}
	}

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
