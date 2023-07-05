package devServer

import (
	"github.com/lizongying/go-crawler/pkg/logger"
	"io"
	"net/http"
)

const UrlHello = "/hello"

type HelloHandler struct {
	logger *logger.Logger
}

func (*HelloHandler) Pattern() string {
	return UrlHello
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HelloHandler")
	defer func() {
		h.logger.Info("exit HelloHandler")
	}()

	h.logger.InfoF("request: %+v", r)
	body, _ := io.ReadAll(r.Body)
	h.logger.InfoF("body: %s", string(body))
}

func NewHelloHandler(logger *logger.Logger) *HelloHandler {
	return &HelloHandler{logger: logger}
}
