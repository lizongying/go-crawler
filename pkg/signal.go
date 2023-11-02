package pkg

type FnCrawlerStarted func(Context)
type FnCrawlerStopped func(Context)
type FnSpiderStarting func(Context)
type FnSpiderStarted func(Context)
type FnSpiderStopping func(Context)
type FnSpiderStopped func(Context)
type FnScheduleStarted func(Context)
type FnScheduleStopped func(Context)
type FnTaskStarted func(Context)
type FnTaskStopped func(Context)
type FnRequestStarted func(Context)
type FnRequestStopped func(Context)
type FnItemStarted func(Context)
type FnItemStopped func(Item)

type Signal interface {
	RegisterCrawlerStarted(FnCrawlerStarted)
	RegisterCrawlerStopped(FnCrawlerStopped)
	RegisterSpiderStarting(FnSpiderStarting)
	RegisterSpiderStarted(FnSpiderStarted)
	RegisterSpiderStopping(FnSpiderStopping)
	RegisterSpiderStopped(FnSpiderStopped)
	RegisterScheduleStarted(FnScheduleStarted)
	RegisterScheduleStopped(FnScheduleStopped)
	RegisterTaskStarted(FnTaskStarted)
	RegisterTaskStopped(FnTaskStopped)
	RegisterRequestStarted(FnRequestStarted)
	RegisterRequestStopped(FnRequestStopped)
	RegisterItemStarted(FnItemStarted)
	RegisterItemStopped(FnItemStopped)
	CrawlerStarted(Context)
	CrawlerStopped(Context)
	SpiderStarting(Context)
	SpiderStarted(Context)
	SpiderStopping(Context)
	SpiderStopped(Context)
	ScheduleStarted(Context)
	ScheduleStopped(Context)
	TaskStarted(Context)
	TaskStopped(Context)
	RequestStarted(Context)
	RequestStopped(Context)
	ItemStarted(Context)
	ItemStopped(Item)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
