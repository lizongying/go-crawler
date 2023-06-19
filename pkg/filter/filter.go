package filter

import (
	"github.com/lizongying/go-crawler/pkg"
	"sync"
)

type Filter struct {
	ids    sync.Map
	logger pkg.Logger
}

func (f *Filter) Exists(uniqueKey any) bool {
	_, ok := f.ids.Load(uniqueKey)
	return ok
}

func (f *Filter) ExistsOrStore(uniqueKey any) bool {
	_, ok := f.ids.LoadOrStore(uniqueKey, struct{}{})
	return ok
}

func (f *Filter) Clean() {
	f.ids.Range(func(key, _ any) bool {
		f.ids.Delete(key)
		return true
	})
}

func (f *Filter) FromCrawler(crawler pkg.Crawler) pkg.Filter {
	if f == nil {
		return new(Filter).FromCrawler(crawler)
	}

	f.logger = crawler.GetLogger()

	return f
}
