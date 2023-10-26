package spider

import (
	"encoding/json"
	"github.com/lizongying/go-crawler/pkg"
	"github.com/lizongying/go-crawler/pkg/utils"
	"sync/atomic"
	"time"
)

type Spider struct {
	pkg.SpiderStatus `json:"last_status,omitempty"`
	Spider           string               `json:"spider,omitempty"`
	Schedule         uint32               `json:"schedule,omitempty"`
	Task             uint32               `json:"task,omitempty"`
	Record           uint32               `json:"record,omitempty"`
	LastTask         pkg.StatisticsTask   `json:"last_task,omitempty"`
	LastRecord       pkg.StatisticsRecord `json:"last_record,omitempty"`
	LastRunAt        utils.Timestamp      `json:"last_run_at,omitempty"`
	LastFinishAt     utils.Timestamp      `json:"last_finish_at,omitempty"`
}

func (s *Spider) GetSpider() string {
	return s.Spider
}
func (s *Spider) IncTask() {
	atomic.AddUint32(&s.Task, 1)
}
func (s *Spider) IncRecord() {
	atomic.AddUint32(&s.Record, 1)
}
func (s *Spider) WithLastRunAt(t time.Time) pkg.StatisticsSpider {
	s.LastRunAt = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) WithLastFinishAt(t time.Time) pkg.StatisticsSpider {
	s.LastFinishAt = utils.Timestamp{
		Time: t,
	}
	return s
}
func (s *Spider) WithStatus(status pkg.SpiderStatus) pkg.StatisticsSpider {
	s.SpiderStatus = status
	return s
}
func (s *Spider) Marshal() (bytes []byte, err error) {
	bytes, err = json.Marshal(s)
	if err != nil {
		return
	}
	return
}
