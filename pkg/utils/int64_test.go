package utils

import (
	"testing"
)

func TestUint64_MarshalJSON(t *testing.T) {
	var u *Uint64
	expected := `""`
	result, err := u.MarshalJSON()
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if string(result) != expected {
		t.Errorf("Expected %s but got %s", expected, string(result))
	}

	u = &Uint64{12345}
	expected = `"12345"`
	result, err = u.MarshalJSON()
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if string(result) != expected {
		t.Errorf("Expected %s but got %s", expected, string(result))
	}
}

func TestUint64_UnmarshalJSON(t *testing.T) {
	u := &Uint64{}
	err := u.UnmarshalJSON([]byte(`""`))
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	err = u.UnmarshalJSON([]byte(`"12345"`))
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	if u.u64 != 12345 {
		t.Errorf("Expected 12345 but got %d", u.u64)
	}

	err = u.UnmarshalJSON([]byte(`"abc"`))
	if err == nil {
		t.Errorf("Expected error but got nil")
	}
}

func uint64Ptr(u uint64) *uint64 {
	return &u
}
