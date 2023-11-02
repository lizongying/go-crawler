package api

import (
	"context"
	"errors"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/statistics/task"
	"net/http"
	"time"
)

const UrlSpiderRun = "/spider/run"

type RouteSpiderRun struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpiderRun) Pattern() string {
	return UrlSpiderRun
}

func (h *RouteSpiderRun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req pkg.ReqSpiderStart
	h.BindJson(w, r, &req)

	if req.Name == "" {
		err := errors.New("name empty")
		h.OutJson(w, 1, err.Error(), nil)
		return
	}
	if req.Func == "" {
		req.Func = "Test"
	}
	if req.Mode == pkg.ScheduleModeUnknown {
		req.Mode = pkg.ScheduleModeOnce
	}

	c := context.Background()
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		c, cancel = context.WithTimeout(c, time.Duration(req.Timeout)*time.Second)
		defer cancel()
	}

	taskId, err := h.crawler.Run(c, req.Name, req.Func, req.Args, req.Mode, req.Spec)
	if err != nil {
		h.OutJson(w, 1, err.Error(), nil)
		return
	}

	spider := task.Task{Id: taskId}
	h.OutJson(w, 0, "", spider)
}

func (h *RouteSpiderRun) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiderRun).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
