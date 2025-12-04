package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	fmt.Println("\nPart one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(lines []string) int {
	grid := g.CreateFromLines(lines)
	accessibleRolls := 0

	for r, row := range grid {
		for c, cell := range row {
			if cell != '@' {
				continue
			}

			adjacentRolls := checkAdjacentRolls(grid, r, c)

			if adjacentRolls < 4 {
				accessibleRolls++
			}
		}
	}

	return accessibleRolls
}

func partTwo(lines []string) int {
	grid := g.CreateFromLines(lines)
	done := false
	rollsRemoved := 0

	// Run multiple passes until no more rolls can be removed
	for !done {
		done = true

		for r, row := range grid {
			for c, cell := range row {
				if cell != '@' {
					continue
				}

				adjacentRolls := checkAdjacentRolls(grid, r, c)

				if adjacentRolls < 4 {
					grid.Set(r, c, '.')
					rollsRemoved++
					done = false
				}
			}
		}
	}

	return rollsRemoved
}

// Check the eight adjacent cells for rolls
func checkAdjacentRolls(grid g.Grid, r, c int) int {
	adjacentRolls := 0

	for _, dir := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		newRow, newCol := r+dir[0], c+dir[1]

		if !grid.IsInsideBounds(newRow, newCol) {
			continue
		}

		if grid.Get(newRow, newCol) == '@' {
			adjacentRolls++
		}
	}

	return adjacentRolls
}
