package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

type BadGatewayHandler struct {
	logger *logger.Logger
}

func (h *BadGatewayHandler) Pattern() string {
	return "/bad-gateway"
}

func (h *BadGatewayHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusBadGateway)
	_, err := w.Write([]byte(`BadGateway`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewBadGatewayHandler(logger *logger.Logger) *BadGatewayHandler {
	return &BadGatewayHandler{logger: logger}
}