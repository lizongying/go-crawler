package schedule

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Schedules []*Schedule

type Schedule struct {
	pkg.ScheduleStatus `json:"status,omitempty"`
	Id                 string          `json:"id,omitempty"`
	Schedule           string          `json:"schedule,omitempty"`
	Node               string          `json:"node,omitempty"`
	Spider             string          `json:"spider,omitempty"`
	Task               uint32          `json:"task,omitempty"`
	Record             uint32          `json:"record,omitempty"`
	Enable             bool            `json:"enable,omitempty"`
	StartTime          utils.Timestamp `json:"start_time,omitempty"`
	FinishTime         utils.Timestamp `json:"finish_time,omitempty"`
}

func (s *Schedule) WithStatus(status pkg.ScheduleStatus) pkg.StatisticsSchedule {
	s.ScheduleStatus = status
	return s
}
func (s *Schedule) WithId(id string) pkg.StatisticsSchedule {
	s.Id = id
	return s
}
func (s *Schedule) WithSchedule(schedule string) pkg.StatisticsSchedule {
	s.Schedule = schedule
	return s
}
func (s *Schedule) WithNode(node string) pkg.StatisticsSchedule {
	s.Node = node
	return s
}
func (s *Schedule) WithSpider(spider string) pkg.StatisticsSchedule {
	s.Spider = spider
	return s
}
func (s *Schedule) IncTask() {
	atomic.AddUint32(&s.Task, 1)
}
func (s *Schedule) DecTask() {
	atomic.AddUint32(&s.Task, ^uint32(0))
}
func (s *Schedule) IncRecord() {
	atomic.AddUint32(&s.Record, 1)
}
func (s *Schedule) DecRecord() {
	atomic.AddUint32(&s.Record, ^uint32(0))
}
func (s *Schedule) WithEnable(enable bool) pkg.StatisticsSchedule {
	s.Enable = enable
	return s
}
func (s *Schedule) WithStartTime(t time.Time) pkg.StatisticsSchedule {
	s.StartTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Schedule) WithFinishTime(t time.Time) pkg.StatisticsSchedule {
	s.FinishTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Schedule) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
