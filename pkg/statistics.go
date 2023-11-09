package pkg

import "time"

type StatisticsNode interface {
	WithId(string) StatisticsNode
	IncSpider()
	DecSpider()
	IncJob()
	DecJob()
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	GetSpider() string
	IncJob()
	DecJob()
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
type StatisticsJob interface {
	WithStatusAndTime(status JobStatus, t time.Time) StatisticsJob
	WithStatus(status JobStatus) StatisticsJob
	WithId(id string) StatisticsJob
	WithSchedule(schedule string) StatisticsJob
	WithCommand(command string) StatisticsJob
	WithNode(node string) StatisticsJob
	WithSpider(spider string) StatisticsJob
	IncTask()
	DecTask()
	IncRecord()
	DecRecord()
	WithEnable(enable bool) StatisticsJob
	WithStartTime(time.Time) StatisticsJob
	WithFinishTime(time.Time) StatisticsJob
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
	GetJob() string
	WithJob(string) StatisticsTask
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
	GetJobs() []StatisticsJob
	GetTasks() []StatisticsTask
	GetRecords() []StatisticsRecord
}
