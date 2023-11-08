package api

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlJobStop = "/job/stop"

type ReqJobStop struct {
	SpiderName string `json:"spider_name"`
	JobId      string `json:"job_id"`
}

type RespJobStop struct {
	SpiderName string `json:"spider_name"`
	JobId      string `json:"job_id"`
}

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
	var req ReqJobStop
	h.BindJson(w, r, &req)

	if req.SpiderName == "" {
		h.OutJson(w, 1, "SpiderName empty", nil)
		return
	}

	if req.JobId == "" {
		h.OutJson(w, 1, "JobId empty", nil)
		return
	}

	if err := h.crawler.KillJob(context.TODO(), req.SpiderName, req.JobId); err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	h.OutJson(w, 0, "", &RespJobStop{})
}

func (h *RouteJobStop) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteJobStop).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
