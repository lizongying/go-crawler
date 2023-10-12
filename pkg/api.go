package pkg

type ReqSpiderStart struct {
	TaskId  string `json:"task_id"` // uuid
	Timeout uint32 `json:"timeout"` // second
	Name    string `json:"name"`
	Func    string `json:"func"`
	Args    string `json:"args"`
	Mode    string `json:"mode"`
}

type ReqSpiderStop struct {
	TaskId string `json:"task_id"` // uuid
}
