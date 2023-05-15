package pkg

import (
	"testing"
)

import "github.com/tidwall/gjson"

func TestGjson(t *testing.T) {
	json := `{"a":{"b":"b","c":"c"},"d":1}`
	value := gjson.Get(json, "a.b")
	t.Log(value.String())
	value2 := gjson.Get(json, "d")
	t.Log(value2.Int())
}
