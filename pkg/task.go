package pkg

type Task interface {
	Scheduler
	GetScheduler() Scheduler
	WithScheduler(Scheduler) Task
	RequestPending(ctx Context, err error)
	RequestRunning(ctx Context, err error)
	RequestStopped(ctx Context, err error)
	ItemPending(ctx Context, err error)
	ItemRunning(ctx Context, err error)
	ItemStopped(ctx Context, err error)
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
	case 1:
		return "pending"
	case 2:
		return "running"
	case 3:
		return "success"
	case 4:
		return "failure"
	default:
		return "unknown"
	}
}
