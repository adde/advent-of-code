package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	pos, partOne, partTwo := 50, 0, 0

	for _, line := range lines {
		var dir string
		var val int
		fmt.Sscanf(line, "%1s%d", &dir, &val)

		for i := 0; i < val; i++ {
			switch dir {
			case "L":
				pos = (pos - 1 + 100) % 100
			case "R":
				pos = (pos + 1) % 100
			}

			if pos == 0 {
				partTwo++
			}
		}

		if pos == 0 {
			partOne++
		}
	}

	fmt.Println("\nPart one:", partOne)
	fmt.Println("Part two:", partTwo)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
