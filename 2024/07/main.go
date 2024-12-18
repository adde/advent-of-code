package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Equation struct {
	testValue int
	numbers   []int
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	equations := []Equation{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")

		equations = append(equations, Equation{
			testValue: u.ToInt(parts[0]),
			numbers:   u.StringsToInts(strings.Split(parts[1], " ")),
		})
	}

	fmt.Println("\nPart one:", partOne(equations))
	fmt.Println("Part two:", partTwo(equations))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(equations []Equation) int {
	sum := 0

	for _, eq := range equations {
		if isTargetPossible(eq.numbers, eq.testValue, []string{"+", "*"}) {
			sum += eq.testValue
		}
	}

	return sum
}

func partTwo(equations []Equation) int {
	sum := 0

	for _, eq := range equations {
		if isTargetPossible(eq.numbers, eq.testValue, []string{"+", "*", "||"}) {
			sum += eq.testValue
		}
	}

	return sum
}

// Check if any combination of operators can produce the target
func isTargetPossible(numbers []int, target int, operators []string) bool {
	return getPossibleOpCombos(numbers, target, operators, 0, numbers[0], 1)
}

// Explore all possible operator combinations recursively
func getPossibleOpCombos(numbers []int, target int, operators []string, operatorIndex int, currentValue int, numberIndex int) bool {
	// If we've used all numbers, check if result matches target
	if numberIndex == len(numbers) {
		return currentValue == target
	}

	// Try all possible operators
	for _, op := range operators {
		var nextValue int

		switch op {
		case "+":
			nextValue = currentValue + numbers[numberIndex]
		case "*":
			nextValue = currentValue * numbers[numberIndex]
		case "||":
			// String concatenate and convert back to int
			nextValue = u.ToInt(fmt.Sprintf("%d%d", currentValue, numbers[numberIndex]))
		}

		// Recursively try the next number with this combination
		if getPossibleOpCombos(numbers, target, operators, operatorIndex+1, nextValue, numberIndex+1) {
			return true
		}
	}

	return false
}
