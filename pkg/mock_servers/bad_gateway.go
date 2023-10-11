package mock_servers

import (
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlBadGateway = "/bad-gateway"

type RouteBadGateway struct {
	logger pkg.Logger
}

func (h *RouteBadGateway) Pattern() string {
	return UrlBadGateway
}

func (h *RouteBadGateway) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Debug("into HandlerBadGateway")
	defer func() {
		h.logger.Debug("exit HandlerBadGateway")
	}()

	w.WriteHeader(http.StatusBadGateway)
	_, err := w.Write([]byte(`BadGateway`))
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewRouteBadGateway(logger pkg.Logger) pkg.Route {
	return &RouteBadGateway{logger: logger}
}
