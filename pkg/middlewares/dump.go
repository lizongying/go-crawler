package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"net/http/httputil"
)

type DumpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DumpMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
	m.logger.InfoF("request: %+v", request)
	return
}

func (m *DumpMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	b, _ := httputil.DumpResponse(response.GetResponse(), false)
	m.logger.DebugF("response: \n%s", string(b))
	return
}

func (m *DumpMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(DumpMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
