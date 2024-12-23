package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
)

const (
	START           = 'S'
	END             = 'E'
	WALL            = '#'
	CHEAT_THRESHOLD = 100
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)

	fmt.Println("\nPart one:", partOne(grid))
	fmt.Println("Part two:", partTwo(grid))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(grid g.Grid) int {
	return getCheats(grid, getDistances(grid), 2)
}

func partTwo(grid g.Grid) int {
	return getCheats(grid, getDistances(grid), 20)
}

// Get the distances from the start cell to all other cells in the grid
func getDistances(grid g.Grid) [][]int {
	row, col := grid.Find(START)
	rows, cols := grid.RowLen(), grid.ColumnLen()

	dists := make([][]int, rows)
	for i := range dists {
		dists[i] = make([]int, cols)
	}

	// Initialize all distances to -1
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			dists[row][col] = -1
		}
	}

	// Set the distance from the start cell to itself to 0
	dists[row][col] = 0

	// Check until we reach the end cell
	for grid[row][col] != END {
		for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			newRow, newCol := row+dir[0], col+dir[1]

			if !grid.IsInsideBounds(newRow, newCol) ||
				grid.Get(newRow, newCol) == WALL ||
				dists[newRow][newCol] != -1 {
				continue
			}

			dists[newRow][newCol] = dists[row][col] + 1
			row, col = newRow, newCol
		}
	}

	return dists
}

// Get the number of cheats that would save an amount of time
func getCheats(grid g.Grid, dists [][]int, cheatLen int) int {
	count := 0

	// Check all cells
	for row := 0; row < grid.RowLen(); row++ {
		for col := 0; col < grid.ColumnLen(); col++ {
			if grid[row][col] == WALL {
				continue
			}

			// Check for possible cheats from the current cell
			for cheat := 2; cheat <= cheatLen; cheat++ {
				for dirRow := 0; dirRow <= cheat; dirRow++ {
					dirCol := cheat - dirRow

					// Define the four possible coordinate combinations
					coordinates := map[[2]int]struct{}{
						{row + dirRow, col + dirCol}: {},
						{row + dirRow, col - dirCol}: {},
						{row - dirRow, col + dirCol}: {},
						{row - dirRow, col - dirCol}: {},
					}

					// Check each coordinate pair
					for coord := range coordinates {
						nextRow, nextCol := coord[0], coord[1]

						if !grid.IsInsideBounds(nextRow, nextCol) || grid[nextRow][nextCol] == WALL {
							continue
						}

						if dists[row][col]-dists[nextRow][nextCol] >= CHEAT_THRESHOLD+cheat {
							count++
						}
					}
				}
			}
		}
	}

	return count
}
