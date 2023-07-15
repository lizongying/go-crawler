package pkg

type Spider interface {
	Crawler
	GetName() string
	SetName(string) Spider
	GetHost() string
	SetHost(string) Spider
	SetCallbacks(callbacks map[string]Callback)
	SetErrbacks(errbacks map[string]Errback)
}

type NewSpider func(Spider) (Spider, error)
