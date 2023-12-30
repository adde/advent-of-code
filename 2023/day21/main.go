package main

import (
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

var grid [][]rune

func main() {
	startTime := time.Now()

	// Populate the grid from input
	parseInput(utils.ReadLines("input.txt"))

	// Find the starting position marked 'S' in the grid
	startX, startY := getStartPos('S')

	fmt.Println("\nGarden plots reached(part one):", gardenPlotsReached(startY, startX, 64))
	fmt.Println("Garden plots reached(part two):", gardenPlotsReachedP2(startX, startY, 26501365))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) {
	row := 0

	for _, line := range lines {
		grid = append(grid, make([]rune, len(line)))

		for col, val := range line {
			grid[row][col] = val
		}

		row++
	}
}

func gardenPlotsReached(startRow, startCol, steps int) int {
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	reached := make(map[string]bool)
	seen := make(map[string]bool)
	q := make([][3]int, 0)
	q = append(q, [3]int{startRow, startCol, steps})

	for len(q) > 0 {
		current := q[0]
		cr, cc, cs := current[0], current[1], current[2]
		q = q[1:]

		if cs%2 == 0 {
			reached[fmt.Sprintf("%d,%d", cr, cc)] = true
		}

		if cs == 0 {
			continue
		}

		for _, dir := range directions {
			nr, nc := cr+dir[0], cc+dir[1]
			sk := fmt.Sprintf("%d,%d", nr, nc)

			if !isInBoundary(nc, nr) || grid[nr][nc] == '#' {
				continue
			}

			next := [3]int{nr, nc, cs - 1}
			if _, ok := seen[sk]; ok {
				continue
			}

			seen[sk] = true
			q = append(q, next)
		}
	}

	return len(reached)
}

func gardenPlotsReachedP2(startX, startY, steps int) int {
	gridSize := len(grid)
	gridWidth := steps/gridSize - 1

	// Get the number of grids that are odd and even
	oddGrids := ((gridWidth/2)*2 + 1) * ((gridWidth/2)*2 + 1)
	evenGrids := (((gridWidth + 1) / 2) * 2) * (((gridWidth + 1) / 2) * 2)

	// Get the number of plots reached in odd and even grids
	oddPlotsReached := gardenPlotsReached(startY, startX, gridSize*2+1)
	evenPlotsReached := gardenPlotsReached(startY, startX, gridSize*2)

	// Get the number of plots reached in the corners of the grid
	cornerT := gardenPlotsReached(gridSize-1, startX, gridSize-1)
	cornerR := gardenPlotsReached(startY, 0, gridSize-1)
	cornerB := gardenPlotsReached(0, startX, gridSize-1)
	cornerL := gardenPlotsReached(startY, gridSize-1, gridSize-1)

	// Get the number of plots reached in the corners of the small grids
	smallTR := gardenPlotsReached(gridSize-1, 0, gridSize/2-1)
	smallTL := gardenPlotsReached(gridSize-1, gridSize-1, gridSize/2-1)
	smallBR := gardenPlotsReached(0, 0, gridSize/2-1)
	smallBL := gardenPlotsReached(0, gridSize-1, gridSize/2-1)

	// Get the number of plots reached in the corners of the large grids
	largeTR := gardenPlotsReached(gridSize-1, 0, gridSize*3/2-1)
	largeTL := gardenPlotsReached(gridSize-1, gridSize-1, gridSize*3/2-1)
	largeBR := gardenPlotsReached(0, 0, gridSize*3/2-1)
	largeBL := gardenPlotsReached(0, gridSize-1, gridSize*3/2-1)

	return oddGrids*oddPlotsReached +
		evenGrids*evenPlotsReached +
		cornerT + cornerR + cornerB + cornerL +
		((gridWidth + 1) * (smallTR + smallTL + smallBR + smallBL)) +
		(gridWidth * (largeTR + largeTL + largeBR + largeBL))
}

func getStartPos(char rune) (int, int) {
	startX, startY := 0, 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == char {
				startX, startY = x, y
				break
			}
		}
	}

	return startX, startY
}

func isInBoundary(x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}
