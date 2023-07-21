package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"io"
	"net/http"
)

const UrlHello = "/hello"

type HandlerHello struct {
	logger *logger.Logger
}

func (*HandlerHello) Pattern() string {
	return UrlHello
}

func (h *HandlerHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerHello")
	defer func() {
		h.logger.Info("exit HandlerHello")
	}()

	h.logger.InfoF("request: %+v", r)
	body, _ := io.ReadAll(r.Body)
	h.logger.InfoF("body: %s", string(body))
}

func NewHandlerHello(logger *logger.Logger) *HandlerHello {
	return &HandlerHello{logger: logger}
}
