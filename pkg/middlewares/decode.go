package middlewares

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"strings"
)

type DecodeMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DecodeMiddleware) ProcessResponse(_ context.Context, response *pkg.Response) (err error) {
	contentType := strings.ToUpper(response.Header.Get("Content-Type"))
	if strings.Contains(contentType, "=BIG5") {
		decoder := traditionalchinese.Big5.NewDecoder()
		response.BodyBytes, err = decoder.Bytes(response.BodyBytes)
	}
	if strings.Contains(contentType, "=GBK") {
		decoder := simplifiedchinese.GBK.NewDecoder()
		response.BodyBytes, err = decoder.Bytes(response.BodyBytes)
	}
	if strings.Contains(contentType, "=GB18030") {
		decoder := simplifiedchinese.GB18030.NewDecoder()
		response.BodyBytes, err = decoder.Bytes(response.BodyBytes)
	}
	if strings.Contains(contentType, "=GB2312") {
		decoder := simplifiedchinese.HZGB2312.NewDecoder()
		response.BodyBytes, err = decoder.Bytes(response.BodyBytes)
	}

	if err != nil {
		m.logger.Error(err)
	}
	return
}

func (m *DecodeMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DecodeMiddleware).FromCrawler(spider)
	}

	m.logger = spider.GetLogger()
	return m
}
