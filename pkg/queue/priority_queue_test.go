package queue

import (
	"fmt"
	"testing"
)

func TestNewPriorityQueue(t *testing.T) {
	items := map[string]int64{
		"b": 1, "banana": 3, "apple": 2, "pear": 4, "p": 5, "q": 6,
	}

	p := NewPriorityQueue(3)
	for value, priority := range items {
		item := NewItem(value, priority)
		p.Push(item)
	}

	item := NewItem("orange", 1)
	p.Push(item)
	if err := p.update(item, item.value, 5); err != nil {
		fmt.Println(err)
	}

	for p.Len() > 0 {
		it := p.PopItem()
		fmt.Printf("%.2d:%s ", it.Priority(), it.Value())
	}
}
