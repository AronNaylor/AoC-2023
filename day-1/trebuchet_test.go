package main

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
		{"Line 979", "jjhxddmg5mqxqbgfivextlcpnvtwothreetwonerzk", 51},
		{"Line 993", "fivebml9gjvtlfctwo", 52},
		{"Line 839", "126", 16},
		{"Line 811", "ffive5x", 55},
		{"Line 765", "rrkshvnsixknbxdfjhq4nineeightvkhqnr", 68},
		{"Line 631", "8m", 88},
		{"Random test", "three2fiveonexrllxsvfive", 35},
		{"Real line 1", "mxmkjvgsdzfhseightonetwoeight7", 87},
		{"Random test", "atwonea", 21},
		{"Random test", "eightninenineeight", 88},
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
