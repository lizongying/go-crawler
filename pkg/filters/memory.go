package filters

import (
	"github.com/lizongying/go-crawler/pkg"
	"sync"
)

type MemoryFilter struct {
	ids    sync.Map
	logger pkg.Logger
}

func (f *MemoryFilter) IsExist(_ pkg.Context, uniqueKey any) (ok bool, err error) {
	_, ok = f.ids.Load(uniqueKey)
	return
}

func (f *MemoryFilter) Store(_ pkg.Context, uniqueKey any) (err error) {
	f.ids.Store(uniqueKey, struct{}{})
	return
}

func (f *MemoryFilter) Clean(_ pkg.Context) (err error) {
	f.ids.Range(func(key, _ any) bool {
		f.ids.Delete(key)
		return true
	})
	return
}

func (f *MemoryFilter) FromSpider(spider pkg.Spider) (filter pkg.Filter, err error) {
	if f == nil {
		return new(MemoryFilter).FromSpider(spider)
	}

	f.logger = spider.GetLogger()

	return f, nil
}
