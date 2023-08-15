package mockServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlCustom = "/custom"

type RouteCustom struct {
	logger pkg.Logger
}

func (h *RouteCustom) Pattern() string {
	return UrlCustom
}

func (h *RouteCustom) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`Custom`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewRouteCustom(logger pkg.Logger) pkg.Route {
	return &RouteCustom{logger: logger}
}
