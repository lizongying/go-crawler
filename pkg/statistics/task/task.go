package task

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/statistics/node"
	"github.com/lizongying/go-crawler/pkg/statistics/schedule"
	"github.com/lizongying/go-crawler/pkg/statistics/spider"
	"sync/atomic"
	"time"
)

type Tasks []*Task

type Status uint8

const (
	StatusUnknown = iota
	StatusPending
	StatusRunning
	StatusSuccess
	StatusError
)

func (s *Status) String() string {
	switch *s {
	case 1:
		return "pending"
	case 2:
		return "running"
	case 3:
		return "success"
	case 4:
		return "error"
	default:
		return "unknown"
	}
}

type Task struct {
	Status
	Id       string
	Spider   *spider.Spider
	Schedule *schedule.Schedule
	Node     *node.Node
	Record   uint32
	Started  time.Time
	Finished time.Time
}

func (t *Task) GetId() string {
	return t.Id
}
func (t *Task) IncRecord() {
	atomic.AddUint32(&t.Record, 1)
}
func (t *Task) SetStarted(dateTime time.Time) {
	t.Started = dateTime
}
func (t *Task) SetFinished(dateTime time.Time) {
	t.Started = dateTime
}
func (t *Task) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(t)
	if err != nil {
		return
	}
	return
}
