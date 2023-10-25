package record

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg/utils"
	"time"
)

type Record struct {
	Spider   string          `json:"spider,omitempty"`
	Schedule string          `json:"schedule,omitempty"`
	TaskId   string          `json:"task_id,omitempty"`
	Meta     string          `json:"meta,omitempty"`
	Data     string          `json:"data,omitempty"`
	SaveTime utils.Timestamp `json:"save_time,omitempty"`
}

func (r *Record) WithSpider(spider string) *Record {
	r.Spider = spider
	return r
}
func (r *Record) WithSchedule(schedule string) *Record {
	r.Schedule = schedule
	return r
}
func (r *Record) WithTaskId(taskId string) *Record {
	r.TaskId = taskId
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
