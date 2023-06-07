package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnyEq(t *testing.T) {
	assert.Equal(t, false, AnyEq([]int{1, 3}, []int{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int8{1, 3}, []int8{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int16{1, 3}, []int16{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int32{1, 3}, []int32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int64{1, 3}, []int64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint{1, 3}, []uint{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint8{1, 3}, []uint8{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint16{1, 3}, []uint16{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint32{1, 3}, []uint32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint64{1, 3}, []uint64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]float32{1, 3}, []float32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]float64{1, 3}, []float64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]string{"1", "3"}, []string{"1", "2", "3"}))
	assert.Equal(t, false, AnyEq([]bool{false}, []bool{true, false}))

	assert.Equal(t, false, AnyEq([]int{1, 3, 2}, []int{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int8{1, 3, 2}, []int8{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int16{1, 3, 2}, []int16{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int32{1, 3, 2}, []int32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]int64{1, 3, 2}, []int64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint{1, 3, 2}, []uint{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint8{1, 3, 2}, []uint8{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint16{1, 3, 2}, []uint16{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint32{1, 3, 2}, []uint32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]uint64{1, 3, 2}, []uint64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]float32{1, 3, 2}, []float32{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]float64{1, 3, 2}, []float64{1, 2, 3}))
	assert.Equal(t, false, AnyEq([]string{"1", "3", "2"}, []string{"1", "2", "3"}))
	assert.Equal(t, false, AnyEq([]bool{false, false}, []bool{true, false}))

	assert.Equal(t, true, AnyEq([]int{1, 2, 3}, []int{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]int8{1, 2, 3}, []int8{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]int16{1, 2, 3}, []int16{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]int32{1, 2, 3}, []int32{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]int64{1, 2, 3}, []int64{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]uint{1, 2, 3}, []uint{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]uint8{1, 2, 3}, []uint8{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]uint16{1, 2, 3}, []uint16{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]uint32{1, 2, 3}, []uint32{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]uint64{1, 2, 3}, []uint64{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]float32{1, 2, 3}, []float32{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]float64{1, 2, 3}, []float64{1, 2, 3}))
	assert.Equal(t, true, AnyEq([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
	assert.Equal(t, true, AnyEq([]bool{true, false}, []bool{true, false}))
}
