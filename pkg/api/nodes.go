package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlNodes = "/nodes"

type RouteNodes struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteNodes) Pattern() string {
	return UrlNodes
}

func (h *RouteNodes) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nodes := h.crawler.GetStatistics().GetNodes()
	//for _, v := range nodes {
	//	fmt.Println(v)
	//	bs, err := v.Marshal()
	//	if err != nil {
	//		h.OutJson(w, 1, err.Error(), nil)
	//		return
	//	}
	//
	//}
	h.OutJson(w, 0, "", nodes)
}

func (h *RouteNodes) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteNodes).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
