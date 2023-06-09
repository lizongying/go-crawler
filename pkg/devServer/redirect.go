package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlRedirect = "/redirect"

type RedirectHandler struct {
	logger pkg.Logger
}

func (h *RedirectHandler) Pattern() string {
	return UrlRedirect
}

func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into RedirectHandler")
	defer func() {
		h.logger.Info("exit RedirectHandler")
	}()

	//http.Redirect(w, r, UrlOk, http.StatusFound)

	w.Header().Set("Location", UrlOk)
	w.WriteHeader(http.StatusFound)
}

func NewRedirectHandler(logger pkg.Logger) *RedirectHandler {
	return &RedirectHandler{logger: logger}
}
