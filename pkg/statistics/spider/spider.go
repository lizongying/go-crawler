package spider

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Spider struct {
	pkg.SpiderStatus   `json:"status,omitempty"`
	Id                 *utils.Uint64        `json:"id,omitempty"`
	Node               string               `json:"node,omitempty"`
	Spider             string               `json:"spider,omitempty"`
	Funcs              []string             `json:"funcs,omitempty"`
	Job                uint32               `json:"job,omitempty"`
	Task               uint32               `json:"task,omitempty"`
	Request            uint32               `json:"request,omitempty"`
	Record             uint32               `json:"record,omitempty"`
	StartTime          utils.Timestamp      `json:"start_time,omitempty"`
	FinishTime         utils.Timestamp      `json:"finish_time,omitempty"`
	LastTaskId         string               `json:"last_task_id,omitempty"`
	LastTaskStatus     pkg.TaskStatus       `json:"last_task_status,omitempty"`
	LastTaskStartTime  utils.Timestamp      `json:"last_task_start_time,omitempty"`
	LastTaskFinishTime utils.Timestamp      `json:"last_task_finish_time,omitempty"`
	UpdateTime         utils.Timestamp      `json:"update_time,omitempty"`
	StatusList         *queue.PriorityQueue `json:"status_list,omitempty"`
	StopReason         string               `json:"stop_reason,omitempty"`
}

func (s *Spider) WithStatusAndTime(status pkg.SpiderStatus, t time.Time) pkg.StatisticsSpider {
	s.withStatus(status)
	s.withUpdateTime(t)
	switch status {
	case pkg.SpiderStatusRunning:
		s.withStartTime(t)
	case pkg.SpiderStatusStopped:
		s.withFinishTime(t)
	}

	if s.StatusList == nil {
		s.StatusList = queue.NewPriorityQueue(10)
	}
	s.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return s
}
func (s *Spider) GetId() uint64 {
	return s.Id.Uint64()
}
func (s *Spider) WithId(id uint64) pkg.StatisticsSpider {
	s.Id = utils.NewUint64(id)
	return s
}
func (s *Spider) GetSpider() string {
	return s.Spider
}
func (s *Spider) WithSpider(spider string) pkg.StatisticsSpider {
	s.Spider = spider
	return s
}
func (s *Spider) GetFuncs() []string {
	return s.Funcs
}
func (s *Spider) WithFuncs(funcs []string) pkg.StatisticsSpider {
	s.Funcs = funcs
	return s
}
func (s *Spider) GetNode() string {
	return s.Node
}
func (s *Spider) WithNode(node string) pkg.StatisticsSpider {
	s.Node = node
	return s
}
func (s *Spider) IncJob() {
	atomic.AddUint32(&s.Job, 1)
}
func (s *Spider) DecJob() {
	atomic.AddUint32(&s.Job, ^uint32(0))
}
func (s *Spider) IncTask() {
	atomic.AddUint32(&s.Task, 1)
}
func (s *Spider) DecTask() {
	atomic.AddUint32(&s.Task, ^uint32(0))
}
func (s *Spider) IncRequest() {
	atomic.AddUint32(&s.Request, 1)
}
func (s *Spider) DecRequest() {
	atomic.AddUint32(&s.Request, ^uint32(0))
}
func (s *Spider) IncRecord() {
	atomic.AddUint32(&s.Record, 1)
}
func (s *Spider) DecRecord() {
	atomic.AddUint32(&s.Record, ^uint32(0))
}
func (s *Spider) GetStatus() pkg.SpiderStatus {
	return s.SpiderStatus
}
func (s *Spider) withStatus(status pkg.SpiderStatus) pkg.StatisticsSpider {
	s.SpiderStatus = status
	return s
}
func (s *Spider) withStartTime(t time.Time) pkg.StatisticsSpider {
	s.StartTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) withFinishTime(t time.Time) pkg.StatisticsSpider {
	s.FinishTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) withUpdateTime(t time.Time) pkg.StatisticsSpider {
	s.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) GetLastTaskId() string {
	return s.LastTaskId
}
func (s *Spider) WithLastTaskId(id string) pkg.StatisticsSpider {
	s.LastTaskId = id
	return s
}
func (s *Spider) WithLastTaskStatus(status pkg.TaskStatus) pkg.StatisticsSpider {
	s.LastTaskStatus = status
	return s
}
func (s *Spider) WithLastTaskStartTime(t time.Time) pkg.StatisticsSpider {
	s.LastTaskStartTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) WithLastTaskFinishTime(t time.Time) pkg.StatisticsSpider {
	s.LastTaskFinishTime = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
