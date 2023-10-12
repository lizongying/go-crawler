package api

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
)

const UrlSpider = "/spider"

type Req struct {
	Name string
}
type Spider struct {
	Name string `json:"name"`
}
type RouteSpider struct {
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSpider) Pattern() string {
	return UrlSpider
}

func (h *RouteSpider) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req Req
	if err = json.Unmarshal(body, &req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	spider := Spider{Name: "test1"}
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

func (h *RouteSpider) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSpider).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
