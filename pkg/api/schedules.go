package api

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlSchedules = "/schedules"

type RouteSchedules struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
}

func (h *RouteSchedules) Pattern() string {
	return UrlSchedules
}

func (h *RouteSchedules) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	schedules := h.crawler.GetStatistics().GetSchedules()
	//for _, v := range nodes {
	//	fmt.Println(v)
	//	bs, err := v.Marshal()
	//	if err != nil {
	//		h.OutJson(w, 1, err.Error(), nil)
	//		return
	//	}
	//
	//}
	h.OutJson(w, 0, "", schedules)
}

func (h *RouteSchedules) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteSchedules).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	return h
}
