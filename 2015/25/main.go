package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

const (
	FIRST_CODE = 20151125
	MULTIPLIER = 252533
	DIVIDER    = 33554393
)

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	numbers := u.GetIntsFromString(file, true)

	fmt.Println("\nPart one:", findCode(numbers[0], numbers[1]))
	fmt.Println("Part two: -")
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func findCode(inputRow, inputCol int) int {
	code := FIRST_CODE
	row, col := 1, 1

	for {
		// Traverse tne grid diagonally,
		// from the bottom left to the top right
		if row == 1 {
			row = col + 1
			col = 1
		} else {
			row--
			col++
		}

		// Calculate the next code
		code = (code * MULTIPLIER) % DIVIDER

		if row == inputRow && col == inputCol {
			return code
		}
	}
}
