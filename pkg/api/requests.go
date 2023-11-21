package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlRequests = "/requests"

type RouteRequests struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteRequests) Pattern() string {
	return UrlRequests
}

func (h *RouteRequests) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	requests := h.crawler.GetStatistics().GetRequests()
	h.OutJson(w, 0, "", requests)
}

func (h *RouteRequests) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteRequests).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
