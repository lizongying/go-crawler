package record

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Record struct {
	Id        string          `json:"id,omitempty"`
	UniqueKey string          `json:"unique_key,omitempty"`
	Node      string          `json:"node,omitempty"`
	Spider    string          `json:"spider,omitempty"`
	Job       string          `json:"job,omitempty"`
	Task      string          `json:"task,omitempty"`
	Meta      string          `json:"meta,omitempty"`
	Data      string          `json:"data,omitempty"`
	SaveTime  utils.Timestamp `json:"save_time,omitempty"`
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
func (r *Record) WithSaveTime(t time.Time) *Record {
	r.SaveTime = utils.Timestamp{
		Time: t,
	}
	return r
}
func (r *Record) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(r)
	if err != nil {
		return
	}
	return
}
