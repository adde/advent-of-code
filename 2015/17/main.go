package main

import (
	"fmt"
	"math"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	containers := []int{}

	for _, line := range lines {
		containers = append(containers, u.ToInt(line))
	}

	combinations := findAllCombinations(containers, 150)

	fmt.Println("\nPart one:", partOne(combinations))
	fmt.Println("Part two:", partTwo(combinations))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(combinations map[int]int) int {
	totalCombinations := 0

	for _, count := range combinations {
		totalCombinations += count
	}

	return totalCombinations
}

func partTwo(combinations map[int]int) int {
	minContainers := math.MaxInt32

	for containers := range combinations {
		if containers < minContainers {
			minContainers = containers
		}
	}

	return combinations[minContainers]
}

// Find all possible combinations of containers that can hold the target amount
func findAllCombinations(containers []int, target int) map[int]int {
	combinations := make(map[int]int)

	var findCombo func(remaining int, startIdx int, used []bool, usedCount int)

	findCombo = func(remaining int, startIdx int, used []bool, usedCount int) {
		// Base case: found a valid combination
		if remaining == 0 {
			combinations[usedCount]++
			return
		}
		// Invalid cases
		if remaining < 0 || startIdx >= len(containers) {
			return
		}

		// Try combinations starting from startIdx to avoid duplicates
		for i := startIdx; i < len(containers); i++ {
			if !used[i] {
				used[i] = true
				findCombo(remaining-containers[i], i+1, used, usedCount+1)
				used[i] = false
			}
		}
	}

	used := make([]bool, len(containers))
	findCombo(target, 0, used, 0)
	return combinations
}
