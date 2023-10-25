package schedule

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"time"
)

type Schedules []*Schedule

type Status uint8

const (
	StatusUnknown = iota
	StatusOnline
	StatusOffline
)

func (s *Status) String() string {
	switch *s {
	case 1:
		return "online"
	case 2:
		return "offline"
	default:
		return "unknown"
	}
}

type Schedule struct {
	Status
	Schedule string
	Spider   pkg.StatisticsSpider
	Enable   bool
	Started  time.Time
	Finished time.Time
}

func (s *Schedule) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
