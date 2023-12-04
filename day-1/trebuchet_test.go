package main

// 1abc2
// pqr3stu8vwx
// a1b2c3d4e5f
// treb7uchet
// In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

//two1nine
//eightwothree
//abcone2threexyz
//xtwone3four
//4nineeightseven2
//zoneight234
//7pqrstsixteen
//In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

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
		{"Line 5", "two1nine", 29},
		{"Line 6", "eightwothree", 83},
		{"Line 7", "abcone2threexyz", 13},
		{"Line 8", "xtwone3four", 24},
		{"Line 9", "4nineeightseven2", 42},
		{"Line 10", "zoneight234", 14},
		{"Line 11", "7pqrstsixteen", 76},
		{"Line 12", "twonine4cjqln", 24},
		{"Line 13", "fivenhcvbntlcfthreemsktzr9two", 52},
		{"Line 14", "sevenfourfour99seven8", 78},
		{"Line 15", "3one6", 36},
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

	// Define test cases
	trebuchetTestCases := []struct {
		name  string
		given string
		want  int
	}{
		{"Part 1", "test.txt", 142},
		{"Part 2", "test2.txt", 281},
	}

	for _, tc := range trebuchetTestCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := trebuchet(tc.given)

			if err != nil {
				t.Errorf("Error recieved getting file total %e", err)
			}

			if got != tc.want {
				t.Errorf("Got %d, want %d", got, tc.want)
			}
		})
	}
}
