package pkg

import "time"

type StatisticsCrawler interface {
	WithStatusAndTime(status CrawlerStatus, t time.Time) StatisticsCrawler
	WithId(string) StatisticsCrawler
	IncSpider()
	DecSpider()
	IncJob()
	DecJob()
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncItem()
	DecItem()
	Marshal() (bytes []byte, err error)
}
type StatisticsSpider interface {
	WithStatusAndTime(status SpiderStatus, t time.Time) StatisticsSpider
	WithId(id uint64) StatisticsSpider
	GetSpider() string
	WithSpider(spider string) StatisticsSpider
	WithFuncs(funcs []string) StatisticsSpider
	WithCrawler(crawler string) StatisticsSpider
	IncJob()
	DecJob()
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncItem()
	DecItem()
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
	WithCrawler(crawler string) StatisticsJob
	WithSpider(spider string) StatisticsJob
	IncTask()
	DecTask()
	IncRequest()
	DecRequest()
	IncItem()
	DecItem()
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
	IncItem()
	DecItem()
	GetCrawler() string
	WithCrawler(string) StatisticsTask
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
type StatisticsRequest interface {
	Marshal() (bytes []byte, err error)
}
type StatisticsItem interface {
	Marshal() (bytes []byte, err error)
}
type Statistics interface {
	GetCrawlers() []StatisticsCrawler
	GetSpiders() []StatisticsSpider
	GetJobs() []StatisticsJob
	GetTasks() []StatisticsTask
	GetRequests() []StatisticsRequest
	GetItems() []StatisticsItem
}
