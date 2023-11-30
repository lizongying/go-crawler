package item

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/queue"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Item struct {
	Id         string               `json:"id,omitempty"`
	UniqueKey  string               `json:"unique_key,omitempty"`
	Crawler    string               `json:"crawler,omitempty"`
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

func (r *Item) WithId(id string) *Item {
	r.Id = id
	return r
}
func (r *Item) WithUniqueKey(uniqueKey string) *Item {
	r.UniqueKey = uniqueKey
	return r
}
func (r *Item) WithCrawler(crawler string) *Item {
	r.Crawler = crawler
	return r
}
func (r *Item) WithSpider(spider string) *Item {
	r.Spider = spider
	return r
}
func (r *Item) WithJob(job string) *Item {
	r.Job = job
	return r
}
func (r *Item) WithTask(task string) *Item {
	r.Task = task
	return r
}
func (r *Item) WithMeta(meta string) *Item {
	r.Meta = meta
	return r
}
func (r *Item) WithData(data string) *Item {
	r.Data = data
	return r
}
func (r *Item) GetStatus() pkg.ItemStatus {
	return r.Status
}
func (r *Item) WithStatus(status pkg.ItemStatus) *Item {
	r.Status = status
	return r
}
func (r *Item) WithStartTime(t time.Time) *Item {
	r.StartTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Item) WithFinishTime(t time.Time) *Item {
	r.FinishTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Item) WithUpdateTime(t time.Time) *Item {
	r.UpdateTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Item) AddStatusList(status pkg.SpiderStatus, t time.Time) *Item {
	if r.StatusList == nil {
		r.StatusList = queue.NewPriorityQueue(10)
	}
	r.StatusList.Push(queue.NewItem(status, t.UnixNano()))
	return r
}
func (r *Item) WithStopReason(stopReason string) *Item {
	r.StopReason = stopReason
	return r
}
func (r *Item) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	if err != nil {
		return
	}
	return
}
