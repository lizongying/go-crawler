package middlewares

import (
	"bytes"
	"compress/flate"
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"io"
)

type CompressMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CompressMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	if response.Header.Get("Content-Encoding") == "deflate" {
		reader := flate.NewReader(bytes.NewReader(response.BodyBytes))
		defer func() {
			e := reader.Close()
			if e != nil {
				err = errors.Join(err, e)
				m.logger.Error(err)
			}
		}()

		response.BodyBytes, err = io.ReadAll(reader)
		if err != nil {
			m.logger.Error(err)
		}
	}

	return
}

func (m *CompressMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CompressMiddleware).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	return m
}
