package pkg

import (
	"fmt"
	"strings"
)

type JobStatus uint8

const (
	JobStatusUnknown = iota
	JobStatusReady
	JobStatusStarting
	JobStatusRunning
	JobStatusIdle
	JobStatusStopping
	JobStatusSuccess
	JobStatusFailure
)

func (s JobStatus) String() string {
	switch s {
	case JobStatusReady:
		return "ready"
	case JobStatusStarting:
		return "starting"
	case JobStatusRunning:
		return "running"
	case JobStatusIdle:
		return "idle"
	case JobStatusStopping:
		return "stopping"
	case JobStatusSuccess:
		return "success"
	case JobStatusFailure:
		return "failure"
	default:
		return "unknown"
	}
}

type JobMode uint8

const (
	JobModeUnknown = iota
	JobModeOnce
	JobModeLoop
	JobModeCron
)

func (s *JobMode) String() string {
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
func (s *JobMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", *s)), nil
}
func (s *JobMode) UnmarshalJSON(bytes []byte) error {
	switch string(bytes) {
	case "1":
		*s = JobModeOnce
	case "2":
		*s = JobModeLoop
	case "3":
		*s = JobModeCron
	default:
		*s = JobModeUnknown
	}
	return nil
}
func JobModeFromString(name string) JobMode {
	switch strings.ToLower(name) {
	case "1":
		return JobModeOnce
	case "2":
		return JobModeLoop
	case "3":
		return JobModeLoop
	case "once":
		return JobModeOnce
	case "loop":
		return JobModeLoop
	case "cron":
		return JobModeCron
	default:
		return JobModeUnknown
	}
}
