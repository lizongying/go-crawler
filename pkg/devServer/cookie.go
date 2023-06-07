package devServer

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCookie = "/cookie"

type CookieHandler struct {
	logger pkg.Logger
}

func (h *CookieHandler) Pattern() string {
	return UrlHttpAuth
}

func (h *CookieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into CookieHandler")

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Header: %v", r.Header)))
	if err != nil {
		h.logger.Error(err)
		return
	}

	_, _ = fmt.Fprint(w, "Hello, authenticated user!")
}

func NewCookieHandler(logger pkg.Logger) *CookieHandler {
	return &CookieHandler{logger: logger}
}
