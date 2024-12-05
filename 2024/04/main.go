package main

import (
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := utils.ReadLines("input.txt")
	grid := make([][]rune, 0)

	r := 0
	for _, line := range lines {
		grid = append(grid, make([]rune, len(line)))
		for c, v := range line {
			grid[r][c] = v
		}
		r++
	}

	fmt.Println("\nPart one:", partOne(grid))
	fmt.Println("Part two:", partTwo(grid))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(grid [][]rune) int {
	sum := 0
	searchTerm := []rune{'X', 'M', 'A', 'S'}
	rotatedGrid := grid

	// Rotate the grid 4 times to check all possible directions
	for i := 0; i < 4; i++ {
		for _, row := range rotatedGrid {
			sum += countRowOccurrences(row, searchTerm)
		}

		if diaOcc := countDiagonalOccurrences(rotatedGrid, searchTerm); diaOcc > 0 {
			sum += diaOcc
		}

		rotatedGrid = rotateGrid(rotatedGrid)
	}

	return sum
}

func partTwo(grid [][]rune) int {
	sum := 0
	searchTerm := []rune{'M', 'A', 'S'}
	rows, cols := len(grid), len(grid[0])

	// Find all 'A' chars in the grid
	// and check if 'MAS' is found in the diagonals
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] != 'A' {
				continue
			}

			if (checkDiagonal(grid, r-1, c-1, 1, 1, searchTerm) ||
				checkDiagonal(grid, r+1, c+1, -1, -1, searchTerm)) &&
				(checkDiagonal(grid, r-1, c+1, 1, -1, searchTerm) ||
					checkDiagonal(grid, r+1, c-1, -1, 1, searchTerm)) {
				sum++
			}
		}
	}

	return sum
}

func rotateGrid(grid [][]rune) [][]rune {
	n := len(grid)

	rotated := make([][]rune, n)
	for i := range rotated {
		rotated[i] = make([]rune, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			rotated[j][n-i-1] = grid[i][j]
		}
	}

	return rotated
}

func countRowOccurrences(row []rune, searchTerm []rune) int {
	sum := 0

	for i, val := range row {
		if string(val) == string(searchTerm[0]) {
			hit := true

			for j := 1; j < len(searchTerm); j++ {
				if i+j >= len(row) || row[i+j] != searchTerm[j] {
					hit = false
					break
				}
			}

			if hit {
				sum++
			}
		}
	}

	return sum
}

func countDiagonalOccurrences(grid [][]rune, searchTerm []rune) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	termLen := len(searchTerm)
	count := 0

	// Check diagonal ↘ (top-left to bottom-right)
	for r := 0; r <= rows-termLen; r++ {
		for c := 0; c <= cols-termLen; c++ {
			if checkDiagonal(grid, r, c, 1, 1, searchTerm) {
				count++
			}
		}
	}

	return count
}

func checkDiagonal(grid [][]rune, r, c, dr, dc int, searchTerm []rune) bool {
	for i := 0; i < len(searchTerm); i++ {
		if grid[r+i*dr][c+i*dc] != searchTerm[i] {
			return false
		}
	}

	return true
}
