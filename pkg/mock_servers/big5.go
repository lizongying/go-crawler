package mock_servers

import (
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/text/encoding/traditionalchinese"
	"net/http"
)

const UrlBig5 = "/big5"

type RouteBig5 struct {
	logger pkg.Logger
}

func (h *RouteBig5) Pattern() string {
	return UrlBig5
}

func (h *RouteBig5) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
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

func NewRouteBig5(logger pkg.Logger) pkg.Route {
	return &RouteBig5{logger: logger}
}
