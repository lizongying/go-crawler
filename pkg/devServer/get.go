package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"net/http/httputil"
)

const UrlGet = "/get"

type HandlerGet struct {
	logger pkg.Logger
}

func (h *HandlerGet) Pattern() string {
	return UrlGet
}

func (h *HandlerGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerGet")
	defer func() {
		h.logger.Info("exit HandlerGet")
	}()

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		h.logger.Error(err)
		return
	}

	_, _ = w.Write(reqDump)
}

func NewHandlerGet(logger pkg.Logger) *HandlerGet {
	return &HandlerGet{logger: logger}
}
