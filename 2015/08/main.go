package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	fmt.Println("\nPart one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(lines []string) int {
	ans := 0

	for _, line := range lines {
		lineLen := len(line)
		memLen := 0

		for i := 1; i < len(line)-1; i++ {
			memLen++

			if line[i] == '\\' && line[i+1] == 'x' {
				i += 3
			} else if line[i] == '\\' {
				i++
			}
		}

		ans += lineLen - memLen
	}

	return ans
}

func partTwo(lines []string) int {
	ans := 0

	for _, line := range lines {
		lineLen := len(line)
		encodedLen := 0

		for i := 0; i < len(line); i++ {
			encodedLen++

			if (line[i] == '\\' || line[i] == '"') &&
				i != 0 && i != len(line)-1 {
				encodedLen += 1
			}
		}

		encodedLen += 4
		ans += encodedLen - lineLen
	}

	return ans
}
