package pkg

import "time"

type StatisticsNode interface {
	WithStatusAndTime(status CrawlerStatus, t time.Time) StatisticsNode
	WithId(string) StatisticsNode
	IncSpider()
	DecSpider()
	IncJob()
	DecJob()
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncRecord()
	DecRecord()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	WithStatusAndTime(status SpiderStatus, t time.Time) StatisticsSpider
	WithId(id uint64) StatisticsSpider
	GetSpider() string
	WithSpider(spider string) StatisticsSpider
	WithFuncs(funcs []string) StatisticsSpider
	WithNode(node string) StatisticsSpider
	IncJob()
	DecJob()
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncRecord()
	DecRecord()
	GetLastTaskId() string
	WithLastTaskId(string) StatisticsSpider
	WithLastTaskStatus(TaskStatus) StatisticsSpider
	WithLastTaskStartTime(time.Time) StatisticsSpider
	WithLastTaskFinishTime(time.Time) StatisticsSpider
	Marshal() (bytes []byte, err error)
}
type StatisticsJob interface {
	WithStatusAndTime(status JobStatus, t time.Time) StatisticsJob
	WithId(id string) StatisticsJob
	WithSchedule(schedule string) StatisticsJob
	WithCommand(command string) StatisticsJob
	WithNode(node string) StatisticsJob
	WithSpider(spider string) StatisticsJob
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncRecord()
	DecRecord()
	WithEnable(enable bool) StatisticsJob
	WithStopReason(stopReason string) StatisticsJob
	Marshal() (bytes []byte, err error)
}
type StatisticsTask interface {
	WithStatus(status TaskStatus) StatisticsTask
	GetId() string
	WithId(id string) StatisticsTask
	IncRequest()
	DecRequest()
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
	WithUpdateTime(updateTime time.Time) StatisticsTask
	WithStopReason(stopReason string) StatisticsTask
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
