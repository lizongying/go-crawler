package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlSpider = "/spider"

type Req struct {
	Name string `json:"name"`
}
type Spider struct {
	Name string `json:"name"`
}
type RouteSpider struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpider) Pattern() string {
	return UrlSpider
}

func (h *RouteSpider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req Req
	h.BindJson(w, r, &req)

	var spider Spider
	if req.Name == "" {
		for _, v := range h.crawler.GetSpiders() {
			if v.Name() == req.Name {
				spider = Spider{
					Name: v.Name(),
				}
				break
			}
		}
	}

	h.OutJson(w, 0, "", spider)
}

func (h *RouteSpider) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpider).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
