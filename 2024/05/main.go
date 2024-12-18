package main

import (
	"fmt"
	"slices"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	rules := make(map[[2]int]int)
	updates := make([][]int, 0)
	addingRules := true

	for _, line := range lines {
		if line == "" {
			addingRules = false
			continue
		}

		if addingRules {
			parts := strings.Split(line, "|")
			rules[[2]int{u.ToInt(parts[0]), u.ToInt(parts[1])}] = 0
		} else {
			parts := strings.Split(line, ",")
			updates = append(updates, u.StringsToInts(parts))
		}
	}

	fmt.Println("\nPart one:", partOne(rules, updates))
	fmt.Println("Part two:", partTwo(rules, updates))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(rules map[[2]int]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		if isCorrectOrder(update, rules) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func partTwo(rules map[[2]int]int, updates [][]int) int {
	sum := 0

	for _, update := range updates {
		// Find the incorrect ordered updates
		if ncUpdate := getIncorrectUpdate(update, rules); ncUpdate != nil {
			// Sort the update using the rules map
			slices.SortFunc(ncUpdate, func(a, b int) int {
				if _, ok := rules[[2]int{a, b}]; ok {
					return 1
				}
				return -1
			})

			sum += ncUpdate[len(ncUpdate)/2]
		}
	}

	return sum
}

func isCorrectOrder(update []int, rules map[[2]int]int) bool {
	for i := 1; i < len(update); i++ {
		if _, ok := rules[[2]int{update[i], update[i-1]}]; ok {
			return false
		}
	}

	return true
}

func getIncorrectUpdate(update []int, rules map[[2]int]int) []int {
	for i := 1; i < len(update); i++ {
		if _, ok := rules[[2]int{update[i], update[i-1]}]; ok {
			return slices.Clone(update)
		}
	}

	return nil
}
