package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net/http"
)

const UrlGbk = "/gbk"

type GbkHandler struct {
	logger pkg.Logger
}

func (h *GbkHandler) Pattern() string {
	return UrlGbk
}

func (h *GbkHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into GbkHandler")
	defer func() {
		h.logger.Info("exit GbkHandler")
	}()

	encoder := simplifiedchinese.GBK.NewEncoder()
	gbkBytes, _ := encoder.Bytes([]byte("汉字GBK"))

	w.Header().Set("Content-Type", "text/plain; charset=GBK")

	_, _ = w.Write(gbkBytes)
}

func NewGbkHandler(logger pkg.Logger) *GbkHandler {
	return &GbkHandler{logger: logger}
}
