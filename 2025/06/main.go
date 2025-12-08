package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	fmt.Println("\nPart one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(lines []string) int {
	grandTotal := 0
	worksheet := make([][]int, 0)
	ops := lines[len(lines)-1]
	ops = strings.ReplaceAll(ops, " ", "")

	// Build worksheet column-wise
	for _, line := range lines {
		numbers := u.GetIntsFromString(line, true)

		for i, num := range numbers {
			if len(worksheet) <= i {
				worksheet = append(worksheet, []int{})
			}

			worksheet[i] = append(worksheet[i], num)
		}
	}

	// Calculate totals per column
	for i, col := range worksheet {
		op := ops[i]
		total := 0
		if op == '*' {
			total = 1
		}

		for _, num := range col {
			if op == '+' {
				total += num
			} else {
				total *= num
			}
		}

		grandTotal += total
	}

	return grandTotal
}

func partTwo(lines []string) int {
	total, grandTotal := 0, 0
	lineLen := len(lines[0])
	numbers := []int{}
	op := '+'

	// Read columns backwards
	for i := lineLen - 1; i >= 0; i-- {
		numStr := ""

		// Read each line for this column
		for _, line := range lines {
			if line[i] == ' ' {
				continue
			}

			if line[i] == '+' || line[i] == '*' {
				op = rune(line[i])
				continue
			}

			numStr += string(line[i])
		}

		// If no number found or at the end of the column, calculate total
		if numStr == "" || i == 0 {
			if i == 0 {
				num := u.ToInt(numStr)
				numbers = append(numbers, num)
			}

			if op == '+' {
				for _, n := range numbers {
					total += n
				}
			} else {
				if total == 0 {
					total = 1
				}
				for _, n := range numbers {
					total *= n
				}
			}

			grandTotal += total

			// Reset for next column
			numbers = []int{}
			total = 0

			continue
		}

		// Add number to the current column's list
		num := u.ToInt(numStr)
		numbers = append(numbers, num)
	}

	return grandTotal
}
