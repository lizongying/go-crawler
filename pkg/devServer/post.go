package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"net/http/httputil"
)

const UrlPost = "/post"

type PostHandler struct {
	logger pkg.Logger
}

func (h *PostHandler) Pattern() string {
	return UrlPost
}

func (h *PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into PostHandler")
	defer func() {
		h.logger.Info("exit PostHandler")
	}()

	body, _ := io.ReadAll(r.Body)
	h.logger.InfoF("body: %s", string(body))

	reqDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		h.logger.Error(err)
		return
	}

	_, _ = w.Write(reqDump)
}

func NewPostHandler(logger pkg.Logger) *PostHandler {
	return &PostHandler{logger: logger}
}
