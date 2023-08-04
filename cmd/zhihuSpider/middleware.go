package main

import (
	"errors"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlDetail string
}

func (m *Middleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	switch request.GetExtraName() {
	case "ExtraDetail":
		var extra ExtraDetail
		err = request.UnmarshalExtra(&extra)
		if err != nil {
			m.logger.Error(err)
			return
		}
		id := extra.Id
		if id == 0 {
			err = errors.New("id 0")
			m.logger.Error(err)
			return
		}
		request.SetUrl(fmt.Sprintf(m.urlDetail, id))
	}
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	return
}

func (m *Middleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.logger = spider.GetLogger()
	m.urlDetail = "https://www.zhihu.com/question/%d"
	return m
}
