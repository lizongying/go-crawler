package spider

import (
	"context"
	"fmt"
	"go.uber.org/ratelimit"
	"golang.org/x/time/rate"
	"testing"
	"time"
)

func TestRate(t *testing.T) {
	m := make(map[string]*rate.Limiter)
	m["a"] = rate.NewLimiter(rate.Every(time.Second)*2, 2)
	go func() {
		time.AfterFunc(time.Second*5, func() {
			m["a"].SetBurst(1)
			m["a"].SetLimit(rate.Every(time.Second) * 1)
		})
	}()

	ctx := context.Background()
	for i := 0; i < 30; i++ {
		err := m["a"].Wait(ctx)
		if err != nil {
			t.Log(err)
			return
		}
		t.Log(i, time.Now())
	}
}

func TestUberRate(t *testing.T) {
	rl := ratelimit.New(1, ratelimit.WithoutSlack)

	go func() {
		time.AfterFunc(time.Second*5, func() {
			rl = ratelimit.New(2, ratelimit.WithoutSlack)
		})
	}()

	last := time.Now()
	for i := 0; i < 10; i++ {
		rl.Take()
		cur := time.Now()
		fmt.Println("last", cur.Sub(last))
		last = cur
	}
}
