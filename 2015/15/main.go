package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Ingredient struct {
	name                                            string
	capacity, durability, flavor, texture, calories int
}

var ingredients []Ingredient

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		parts := strings.Split(line, ":")
		numbers := u.GetIntsFromString(parts[1], true)

		ingredients = append(ingredients, Ingredient{
			name:       parts[0],
			capacity:   numbers[0],
			durability: numbers[1],
			flavor:     numbers[2],
			texture:    numbers[3],
			calories:   numbers[4],
		})
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	amounts := make([]int, len(ingredients))
	return findBestScore(ingredients, amounts, 100, 0, false)
}

func partTwo() int {
	amounts := make([]int, len(ingredients))
	return findBestScore(ingredients, amounts, 100, 0, true)
}

func findBestScore(ingredients []Ingredient, amounts []int, remaining int, currentIndex int, checkCalories bool) int {
	if currentIndex == len(ingredients)-1 {
		amounts[currentIndex] = remaining
		return calculateScore(ingredients, amounts, checkCalories)
	}

	bestScore := 0
	for amount := 0; amount <= remaining; amount++ {
		amounts[currentIndex] = amount
		score := findBestScore(ingredients, amounts, remaining-amount, currentIndex+1, checkCalories)
		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

func calculateScore(ingredients []Ingredient, amounts []int, checkCalories bool) int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	calories := 0

	for i, ing := range ingredients {
		capacity += ing.capacity * amounts[i]
		durability += ing.durability * amounts[i]
		flavor += ing.flavor * amounts[i]
		texture += ing.texture * amounts[i]
		calories += ing.calories * amounts[i]
	}

	// If checking calories and not exactly 500, return 0
	if checkCalories && calories != 500 {
		return 0
	}

	// If any property is negative, return 0
	if capacity <= 0 || durability <= 0 || flavor <= 0 || texture <= 0 {
		return 0
	}

	return capacity * durability * flavor * texture
}
