package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlHome = "/"

type RouteHome struct {
	http.Handler
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteHome) Pattern() string {
	return UrlHome
}

func (h *RouteHome) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteHome).FromCrawler(crawler)
	}

	files, _ := fs.Sub(static.Api, "api")
	h.Handler = http.StripPrefix(UrlHome, http.FileServer(http.FS(files)))
	h.crawler = crawler
	h.logger = crawler.GetLogger()
	return h
}
