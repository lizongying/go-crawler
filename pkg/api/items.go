package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlItems = "/items"

type RouteItems struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteItems) Pattern() string {
	return UrlItems
}

func (h *RouteItems) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	records := h.crawler.GetStatistics().GetItems()
	h.OutJsonGzip(w, 0, "", records)
}

func (h *RouteItems) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteItems).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
