package utils

import (
	"testing"
)

func TestStruct2JsonKV(t *testing.T) {
	type A struct {
		B int
	}
	a := A{B: 1}

	k, v := Struct2JsonKV(a)
	t.Log(k, v)
}
