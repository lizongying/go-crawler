package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlRedirect = "/redirect"

type HandlerRedirect struct {
	logger pkg.Logger
}

func (h *HandlerRedirect) Pattern() string {
	return UrlRedirect
}

func (h *HandlerRedirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerRedirect")
	defer func() {
		h.logger.Info("exit HandlerRedirect")
	}()

	//http.Redirect(w, r, UrlOk, http.StatusFound)

	w.Header().Set("Location", UrlOk)
	w.WriteHeader(http.StatusFound)
}

func NewHandlerRedirect(logger pkg.Logger) *HandlerRedirect {
	return &HandlerRedirect{logger: logger}
}
