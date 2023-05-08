package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCustom = "/custom"

type CustomHandler struct {
	logger pkg.Logger
}

func (h *CustomHandler) Pattern() string {
	return UrlCustom
}

func (h *CustomHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`Custom`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewCustomHandler(logger pkg.Logger) *CustomHandler {
	return &CustomHandler{logger: logger}
}
