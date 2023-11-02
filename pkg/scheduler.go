package pkg

type SchedulerType string

const (
	SchedulerUnknown SchedulerType = ""
	SchedulerMemory  SchedulerType = "memory"
	SchedulerRedis   SchedulerType = "redis"
	SchedulerKafka   SchedulerType = "kafka"
)

type Scheduler interface {
	YieldItem(Context, Item) error
	Request(Context, Request) (Response, error)
	YieldRequest(Context, Request) error
	YieldExtra(Context, any) error
	GetExtra(Context, any) error
	StartScheduler(Context) error
	StopScheduler(Context) error
}
