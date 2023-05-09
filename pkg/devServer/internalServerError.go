package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

const UrlInternalServerError = "/internal-server-error"

type InternalServerErrorHandler struct {
	logger *logger.Logger
}

func (*InternalServerErrorHandler) Pattern() string {
	return UrlInternalServerError
}

func (h *InternalServerErrorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(`InternalServerError`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewInternalServerErrorHandler(logger *logger.Logger) *InternalServerErrorHandler {
	return &InternalServerErrorHandler{logger: logger}
}
