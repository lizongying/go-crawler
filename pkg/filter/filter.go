package filter

import (
	"sync"
)

type Filter struct {
	ids sync.Map
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
