package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlJobs = "/jobs"

type RouteJobs struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteJobs) Pattern() string {
	return UrlJobs
}

func (h *RouteJobs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	jobs := h.crawler.GetStatistics().GetJobs()
	h.OutJson(w, 0, "", jobs)
}

func (h *RouteJobs) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteJobs).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
