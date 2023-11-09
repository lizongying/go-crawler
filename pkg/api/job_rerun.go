package api

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlJobRerun = "/job/rerun"

type ReqJobRerun struct {
	SpiderName string `json:"spider_name"`
	JobId      string `json:"job_id"`
}

type RespJobRerun struct {
	SpiderName string `json:"spider_name"`
	JobId      string `json:"job_id"`
}

type RouteJobRerun struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteJobRerun) Pattern() string {
	return UrlJobRerun
}

func (h *RouteJobRerun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req ReqJobRerun
	h.BindJson(w, r, &req)

	if req.SpiderName == "" {
		h.OutJson(w, 1, "SpiderName empty", nil)
		return
	}

	if req.JobId == "" {
		h.OutJson(w, 1, "JobId empty", nil)
		return
	}

	if err := h.crawler.RerunJob(context.TODO(), req.SpiderName, req.JobId); err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	h.OutJson(w, 0, "", &RespJobRerun{})
}

func (h *RouteJobRerun) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteJobRerun).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
