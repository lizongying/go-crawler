package request

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Request struct {
	Id         string               `json:"id,omitempty"`
	UniqueKey  string               `json:"unique_key,omitempty"`
	Crawler    string               `json:"crawler,omitempty"`
	Spider     string               `json:"spider,omitempty"`
	Job        string               `json:"job,omitempty"`
	Task       string               `json:"task,omitempty"`
	Meta       string               `json:"meta,omitempty"`
	Data       string               `json:"data,omitempty"`
	Status     pkg.RequestStatus    `json:"status,omitempty"`
	StartTime  utils.Timestamp      `json:"start_time,omitempty"`
	FinishTime utils.Timestamp      `json:"finish_time,omitempty"`
	UpdateTime utils.Timestamp      `json:"update_time,omitempty"`
	StatusList *queue.PriorityQueue `json:"status_list,omitempty"`
	StopReason string               `json:"stop_reason,omitempty"`
}

func (r *Request) WithId(id string) *Request {
	r.Id = id
	return r
}
func (r *Request) WithUniqueKey(uniqueKey string) *Request {
	r.UniqueKey = uniqueKey
	return r
}
func (r *Request) WithCrawler(crawler string) *Request {
	r.Crawler = crawler
	return r
}
func (r *Request) WithSpider(spider string) *Request {
	r.Spider = spider
	return r
}
func (r *Request) WithJob(job string) *Request {
	r.Job = job
	return r
}
func (r *Request) WithTask(task string) *Request {
	r.Task = task
	return r
}
func (r *Request) WithMeta(meta string) *Request {
	r.Meta = meta
	return r
}
func (r *Request) WithData(data string) *Request {
	r.Data = data
	return r
}
func (r *Request) GetStatus() pkg.RequestStatus {
	return r.Status
}
func (r *Request) WithStatus(status pkg.RequestStatus) *Request {
	r.Status = status
	return r
}
func (r *Request) WithStartTime(t time.Time) *Request {
	r.StartTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Request) WithFinishTime(t time.Time) *Request {
	r.FinishTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Request) WithUpdateTime(t time.Time) *Request {
	r.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Request) AddStatusList(status pkg.SpiderStatus, t time.Time) *Request {
	if r.StatusList == nil {
		r.StatusList = queue.NewPriorityQueue(10)
	}
	r.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return r
}
func (r *Request) WithStopReason(stopReason string) *Request {
	r.StopReason = stopReason
	return r
}
func (r *Request) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	if err != nil {
		return
	}
	return
}
