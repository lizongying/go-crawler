package mockServers

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlRedirect = "/redirect"

type RouteRedirect struct {
	logger pkg.Logger
}

func (h *RouteRedirect) Pattern() string {
	return UrlRedirect
}

func (h *RouteRedirect) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerRedirect")
	defer func() {
		h.logger.Info("exit HandlerRedirect")
	}()

	//http.Redirect(w, r, UrlOk, http.StatusFound)

	w.Header().Set("Location", UrlOk)
	w.WriteHeader(http.StatusFound)
}

func NewRouteRedirect(logger pkg.Logger) pkg.Route {
	return &RouteRedirect{logger: logger}
}
