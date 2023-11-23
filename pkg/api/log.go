package api

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlLog = "/log"

type RouteLog struct {
	Request
	Response
	crawler pkg.Crawler
	logger  pkg.Logger
	stream  pkg.Stream
}

func (h *RouteLog) Pattern() string {
	return UrlLog
}

func (h *RouteLog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	if r.URL.Query().Get("task_id") == "" {
		http.Error(w, "TaskId empty", http.StatusInternalServerError)
		return
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming not supported", http.StatusInternalServerError)
		return
	}

	logChannel := make(chan []byte, 100)
	h.stream.Register("", logChannel)
	defer func() {
		h.stream.Unregister("")
		close(logChannel)
	}()

	for {
		var message []byte
		select {
		case message, ok = <-logChannel:
			if !ok {
				return
			}
			_, _ = fmt.Fprintf(w, "data: %s\n\n", string(message))
			flusher.Flush()
		case <-r.Context().Done():
			return
		}
	}
}

func (h *RouteLog) FromCrawler(crawler pkg.Crawler) pkg.Route {
	if h == nil {
		return new(RouteLog).FromCrawler(crawler)
	}

	h.logger = crawler.GetLogger()
	h.crawler = crawler
	h.stream = crawler.GetStream()
	return h
}
