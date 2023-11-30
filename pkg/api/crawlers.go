package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCrawlers = "/crawlers"

type RouteCrawlers struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteCrawlers) Pattern() string {
	return UrlCrawlers
}

func (h *RouteCrawlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nodes := h.crawler.GetStatistics().GetCrawlers()
	h.OutJson(w, 0, "", nodes)
}

func (h *RouteCrawlers) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteCrawlers).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
