package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getLineTotal(line string) (int, error) {
	var digits []rune

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
