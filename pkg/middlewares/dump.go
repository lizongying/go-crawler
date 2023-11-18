package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http/httputil"
)

type DumpMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DumpMiddleware) ProcessRequest(ctx pkg.Context, request pkg.Request) (err error) {
	bs, _ := request.Marshal()
	m.logger.Info(m.GetSpider().Name(), ctx.GetTask().GetId(), ctx.GetRequest().GetId(), "request:", string(bs))
	return
}

func (m *DumpMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
	if response.GetResponse() == nil {
		m.logger.Debug("response nil")
		return
	}

	b, _ := httputil.DumpResponse(response.GetResponse(), false)
	m.logger.Debugf("response: \n%s", string(b))
	return
}

func (m *DumpMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DumpMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
