package pkg

type Signal interface {
	SetSpider(Spider)
	RegisterSpiderOpened(func(Spider))
	RegisterSpiderClosed(func(Spider))
	SpiderOpened()
	SpiderClosed()
	FromCrawler(Crawler) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
