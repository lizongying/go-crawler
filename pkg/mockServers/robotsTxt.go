package mockServers

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlRobotsTxt = "/robots.txt"

type RouteRobotsTxt struct {
	http.Handler
	logger pkg.Logger
}

func (h *RouteRobotsTxt) Pattern() string {
	return UrlRobotsTxt
}

func NewRouteRobotsTxt(logger pkg.Logger) pkg.Route {
	files, _ := fs.Sub(static.Statics, "statics")
	return &RouteRobotsTxt{
		Handler: http.FileServer(http.FS(files)),
		logger:  logger,
	}
}
