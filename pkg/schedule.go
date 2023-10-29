package pkg

type ScheduleStatus uint8

const (
	ScheduleStatusUnknown = iota
	ScheduleStatusStarted
	ScheduleStatusPending
	ScheduleStatusRunning
	ScheduleStatusStopped
)

func (s *ScheduleStatus) String() string {
	switch *s {
	case 1:
		return "started"
	case 2:
		return "pending"
	case 3:
		return "running"
	case 4:
		return "stopped"
	default:
		return "unknown"
	}
}
