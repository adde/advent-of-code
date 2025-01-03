package main

import (
	"fmt"
	"sort"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type PackageGroup struct {
	packages []int
	qe       int64
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	weights := make([]int, 0)

	for _, line := range lines {
		weights = append(weights, u.ToInt(line))
	}

	fmt.Println("\nPart one:", findIdealConfiguration(weights, 3))
	fmt.Println("Part two:", findIdealConfiguration(weights, 4))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func findIdealConfiguration(weights []int, numGroups int) int64 {
	total := 0
	for _, w := range weights {
		total += w
	}

	if total%numGroups != 0 {
		panic("Total weight must be divisible by number of groups")
	}

	targetWeight := total / numGroups

	// Sort weights in descending order to optimize search
	sort.Sort(sort.Reverse(sort.IntSlice(weights)))

	// Try different group sizes starting from smallest possible
	for size := 1; size <= len(weights); size++ {
		var validGroups []PackageGroup
		findValidGroups(weights, size, targetWeight, 0, []int{}, &validGroups)

		if len(validGroups) > 0 {
			// Sort groups by quantum entanglement
			sort.Slice(validGroups, func(i, j int) bool {
				return validGroups[i].qe < validGroups[j].qe
			})

			// For each valid first group, verify we can split remaining packages
			for _, group := range validGroups {
				remaining := removePackages(weights, group.packages)
				if canSplitIntoGroups(remaining, numGroups-1, targetWeight) {
					return group.qe // Found valid configuration
				}
			}
		}
	}

	panic("No valid configuration found")
}

func findValidGroups(weights []int, size, target, start int, current []int, results *[]PackageGroup) {
	sum := 0
	for _, w := range current {
		sum += w
	}

	if sum > target {
		return
	}

	if len(current) == size {
		if sum == target {
			group := PackageGroup{
				packages: make([]int, len(current)),
				qe:       1,
			}
			copy(group.packages, current)
			for _, w := range current {
				group.qe *= int64(w)
			}
			*results = append(*results, group)
		}
		return
	}

	for i := start; i < len(weights); i++ {
		current = append(current, weights[i])
		findValidGroups(weights, size, target, i+1, current, results)
		current = current[:len(current)-1]
	}
}

func canSplitIntoGroups(weights []int, numGroups, targetWeight int) bool {
	if numGroups == 1 {
		sum := 0
		for _, w := range weights {
			sum += w
		}
		return sum == targetWeight
	}

	var validFirstGroups [][]int
	findAllValidGroups(weights, targetWeight, 0, []int{}, &validFirstGroups)

	for _, group := range validFirstGroups {
		remaining := removePackages(weights, group)
		if canSplitIntoGroups(remaining, numGroups-1, targetWeight) {
			return true
		}
	}

	return false
}

func findAllValidGroups(weights []int, target, start int, current []int, results *[][]int) {
	sum := 0
	for _, w := range current {
		sum += w
	}

	if sum == target {
		groupCopy := make([]int, len(current))
		copy(groupCopy, current)
		*results = append(*results, groupCopy)
		return
	}

	if sum > target {
		return
	}

	for i := start; i < len(weights); i++ {
		current = append(current, weights[i])
		findAllValidGroups(weights, target, i+1, current, results)
		current = current[:len(current)-1]
	}
}

func removePackages(weights []int, toRemove []int) []int {
	result := make([]int, 0)
	weightMap := make(map[int]int)

	for _, w := range weights {
		weightMap[w]++
	}

	for _, w := range toRemove {
		weightMap[w]--
	}

	for w, count := range weightMap {
		for i := 0; i < count; i++ {
			result = append(result, w)
		}
	}

	return result
}
