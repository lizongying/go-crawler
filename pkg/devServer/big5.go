package devServer

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/traditionalchinese"
	"net/http"
)

const UrlBig5 = "/big5"

type Big5Handler struct {
	logger pkg.Logger
}

func (h *Big5Handler) Pattern() string {
	return UrlBig5
}

func (h *Big5Handler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	h.logger.Info("into Big5Handler")
	defer func() {
		h.logger.Info("exit Big5Handler")
	}()

	encoder := traditionalchinese.Big5.NewEncoder()
	big5Bytes, err := encoder.Bytes([]byte("漢字BIG5"))
	if err != nil {
		h.logger.Error(err)
	}

	w.Header().Set("Content-Type", "text/plain; charset=Big5")

	_, _ = w.Write(big5Bytes)
}

func NewBig5Handler(logger pkg.Logger) *Big5Handler {
	return &Big5Handler{logger: logger}
}
