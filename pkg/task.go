package pkg

type Task struct {
	TaskId   string `json:"task_id,omitempty"`
	Deadline int    `json:"deadline,omitempty"`
}
