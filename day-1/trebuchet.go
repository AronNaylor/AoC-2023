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
	stringDigits := map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3",
		"four": "4", "five": "5", "six": "6", "seven": "7",
		"eight": "8", "nine": "9"}

	stringDigitRegex := regexp.MustCompile(`(zero|one|two|three|four|five|six|seven|eight|nine)`)

	var digits []rune

	// Helper func to replace in string
	replaceInString := func(match string) string {
		if val, ok := stringDigits[match]; ok {
			return val
		}
		return match
	}

	transformedLine := stringDigitRegex.ReplaceAllStringFunc(line, replaceInString)

	for _, val := range transformedLine {
		isDigit := unicode.IsDigit(val)

		if isDigit {
			digits = append(digits, val)
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
