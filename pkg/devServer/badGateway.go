package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

const UrlBadGateway = "/bad-gateway"

type BadGatewayHandler struct {
	logger *logger.Logger
}

func (h *BadGatewayHandler) Pattern() string {
	return UrlBadGateway
}

func (h *BadGatewayHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into BadGatewayHandler")
	defer func() {
		h.logger.Info("exit BadGatewayHandler")
	}()

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
