package mockServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
)

const UrlHello = "/hello"

type RouteHello struct {
	logger pkg.Logger
}

func (h *RouteHello) Pattern() string {
	return UrlHello
}

func (h *RouteHello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerHello")
	defer func() {
		h.logger.Info("exit HandlerHello")
	}()

	h.logger.InfoF("request: %+v", r)
	body, _ := io.ReadAll(r.Body)
	h.logger.InfoF("body: %s", string(body))
}

func NewRouteHello(logger pkg.Logger) pkg.Route {
	return &RouteHello{logger: logger}
}
