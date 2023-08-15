package mockServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlFile = "/statics/"

type RouteFile struct {
	http.Handler
	logger pkg.Logger
}

func (h *RouteFile) Pattern() string {
	return UrlFile
}

// NewHandlerFile curl -k -v -s https://localhost:8081/statics/images/th.jpeg
func NewRouteFile(logger pkg.Logger) pkg.Route {
	files, _ := fs.Sub(static.Statics, "statics")
	return &RouteFile{
		Handler: http.StripPrefix(UrlFile, http.FileServer(http.FS(files))),
		logger:  logger,
	}
}
