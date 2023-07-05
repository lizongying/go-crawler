package pkg

type Spider interface {
	Crawler
	GetName() string
	SetName(string)
	GetHost() string
	SetHost(string)
	SetCallbacks(callbacks map[string]Callback)
	SetErrbacks(errbacks map[string]Errback)
}

type NewSpider func(Spider) (Spider, error)
