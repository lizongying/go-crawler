package test_must_ok_spider

type ExtraOk struct {
	Count int `json:"count,omitempty"`
}

type DataOk struct {
	TaskId string `json:"task_id,omitempty"`
	Count  int    `json:"count,omitempty"`
}
