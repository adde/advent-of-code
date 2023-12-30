package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	sumP1 := 0
	sumP2 := 0

	numbersMap := map[int][][]int{}
	specialsMap := map[int][][]int{}

	reN := regexp.MustCompile(`\d+`)
	reS := regexp.MustCompile(`[^a-zA-Z0-9.\n]`)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		line := scanner.Text()

		matchesNumbers := reN.FindAllStringIndex(line, -1)
		matchesSpecial := reS.FindAllStringIndex(line, -1)

		for _, match := range matchesSpecial {
			start, end := match[0], match[1]
			special := line[start:end]
			isGear := 0
			if special == "*" {
				isGear = 1
			}
			specialsMap[counter] = append(specialsMap[counter], []int{start - 1, end, isGear})
		}

		for _, match := range matchesNumbers {
			start, end := match[0], match[1]
			number := line[start:end]
			numbersMap[counter] = append(numbersMap[counter], []int{start, end, getInt(number)})
		}

		counter++
	}

	// Get all engine parts
	for lineNumber, line := range numbersMap {
		for _, number := range line {
			specialLines := []int{lineNumber - 1, lineNumber, lineNumber + 1}

		outer:
			for _, specialLine := range specialLines {
				if specialsMap[specialLine] != nil {
					for _, special := range specialsMap[specialLine] {
						if isRangeOverlap(number[0], number[1], special[0], special[1]) {
							sumP1 += number[2]
							break outer
						}
					}
				}
			}
		}
	}

	// Get gear ratio
	for lineNumber, line := range specialsMap {
		for _, special := range line {
			numberLines := []int{lineNumber - 1, lineNumber, lineNumber + 1}

			// Check if special character is a gear
			if special[2] == 1 {
				matches := []int{}
				for _, numberLine := range numberLines {
					if numbersMap[numberLine] != nil {
						for _, number := range numbersMap[numberLine] {
							if isRangeOverlap(number[0], number[1], special[0], special[1]) {
								matches = append(matches, number[2])
							}
						}
					}
				}

				if len(matches) == 2 {
					sumP2 += matches[0] * matches[1]
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

// Convert string to int
func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// Check overlapping ranges
func isRangeOverlap(start1, end1, start2, end2 int) bool {
	return start1 <= end2 && end1 > start2 || start2 < end1 && end2 >= start1
}
