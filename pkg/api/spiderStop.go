package api

import (
	"context"
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
)

const UrlSpiderStop = "/spider/run"

type ReqSpiderStop struct {
	Name string
	Func string
	Args string
}

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

	var req ReqSpiderStop
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.crawler.StartSpider(context.Background(), req.Name, req.Func, req.Args)
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
