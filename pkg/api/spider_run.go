package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lizongying/go-crawler/pkg"
	crawlerContext "github.com/lizongying/go-crawler/pkg/context"
	"io"
	"net/http"
	"time"
)

const UrlSpiderRun = "/spider/run"

type RouteSpiderRun struct {
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpiderRun) Pattern() string {
	return UrlSpiderRun
}

func (h *RouteSpiderRun) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req pkg.ReqSpiderStart
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.TaskId == "" {
		req.TaskId = uuid.New().String()
	}
	if req.Mode == "" {
		req.Mode = "once"
	}

	c := context.Background()
	if req.Timeout > 0 {
		var cancel context.CancelFunc
		c, cancel = context.WithTimeout(c, time.Duration(req.Timeout)*time.Second)
		defer cancel()
	}
	ctx := new(crawlerContext.Context).
		WithGlobalContext(c).
		WithTaskId(req.TaskId).
		WithSpiderName(req.Name).
		WithStartFunc(req.Func).
		WithArgs(req.Args).
		WithMode(req.Mode)
	err = h.crawler.SpiderStart(ctx)
	if err != nil {
		h.Json(w, 1, err.Error(), nil)
		return
	}

	spider := Spider{Name: req.Name}
	h.Json(w, 0, "", spider)
}

func (h *RouteSpiderRun) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiderRun).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
