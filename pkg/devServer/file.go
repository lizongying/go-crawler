package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlFile = "/statics/"

type HandlerFile struct {
	http.Handler
	logger pkg.Logger
}

func (h *HandlerFile) Pattern() string {
	return UrlFile
}

// NewHandlerFile curl -k -v -s https://localhost:8081/statics/images/th.jpeg
func NewHandlerFile(logger pkg.Logger) *HandlerFile {
	files, _ := fs.Sub(static.Statics, "statics")
	return &HandlerFile{
		Handler: http.StripPrefix(UrlFile, http.FileServer(http.FS(files))),
		logger:  logger,
	}
}
