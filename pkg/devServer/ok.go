package devServer

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlOk = "/ok"

type OkHandler struct {
	logger pkg.Logger
}

func (h *OkHandler) Pattern() string {
	return UrlOk
}

func (h *OkHandler) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	h.logger.Info("into OkHandler")
	defer func() {
		h.logger.Info("exit OkHandler")
	}()

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Header: %v", request.Header)))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewOkHandler(logger pkg.Logger) *OkHandler {
	return &OkHandler{logger: logger}
}
