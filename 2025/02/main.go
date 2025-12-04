package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	line := u.ReadAll("input.txt")
	ranges := strings.Split(line, ",")

	fmt.Println("\nPart one:", partOne(ranges))
	fmt.Println("Part two:", partTwo(ranges))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(ranges []string) int {
	total := 0

	for _, r := range ranges {
		bounds := u.GetIntsFromString(r, false)
		start, end := bounds[0], bounds[1]

		for id := start; id <= end; id++ {
			idStr := fmt.Sprintf("%d", id)

			// Skip if length is odd
			if len(idStr)%2 != 0 {
				continue
			}

			half := len(idStr) / 2
			firstHalf := idStr[:half]
			secondHalf := idStr[half:]

			// If both halves are equal, add invalid ID to total
			if firstHalf == secondHalf {
				total += id
			}
		}
	}

	return total
}

func partTwo(ranges []string) int {
	total := 0

	for _, r := range ranges {
		bounds := u.GetIntsFromString(r, false)
		start, end := bounds[0], bounds[1]

		for id := start; id <= end; id++ {
			if getRepeatingPattern(id) {
				total += id
			}
		}
	}

	return total
}

func getRepeatingPattern(n int) bool {
	idStr := strconv.Itoa(n)
	idStrLen := len(idStr)

	// Check all divisors of the ID length (excluding the full length)
	for pLen := 1; pLen <= idStrLen/2; pLen++ {
		if idStrLen%pLen == 0 {
			pattern := idStr[:pLen]

			if strings.Repeat(pattern, idStrLen/pLen) == idStr {
				return true
			}
		}
	}

	return false
}
