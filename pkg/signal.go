package pkg

type FnCrawlerStarted func(Crawler)
type FnCrawlerStopped func(Crawler)
type FnSpiderStarting func(Spider)
type FnSpiderStarted func(Spider)
type FnSpiderStopping func(Spider)
type FnSpiderStopped func(Spider)
type FnTaskStarted func(Context)
type FnTaskStopped func(Context)
type FnItemSaved func(ItemWithContext)
type FnScheduled func(Context)

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
	RegisterScheduled(FnScheduled)
	CrawlerStarted(Crawler)
	CrawlerStopped(Crawler)
	SpiderStarting(Spider)
	SpiderStarted(Spider)
	SpiderStopping(Spider)
	SpiderStopped(Spider)
	TaskStarted(Context)
	TaskStopped(Context)
	ItemSaved(ItemWithContext)
	Scheduled(Context)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
