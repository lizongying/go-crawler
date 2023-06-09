package devServer

import (
	"compress/gzip"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlGzip = "/gzip"

type GzipHandler struct {
	logger pkg.Logger
}

func (h *GzipHandler) Pattern() string {
	return UrlGzip
}

func (h *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into GzipHandler")
	defer func() {
		h.logger.Info("exit GzipHandler")
	}()

	w.Header().Set("Content-Encoding", "gzip")

	gw := gzip.NewWriter(w)
	defer func() {
		_ = gw.Close()
	}()

	_, _ = gw.Write([]byte("Hello, Gzip!"))
}

func NewGzipHandler(logger pkg.Logger) *GzipHandler {
	return &GzipHandler{logger: logger}
}
