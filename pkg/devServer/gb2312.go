package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGb2312 = "/gb2312"

type HandlerGb2312 struct {
	logger pkg.Logger
}

func (h *HandlerGb2312) Pattern() string {
	return UrlGb2312
}

func (h *HandlerGb2312) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into HandlerGb2312")
	defer func() {
		h.logger.Info("exit HandlerGb2312")
	}()

	encoder := simplifiedchinese.HZGB2312.NewEncoder()
	gb2312Bytes, _ := encoder.Bytes([]byte("汉字GB2312"))

	w.Header().Set("Content-Type", "text/plain; charset=GB2312")

	_, _ = w.Write(gb2312Bytes)
}

func NewHandlerGb2312(logger pkg.Logger) *HandlerGb2312 {
	return &HandlerGb2312{logger: logger}
}
