package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

const (
	MULTIPLIER = 2024
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	stones := make([]int, 0)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		for _, part := range parts {
			stones = append(stones, u.ToInt(part))
		}
	}

	fmt.Println("\nPart one:", partOne(stones))
	fmt.Println("Part two:", partTwo(stones))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(stones []int) int {
	sum := 0
	cache := make(map[string]int)

	for _, stone := range stones {
		sum += blink(stone, cache, 25)
	}

	return sum
}

func partTwo(stones []int) int {
	sum := 0
	cache := make(map[string]int)

	for _, stone := range stones {
		sum += blink(stone, cache, 75)
	}

	return sum
}

func blink(stone int, cache map[string]int, blinks int) int {
	key := strconv.Itoa(stone) + "," + strconv.Itoa(blinks)
	if val, ok := cache[key]; ok {
		return val
	}

	newVal := -1
	if blinks > 0 {
		f, s := handleStone(stone)
		newVal = blink(f, cache, blinks-1)

		if s != -1 {
			newVal += blink(s, cache, blinks-1)
		}
	} else {
		newVal = 1
	}

	cache[key] = newVal
	return newVal
}

func handleStone(stone int) (int, int) {
	if stone == 0 {
		return 1, -1
	}

	if isEvenNumberOfDigits(stone) {
		f, s := getHalfsFromNumber(stone)
		return f, s
	}

	return stone * MULTIPLIER, -1
}

func getHalfsFromNumber(num int) (int, int) {
	numStr := strconv.Itoa(num)
	half := len(numStr) / 2

	firstHalf, _ := strconv.Atoi(numStr[:half])
	secondHalf, _ := strconv.Atoi(numStr[half:])

	return firstHalf, secondHalf
}

func isEvenNumberOfDigits(num int) bool {
	return len(strconv.Itoa(num))%2 == 0
}
