package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"io"
	"net/http"
	"net/http/httputil"
)

const UrlPost = "/post"

type HandlerPost struct {
	logger pkg.Logger
}

func (h *HandlerPost) Pattern() string {
	return UrlPost
}

func (h *HandlerPost) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("into HandlerPost")
	defer func() {
		h.logger.Info("exit HandlerPost")
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

func NewHandlerPost(logger pkg.Logger) *HandlerPost {
	return &HandlerPost{logger: logger}
}
