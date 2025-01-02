package main

import (
	"fmt"
	"math"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	input := u.ReadAll("input.txt")
	target := u.ToInt(input)

	fmt.Println("\nPart one:", partOne(target))
	fmt.Println("Part two:", partTwo(target))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(targetPresents int) int {
	total := 0

	for i := 1; i < targetPresents; i++ {
		if sumOfFactors(i)*10 >= targetPresents {
			total = i
			break
		}
	}

	return total
}

func partTwo(targetPresents int) int {
	maxHouses := 1000000 // Reasonable upper limit
	deliveryLimit := 50  // Each elf stops after 50 houses
	multiplier := 11     // Each elf delivers 11 times their number

	presents := make([]int, maxHouses+1)

	// For each elf
	for elf := 1; elf <= maxHouses; elf++ {
		// Calculate how many houses this elf will visit
		deliveries := 0

		// Visit each house that's a multiple of this elf's number
		for house := elf; house <= maxHouses && deliveries < deliveryLimit; house += elf {
			presents[house] += elf * multiplier
			deliveries++
		}

		if presents[elf] >= targetPresents {
			return elf
		}
	}

	// Find the first house that gets enough presents
	for house := 1; house <= maxHouses; house++ {
		if presents[house] >= targetPresents {
			return house
		}
	}

	return -1
}

func sumOfFactors(n int) int {
	if n == 1 {
		return 1
	}

	sqrt := int(math.Sqrt(float64(n)))
	sum := 1 + n // Always include 1 and n

	// Only check up to square root
	for i := 2; i <= sqrt; i++ {
		if n%i == 0 {
			sum += i

			// Add the pair factor if it's different
			if i != n/i {
				sum += n / i
			}
		}
	}

	return sum
}
