package pkg

type FnCrawlerOpened func()
type FnCrawlerClosed func()
type FnSpiderStarting func(Spider)
type FnSpiderStarted func(Spider)
type FnSpiderStopping func(Spider)
type FnSpiderStopped func(Spider)
type FnTaskStarted func(Spider)
type FnTaskStopped func(Spider)
type FnItemSaved func(ItemWithContext)

type Signal interface {
	RegisterCrawlerOpened(FnCrawlerOpened)
	RegisterCrawlerClosed(FnCrawlerClosed)
	RegisterSpiderStarting(string, FnSpiderStarting)
	RegisterSpiderStarted(string, FnSpiderStarted)
	RegisterSpiderStopping(string, FnSpiderStopping)
	RegisterSpiderStopped(string, FnSpiderStopped)
	RegisterItemSaved(FnItemSaved)
	CrawlerOpened()
	CrawlerClosed()
	SpiderStarting(spider Spider)
	SpiderStarted(spider Spider)
	SpiderStopping(spider Spider)
	SpiderStopped(spider Spider)
	ItemSaved(itemWithContext ItemWithContext)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
