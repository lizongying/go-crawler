package pkg

type Task interface {
	Scheduler
	GetScheduler() Scheduler
	WithScheduler(Scheduler) Task
	RequestIn()
	RequestOut()
	ItemIn()
	ItemOut()
	MethodIn()
	MethodOut()
}

type TaskStatus uint8

const (
	TaskStatusUnknown = iota
	TaskStatusPending
	TaskStatusRunning
	TaskStatusSuccess
	TaskStatusFailure
)

func (s TaskStatus) String() string {
	switch s {
	case TaskStatusPending:
		return "pending"
	case TaskStatusRunning:
		return "running"
	case TaskStatusSuccess:
		return "success"
	case TaskStatusFailure:
		return "failure"
	default:
		return "unknown"
	}
}
