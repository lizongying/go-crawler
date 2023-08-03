package middlewares

import (
	"bytes"
	"compress/flate"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"io"
)

type CompressMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *CompressMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
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

func (m *CompressMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(CompressMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
