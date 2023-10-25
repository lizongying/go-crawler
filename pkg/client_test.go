package pkg

import (
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	c := ClientUnknown
	expected := "\"\""

	data, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON returned an error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("MarshalJSON returned %s, expected %s", string(data), expected)
	}
}

func TestMarshalJSON_GO(t *testing.T) {
	c := ClientGo
	expected := "\"go\""

	data, err := c.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON returned an error: %v", err)
	}

	if string(data) != expected {
		t.Errorf("MarshalJSON returned %s, expected %s", string(data), expected)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	c := ClientUnknown
	data := []byte("\"go\"")

	err := c.UnmarshalJSON(data)
	if err != nil {
		t.Errorf("UnmarshalJSON returned an error: %v", err)
	}

	if c != ClientGo {
		t.Errorf("UnmarshalJSON set the value to %s, expected %s", c, ClientGo)
	}
}

func TestUnmarshalJSON_InvalidValue(t *testing.T) {
	c := ClientUnknown
	data := []byte("\"invalid\"")

	err := c.UnmarshalJSON(data)
	if err == nil {
		t.Error("UnmarshalJSON did not return an error for invalid value")
	}

	if c != ClientUnknown {
		t.Errorf("UnmarshalJSON set the value to %s, expected %s", c, ClientUnknown)
	}
}
