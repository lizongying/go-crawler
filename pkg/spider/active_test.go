package spider

import (
	"testing"
	"time"
)

func TestActive_SetActive(t *testing.T) {
	active := NewActive(time.Second * 2)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 1)
		t.Log(active.active)
		active.SetActive()
	}
	time.Sleep(time.Second * 3)
	t.Log(active.active)
}
