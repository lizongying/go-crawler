package middlewares

import (
	"github.com/lizongying/go-crawler/pkg"
)

type ChromeMiddleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
}

func (m *ChromeMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	if request.Headers() == nil {
		request.SetHeader("Accept", "*/*")
		request.SetHeader("Cache-Control", "no-cache")
		request.SetHeader("Content-Type", "text/plain;charset=UTF-8")
		request.SetHeader("Pragma", "no-cache")
		request.SetHeader("Sec-Ch-Ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
		request.SetHeader("Sec-Ch-Ua-Mobile", "?0")
		request.SetHeader("Sec-Ch-Ua-Platform", "\"macOS\"")
		request.SetHeader("Sec-Fetch-Dest", "empty")
		request.SetHeader("Sec-Fetch-Mode", "no-cors")
		request.SetHeader("Sec-Fetch-Site", "same-site")
		request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")
	}

	return
}

func (m *ChromeMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(ChromeMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	return m
}
