package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	r "github.com/adde/advent-of-code/utils/ranges"
)

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	lines := strings.Split(file, "\n\n")
	freshLines := strings.Split(lines[0], "\n")
	ingredientLines := strings.Split(lines[1], "\n")
	freshIDs := make(r.Ranges, len(freshLines))
	ingredientIDs := make([]int, len(ingredientLines))

	for i, v := range ingredientLines {
		ingredientIDs[i] = u.ToInt(v)
	}

	for i, v := range freshLines {
		parts := strings.Split(v, "-")
		freshIDs[i] = r.Range{
			Start: u.ToInt(parts[0]),
			End:   u.ToInt(parts[1]),
		}
	}

	fmt.Println("\nPart one:", partOne(freshIDs, ingredientIDs))
	fmt.Println("Part two:", partTwo(freshIDs))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(freshIDs r.Ranges, ingredientIDs []int) int {
	total := 0

	for _, ingredientID := range ingredientIDs {
		for _, freshID := range freshIDs {
			if ingredientID >= freshID.Start && ingredientID <= freshID.End {
				total++
				break
			}
		}
	}

	return total
}

func partTwo(freshIDs r.Ranges) int {
	total := 0

	// Merge overlapping fresh ID ranges
	freshIDs.Sort()
	mergedIDs := freshIDs.Merge()

	for _, freshID := range mergedIDs {
		total += freshID.End - freshID.Start + 1
	}

	return total
}
