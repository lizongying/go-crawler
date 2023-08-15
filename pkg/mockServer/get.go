package mockServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"net/http/httputil"
)

const UrlGet = "/get"

type RouteGet struct {
	logger pkg.Logger
}

func (h *RouteGet) Pattern() string {
	return UrlGet
}

func (h *RouteGet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

func NewRouteGet(logger pkg.Logger) pkg.Route {
	return &RouteGet{logger: logger}
}
