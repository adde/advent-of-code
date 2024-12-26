package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	data := u.ReadAll("input.txt")

	floor := 0
	basement := 0

	for i, c := range data {
		if c == '(' {
			floor++
		} else if c == ')' {
			floor--
		}
		if floor == -1 && basement == 0 {
			basement = i + 1
		}
	}

	fmt.Println("\nPart one:", floor)
	fmt.Println("Part two:", basement)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
