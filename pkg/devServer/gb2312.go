package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGb2312 = "/gb2312"

type Gb2312Handler struct {
	logger pkg.Logger
}

func (h *Gb2312Handler) Pattern() string {
	return UrlGb2312
}

func (h *Gb2312Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into Gb2312Handler")
	defer func() {
		h.logger.Info("exit Gb2312Handler")
	}()

	encoder := simplifiedchinese.HZGB2312.NewEncoder()
	gb2312Bytes, _ := encoder.Bytes([]byte("汉字GB2312"))

	w.Header().Set("Content-Type", "text/plain; charset=GB2312")

	_, _ = w.Write(gb2312Bytes)
}

func NewGb2312Handler(logger pkg.Logger) *Gb2312Handler {
	return &Gb2312Handler{logger: logger}
}
