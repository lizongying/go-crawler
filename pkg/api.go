package pkg

type ReqSpiderStart struct {
	Name    string       `json:"name"`
	Func    string       `json:"func"`
	Args    string       `json:"args,omitempty"`
	Mode    ScheduleMode `json:"mode,omitempty"`
	Spec    string       `json:"spec,omitempty"`
	Timeout uint32       `json:"timeout,omitempty"` // second
}

type ReqSpiderStop struct {
	TaskId string `json:"task_id"` // uuid
}
