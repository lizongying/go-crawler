package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGb18030 = "/gb18030"

type Gb18030Handler struct {
	logger pkg.Logger
}

func (h *Gb18030Handler) Pattern() string {
	return UrlGb18030
}

func (h *Gb18030Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into Gb18030Handler")
	defer func() {
		h.logger.Info("exit Gb18030Handler")
	}()

	encoder := simplifiedchinese.GB18030.NewEncoder()
	gb18030Bytes, _ := encoder.Bytes([]byte("汉字GB18030"))

	w.Header().Set("Content-Type", "text/plain; charset=GB18030")

	_, _ = w.Write(gb18030Bytes)
}

func NewGb18030Handler(logger pkg.Logger) *Gb18030Handler {
	return &Gb18030Handler{logger: logger}
}
