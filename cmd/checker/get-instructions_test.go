package main

import (
	"strings"
	"testing"
)

func TestGetInstructions(t *testing.T) {
	// Create a mock reader with some data
	reader := strings.NewReader("ra\nrra\npb\nrrr\npa\n")

	// Call getInstructions
	instructions, err := getInstructions(reader, false)

	// Check for errors
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Check that the instructions slice contains the expected values
	expected := []string{"ra", "rra", "pb", "rrr", "pa"}
	for i, instruction := range instructions {
		if instruction != expected[i] {
			t.Errorf("expected %q, got %q", expected[i], instruction)
		}
	}
}
