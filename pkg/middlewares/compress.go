package middlewares

import (
	"bytes"
	"compress/flate"
	"github.com/lizongying/go-crawler/pkg"
	"io"
)

type CompressMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CompressMiddleware) ProcessRequest(c *pkg.Context) (err error) {
	m.logger.Debug("enter ProcessRequest")
	defer func() {
		m.logger.Debug("exit ProcessRequest")
	}()

	err = c.NextRequest()
	if err != nil {
		m.logger.Debug(err)
		return
	}

	// request := c.Request
	// request.Header.Set("Accept-Encoding", "deflate")

	return
}

func (m *CompressMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	r := c.Response

	if r.Header.Get("Content-Encoding") == "deflate" {
		reader := flate.NewReader(bytes.NewReader(r.BodyBytes))
		defer func() {
			_ = reader.Close()
		}()

		r.BodyBytes, _ = io.ReadAll(reader)
	}

	err = c.NextResponse()
	return
}

func (m *CompressMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CompressMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
