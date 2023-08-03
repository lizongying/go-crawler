package devServer

import (
	"compress/gzip"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlGzip = "/gzip"

type RouteGzip struct {
	logger pkg.Logger
}

func (h *RouteGzip) Pattern() string {
	return UrlGzip
}

func (h *RouteGzip) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerGzip")
	defer func() {
		h.logger.Info("exit HandlerGzip")
	}()

	w.Header().Set("Content-Encoding", "gzip")

	gw := gzip.NewWriter(w)
	defer func() {
		_ = gw.Close()
	}()

	_, _ = gw.Write([]byte("Hello, Gzip!"))
}

func NewRouteGzip(logger pkg.Logger) pkg.Route {
	return &RouteGzip{logger: logger}
}
