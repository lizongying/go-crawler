package api

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/statistics/job"
	"net/http"
	"time"
)

const UrlJobRun = "/job/run"

type RouteJobRun struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteJobRun) Pattern() string {
	return UrlJobRun
}

func (h *RouteJobRun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req pkg.ReqJobStart
	h.BindJson(w, r, &req)

	if req.Name == "" {
		err := errors.New("name empty")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}
	if req.Func == "" {
		req.Func = "Test"
	}
	if req.Mode == pkg.JobModeUnknown {
		req.Mode = pkg.JobModeOnce
	}

	c := context.Background()
	if req.Timeout > 0 {
		c, _ = context.WithTimeout(c, time.Duration(req.Timeout)*time.Second)
	}

	jobId, err := h.crawler.RunJob(c, req.Name, req.Func, req.Args, req.Mode, req.Spec)
	if err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	h.OutJson(w, 0, "", &job.Job{Id: jobId})
}

func (h *RouteJobRun) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteJobRun).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
