package middlewares

import (
	"context"
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/request"
	"github.com/lizongying/go-crawler/pkg/utils"
	"github.com/temoto/robotstxt"
	"net/url"
)

type RobotsTxtMiddleware struct {
	pkg.UnimplementedMiddleware
	spider    pkg.Spider
	logger    pkg.Logger
	userAgent string
	group     *robotstxt.Group
	ignoreUrl []string
}

func (m *RobotsTxtMiddleware) SpiderOpened() {
	host := m.spider.GetHost()
	if host == "" {
		m.logger.Warn("host is emtpy")
		return
	}

	ctx := pkg.Context{
		Spider: m.spider,
	}
	r, e := m.spider.Request(ctx, request.NewRequest().SetUrl(fmt.Sprintf("%s/robots.txt", host)).SetSkipMiddleware(true))
	if e != nil {
		m.logger.Error(e)
		return
	}
	robots, err := robotstxt.FromBytes(r.BodyBytes())
	if err != nil {
		return
	}
	m.group = robots.FindGroup(m.userAgent)
}

func (m *RobotsTxtMiddleware) Start(ctx context.Context, spider pkg.Spider) (err error) {
	err = m.UnimplementedMiddleware.Start(ctx, spider)
	spider.GetSignal().RegisterSpiderOpened(m.SpiderOpened)
	return
}

func (m *RobotsTxtMiddleware) ProcessRequest(_ pkg.Context, request pkg.Request) (err error) {
	if m.group == nil {
		return
	}

	u, _ := url.Parse(request.GetUrl())
	if utils.InSlice(u.Path, m.ignoreUrl) {
		return
	}

	allow := m.group.Test(u.Path)
	if !allow {
		err = pkg.ErrNotAllowRequest
		return
	}

	return
}

func (m *RobotsTxtMiddleware) FromSpider(spider pkg.Spider) pkg.Middleware {
	if m == nil {
		return new(RobotsTxtMiddleware).FromSpider(spider)
	}

	m.UnimplementedMiddleware.FromSpider(spider)
	m.spider = spider
	m.logger = spider.GetLogger()
	m.userAgent = "*"
	m.ignoreUrl = []string{"/robots.txt"}
	return m
}
