package main

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/url"
)

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger

	urlDetail string
}

func (m *Middleware) ProcessRequest(_ context.Context, request *pkg.Request) (err error) {
	var extra ExtraDetail
	e := request.GetExtra(&extra)
	if e != nil {
		m.logger.Warn(e)
		return
	}
	keyword := extra.Keyword
	itemId := extra.ItemId
	if itemId != "" {
		itemId = fmt.Sprintf("/%s", itemId)
	}
	request.Url = fmt.Sprintf(m.urlDetail, url.QueryEscape(keyword), itemId)
	m.logger.Info("request.Url", request.Url)
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	return
}

func (m *Middleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(Middleware).FromCrawler(crawler)
	}
	m.logger = crawler.GetLogger()
	m.urlDetail = "https://baike.baidu.com/item/%s%s"
	return m
}
