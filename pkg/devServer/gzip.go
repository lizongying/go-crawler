package devServer

import (
	"compress/gzip"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlGzip = "/gzip"

type HandlerGzip struct {
	logger pkg.Logger
}

func (h *HandlerGzip) Pattern() string {
	return UrlGzip
}

func (h *HandlerGzip) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func NewHandlerGzip(logger pkg.Logger) *HandlerGzip {
	return &HandlerGzip{logger: logger}
}
