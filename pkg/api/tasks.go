package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlTasks = "/tasks"

type RouteTasks struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteTasks) Pattern() string {
	return UrlTasks
}

func (h *RouteTasks) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tasks := h.crawler.GetStatistics().GetTasks()
	h.OutJson(w, 0, "", tasks)
}

func (h *RouteTasks) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteTasks).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
