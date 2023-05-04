package httpServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlOk = "/ok"

type OkHandler struct {
	logger pkg.Logger
}

func (h *OkHandler) Pattern() string {
	return UrlOk
}

func (h *OkHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`Ok`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewOkHandler(logger pkg.Logger) *OkHandler {
	return &OkHandler{logger: logger}
}
