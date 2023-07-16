package pkg

type Spider interface {
	Crawler
	GetName() string
	SetName(string) Spider
	GetHost() string
	SetHost(string) Spider
	SetCallBacks(callbacks map[string]CallBack)
	SetErrBacks(errbacks map[string]ErrBack)
}

type NewSpider func(Spider) (Spider, error)
