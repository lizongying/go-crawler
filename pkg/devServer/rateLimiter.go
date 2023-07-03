package devServer

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/logger"
	"golang.org/x/time/rate"
	"net/http"
)

type Message struct {
	Data string `json:"data"`
}

const UrlRateLimiter = "/rate-limiter"

type RateLimiterHandler struct {
	logger  *logger.Logger
	limiter *rate.Limiter
}

func (*RateLimiterHandler) Pattern() string {
	return UrlRateLimiter
}

func (h *RateLimiterHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	message := Message{
		Data: "Success",
	}
	if !h.limiter.Allow() {
		message.Data = "Rate limit"
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		h.logger.Error(err)
		return
	}
}

func NewRateLimiterHandler(logger *logger.Logger) *RateLimiterHandler {
	return &RateLimiterHandler{
		logger:  logger,
		limiter: rate.NewLimiter(3, 6), // rate 3/s
	}
}
