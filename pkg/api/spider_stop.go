package api

import (
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"net/http"
)

const UrlSpiderStop = "/spider/run"

type RouteSpiderStop struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpiderStop) Pattern() string {
	return UrlSpiderStop
}

func (h *RouteSpiderStop) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req pkg.ReqSpiderStop
	h.BindJson(w, r, &req)

	if req.TaskId == "" {
		h.OutJson(w, 1, "TaskId empty", nil)
		return
	}

	ctx := new(crawlerContext.Context).WithTaskId(req.TaskId)
	err := h.crawler.SpiderStop(ctx)
	if err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	spider := Spider{Name: ""}
	h.OutJson(w, 0, "", spider)
}

func (h *RouteSpiderStop) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiderStop).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
