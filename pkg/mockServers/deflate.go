package mockServers

import (
	"compress/flate"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlDeflate = "/deflate"

type RouteDeflate struct {
	logger pkg.Logger
}

func (h *RouteDeflate) Pattern() string {
	return UrlDeflate
}

func (h *RouteDeflate) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerDeflate")
	defer func() {
		h.logger.Info("exit HandlerDeflate")
	}()

	w.Header().Set("Content-Encoding", "deflate")

	fw, _ := flate.NewWriter(w, flate.BestCompression)
	defer func() {
		_ = fw.Close()
	}()

	_, _ = fw.Write([]byte("Hello, Deflate!"))
}

func NewRouteDeflate(logger pkg.Logger) pkg.Route {
	return &RouteDeflate{logger: logger}
}
