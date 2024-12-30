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
	grid := g.CreateFromLines(lines)

	fmt.Println("\nPart one:", animateLights(grid, 100, false))
	fmt.Println("Part two:", animateLights(grid, 100, true))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func animateLights(grid g.Grid, steps int, isCornersAlwaysOn bool) int {
	grid = grid.Copy()
	if isCornersAlwaysOn {
		turnCornerLightsOn(grid)
	}

	for i := 0; i < steps; i++ {
		next := grid.Copy()

		// Check if light should be turned on or off
		for r, row := range grid {
			for c, val := range row {
				on := countNeighboursOn(r, c, grid)

				if val == '.' && on == 3 {
					next.Set(r, c, '#')
				} else if val == '#' && on != 2 && on != 3 {
					next.Set(r, c, '.')
				}
			}
		}

		grid = next
		if isCornersAlwaysOn {
			turnCornerLightsOn(grid)
		}
	}

	// Count lights that are on
	count := 0
	for _, row := range grid {
		for _, val := range row {
			if val == '#' {
				count++
			}
		}
	}

	return count
}

// Count the number of neighbouring lights that are on
func countNeighboursOn(r, c int, grid g.Grid) int {
	count := 0

	for nr := r - 1; nr <= r+1; nr++ {
		for nc := c - 1; nc <= c+1; nc++ {
			if nr == r && nc == c {
				continue
			}
			if grid.IsInsideBounds(nr, nc) && grid.Get(nr, nc) == '#' {
				count++
			}
		}
	}

	return count
}

// Turn on the corner lights
func turnCornerLightsOn(grid g.Grid) {
	grid.Set(0, 0, '#')
	grid.Set(0, grid.ColumnLen()-1, '#')
	grid.Set(grid.RowLen()-1, 0, '#')
	grid.Set(grid.RowLen()-1, grid.ColumnLen()-1, '#')
}
