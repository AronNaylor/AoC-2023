package main

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

import (
	"testing"
)

func TestTrebuchet(t *testing.T) {

	// Define test cases
	lineTestCases := []struct {
		name  string
		given string
		want  int
	}{
		{"Line 1", "1abc2", 12},
		{"Line 2", "pqr3stu8vwx", 38},
		{"Line 3", "a1b2c3d4e5f", 15},
		{"Line 4", "treb7uchet", 77},
	}

	for _, tc := range lineTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := getLineTotal(tc.given)

			if err != nil {
				t.Errorf("Received an error: %s", err)
			}

			if got != tc.want {
				t.Errorf("Got %d, want %d", got, tc.want)
			}
		})
	}

	t.Run("Gets file total", func(t *testing.T) {
		want := 142
		got, err := trebuchet("test.txt")

		if err != nil {
			t.Errorf("Error recieved getting file total %e", err)
		}

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})
}
