package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGb18030 = "/gb18030"

type HandlerGb18030 struct {
	logger pkg.Logger
}

func (h *HandlerGb18030) Pattern() string {
	return UrlGb18030
}

func (h *HandlerGb18030) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into HandlerGb18030")
	defer func() {
		h.logger.Info("exit HandlerGb18030")
	}()

	encoder := simplifiedchinese.GB18030.NewEncoder()
	gb18030Bytes, _ := encoder.Bytes([]byte("汉字GB18030"))

	w.Header().Set("Content-Type", "text/plain; charset=GB18030")

	_, _ = w.Write(gb18030Bytes)
}

func NewHandlerGb18030(logger pkg.Logger) *HandlerGb18030 {
	return &HandlerGb18030{logger: logger}
}
