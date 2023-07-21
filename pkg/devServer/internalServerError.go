package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

const UrlInternalServerError = "/internal-server-error"

type HandlerInternalServerError struct {
	logger *logger.Logger
}

func (*HandlerInternalServerError) Pattern() string {
	return UrlInternalServerError
}

func (h *HandlerInternalServerError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerInternalServerError")
	defer func() {
		h.logger.Info("exit HandlerInternalServerError")
	}()

	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(`InternalServerError`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewHandlerInternalServerError(logger *logger.Logger) *HandlerInternalServerError {
	return &HandlerInternalServerError{logger: logger}
}
