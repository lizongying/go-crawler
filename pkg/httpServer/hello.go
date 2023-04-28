package httpServer

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg/logger"
	"io"
	"net/http"
)

type HelloHandler struct {
	logger *logger.Logger
}

func (*HelloHandler) Pattern() string {
	return "/hello"
}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("Failed to read request")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if _, err = fmt.Fprintf(w, "Hello, %s\n", body); err != nil {
		h.logger.Error("Failed to write response")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func NewHelloHandler(logger *logger.Logger) *HelloHandler {
	return &HelloHandler{logger: logger}
}
