package pkg

import (
	"testing"
)

func TestScheduleMode_String(t *testing.T) {
	tests := []struct {
		mode     ScheduleMode
		expected string
	}{
		{ScheduleModeOnce, "once"},
		{ScheduleModeLoop, "loop"},
		{ScheduleModeCron, "cron"},
		{ScheduleModeUnknown, "unknown"},
	}

	for _, test := range tests {
		result := test.mode.String()
		if result != test.expected {
			t.Errorf("Expected %s, but got %s", test.expected, result)
		}
	}
}

func TestScheduleMode_MarshalJSON(t *testing.T) {
	tests := []struct {
		mode     ScheduleMode
		expected string
	}{
		{ScheduleModeOnce, "1"},
		{ScheduleModeLoop, "2"},
		{ScheduleModeCron, "3"},
		{ScheduleModeUnknown, "0"},
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

func TestScheduleMode_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected ScheduleMode
	}{
		{"1", ScheduleModeOnce},
		{"2", ScheduleModeLoop},
		{"3", ScheduleModeCron},
		{"0", ScheduleModeUnknown},
	}

	for _, test := range tests {
		var mode ScheduleMode
		err := mode.UnmarshalJSON([]byte(test.input))
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if mode != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, mode)
		}
	}
}
