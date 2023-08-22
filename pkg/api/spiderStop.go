package api

import (
	"context"
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"time"
)

const UrlSpiderStop = "/spider/run"

type RouteSpiderStop struct {
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpiderStop) Pattern() string {
	return UrlSpiderStop
}

func (h *RouteSpiderStop) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req pkg.ReqSpiderStop
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.TaskId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	if req.Timeout != 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(req.Timeout)*time.Second)
		defer cancel()
	}
	err = h.crawler.SpiderStop(context.Background(), req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	spider := Spider{Name: ""}
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

	w.WriteHeader(http.StatusOK)
}

func (h *RouteSpiderStop) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpiderStop).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
