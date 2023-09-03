package mockServers

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"golang.org/x/time/rate"
	"net/http"
)

type Message struct {
	Data string `json:"data"`
}

const UrlRateLimiter = "/rate-limiter"

type RouteRateLimiter struct {
	logger  pkg.Logger
	limiter *rate.Limiter
}

func (h *RouteRateLimiter) Pattern() string {
	return UrlRateLimiter
}

func (h *RouteRateLimiter) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
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

func NewRouteRateLimiter(logger pkg.Logger) pkg.Route {
	return &RouteRateLimiter{
		logger:  logger,
		limiter: rate.NewLimiter(3, 6), // rate 3/s
	}
}
