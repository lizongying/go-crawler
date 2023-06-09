package devServer

import (
	"compress/flate"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlDeflate = "/deflate"

type DeflateHandler struct {
	logger pkg.Logger
}

func (h *DeflateHandler) Pattern() string {
	return UrlDeflate
}

func (h *DeflateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into DeflateHandler")
	defer func() {
		h.logger.Info("exit DeflateHandler")
	}()

	w.Header().Set("Content-Encoding", "deflate")

	fw, _ := flate.NewWriter(w, flate.BestCompression)
	defer func() {
		_ = fw.Close()
	}()

	_, _ = fw.Write([]byte("Hello, Deflate!"))
}

func NewDeflateHandler(logger pkg.Logger) *DeflateHandler {
	return &DeflateHandler{logger: logger}
}
