package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max([]int{1, 2, 3}...))
	assert.Equal(t, int8(3), Max([]int8{1, 2, 3}...))
	assert.Equal(t, int16(3), Max([]int16{1, 2, 3}...))
	assert.Equal(t, int32(3), Max([]int32{1, 2, 3}...))
	assert.Equal(t, int64(3), Max([]int64{1, 2, 3}...))
	assert.Equal(t, uint(3), Max([]uint{1, 2, 3}...))
	assert.Equal(t, uint8(3), Max([]uint8{1, 2, 3}...))
	assert.Equal(t, uint16(3), Max([]uint16{1, 2, 3}...))
	assert.Equal(t, uint32(3), Max([]uint32{1, 2, 3}...))
	assert.Equal(t, uint64(3), Max([]uint64{1, 2, 3}...))
	assert.Equal(t, float32(3), Max([]float32{1, 2, 3}...))
	assert.Equal(t, float64(3), Max([]float64{1, 2, 3}...))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min([]int{1, 2, 3}...))
	assert.Equal(t, int8(1), Min([]int8{1, 2, 3}...))
	assert.Equal(t, int16(1), Min([]int16{1, 2, 3}...))
	assert.Equal(t, int32(1), Min([]int32{1, 2, 3}...))
	assert.Equal(t, int64(1), Min([]int64{1, 2, 3}...))
	assert.Equal(t, uint(1), Min([]uint{1, 2, 3}...))
	assert.Equal(t, uint8(1), Min([]uint8{1, 2, 3}...))
	assert.Equal(t, uint16(1), Min([]uint16{1, 2, 3}...))
	assert.Equal(t, uint32(1), Min([]uint32{1, 2, 3}...))
	assert.Equal(t, uint64(1), Min([]uint64{1, 2, 3}...))
	assert.Equal(t, float32(1), Min([]float32{1, 2, 3}...))
	assert.Equal(t, float64(1), Min([]float64{1, 2, 3}...))
}
