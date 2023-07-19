package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"net/http/httputil"
)

const UrlGet = "/get"

type GetHandler struct {
	logger pkg.Logger
}

func (h *GetHandler) Pattern() string {
	return UrlGet
}

func (h *GetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into GetHandler")
	defer func() {
		h.logger.Info("exit GetHandler")
	}()

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		h.logger.Error(err)
		return
	}

	_, _ = w.Write(reqDump)
}

func NewGetHandler(logger pkg.Logger) *GetHandler {
	return &GetHandler{logger: logger}
}
