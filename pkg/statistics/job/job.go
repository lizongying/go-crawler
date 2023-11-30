package job

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Jobs []*Job

type Job struct {
	pkg.JobStatus `json:"status,omitempty"`
	Id            string               `json:"id,omitempty"`
	Schedule      string               `json:"schedule,omitempty"`
	Command       string               `json:"command,omitempty"`
	Crawler       string               `json:"crawler,omitempty"`
	Spider        string               `json:"spider,omitempty"`
	Task          uint32               `json:"task,omitempty"`
	Request       uint32               `json:"request,omitempty"`
	Item          uint32               `json:"item,omitempty"`
	Enable        bool                 `json:"enable,omitempty"`
	StartTime     utils.Timestamp      `json:"start_time,omitempty"`
	FinishTime    utils.Timestamp      `json:"finish_time,omitempty"`
	UpdateTime    utils.Timestamp      `json:"update_time,omitempty"`
	StatusList    *queue.PriorityQueue `json:"status_list,omitempty"`
	StopReason    string               `json:"stop_reason,omitempty"`
}

func (s *Job) WithStatusAndTime(status pkg.JobStatus, t time.Time) pkg.StatisticsJob {
	s.withStatus(status)
	s.withUpdateTime(t)
	switch status {
	case pkg.JobStatusRunning:
		s.withStartTime(t)
	case pkg.JobStatusSuccess:
		s.withFinishTime(t)
	case pkg.JobStatusFailure:
		s.withFinishTime(t)
	}

	if s.StatusList == nil {
		s.StatusList = queue.NewPriorityQueue(10)
	}
	s.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return s
}
func (s *Job) withStatus(status pkg.JobStatus) pkg.StatisticsJob {
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
func (s *Job) WithCrawler(crawler string) pkg.StatisticsJob {
	s.Crawler = crawler
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
func (s *Job) IncRequest() {
	atomic.AddUint32(&s.Request, 1)
}
func (s *Job) DecRequest() {
	atomic.AddUint32(&s.Request, ^uint32(0))
}
func (s *Job) IncItem() {
	atomic.AddUint32(&s.Item, 1)
}
func (s *Job) DecItem() {
	atomic.AddUint32(&s.Item, ^uint32(0))
}
func (s *Job) WithEnable(enable bool) pkg.StatisticsJob {
	s.Enable = enable
	return s
}
func (s *Job) withStartTime(t time.Time) pkg.StatisticsJob {
	s.StartTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Job) withFinishTime(t time.Time) pkg.StatisticsJob {
	s.FinishTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Job) withUpdateTime(t time.Time) pkg.StatisticsJob {
	s.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Job) GetStopReason() string {
	return s.StopReason
}
func (s *Job) WithStopReason(stopReason string) pkg.StatisticsJob {
	s.StopReason = stopReason
	return s
}
func (s *Job) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
