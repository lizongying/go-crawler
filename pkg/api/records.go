package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlRecords = "/records"

type RouteRecords struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteRecords) Pattern() string {
	return UrlRecords
}

func (h *RouteRecords) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	records := h.crawler.GetStatistics().GetRecords()
	h.OutJsonGzip(w, 0, "", records)
}

func (h *RouteRecords) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteRecords).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
