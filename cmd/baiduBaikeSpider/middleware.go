package main

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/url"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlDetail string
}

func (m *Middleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request
	_, ok := request.Extra.(*ExtraDetail)
	if ok {
		extra := request.Extra.(*ExtraDetail)
		keyword := extra.Keyword
		itemId := extra.ItemId
		if itemId != "" {
			itemId = fmt.Sprintf("/%s", itemId)
		}
		request.Url = fmt.Sprintf(m.urlDetail, url.QueryEscape(keyword), itemId)
		m.logger.Info("request.Url", request.Url)
		request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
	}

	err = c.NextRequest()
	return
}

func (m *Middleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromCrawler(spider)
	}
	m.logger = spider.GetLogger()
	m.urlDetail = "https://baike.baidu.com/item/%s%s"
	return m
}
