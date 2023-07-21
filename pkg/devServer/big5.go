package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/traditionalchinese"
	"net/http"
)

const UrlBig5 = "/big5"

type HandlerBig5 struct {
	logger pkg.Logger
}

func (h *HandlerBig5) Pattern() string {
	return UrlBig5
}

func (h *HandlerBig5) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into HandlerBig5")
	defer func() {
		h.logger.Info("exit HandlerBig5")
	}()

	encoder := traditionalchinese.Big5.NewEncoder()
	big5Bytes, err := encoder.Bytes([]byte("漢字BIG5"))
	if err != nil {
		h.logger.Error(err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=Big5")

	_, _ = w.Write(big5Bytes)
}

func NewHandlerBig5(logger pkg.Logger) *HandlerBig5 {
	return &HandlerBig5{logger: logger}
}
