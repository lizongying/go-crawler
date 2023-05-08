package devServer

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/logger"
	"golang.org/x/time/rate"
	"net/http"
)

type Message struct {
	Response    string `json:"response"`
	Description string `json:"description"`
}

type RateLimiterHandler struct {
	logger  *logger.Logger
	limiter *rate.Limiter
}

func (*RateLimiterHandler) Pattern() string {
	return "/rate-limiter"
}

func (h *RateLimiterHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	if !h.limiter.Allow() {
		_, _ = w.Write([]byte("rate limit exceeded "))
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	message := Message{
		Response:    "Successful",
		Description: "You've successfully hit the API endpoint",
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		return
	}
}

func NewRateLimiterHandler(logger *logger.Logger) *RateLimiterHandler {
	return &RateLimiterHandler{
		logger:  logger,
		limiter: rate.NewLimiter(3, 6), // max of 6 requests and then three more requests per second
	}
}
