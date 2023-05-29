package main

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/url"
)

const Video = "EgIQAQ%253D%253D"

type Middleware struct {
	pkg.UnimplementedMiddleware
	logger pkg.Logger
	spider pkg.Spider

	urlSearch    string
	urlSearchApi string
	urlUserApi   string
	urlVideos    string

	apiKey string
}

func (m *Middleware) SpiderStart(_ context.Context, spider pkg.Spider) (err error) {
	m.spider = spider
	return
}

func (m *Middleware) ProcessRequest(c *pkg.Context) (err error) {
	request := c.Request
	_, ok := request.Extra.(*ExtraSearch)
	if ok {
		extra := request.Extra.(*ExtraSearch)
		keyword := url.QueryEscape(extra.Keyword)
		request.Url = fmt.Sprintf(m.urlSearch, keyword)
		if extra.Sp == Video {
			request.Url += fmt.Sprintf("&sp=%s", Video)
		}
	}
	_, ok = request.Extra.(*ExtraSearchApi)
	if ok {
		extra := request.Extra.(*ExtraSearchApi)
		request.Method = "POST"
		request.Url = fmt.Sprintf(m.urlSearchApi, m.apiKey)
		request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"continuation":"%s"}`, extra.NextPageToken)
	}
	_, ok = request.Extra.(*ExtraVideos)
	if ok {
		extra := request.Extra.(*ExtraVideos)
		request.Url = fmt.Sprintf(m.urlVideos, extra.Id)
	}
	_, ok = request.Extra.(*ExtraUserApi)
	if ok {
		extra := request.Extra.(*ExtraUserApi)
		request.Method = "POST"
		request.Url = fmt.Sprintf(m.urlUserApi, m.apiKey)
		request.BodyStr = fmt.Sprintf(`{"context":{"client":{"hl":"en","gl":"US","clientName":"WEB","clientVersion":"2.20230327.01.00"}},"browseId":"%s"}`, extra.Key)
	}
	request.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")

	err = c.NextRequest()
	return
}

func (m *Middleware) FromCrawler(spider pkg.Spider) pkg.Middleware {
	m.logger = spider.GetLogger()
	m.urlSearch = "https://www.youtube.com/results?search_query=%s"
	m.urlSearchApi = "https://www.youtube.com/youtubei/v1/search?key=%s"
	m.urlUserApi = "https://www.youtube.com/youtubei/v1/browse?key=%s"
	m.urlVideos = "https://www.youtube.com/@%s/videos"
	m.apiKey = "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"
	return m
}

func NewMiddleware() pkg.Middleware {
	return &Middleware{}
}
