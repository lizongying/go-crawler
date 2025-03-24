package pkg

type FnCrawlerChanged func(Context) error
type FnSpiderChanged func(Context) error
type FnJobChanged func(Context) error
type FnTaskChanged func(Task) error
type FnRequestChanged func(Request) error
type FnItemChanged func(Item) error

type Signal interface {
	RegisterCrawlerChanged(FnCrawlerChanged)
	RegisterSpiderChanged(FnSpiderChanged)
	RegisterJobChanged(FnJobChanged)
	RegisterTaskChanged(FnTaskChanged)
	RegisterRequestChanged(FnRequestChanged)
	RegisterItemChanged(FnItemChanged)
	CrawlerChanged(Context)
	SpiderChanged(Context)
	JobChanged(Context)
	TaskChanged(Task)
	RequestChanged(Request)
	ItemChanged(Item)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
