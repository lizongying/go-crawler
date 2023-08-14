package pkg

type SignalFn func()

type Signal interface {
	RegisterSpiderOpened(SignalFn)
	RegisterSpiderClosed(SignalFn)
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
