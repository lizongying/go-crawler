package context

import (
	"github.com/lizongying/go-crawler/pkg"
)

type Json struct {
	TaskId string `json:"task_id"`
}

func (j *Json) ToContext() (ctx pkg.Context) {
	ctx = &Context{
		TaskId: j.TaskId,
	}
	return
}
