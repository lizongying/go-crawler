package pkg

import (
	"testing"
)

func TestJobMode_String(t *testing.T) {
	tests := []struct {
		mode     JobMode
		expected string
	}{
		{JobModeOnce, "once"},
		{JobModeLoop, "loop"},
		{JobModeCron, "cron"},
		{JobModeUnknown, "unknown"},
	}

	for _, test := range tests {
		result := test.mode.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestJobMode_MarshalJSON(t *testing.T) {
	tests := []struct {
		mode     JobMode
		expected string
	}{
		{JobModeOnce, "1"},
		{JobModeLoop, "2"},
		{JobModeCron, "3"},
		{JobModeUnknown, "0"},
	}

	for _, test := range tests {
		result, err := test.mode.MarshalJSON()
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if string(result) != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, string(result))
		}
	}
}

func TestJobMode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected JobMode
	}{
		{"1", JobModeOnce},
		{"2", JobModeLoop},
		{"3", JobModeCron},
		{"0", JobModeUnknown},
	}

	for _, test := range tests {
		var mode JobMode
		err := mode.UnmarshalJSON([]byte(test.input))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if mode != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, mode)
		}
	}
}
