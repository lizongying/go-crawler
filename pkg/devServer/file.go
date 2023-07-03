package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/static"
	"io/fs"
	"net/http"
)

const UrlFile = "/statics/"

type FileHandler struct {
	http.Handler
	logger pkg.Logger
}

func (h *FileHandler) Pattern() string {
	return UrlFile
}

// NewFileHandler curl -k -v -s https://localhost:8081/statics/images/th.jpeg
func NewFileHandler(logger pkg.Logger) *FileHandler {
	files, _ := fs.Sub(static.Statics, "statics")
	return &FileHandler{
		Handler: http.StripPrefix(UrlFile, http.FileServer(http.FS(files))),
		logger:  logger,
	}
}
