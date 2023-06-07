package utils

import (
	"sort"
)

type Sortable interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string
}

type DescSlice[T Sortable] []T

func (x DescSlice[T]) Len() int           { return len(x) }
func (x DescSlice[T]) Less(i, j int) bool { return x[i] > x[j] }
func (x DescSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x DescSlice[T]) Sort() { sort.Sort(x) }

type AscSlice[T Sortable] []T

func (x AscSlice[T]) Len() int           { return len(x) }
func (x AscSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x AscSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x AscSlice[T]) Sort() { sort.Sort(x) }

func DescSort[T Sortable](s []T) {
	DescSlice[T](s).Sort()
}

func AscSort[T Sortable](s []T) {
	AscSlice[T](s).Sort()
}
