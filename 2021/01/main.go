package main

import (
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	input := make([]int, 0)

	for _, line := range lines {
		value := utils.ToInt(line)
		input = append(input, value)
	}

	incOne := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			incOne++
		}
	}

	incTwo := 0
	for i := 3; i < len(input); i++ {
		if input[i]+input[i-1]+input[i-2] > input[i-1]+input[i-2]+input[i-3] {
			incTwo++
		}
	}

	fmt.Println("\nPart one:", incOne)
	fmt.Println("Part two:", incTwo)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
