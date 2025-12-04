package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	fmt.Println("\nPart one:", getJoltage(lines, 2))
	fmt.Println("Part two:", getJoltage(lines, 12))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func getJoltage(lines []string, numberOfDigits int) int {
	joltage := 0

	for _, line := range lines {
		digits := []int{}
		remainingLen, currentIndex, number := numberOfDigits, 0, 0

		// Pick the largest digits in order
		for index := 0; index < numberOfDigits; index++ {
			maxDigit, maxIndex := findMaxDigit(line, currentIndex, remainingLen)
			digits = append(digits, maxDigit)
			currentIndex = maxIndex + 1
			remainingLen--
		}

		// Build the number from the digits
		for _, digit := range digits {
			number = number*10 + digit
		}

		joltage += number
	}

	return joltage
}

func findMaxDigit(line string, startIndex int, remainingLen int) (int, int) {
	maxDigit, maxIndex := -1, -1

	for index := startIndex; index <= len(line)-remainingLen; index++ {
		digit := int(line[index] - '0')

		if digit > maxDigit {
			maxDigit = digit
			maxIndex = index
		}
	}

	return maxDigit, maxIndex
}
