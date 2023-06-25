package filter

import (
	"context"
	"github.com/lizongying/go-crawler/pkg"
	"sync"
)

type MemoryFilter struct {
	ids    sync.Map
	logger pkg.Logger
}

func (f *MemoryFilter) IsExist(_ context.Context, uniqueKey any) (ok bool, err error) {
	_, ok = f.ids.Load(uniqueKey)
	return
}

func (f *MemoryFilter) Store(_ context.Context, uniqueKey any) (err error) {
	f.ids.Store(uniqueKey, struct{}{})
	return
}

func (f *MemoryFilter) Clean(_ context.Context) (err error) {
	f.ids.Range(func(key, _ any) bool {
		f.ids.Delete(key)
		return true
	})
	return
}

func (f *MemoryFilter) FromCrawler(crawler pkg.Crawler) pkg.Filter {
	if f == nil {
		return new(MemoryFilter).FromCrawler(crawler)
	}

	f.logger = crawler.GetLogger()

	return f
}
