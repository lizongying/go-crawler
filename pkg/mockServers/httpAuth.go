package mockServers

import (
	"fmt"
	"github.com/lizongying/go-crawler/pkg"
	"net/http"
)

const UrlHttpAuth = "/http-auth"
const username = "username"
const password = "password"

type RouteHttpAuth struct {
	logger pkg.Logger
}

func (h *RouteHttpAuth) Pattern() string {
	return UrlHttpAuth
}

func (h *RouteHttpAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerHttpAuth")
	_, err := w.Write([]byte(fmt.Sprintf("Header: %v", r.Header)))
	if err != nil {
		h.logger.Error(err)
		return
	}

	user, pass, ok := r.BasicAuth()
	if !ok || user != username || pass != password {
		// 身份验证失败，返回401 Unauthorized状态码
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = fmt.Fprint(w, "Unauthorized access")
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Hello, authenticated user!")
}

func NewRouteHttpAuth(logger pkg.Logger) pkg.Route {
	return &RouteHttpAuth{logger: logger}
}
