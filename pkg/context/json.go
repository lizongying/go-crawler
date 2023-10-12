package context

import (
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type Json struct {
	TaskId     string `json:"task_id,omitempty"`
	Deadline   int64  `json:"deadline,omitempty"`
	SpiderName string `json:"spider_name,omitempty"`
	StartFunc  string `json:"start_func,omitempty"`
	Args       string `json:"args,omitempty"`
	Mode       string `json:"mode,omitempty"`
}

func (j *Json) ToContext() (ctx pkg.Context) {
	ctx = &Context{
		taskId:     j.TaskId,
		deadline:   time.Unix(0, j.Deadline),
		spiderName: j.SpiderName,
		startFunc:  j.StartFunc,
		args:       j.Args,
		mode:       j.Mode,
	}
	return
}
