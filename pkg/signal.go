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
type FnItemSaved func(ItemWithContext)

type Signal interface {
	RegisterCrawlerStarted(FnCrawlerStarted)
	RegisterCrawlerStopped(FnCrawlerStopped)
	RegisterSpiderStarting(FnSpiderStarting)
	RegisterSpiderStarted(FnSpiderStarted)
	RegisterSpiderStopping(FnSpiderStopping)
	RegisterSpiderStopped(FnSpiderStopped)
	RegisterTaskStarted(FnTaskStarted)
	RegisterTaskStopped(FnTaskStopped)
	RegisterItemSaved(FnItemSaved)
	RegisterScheduleStarted(FnScheduleStarted)
	RegisterScheduleStopped(FnScheduleStopped)
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
	ItemSaved(ItemWithContext)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
