package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlHello = "/hello"

type RouteHello struct {
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteHello) Pattern() string {
	return UrlHello
}

func (h *RouteHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("hello"))
}
func (h *RouteHello) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteHello).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
