package mock_servers

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
	"time"
)

const UrlTimeout = "/timeout"

type RouteTimeout struct {
	logger pkg.Logger
}

func (h *RouteTimeout) Pattern() string {
	return UrlTimeout
}

func (h *RouteTimeout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerTimeout")
	defer func() {
		h.logger.Info("exit HandlerTimeout")
	}()

	time.Sleep(5 * time.Second)

	w.WriteHeader(http.StatusBadGateway)
	_, _ = w.Write([]byte("ok"))
}

func NewRouteTimeout(logger pkg.Logger) pkg.Route {
	return &RouteTimeout{logger: logger}
}
