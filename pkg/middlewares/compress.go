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

func (m *CompressMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	if response.GetHeader("Content-Encoding") == "deflate" {
		reader := flate.NewReader(bytes.NewReader(response.GetBodyBytes()))
		defer func() {
			e := reader.Close()
			if e != nil {
				err = errors.Join(err, e)
				m.logger.Error(err)
			}
		}()

		var bodyBytes []byte
		bodyBytes, err = io.ReadAll(reader)
		if err != nil {
			m.logger.Error(err)
			return
		}
		response.SetBodyBytes(bodyBytes)
	}

	return
}

func (m *CompressMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(CompressMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
