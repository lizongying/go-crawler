package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlRobotsTxt = "/robots.txt"

type HandlerRobotsTxt struct {
	http.Handler
	logger pkg.Logger
}

func (h *HandlerRobotsTxt) Pattern() string {
	return UrlRobotsTxt
}

func NewHandlerRobotsTxt(logger pkg.Logger) *HandlerRobotsTxt {
	files, _ := fs.Sub(static.Statics, "statics")
	return &HandlerRobotsTxt{
		Handler: http.FileServer(http.FS(files)),
		logger:  logger,
	}
}
