package api

import (
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"net/http"
)

const UrlJobStop = "/job/stop"

type RouteJobStop struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteJobStop) Pattern() string {
	return UrlJobStop
}

func (h *RouteJobStop) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req pkg.ReqJobStop
	h.BindJson(w, r, &req)

	if req.Id == "" {
		h.OutJson(w, 1, "TaskId empty", nil)
		return
	}

	ctx := new(crawlerContext.Context).WithTaskId(req.Id)
	err := h.crawler.SpiderStop(ctx)
	if err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	spider := Spider{Name: ""}
	h.OutJson(w, 0, "", spider)
}

func (h *RouteJobStop) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteJobStop).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
