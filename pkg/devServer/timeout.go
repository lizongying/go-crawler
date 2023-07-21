package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"time"
)

const UrlTimeout = "/timeout"

type HandlerTimeout struct {
	logger pkg.Logger
}

func (h *HandlerTimeout) Pattern() string {
	return UrlTimeout
}

func (h *HandlerTimeout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerTimeout")
	defer func() {
		h.logger.Info("exit HandlerTimeout")
	}()

	time.Sleep(5 * time.Second)

	w.WriteHeader(http.StatusBadGateway)
	_, _ = w.Write([]byte("ok"))
}

func NewHandlerTimeout(logger pkg.Logger) *HandlerTimeout {
	return &HandlerTimeout{logger: logger}
}
