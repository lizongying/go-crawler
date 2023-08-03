package pkg

type Signal interface {
	RegisterSpiderOpened(func(Spider))
	RegisterSpiderClosed(func(Spider))
	SpiderOpened()
	SpiderClosed()
	FromSpider(Spider) Signal
}

type SignalType string

const (
	SignalUnknown      SignalType = ""
	SignalSpiderOpened SignalType = "opened"
	SignalSpiderClosed SignalType = "closed"
)
