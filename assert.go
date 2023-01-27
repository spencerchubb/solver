package solver

import (
	"testing"
)

func assertEqual[T comparable](t *testing.T, expected, actual T) {
	if expected != actual {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func assertArraysEqual[T comparable](t *testing.T, expected, actual []T) {
	if len(expected) != len(actual) {
		t.Errorf("Expected length of %d, got length of %d", len(expected), len(actual))
	}

	for i := 0; i < len(expected); i++ {
		if expected[i] != actual[i] {
			t.Errorf("Expected %v, got %v", expected, actual)
			return
		}
	}
}
