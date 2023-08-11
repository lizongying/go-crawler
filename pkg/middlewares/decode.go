package middlewares

import (
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

func (m *DecodeMiddleware) ProcessResponse(_ pkg.Context, response pkg.Response) (err error) {
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
		bodyBytes, err = decoder.Bytes(response.BodyBytes())
		if err != nil {
			m.logger.Error(err)
			return
		}
		response.SetBodyBytes(bodyBytes)
	}

	return
}

func (m *DecodeMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DecodeMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
