package pkg

type Spider interface {
	Crawler
	GetName() string
	SetName(string)
}

type NewSpider func(Spider) (Spider, error)
