package devServer

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlOk = "/ok"

type HandlerOk struct {
	logger pkg.Logger
}

func (h *HandlerOk) Pattern() string {
	return UrlOk
}

func (h *HandlerOk) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	h.logger.Debug("into HandlerOk")
	defer func() {
		h.logger.Debug("exit HandlerOk")
	}()

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Header: %v", request.Header)))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewHandlerOk(logger pkg.Logger) *HandlerOk {
	return &HandlerOk{logger: logger}
}
