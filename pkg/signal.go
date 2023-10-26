package pkg

type FnCrawlerStarted func(Crawler)
type FnCrawlerStopped func(Crawler)
type FnSpiderStarting func(Spider)
type FnSpiderStarted func(Spider)
type FnSpiderStopping func(Spider)
type FnSpiderStopped func(Spider)
type FnTaskStarted func(Spider)
type FnTaskStopped func(Spider)
type FnItemSaved func(ItemWithContext)

type Signal interface {
	RegisterCrawlerStarted(FnCrawlerStarted)
	RegisterCrawlerStopped(FnCrawlerStopped)
	RegisterSpiderStarting(FnSpiderStarting)
	RegisterSpiderStarted(FnSpiderStarted)
	RegisterSpiderStopping(FnSpiderStopping)
	RegisterSpiderStopped(FnSpiderStopped)
	RegisterItemSaved(FnItemSaved)
	CrawlerStarted(Crawler)
	CrawlerStopped(Crawler)
	SpiderStarting(Spider)
	SpiderStarted(Spider)
	SpiderStopping(Spider)
	SpiderStopped(Spider)
	TaskStarted(Spider)
	TaskStopped(Spider)
	ItemSaved(ItemWithContext)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
