package queue

import (
	"container/heap"
	"errors"
)

type Item struct {
	value    any
	priority int64
	index    int // for update
}

func (i *Item) Value() any {
	return i.value
}
func (i *Item) Priority() int64 {
	return i.priority
}
func NewItem(value any, priority int64) *Item {
	return &Item{
		value:    value,
		priority: priority,
		index:    -1,
	}
}

type PriorityQueue struct {
	data    []*Item
	maxSize uint32
}

func (p *PriorityQueue) Len() int { return len(p.data) }
func (p *PriorityQueue) Less(i, j int) bool {
	return p.data[i].priority < p.data[j].priority
}
func (p *PriorityQueue) Swap(i, j int) {
	p.data[i], p.data[j] = p.data[j], p.data[i]
	p.data[i].index = i
	p.data[j].index = j
}
func (p *PriorityQueue) Push(x any) {
	item := x.(*Item)

	if uint32(len(p.data)) < p.maxSize {
		p.data = append(p.data, item)
		idx := len(p.data) - 1
		item.index = idx
		heap.Fix(p, idx)
	} else if item.priority > p.data[0].priority {
		p.data[0] = item
		item.index = 0
		heap.Fix(p, 0)
	}
}
func (p *PriorityQueue) Pop() any {
	if len(p.data) == 0 {
		panic("Heap is empty")
	}

	item := p.data[0]
	lastIdx := len(p.data) - 1
	p.data[0], p.data[lastIdx] = p.data[lastIdx], p.data[0]
	p.data = p.data[:lastIdx]
	heap.Fix(p, 0)

	return item
}
func (p *PriorityQueue) PopItem() *Item {
	return p.Pop().(*Item)
}
func (p *PriorityQueue) GetItemN(n int) (items []*Item, err error) {
	if n > p.Len() {
		err = errors.New("too much")
		return
	}
	if n == -1 {
		items = p.data
		return
	}
	items = p.data[:n]
	return
}
func (p *PriorityQueue) update(item *Item, value any, priority int64) (err error) {
	if item.index == -1 {
		err = errors.New("not exists")
		return
	}

	item.value = value
	item.priority = priority
	heap.Fix(p, item.index)
	return
}

func NewPriorityQueue(maxSize uint32) *PriorityQueue {
	return &PriorityQueue{
		data:    make([]*Item, 0),
		maxSize: maxSize,
	}
}
