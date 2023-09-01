package mockServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlHtml = "/html/"

type RouteHtml struct {
	http.Handler
	logger pkg.Logger
}

func (h *RouteHtml) Pattern() string {
	return UrlHtml
}

// NewHandlerFile curl -k -v -s https://localhost:8081/html/a.html
func NewRouteHtml(logger pkg.Logger) pkg.Route {
	files, _ := fs.Sub(static.Statics, "html")
	return &RouteFile{
		Handler: http.StripPrefix(UrlHtml, http.FileServer(http.FS(files))),
		logger:  logger,
	}
}
