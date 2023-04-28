package httpServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

type NoLimitHandler struct {
	logger *logger.Logger
}

func (*NoLimitHandler) Pattern() string {
	return "/no-limit"
}

func (h *NoLimitHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	//w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusBadGateway)

	_, err := w.Write([]byte(`hi`))
	if err != nil {
		return
	}
}

func NewNoLimitHandler(logger *logger.Logger) *NoLimitHandler {
	return &NoLimitHandler{logger: logger}
}
