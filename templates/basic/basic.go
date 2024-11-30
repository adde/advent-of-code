package main

import (
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")

	for _, line := range lines {
		fmt.Println(line)
	}

	fmt.Println("\nPart one:")
	fmt.Println("Part two:")
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
