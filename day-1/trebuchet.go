package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func getLineTotal(line string) (int, error) {
	stringDigits := map[string]string{
		"zero": "z0o", "one": "o1e", "two": "t2o", "three": "t3e",
		"four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n",
		"eight": "e8t", "nine": "n9e"}

	var digits []rune

	for strDigit, replacement := range stringDigits {
		line = strings.ReplaceAll(line, strDigit, replacement)
	}

	for _, val := range line {
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
