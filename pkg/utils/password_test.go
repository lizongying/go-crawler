package utils

import (
	"fmt"
	"github.com/lizongying/go-gua64/gua64"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	// Positive test case
	password := "znU2LtswYWW8kbf5"
	hashedPassword, err := GeneratePassword(password)
	fmt.Println(hashedPassword)
	g := gua64.NewGua64()
	fmt.Println(string(g.Decode(hashedPassword)))
	if err != nil {
		t.Errorf("Error generating password hash: %v", err)
	}
	if len(hashedPassword) == 0 {
		t.Errorf("Generated password hash is empty")
	}

	// Negative test case
	password = ""
	hashedPassword, err = GeneratePassword(password)
	if err == nil {
		t.Errorf("Expected error for empty password")
	}
}

func TestComparePassword(t *testing.T) {
	// Positive test case
	password := "password123"
	hashedPassword, _ := GeneratePassword(password)
	match := ComparePassword(password, hashedPassword)
	if !match {
		t.Errorf("Password and hashed password do not match")
	}

	// Negative test case
	password = "password123"
	invalidHashedPassword := "invalidhashedpassword"
	match = ComparePassword(password, invalidHashedPassword)
	if match {
		t.Errorf("Expected password and invalid hashed password to not match")
	}
}
