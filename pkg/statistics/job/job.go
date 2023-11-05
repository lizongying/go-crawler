package job

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Jobs []*Job

type Job struct {
	pkg.JobStatus `json:"status,omitempty"`
	Id            string          `json:"id,omitempty"`
	Schedule      string          `json:"schedule,omitempty"`
	Command       string          `json:"command,omitempty"`
	Node          string          `json:"node,omitempty"`
	Spider        string          `json:"spider,omitempty"`
	Task          uint32          `json:"task,omitempty"`
	Record        uint32          `json:"record,omitempty"`
	Enable        bool            `json:"enable,omitempty"`
	StartTime     utils.Timestamp `json:"start_time,omitempty"`
	FinishTime    utils.Timestamp `json:"finish_time,omitempty"`
}

func (s *Job) WithStatus(status pkg.JobStatus) pkg.StatisticsJob {
	s.JobStatus = status
	return s
}
func (s *Job) WithId(id string) pkg.StatisticsJob {
	s.Id = id
	return s
}
func (s *Job) WithSchedule(schedule string) pkg.StatisticsJob {
	s.Schedule = schedule
	return s
}
func (s *Job) WithCommand(command string) pkg.StatisticsJob {
	s.Command = command
	return s
}
func (s *Job) WithNode(node string) pkg.StatisticsJob {
	s.Node = node
	return s
}
func (s *Job) WithSpider(spider string) pkg.StatisticsJob {
	s.Spider = spider
	return s
}
func (s *Job) IncTask() {
	atomic.AddUint32(&s.Task, 1)
}
func (s *Job) DecTask() {
	atomic.AddUint32(&s.Task, ^uint32(0))
}
func (s *Job) IncRecord() {
	atomic.AddUint32(&s.Record, 1)
}
func (s *Job) DecRecord() {
	atomic.AddUint32(&s.Record, ^uint32(0))
}
func (s *Job) WithEnable(enable bool) pkg.StatisticsJob {
	s.Enable = enable
	return s
}
func (s *Job) WithStartTime(t time.Time) pkg.StatisticsJob {
	s.StartTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Job) WithFinishTime(t time.Time) pkg.StatisticsJob {
	s.FinishTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Job) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
