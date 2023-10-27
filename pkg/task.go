package pkg

type TaskStatus uint8

const (
	TaskStatusUnknown = iota
	TaskStatusPending
	TaskStatusRunning
	TaskStatusSuccess
	TaskStatusError
)

func (s *TaskStatus) String() string {
	switch *s {
	case 1:
		return "pending"
	case 2:
		return "running"
	case 3:
		return "success"
	case 4:
		return "error"
	default:
		return "unknown"
	}
}
