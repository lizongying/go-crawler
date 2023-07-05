package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlRobotsTxt = "/robots.txt"

type RobotsTxtHandler struct {
	http.Handler
	logger pkg.Logger
}

func (h *RobotsTxtHandler) Pattern() string {
	return UrlRobotsTxt
}

func NewRobotsTxtHandler(logger pkg.Logger) *RobotsTxtHandler {
	files, _ := fs.Sub(static.Statics, "statics")
	return &RobotsTxtHandler{
		Handler: http.FileServer(http.FS(files)),
		logger:  logger,
	}
}
