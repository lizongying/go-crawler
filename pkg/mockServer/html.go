package mockServer

import (
	"fmt"
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

// NewRouteHtml curl -k -v -s https://localhost:8081/html/index.html
func NewRouteHtml(logger pkg.Logger) pkg.Route {
	fmt.Println(111)
	files, _ := fs.Sub(static.Html, "html")
	return &RouteHtml{
		Handler: http.StripPrefix(UrlHtml, http.FileServer(http.FS(files))),
		logger:  logger,
	}
}
