package pkg

type FnCrawlerChanged func(Context)
type FnSpiderChanged func(Context)
type FnJobChanged func(Context)
type FnTaskStarted func(Context)
type FnTaskStopped func(Context)
type FnRequestStarted func(Context)
type FnRequestStopped func(Context)
type FnItemStarted func(Context)
type FnItemStopped func(Item)

type Signal interface {
	RegisterCrawlerChanged(FnCrawlerChanged)
	RegisterSpiderChanged(FnSpiderChanged)
	RegisterJobChanged(FnJobChanged)
	RegisterTaskStarted(FnTaskStarted)
	RegisterTaskStopped(FnTaskStopped)
	RegisterRequestStarted(FnRequestStarted)
	RegisterRequestStopped(FnRequestStopped)
	RegisterItemStarted(FnItemStarted)
	RegisterItemStopped(FnItemStopped)
	CrawlerChanged(Context)
	SpiderChanged(Context)
	JobChanged(Context)
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
