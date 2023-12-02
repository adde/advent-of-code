package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	patternP1 := "^\\D*?(\\d)(?:.*(\\d))?\\D*$"
	regexP1 := regexp.MustCompile(patternP1)
	patternP2 := "^\\D*?(\\d|one|two|three|four|five|six|seven|eight|nine)(?:.*(\\d|one|two|three|four|five|six|seven|eight|nine))?\\D*$"
	regexP2 := regexp.MustCompile(patternP2)

	sumP1 := 0
	sumP2 := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numP1 := getCombinedNumbers(line, regexP1)
		sumP1 += numP1
		numP2 := getCombinedNumbers(line, regexP2)
		sumP2 += numP2
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func getCombinedNumbers(line string, regex *regexp.Regexp) int {
	matches := regex.FindStringSubmatch(line)
	matchOne := getNumber(matches[1])
	matchTwo := matchOne
	if matches[2] != "" {
		matchTwo = getNumber(matches[2])
	}

	num, err := strconv.Atoi(matchOne + matchTwo)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	return num
}

func getNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}
