package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	wrappingPaper := 0
	ribbon := 0

	for _, line := range lines {
		parts := strings.Split(line, "x")
		length, width, height := u.ToInt(parts[0]), u.ToInt(parts[1]), u.ToInt(parts[2])

		// Part one
		extra := min(length*width, width*height, height*length)
		wrappingPaper += (2 * length * width) + (2 * width * height) + (2 * height * length) + extra

		// Part two
		sides := []int{length, width, height}
		sort.Ints(sides)
		ribbon += (sides[0] * 2) + (sides[1] * 2) + (length * width * height)
	}

	fmt.Println("\nPart one:", wrappingPaper)
	fmt.Println("Part two:", ribbon)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
