package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"net/http"
)

const UrlBadGateway = "/bad-gateway"

type HandlerBadGateway struct {
	logger *logger.Logger
}

func (h *HandlerBadGateway) Pattern() string {
	return UrlBadGateway
}

func (h *HandlerBadGateway) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into HandlerBadGateway")
	defer func() {
		h.logger.Info("exit HandlerBadGateway")
	}()

	w.WriteHeader(http.StatusBadGateway)
	_, err := w.Write([]byte(`BadGateway`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewHandlerBadGateway(logger *logger.Logger) *HandlerBadGateway {
	return &HandlerBadGateway{logger: logger}
}
