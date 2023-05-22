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

func TestStr2Int(t *testing.T) {
	i, e := Str2Int("20")
	t.Log(i, e)
}

func TestStr2Int8(t *testing.T) {
	i, e := Str2Int8("20a")
	t.Log(i, e)
}
