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

	pattern := "^\\D*?(\\d|one|two|three|four|five|six|seven|eight|nine)(?:.*(\\d|one|two|three|four|five|six|seven|eight|nine))?\\D*$"
	regex := regexp.MustCompile(pattern)

	sum := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regex.FindStringSubmatch(line)
		matchOne := getNumber(matches[1])
		matchTwo := matchOne
		if matches[2] != "" {
			matchTwo = getNumber(matches[2])
		}

		num, err := strconv.Atoi(matchOne + matchTwo)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		sum += num
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Elapsed time:", time.Now().Sub(startTime))
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
