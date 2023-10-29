package pkg

import "time"

type StatisticsNode interface {
	WithId(string) StatisticsNode
	IncSpider()
	DecSpider()
	IncSchedule()
	DecSchedule()
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	GetSpider() string
	IncSchedule()
	DecSchedule()
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	WithStatus(SpiderStatus) StatisticsSpider
	WithStartTime(time.Time) StatisticsSpider
	WithFinishTime(time.Time) StatisticsSpider
	GetLastTaskId() string
	WithLastTaskId(string) StatisticsSpider
	WithLastTaskStatus(TaskStatus) StatisticsSpider
	WithLastTaskStartTime(time.Time) StatisticsSpider
	WithLastTaskFinishTime(time.Time) StatisticsSpider
	Marshal() (bytes []byte, err error)
}
type StatisticsSchedule interface {
	WithStatus(status ScheduleStatus) StatisticsSchedule
	WithId(id string) StatisticsSchedule
	WithSchedule(schedule string) StatisticsSchedule
	WithNode(node string) StatisticsSchedule
	WithSpider(spider string) StatisticsSchedule
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	WithEnable(enable bool) StatisticsSchedule
	WithStartTime(time.Time) StatisticsSchedule
	WithFinishTime(time.Time) StatisticsSchedule
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
	GetSchedule() string
	WithSchedule(string) StatisticsTask
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
	AddTasks(...StatisticsTask)
	AddRecords(...StatisticsRecord)
}
