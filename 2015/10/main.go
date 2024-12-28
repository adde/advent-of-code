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
	digits := u.ReadAll("input.txt")
	ansP1, ansP2 := 0, 0

	for i := 0; i < 50; i++ {
		digits = generateNext(digits)
		if i == 39 {
			ansP1 = len(digits)
		}
	}
	ansP2 = len(digits)

	fmt.Println("\nPart one:", ansP1)
	fmt.Println("Part two:", ansP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func generateNext(digits string) string {
	// Pre-allocate builder with estimated capacity
	// Most sequences grow by ~30% each iteration
	builder := strings.Builder{}
	builder.Grow(len(digits) * 13 / 10)

	for i := 0; i < len(digits); {
		digit := digits[i]
		count := 1

		// Count consecutive occurrences
		for i+1 < len(digits) && digits[i+1] == digit {
			count++
			i++
		}

		// Write count and digit directly to builder
		if count < 10 {
			builder.WriteByte(byte('0' + count))
		} else {
			builder.WriteString(strconv.Itoa(count))
		}
		builder.WriteByte(digit)

		i++
	}

	return builder.String()
}
