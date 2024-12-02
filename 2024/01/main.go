package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")

	distance, similarity := getDistanceAndSimilarity(parseLists(lines))

	fmt.Println("\nPart one:", distance)
	fmt.Println("Part two:", similarity)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func getDistanceAndSimilarity(left, right []int) (int, int) {
	distance := 0
	similarity := 0

	for i := 0; i < len(left); i++ {
		distance += utils.Abs(left[i] - right[i])

		appearances := 0
		for j := 0; j < len(left); j++ {
			if left[i] == right[j] {
				appearances++
			}
		}
		similarity += left[i] * appearances
	}

	return distance, similarity
}

func parseLists(lines []string) ([]int, []int) {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		numbers := strings.Split(line, "   ")
		left = append(left, utils.ToInt(numbers[0]))
		right = append(right, utils.ToInt(numbers[1]))
	}

	sort.Ints(left)
	sort.Ints(right)

	return left, right
}
