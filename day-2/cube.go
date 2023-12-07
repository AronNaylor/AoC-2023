package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// 12 red cubes, 13 green cubes, and 14 blue cubes.
func cube(filepath string) {
	// validGamesTotal := 0
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Error openning file %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	blue := regexp.MustCompile("(\\d\\sblue)")
	red := regexp.MustCompile("(\\d\\sred)")
	green := regexp.MustCompile("(\\d\\sgreen)")

	for scanner.Scan() {
		line := scanner.Text()
		match := blue.FindAllString(line, -1)
		fmt.Println(match)
	}

}

func main() {
	cube("test.txt")
}
