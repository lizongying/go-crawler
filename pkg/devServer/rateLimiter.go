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

type HandlerRateLimiter struct {
	logger  *logger.Logger
	limiter *rate.Limiter
}

func (*HandlerRateLimiter) Pattern() string {
	return UrlRateLimiter
}

func (h *HandlerRateLimiter) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
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

func NewHandlerRateLimiter(logger *logger.Logger) *HandlerRateLimiter {
	return &HandlerRateLimiter{
		logger:  logger,
		limiter: rate.NewLimiter(3, 6), // rate 3/s
	}
}
