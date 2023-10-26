package pkg

import "time"

type StatisticsNode interface {
	WithId(string) StatisticsNode
	IncRecord()
	IncSpider()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	GetSpider() string
	IncRecord()
	WithLastRunAt(time.Time) StatisticsSpider
	WithLastFinishAt(time.Time) StatisticsSpider
	WithStatus(SpiderStatus) StatisticsSpider
	Marshal() (bytes []byte, err error)
}
type StatisticsSchedule interface {
	Marshal() (bytes []byte, err error)
}
type StatisticsTask interface {
	GetId() string
	SetStarted(dateTime time.Time)
	SetFinished(dateTime time.Time)
	Marshal() (bytes []byte, err error)
}
type StatisticsRecord interface {
	Marshal() (bytes []byte, err error)
}
type Statistics interface {
	GetNodes() []StatisticsNode
	GetSpiders() []StatisticsSpider
	GetSchedules() []StatisticsSchedule
	GetTasks() []StatisticsTask
	GetRecords() []StatisticsRecord
	AddSpiders(...Spider)
	AddSchedules(...StatisticsSchedule)
	AddTasks(...StatisticsTask)
	AddRecords(...StatisticsRecord)
}
