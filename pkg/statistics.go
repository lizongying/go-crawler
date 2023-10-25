package pkg

import "time"

type StatisticsNode interface {
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	GetSpider() string
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
	AddNodes(...StatisticsNode)
	AddSpiders(...Spider)
	AddSchedules(...StatisticsSchedule)
	AddTasks(...StatisticsTask)
	AddRecords(...StatisticsRecord)
}
