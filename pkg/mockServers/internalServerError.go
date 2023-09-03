package mockServers

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlInternalServerError = "/internal-server-error"

type RouteInternalServerError struct {
	logger pkg.Logger
}

func (h *RouteInternalServerError) Pattern() string {
	return UrlInternalServerError
}

func (h *RouteInternalServerError) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerInternalServerError")
	defer func() {
		h.logger.Info("exit HandlerInternalServerError")
	}()

	w.WriteHeader(http.StatusInternalServerError)
	_, err := w.Write([]byte(`InternalServerError`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewRouteInternalServerError(logger pkg.Logger) pkg.Route {
	return &RouteInternalServerError{logger: logger}
}
