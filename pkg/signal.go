package pkg

type FnCrawlerChanged func(Context)
type FnSpiderChanged func(Context)
type FnJobChanged func(Context)
type FnTaskChanged func(Context)
type FnRequestChanged func(Context)
type FnItemChanged func(Item)

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
	TaskChanged(Context)
	RequestChanged(Context)
	ItemChanged(Item)
	FromCrawler(crawler Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
