package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCustom = "/custom"

type HandlerCustom struct {
	logger pkg.Logger
}

func (h *HandlerCustom) Pattern() string {
	return UrlCustom
}

func (h *HandlerCustom) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`Custom`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewHandlerCustom(logger pkg.Logger) *HandlerCustom {
	return &HandlerCustom{logger: logger}
}
