package middlewares

import (
	"reflect"
	"testing"
)

func TestZero(t *testing.T) {
	a := 0
	if reflect.ValueOf(any(a)).IsZero() {
		t.Log("zero", a)
	} else {
		t.Log("nzero", a)
	}
	a = 1
	if reflect.ValueOf(any(a)).IsZero() {
		t.Log("zero", a)
	} else {
		t.Log("nzero", a)
	}

	b := ""
	if reflect.ValueOf(any(b)).IsZero() {
		t.Log("zero", b)
	} else {
		t.Log("nzero", b)
	}
	b = "1"
	if reflect.ValueOf(any(b)).IsZero() {
		t.Log("zero", b)
	} else {
		t.Log("nzero", b)
	}
}
