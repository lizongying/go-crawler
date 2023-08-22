package api

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"time"
)

const UrlSpiderRun = "/spider/run"

type RouteSpiderRun struct {
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

	ctx := context.Background()
	if req.Timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(req.Timeout)*time.Second)
		defer cancel()
	}
	err = h.crawler.SpiderStart(ctx, req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	spider := Spider{Name: req.Name}
	var resp []byte
	resp, err = json.Marshal(spider)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (h *RouteSpiderRun) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiderRun).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
