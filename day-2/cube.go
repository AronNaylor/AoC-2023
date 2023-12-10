package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func colourGreaterThanLimit(colour []string, limit int) bool {
	countsRegex := regexp.MustCompile("(\\d+)")
	//fmt.Println(colour)

	for _, v := range colour {
		digitMatch := countsRegex.FindString(v)
		digit, _ := strconv.Atoi(digitMatch)
		//fmt.Println(digit)
		if digit > limit {
			return true
		}
	}
	return false
}

// 12 red cubes, 13 green cubes, and 14 blue cubes.
func cube(filepath string, blueLimit, redLimit, greenLimit int) int {
	validGamesTotal := 0
	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Error openning file %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	game := regexp.MustCompile("(\\w+\\s\\d+:)")
	gameNumber := regexp.MustCompile("\\s(\\d{1,2}):")
	blue := regexp.MustCompile("(\\d+\\sblue)")
	red := regexp.MustCompile("(\\d+\\sred)")
	green := regexp.MustCompile("(\\d+\\sgreen)")

	for scanner.Scan() {
		line := scanner.Text()
		game := game.FindString(line)
		gameNumber := gameNumber.FindString(game)
		blues := blue.FindAllString(line, -1)
		reds := red.FindAllString(line, -1)
		greens := green.FindAllString(line, -1)
		blueGreaterThanLimit := colourGreaterThanLimit(blues, blueLimit)
		redsGreaterThanLimit := colourGreaterThanLimit(reds, redLimit)
		greenGreaterThanLimit := colourGreaterThanLimit(greens, greenLimit)
		gameNumber = strings.TrimPrefix(gameNumber, " ")
		gameNumber = strings.TrimSuffix(gameNumber, ":")

		if !blueGreaterThanLimit && !redsGreaterThanLimit && !greenGreaterThanLimit {
			fmt.Printf("%s (ID: %s) = possible\n", game, gameNumber)
			validDigit, _ := strconv.Atoi(gameNumber)
			validGamesTotal += validDigit
		}
	}
	return validGamesTotal
}

func main() {
	res := cube("input.txt", 14, 12, 13)
	fmt.Println(res)
}
