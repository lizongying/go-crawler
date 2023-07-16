package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"strings"
)

type DecodeMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DecodeMiddleware) ProcessResponse(_ context.Context, response pkg.Response) (err error) {
	contentType := strings.ToUpper(response.GetHeader("Content-Type"))

	var decoder *encoding.Decoder
	if strings.Contains(contentType, "=BIG5") {
		decoder = traditionalchinese.Big5.NewDecoder()
	}
	if strings.Contains(contentType, "=GBK") {
		decoder = simplifiedchinese.GBK.NewDecoder()
	}
	if strings.Contains(contentType, "=GB18030") {
		decoder = simplifiedchinese.GB18030.NewDecoder()
	}
	if strings.Contains(contentType, "=GB2312") {
		decoder = simplifiedchinese.HZGB2312.NewDecoder()
	}

	if decoder != nil {
		var bodyBytes []byte
		bodyBytes, err = decoder.Bytes(response.GetBodyBytes())
		if err != nil {
			m.logger.Error(err)
			return
		}
		response.SetBodyBytes(bodyBytes)
	}

	return
}

func (m *DecodeMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(DecodeMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	return m
}
