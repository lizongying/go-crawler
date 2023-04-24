package spider

import (
	"time"
)

type Active struct {
	active   bool
	time     time.Time
	duration time.Duration
	timer    *time.Timer
}

func (a *Active) SetActive() {
	a.active = true
	now := time.Now()
	a.time = now
	go func() {
		a.timer.Reset(a.duration)
		<-a.timer.C
		a.active = false
	}()
}

func NewActive(duration time.Duration) (active *Active) {
	active = &Active{
		duration: duration,
		timer:    time.NewTimer(duration),
	}

	return
}
