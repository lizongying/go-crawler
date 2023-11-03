package pkg

import (
	"fmt"
	"strings"
)

type ScheduleStatus uint8

const (
	ScheduleStatusUnknown = iota
	ScheduleStatusStarted
	ScheduleStatusStopped
)

func (s *ScheduleStatus) String() string {
	switch *s {
	case 1:
		return "started"
	case 2:
		return "stopped"
	default:
		return "unknown"
	}
}

type ScheduleMode uint8

const (
	ScheduleModeUnknown = iota
	ScheduleModeOnce
	ScheduleModeLoop
	ScheduleModeCron
)

func (s *ScheduleMode) String() string {
	switch *s {
	case 1:
		return "once"
	case 2:
		return "loop"
	case 3:
		return "cron"
	default:
		return "manual"
	}
}
func (s *ScheduleMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", *s)), nil
}
func (s *ScheduleMode) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "1":
		*s = ScheduleModeOnce
	case "2":
		*s = ScheduleModeLoop
	case "3":
		*s = ScheduleModeCron
	default:
		*s = ScheduleModeUnknown
	}
	return nil
}
func ScheduleModeFromString(name string) ScheduleMode {
	switch strings.ToLower(name) {
	case "1":
		return ScheduleModeOnce
	case "2":
		return ScheduleModeLoop
	case "3":
		return ScheduleModeLoop
	case "once":
		return ScheduleModeOnce
	case "loop":
		return ScheduleModeLoop
	case "cron":
		return ScheduleModeCron
	default:
		return ScheduleModeUnknown
	}
}
