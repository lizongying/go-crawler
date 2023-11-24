package queue

import (
	"sync"
)

type GroupQueue struct {
	maxSize uint32
	data    map[string]*PriorityQueue
	mutex   sync.RWMutex
}

func NewGroupQueue(maxSize uint32) *GroupQueue {
	return &GroupQueue{
		maxSize: maxSize,
		data:    make(map[string]*PriorityQueue),
	}
}

func (q *GroupQueue) Enqueue(key string, value any, priority int64) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	v, ok := q.data[key]
	if !ok {
		q.data[key] = NewPriorityQueue(q.maxSize)
		v = q.data[key]
	}
	v.Push(NewItem(value, priority))
}

func (q *GroupQueue) Get(key string) (items []*Item) {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if key == "" {
		for _, v := range q.data {
			item, _ := v.GetItemN(-1)
			items = append(items, item...)
		}
		return
	}
	if v, ok := q.data[key]; ok {
		items, _ = v.GetItemN(-1)
	}
	return
}
func (q *GroupQueue) Size(key string) int {
	q.mutex.RLock()
	defer q.mutex.RUnlock()

	if key == "" {
		var l int
		for _, v := range q.data {
			l += v.Len()
		}
		return l
	}

	return q.data[key].Len()
}
