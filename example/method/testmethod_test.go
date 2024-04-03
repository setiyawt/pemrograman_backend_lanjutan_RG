package main

import (
	"testing"
)

func TestSumPositiveNumber(t *testing.T) {
	t.Run("Return error when a or b is negative", func(t *testing.T) {
		_, err := SumPositiveNumber(-1, 1)
		if err == nil {
			t.Error("Expected error, but got nil")
		}
	})

	t.Run("a and b is positive", func(t *testing.T) {
		result, err := SumPositiveNumber(1, 1)
		if err != nil {
			t.Error("Expected nil, but got error")
		}

		if result != 2 {
			t.Errorf("Expected 2, but got %d", result)
		}
	})
}
