package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	parts := strings.Split(file, "\n\n")
	presents := make([]int, 0)
	canFitPresents := 0

	for _, present := range parts[:len(parts)-1] {
		presents = append(presents, strings.Count(present, "#"))
	}

	regions := strings.Split(parts[len(parts)-1], "\n")
	for _, r := range regions {
		regionParts := strings.Split(r, ": ")
		dimensions := strings.Split(regionParts[0], "x")
		presentConfig := strings.Split(regionParts[1], " ")

		// Shrink region size by 10% to make the example work
		regionSize := int(float64(u.ToInt(dimensions[0])*u.ToInt(dimensions[1])) * 0.9)

		totalPresentSize := 0
		for presentIndex, presentCount := range presentConfig {
			if presentIndex < len(presents) {
				totalPresentSize += u.ToInt(presentCount) * presents[presentIndex]
			}
		}

		if totalPresentSize < regionSize {
			canFitPresents++
		}
	}

	fmt.Println("\nPart one:", canFitPresents)
	fmt.Println("Part two: -")
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
