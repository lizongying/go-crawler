package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlSpiders = "/spiders"

type RouteSpiders struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpiders) Pattern() string {
	return UrlSpiders
}

func (h *RouteSpiders) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	spiders := h.crawler.GetStatistics().GetSpiders()
	h.OutJson(w, 0, "", spiders)
}

func (h *RouteSpiders) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiders).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
