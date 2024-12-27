package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	ansP1, ansP2 := 0, 0

	// Part one
	for _, line := range lines {
		if strings.Contains(line, "ab") ||
			strings.Contains(line, "cd") ||
			strings.Contains(line, "pq") ||
			strings.Contains(line, "xy") {
			continue
		}

		vowels := 0
		for _, c := range line {
			if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
				vowels++
			}
		}
		if vowels < 3 {
			continue
		}

		repeat := false
		for i := 0; i < len(line)-1; i++ {
			if line[i] == line[i+1] {
				repeat = true
				break
			}
		}
		if !repeat {
			continue
		}

		ansP1++
	}

	// Part two
	for _, line := range lines {
		repeat := false
		for i := 0; i < len(line)-2; i++ {
			if strings.Contains(line[i+2:], line[i:i+2]) {
				repeat = true
				break
			}
		}
		if !repeat {
			continue
		}

		repeat = false
		for i := 0; i < len(line)-2; i++ {
			if line[i] == line[i+2] {
				repeat = true
				break
			}
		}
		if !repeat {
			continue
		}

		ansP2++
	}

	fmt.Println("\nPart one:", ansP1)
	fmt.Println("Part two:", ansP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
