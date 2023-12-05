package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

func getLineTotal(line string) (int, error) {
	stringDigits := map[string]rune{
		"zero": '0', "one": '1', "two": '2', "three": '3',
		"four": '4', "five": '5', "six": '6', "seven": '7',
		"eight": '8', "nine": '9'}

	stringDigitRegex := regexp.MustCompile(`zero|one|two|three|four|five|six|seven|eight|nine`)

	var digits []rune

	for idx, val := range line {
		isDigit := unicode.IsDigit(val)

		if isDigit {
			digits = append(digits, val)
		}

		// This is a little bit naive as it will lead to duplicate values
		// but order will be preserved which is important
		match := stringDigitRegex.FindString(line[idx:])
		if match != "" {
			digit := stringDigits[match]
			digits = append(digits, digit)
		}
	}

	lineSum := string(digits[0]) + string(digits[len(digits)-1])
	lineNum, err := strconv.Atoi(lineSum)

	if err != nil {
		fmt.Println("Error:", err)
	}
	return lineNum, nil
}

func trebuchet(filepath string) (int, error) {
	var total int

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSum, err := getLineTotal(line)
		if err != nil {
			log.Fatalf("Error received when getting line total %e", err)
		}
		total += lineSum
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scanning: %s", err)
	}

	return total, nil
}

func main() {
	res, err := trebuchet("./input.txt")

	if err != nil {
		log.Fatalf("Error getting file totals: %e", err)
	}

	fmt.Println(res)
}
