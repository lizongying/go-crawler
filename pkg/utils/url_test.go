package utils

import (
	"net/url"
	"testing"
)

func TestURL_MarshalJSON(t *testing.T) {
	testURL := &Url{
		URL: &url.URL{
			Scheme: "https",
			Host:   "example.com",
			Path:   "/path",
		},
	}

	expectedJSON := []byte(`"https://example.com/path"`)
	json, err := testURL.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(json) != string(expectedJSON) {
		t.Errorf("Expected JSON: %s, but got: %s", expectedJSON, json)
	}
}

func TestURL_UnmarshalJSON(t *testing.T) {
	json := []byte(`"https://example.com/path"`)
	expectedURL := &Url{
		URL: &url.URL{
			Scheme: "https",
			Host:   "example.com",
			Path:   "/path",
		},
	}

	testURL := &Url{}
	err := testURL.UnmarshalJSON(json)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if testURL.String() != expectedURL.String() {
		t.Errorf("Expected URL: %s, but got: %s", expectedURL.String(), testURL.String())
	}
}

func TestURL_UnmarshalJSON_InvalidURL(t *testing.T) {
	json := []byte(`"invalid-url"`)
	testURL := &Url{}

	err := testURL.UnmarshalJSON(json)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}
