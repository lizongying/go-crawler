package utils

import (
	"github.com/google/uuid"
	"testing"
)

func TestStrToUUID(t *testing.T) {
	// Positive test case
	str := "f6d8d99c7a6811ee90869221bc92ca26"
	expectedUUID := uuid.MustParse(str)

	u, err := uuid.Parse(str)
	if err != nil {
		t.Errorf("Failed to parse UUID: %v", err)
	}

	if u != expectedUUID {
		t.Errorf("Expected UUID: %v, got: %v", expectedUUID, u)
	}

	// Negative test case
	invalidStr := "invalid-uuid"
	_, err = uuid.Parse(invalidStr)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestStrToUUID2(t *testing.T) {
	// Positive test case
	str := "564aab647a5d11ee90869221bc92ca26"
	expectedUUID := uuid.MustParse(str)

	u, err := uuid.Parse(str)
	if err != nil {
		t.Errorf("Failed to parse UUID: %v", err)
	}

	if u != expectedUUID {
		t.Errorf("Expected UUID: %v, got: %v", expectedUUID, u)
	}

	// Negative test case
	invalidStr := "invalid-uuid"
	_, err = uuid.Parse(invalidStr)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
