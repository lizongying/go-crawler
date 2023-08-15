package mockServer

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlOk = "/ok"

type RouteOk struct {
	logger pkg.Logger
}

func (h *RouteOk) Pattern() string {
	return UrlOk
}

func (h *RouteOk) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	h.logger.Debug("into HandlerOk")
	defer func() {
		h.logger.Debug("exit HandlerOk")
	}()

	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(fmt.Sprintf("Header: %v", request.Header)))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewRouteOk(logger pkg.Logger) pkg.Route {
	return &RouteOk{logger: logger}
}
