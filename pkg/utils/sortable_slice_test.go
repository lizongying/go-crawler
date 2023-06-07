package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAscSlice_Sort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, func() []int {
		a := []int{1, 3, 2}
		AscSlice[int](a).Sort()
		return a
	}())
	assert.Equal(t, []int8{1, 2, 3}, func() []int8 {
		a := []int8{1, 3, 2}
		AscSlice[int8](a).Sort()
		return a
	}())
	assert.Equal(t, []int16{1, 2, 3}, func() []int16 {
		a := []int16{1, 3, 2}
		AscSlice[int16](a).Sort()
		return a
	}())
	assert.Equal(t, []int32{1, 2, 3}, func() []int32 {
		a := []int32{1, 3, 2}
		AscSlice[int32](a).Sort()
		return a
	}())
	assert.Equal(t, []int64{1, 2, 3}, func() []int64 {
		a := []int64{1, 3, 2}
		AscSlice[int64](a).Sort()
		return a
	}())
	assert.Equal(t, []uint{1, 2, 3}, func() []uint {
		a := []uint{1, 3, 2}
		AscSlice[uint](a).Sort()
		return a
	}())
	assert.Equal(t, []uint8{1, 2, 3}, func() []uint8 {
		a := []uint8{1, 3, 2}
		AscSlice[uint8](a).Sort()
		return a
	}())
	assert.Equal(t, []uint16{1, 2, 3}, func() []uint16 {
		a := []uint16{1, 3, 2}
		AscSlice[uint16](a).Sort()
		return a
	}())
	assert.Equal(t, []uint32{1, 2, 3}, func() []uint32 {
		a := []uint32{1, 3, 2}
		AscSlice[uint32](a).Sort()
		return a
	}())
	assert.Equal(t, []uint64{1, 2, 3}, func() []uint64 {
		a := []uint64{1, 3, 2}
		AscSlice[uint64](a).Sort()
		return a
	}())
	assert.Equal(t, []float32{1, 2, 3}, func() []float32 {
		a := []float32{1, 3, 2}
		AscSlice[float32](a).Sort()
		return a
	}())
	assert.Equal(t, []float64{1, 2.0, 2.1}, func() []float64 {
		a := []float64{1, 2.1, 2.0}
		AscSlice[float64](a).Sort()
		return a
	}())
}

func TestDescSlice_Sort(t *testing.T) {
	assert.Equal(t, []int{3, 2, 1}, func() []int {
		a := []int{1, 3, 2}
		DescSlice[int](a).Sort()
		return a
	}())
	assert.Equal(t, []int8{3, 2, 1}, func() []int8 {
		a := []int8{1, 3, 2}
		DescSlice[int8](a).Sort()
		return a
	}())
	assert.Equal(t, []int16{3, 2, 1}, func() []int16 {
		a := []int16{1, 3, 2}
		DescSlice[int16](a).Sort()
		return a
	}())
	assert.Equal(t, []int32{3, 2, 1}, func() []int32 {
		a := []int32{1, 3, 2}
		DescSlice[int32](a).Sort()
		return a
	}())
	assert.Equal(t, []int64{3, 2, 1}, func() []int64 {
		a := []int64{1, 3, 2}
		DescSlice[int64](a).Sort()
		return a
	}())
	assert.Equal(t, []uint{3, 2, 1}, func() []uint {
		a := []uint{1, 3, 2}
		DescSlice[uint](a).Sort()
		return a
	}())
	assert.Equal(t, []uint8{3, 2, 1}, func() []uint8 {
		a := []uint8{1, 3, 2}
		DescSlice[uint8](a).Sort()
		return a
	}())
	assert.Equal(t, []uint16{3, 2, 1}, func() []uint16 {
		a := []uint16{1, 3, 2}
		DescSlice[uint16](a).Sort()
		return a
	}())
	assert.Equal(t, []uint32{3, 2, 1}, func() []uint32 {
		a := []uint32{1, 3, 2}
		DescSlice[uint32](a).Sort()
		return a
	}())
	assert.Equal(t, []uint64{3, 2, 1}, func() []uint64 {
		a := []uint64{1, 3, 2}
		DescSlice[uint64](a).Sort()
		return a
	}())
	assert.Equal(t, []float32{3, 2, 1}, func() []float32 {
		a := []float32{1, 3, 2}
		DescSlice[float32](a).Sort()
		return a
	}())
	assert.Equal(t, []float64{2.1, 2.0, 1}, func() []float64 {
		a := []float64{1, 2.1, 2.0}
		DescSlice[float64](a).Sort()
		return a
	}())
}

func TestAscSort(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, func() []string {
		a := []string{"1", "3", "2"}
		AscSort(a)
		return a
	}())
}

func TestDescSort(t *testing.T) {
	assert.Equal(t, []string{"3", "2", "1"}, func() []string {
		a := []string{"1", "3", "2"}
		DescSort(a)
		return a
	}())
}
