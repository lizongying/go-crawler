package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimestamp_MarshalJSON(t *testing.T) {
	timestamp := Timestamp{time.Unix(1234567890, 0)}

	expected := []byte("1234567890")
	result, err := json.Marshal(&timestamp)
	if err != nil {
		t.Errorf("Error occurred during MarshalJSON: %v", err)
	}

	if string(result) != string(expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	input := []byte("1234567890")

	expected := Timestamp{time.Unix(1234567890, 0)}
	var result Timestamp

	err := json.Unmarshal(input, &result)
	if err != nil {
		t.Errorf("Error occurred during UnmarshalJSON: %v", err)
	}

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestTimestampNano_MarshalJSON(t *testing.T) {
	timestamp := TimestampNano{time.Now()}

	data, err := json.Marshal(&timestamp)
	if err != nil {
		t.Errorf("Error occurred while marshaling JSON: %v", err)
	}

	expected := fmt.Sprintf(`%d`, timestamp.UnixNano())
	if string(data) != expected {
		t.Errorf("Expected JSON data %s, but got %s", expected, string(data))
	}
}

func TestTimestampNano_UnmarshalJSON(t *testing.T) {
	timestamp := TimestampNano{}

	jsonData := []byte(`1234567890`)
	err := json.Unmarshal(jsonData, &timestamp)
	if err != nil {
		t.Errorf("Error occurred while unmarshaling JSON: %v", err)
	}

	expected := time.Unix(0, 1234567890)
	if !timestamp.Equal(expected) {
		t.Errorf("Expected timestamp %v, but got %v", expected, timestamp)
	}
}

func TestTimestampNano_UnmarshalJSON_InvalidData(t *testing.T) {
	timestamp := TimestampNano{}

	jsonData := []byte(`"invalid"`)
	err := json.Unmarshal(jsonData, &timestamp)
	if err == nil {
		t.Error("Expected error while unmarshaling invalid JSON data, but got nil")
	}
}

func TestTimestampNano_UnmarshalJSON_EmptyData(t *testing.T) {
	timestamp := TimestampNano{}

	jsonData := []byte(`""`)
	err := json.Unmarshal(jsonData, &timestamp)
	if err == nil {
		t.Error("Expected error while unmarshaling empty JSON data, but got nil")
	}
}

func TestDurationSecond_MarshalJSON(t *testing.T) {
	duration := time.Second * 10
	durationSecond := &DurationSecond{duration}

	expected := []byte("10")
	result, err := json.Marshal(durationSecond)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(result) != string(expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestDurationSecond_UnmarshalJSON(t *testing.T) {
	input := []byte("10")
	expected := time.Second * 10

	durationSecond := &DurationSecond{}
	err := json.Unmarshal(input, durationSecond)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if durationSecond.Duration != expected {
		t.Errorf("Expected %s, but got %s", expected, durationSecond.Duration)
	}
}

func TestDurationSecond_UnmarshalJSONInvalidInput(t *testing.T) {
	input := []byte("invalid")

	durationSecond := &DurationSecond{}
	err := json.Unmarshal(input, durationSecond)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestDurationNano_MarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		duration time.Duration
		expected string
	}{
		{
			name:     "positive test case",
			duration: time.Duration(1000000000),
			expected: "1000000000",
		},
		{
			name:     "negative test case",
			duration: time.Duration(-1000000000),
			expected: "-1000000000",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dn := DurationNano{Duration: tc.duration}
			b, err := dn.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, string(b))
		})
	}
}

func TestDurationNano_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected time.Duration
		err      error
	}{
		{
			name:     "positive test case",
			input:    "1000000000",
			expected: time.Duration(1000000000),
			err:      nil,
		},
		{
			name:     "negative test case",
			input:    "-1000000000",
			expected: time.Duration(-1000000000),
			err:      nil,
		},
		{
			name:     "invalid input",
			input:    "invalid",
			expected: 0,
			err:      errors.New("strconv.ParseInt: parsing \"invalid\": invalid syntax"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dn := DurationNano{}
			err := dn.UnmarshalJSON([]byte(tc.input))
			assert.Equal(t, tc.err.Error(), err.Error())
			assert.Equal(t, tc.expected, dn.Duration)
		})
	}
}
