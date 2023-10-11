package mock_servers

import (
	"github.com/andybalholm/brotli"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlBrotli = "/brotli"

type RouteBrotli struct {
	logger pkg.Logger
}

func (h *RouteBrotli) Pattern() string {
	return UrlBrotli
}

func (h *RouteBrotli) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerBrotli")
	defer func() {
		h.logger.Info("exit HandlerBrotli")
	}()

	w.Header().Set("Content-Encoding", "br")

	fw := brotli.NewWriter(w)
	defer func() {
		_ = fw.Close()
	}()

	_, _ = fw.Write([]byte("Hello, Brotli!"))
}

func NewRouteBrotli(logger pkg.Logger) pkg.Route {
	return &RouteBrotli{logger: logger}
}
