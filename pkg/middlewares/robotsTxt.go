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
	logger    pkg.Logger
	userAgent string
	group     *robotstxt.Group
	ignoreUrl []string
}

func (m *RobotsTxtMiddleware) Start(ctx context.Context, crawler pkg.Crawler) (err error) {
	host := crawler.GetSpider().GetHost()
	if host == "" {
		m.logger.Warn("host is emtpy")
		return
	}
	r, e := crawler.GetScheduler().Request(ctx, request.NewRequest().SetUrl(fmt.Sprintf("%s/robots.txt", host)).SetSkipMiddleware(true))
	if e != nil {
		err = e
		m.logger.Error(e)
		return
	}
	robots, err := robotstxt.FromBytes(r.BodyBytes)
	if err != nil {
		return
	}
	m.group = robots.FindGroup(m.userAgent)
	return
}

func (m *RobotsTxtMiddleware) ProcessRequest(_ context.Context, request pkg.Request) (err error) {
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

func (m *RobotsTxtMiddleware) FromCrawler(crawler pkg.Crawler) pkg.Middleware {
	if m == nil {
		return new(RobotsTxtMiddleware).FromCrawler(crawler)
	}

	m.logger = crawler.GetLogger()
	m.userAgent = "*"
	m.ignoreUrl = []string{"/robots.txt"}

	return m
}
