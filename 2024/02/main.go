package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := utils.ReadLines("input.txt")
	reports := make([][]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")
		report := make([]int, 0)

		for _, part := range parts {
			report = append(report, utils.ToInt(part))
		}

		reports = append(reports, report)
	}

	fmt.Println("\nPart one:", partOne(reports))
	fmt.Println("Part two:", partTwo(reports))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(reports [][]int) int {
	sum := 0

	for _, report := range reports {
		if isSafeAndSorted(report) {
			sum++
		}
	}

	return sum
}

func partTwo(reports [][]int) int {
	sum := 0

	for _, report := range reports {
		if isSafeWithDampener(report) {
			sum++
		}
	}

	return sum
}

// Checks if the report becomes safe after removing one element
func isSafeWithDampener(report []int) bool {
	for i := range report {
		reduced := make([]int, len(report)-1)
		copy(reduced[:i], report[:i])
		copy(reduced[i:], report[i+1:])

		if isSafeAndSorted(reduced) {
			return true
		}
	}

	return false
}

// Checks if adjacent elements are safe and the report is sorted
func isSafeAndSorted(report []int) bool {
	if len(report) <= 1 {
		return true
	}

	isAscending, isDescending := true, true
	for i := 1; i < len(report); i++ {
		if !isSafe(report[i-1], report[i]) {
			return false
		}

		if report[i-1] > report[i] {
			isAscending = false
		}
		if report[i-1] < report[i] {
			isDescending = false
		}
	}

	return isAscending || isDescending
}

// Checks if the difference between two numbers is between 1 and 3
func isSafe(a, b int) bool {
	diff := utils.Abs(a - b)
	return diff > 0 && diff <= 3
}
