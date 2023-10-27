package pkg

import "time"

type StatisticsNode interface {
	WithId(string) StatisticsNode
	IncSpider()
	DecSpider()
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	GetSpider() string
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	WithLastRunAt(time.Time) StatisticsSpider
	WithLastFinishAt(time.Time) StatisticsSpider
	WithStatus(SpiderStatus) StatisticsSpider
	Marshal() (bytes []byte, err error)
}
type StatisticsSchedule interface {
	Marshal() (bytes []byte, err error)
}
type StatisticsTask interface {
	WithStatus(status TaskStatus) StatisticsTask
	GetId() string
	WithId(id string) StatisticsTask
	IncRecord()
	DecRecord()
	GetNode() string
	WithNode(string) StatisticsTask
	GetSpider() string
	WithSpider(string) StatisticsTask
	WithStartTime(startTime time.Time) StatisticsTask
	WithFinishTime(finishTime time.Time) StatisticsTask
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
