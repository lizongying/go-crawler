package mockServers

import (
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"net/http/httputil"
)

const UrlPost = "/post"

type RoutePost struct {
	logger pkg.Logger
}

func (h *RoutePost) Pattern() string {
	return UrlPost
}

func (h *RoutePost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerPost")
	defer func() {
		h.logger.Info("exit HandlerPost")
	}()

	body, _ := io.ReadAll(r.Body)
	h.logger.Infof("body: %s", string(body))

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		h.logger.Error(err)
		return
	}

	_, _ = w.Write(reqDump)
}

func NewRoutePost(logger pkg.Logger) pkg.Route {
	return &RoutePost{logger: logger}
}
