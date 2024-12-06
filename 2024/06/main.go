package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
)

const (
	EMPTY    rune = '.'
	GUARD    rune = '^'
	OBSTACLE rune = '#'
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)

	fmt.Println("\nPart one:", partOne(grid))
	fmt.Println("Part two:", partTwo(grid))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(grid [][]rune) int {
	sum := 0
	visited := make(map[[4]int]bool)
	startR, startC := findGuardStartPos(grid)

	// Find the positions that the guard visits to leave the grid
	isGuardInLoop(grid, visited, startR, startC, -1, 0)

	for r, row := range grid {
		for c := range row {
			if visited[[4]int{r, c, 0, 1}] || visited[[4]int{r, c, 0, -1}] ||
				visited[[4]int{r, c, 1, 0}] || visited[[4]int{r, c, -1, 0}] {
				sum++
			}
		}
	}

	return sum
}

func partTwo(grid [][]rune) int {
	sum := 0
	startR, startC := findGuardStartPos(grid)

	for r, row := range grid {
		for c, cell := range row {
			if cell != EMPTY {
				continue
			}

			grid[r][c] = OBSTACLE

			// Check if the guard is stuck in a loop
			visited := make(map[[4]int]bool)
			if isGuardInLoop(grid, visited, startR, startC, -1, 0) {
				sum++
			}

			grid[r][c] = EMPTY
		}
	}

	return sum
}

func isGuardInLoop(grid [][]rune, visited map[[4]int]bool, r, c, dr, dc int) bool {
	if visited[[4]int{r, c, dr, dc}] {
		return true
	}
	visited[[4]int{r, c, dr, dc}] = true

	newR, newC := r+dr, c+dc
	newDr, newDc := dr, dc

	if !g.IsInsideBounds(grid, newR, newC) {
		return false
	}

	if grid[newR][newC] == OBSTACLE {
		// Change direction by turning right 90 degrees
		// When turning, stay at the same position
		newDr, newDc = dc, -dr
		newR, newC = r, c
	} else {
		newR, newC = r+newDr, c+newDc
	}

	return isGuardInLoop(grid, visited, newR, newC, newDr, newDc)
}

func findGuardStartPos(grid [][]rune) (int, int) {
	for r, row := range grid {
		for c, cell := range row {
			if cell == GUARD {
				return r, c
			}
		}
	}

	return -1, -1
}
