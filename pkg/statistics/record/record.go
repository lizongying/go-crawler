package record

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Record struct {
	Id         string               `json:"id,omitempty"`
	UniqueKey  string               `json:"unique_key,omitempty"`
	Node       string               `json:"node,omitempty"`
	Spider     string               `json:"spider,omitempty"`
	Job        string               `json:"job,omitempty"`
	Task       string               `json:"task,omitempty"`
	Meta       string               `json:"meta,omitempty"`
	Data       string               `json:"data,omitempty"`
	Status     pkg.ItemStatus       `json:"status,omitempty"`
	StartTime  utils.Timestamp      `json:"start_time,omitempty"`
	FinishTime utils.Timestamp      `json:"finish_time,omitempty"`
	UpdateTime utils.Timestamp      `json:"update_time,omitempty"`
	StatusList *queue.PriorityQueue `json:"status_list,omitempty"`
	StopReason string               `json:"stop_reason,omitempty"`
}

func (r *Record) WithId(id string) *Record {
	r.Id = id
	return r
}
func (r *Record) WithUniqueKey(uniqueKey string) *Record {
	r.UniqueKey = uniqueKey
	return r
}
func (r *Record) WithNode(node string) *Record {
	r.Node = node
	return r
}
func (r *Record) WithSpider(spider string) *Record {
	r.Spider = spider
	return r
}
func (r *Record) WithJob(job string) *Record {
	r.Job = job
	return r
}
func (r *Record) WithTask(task string) *Record {
	r.Task = task
	return r
}
func (r *Record) WithMeta(meta string) *Record {
	r.Meta = meta
	return r
}
func (r *Record) WithData(data string) *Record {
	r.Data = data
	return r
}
func (r *Record) GetStatus() pkg.ItemStatus {
	return r.Status
}
func (r *Record) WithStatus(status pkg.ItemStatus) *Record {
	r.Status = status
	return r
}
func (r *Record) WithStartTime(t time.Time) *Record {
	r.StartTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Record) WithFinishTime(t time.Time) *Record {
	r.FinishTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Record) WithUpdateTime(t time.Time) *Record {
	r.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Record) AddStatusList(status pkg.SpiderStatus, t time.Time) *Record {
	if r.StatusList == nil {
		r.StatusList = queue.NewPriorityQueue(10)
	}
	r.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return r
}
func (r *Record) WithStopReason(stopReason string) *Record {
	r.StopReason = stopReason
	return r
}
func (r *Record) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	if err != nil {
		return
	}
	return
}
