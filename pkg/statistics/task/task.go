package task

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type WithRecords struct {
	Task     pkg.StatisticsTask
	Requests *queue.GroupQueue
	Records  *queue.GroupQueue
}

type Task struct {
	pkg.TaskStatus `json:"status,omitempty"`
	Id             string               `json:"id,omitempty"`
	Spider         string               `json:"spider,omitempty"`
	Job            string               `json:"job,omitempty"`
	Node           string               `json:"node,omitempty"`
	Request        uint32               `json:"request,omitempty"`
	Record         uint32               `json:"record,omitempty"`
	StartTime      utils.Timestamp      `json:"start_time"`
	FinishTime     utils.Timestamp      `json:"finish_time"`
	UpdateTime     utils.Timestamp      `json:"update_time,omitempty"`
	StatusList     *queue.PriorityQueue `json:"status_list,omitempty"`
	StopReason     string               `json:"stop_reason,omitempty"`
}

func (t *Task) WithStatus(status pkg.TaskStatus) pkg.StatisticsTask {
	t.TaskStatus = status
	return t
}
func (t *Task) GetId() string {
	return t.Id
}
func (t *Task) WithId(id string) pkg.StatisticsTask {
	t.Id = id
	return t
}
func (t *Task) GetNode() string {
	return t.Node
}
func (t *Task) WithNode(node string) pkg.StatisticsTask {
	t.Node = node
	return t
}
func (t *Task) GetSpider() string {
	return t.Spider
}
func (t *Task) WithSpider(spider string) pkg.StatisticsTask {
	t.Spider = spider
	return t
}
func (t *Task) GetJob() string {
	return t.Job
}
func (t *Task) WithJob(job string) pkg.StatisticsTask {
	t.Job = job
	return t
}
func (t *Task) IncRequest() {
	atomic.AddUint32(&t.Request, 1)
}
func (t *Task) DecRequest() {
	atomic.AddUint32(&t.Request, ^uint32(0))
}
func (t *Task) IncRecord() {
	atomic.AddUint32(&t.Record, 1)
}
func (t *Task) DecRecord() {
	atomic.AddUint32(&t.Record, ^uint32(0))
}
func (t *Task) WithStartTime(startTime time.Time) pkg.StatisticsTask {
	t.StartTime = utils.Timestamp{
		Time: startTime,
	}
	return t
}
func (t *Task) WithFinishTime(finishTime time.Time) pkg.StatisticsTask {
	t.FinishTime = utils.Timestamp{
		Time: finishTime,
	}
	return t
}
func (t *Task) WithUpdateTime(updateTime time.Time) pkg.StatisticsTask {
	t.FinishTime = utils.Timestamp{
		Time: updateTime,
	}
	return t
}
func (t *Task) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(t)
	if err != nil {
		return
	}
	return
}
func (t *Task) GetStopReason() string {
	return t.StopReason
}
func (t *Task) WithStopReason(stopReason string) pkg.StatisticsTask {
	t.StopReason = stopReason
	return t
}
