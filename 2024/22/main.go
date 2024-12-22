package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	"github.com/adde/advent-of-code/utils/set"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	buyerNumbers := make([]int, 0)

	for _, line := range lines {
		buyerNumbers = append(buyerNumbers, u.ToInt(line))
	}

	fmt.Println("\nPart one:", partOne(buyerNumbers))
	fmt.Println("Part two:", partTwo(buyerNumbers))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(buyerNumbers []int) int {
	sum := 0

	for _, number := range buyerNumbers {
		secretNumber := number

		for i := 0; i < 2000; i++ {
			secretNumber = generateSecretNumber(secretNumber)
		}

		sum += secretNumber
	}

	return sum
}

func partTwo(buyerNumbers []int) int {
	seqTotal := make(map[[4]int]int)

	for _, number := range buyerNumbers {
		buyer := []int{number % 10}

		secretNumber := number
		for i := 0; i < 2000; i++ {
			secretNumber = generateSecretNumber(secretNumber)
			buyer = append(buyer, secretNumber%10)
		}

		seen := set.New[[4]int]()
		for i := 0; i < len(buyer)-4; i++ {
			seqNumbers := buyer[i : i+5]
			seqDiff := [4]int{
				seqNumbers[1] - seqNumbers[0],
				seqNumbers[2] - seqNumbers[1],
				seqNumbers[3] - seqNumbers[2],
				seqNumbers[4] - seqNumbers[3],
			}

			if seen.Contains(seqDiff) {
				continue
			}
			seen.Add(seqDiff)

			if _, ok := seqTotal[seqDiff]; !ok {
				seqTotal[seqDiff] = 0
			}
			seqTotal[seqDiff] += seqNumbers[4]
		}
	}

	return u.MaxSlice(u.MapValuesToInts(seqTotal))
}

func generateSecretNumber(number int) int {
	number = prune(mix(number, number*64))
	number = prune(mix(number, number/32))
	number = prune(mix(number, number*2048))

	return number
}

func mix(number, value int) int {
	return number ^ value
}

func prune(number int) int {
	return number % 16777216
}
