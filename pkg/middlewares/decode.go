package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"strings"
)

type DecodeMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *DecodeMiddleware) ProcessResponse(c *pkg.Context) (err error) {
	r := c.Response

	if strings.Contains(strings.ToUpper(r.Header.Get("Content-Type")), "=BIG5") {
		decoder := traditionalchinese.Big5.NewDecoder()
		r.BodyBytes, _ = decoder.Bytes(r.BodyBytes)
	}
	if strings.Contains(strings.ToUpper(r.Header.Get("Content-Type")), "=GBK") {
		decoder := simplifiedchinese.GBK.NewDecoder()
		r.BodyBytes, _ = decoder.Bytes(r.BodyBytes)
	}
	if strings.Contains(strings.ToUpper(r.Header.Get("Content-Type")), "=GB18030") {
		decoder := simplifiedchinese.GB18030.NewDecoder()
		r.BodyBytes, _ = decoder.Bytes(r.BodyBytes)
	}
	if strings.Contains(strings.ToUpper(r.Header.Get("Content-Type")), "=GB2312") {
		decoder := simplifiedchinese.HZGB2312.NewDecoder()
		r.BodyBytes, _ = decoder.Bytes(r.BodyBytes)
	}
	err = c.NextResponse()
	return
}

func (m *DecodeMiddleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(DecodeMiddleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	return m
}
