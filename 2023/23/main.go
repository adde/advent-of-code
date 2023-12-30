package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var grid [][]rune
var visited [][]bool
var maxLength int
var dirs = map[rune][][2]int{
	'^': {{-1, 0}},
	'v': {{1, 0}},
	'<': {{0, -1}},
	'>': {{0, 1}},
	'.': {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
}

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	r := 0
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, make([]rune, len(line)))

		for c, v := range line {
			grid[r][c] = v
		}

		r++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part two needs optimization... Takes ~5 minutes to run
	fmt.Println("\nLongest path length(part one):", findLongestPath(false))
	fmt.Println("Longest path length(part two):", findLongestPath(true))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func findLongestPath(removeSlopes bool) int {
	// Remove slopes
	if removeSlopes {
		for r, row := range grid {
			for c, v := range row {
				if v == '<' || v == '>' || v == '^' || v == 'v' {
					grid[r][c] = '.'
				}
			}
		}
	}

	// Find start position
	startRow, startCol := 0, 0
	for c, v := range grid[0] {
		if v == '.' {
			startCol = c
			break
		}
	}

	// Find end position
	endRow, endCol := len(grid)-1, 0
	for c, v := range grid[len(grid)-1] {
		if v == '.' {
			endCol = c
			break
		}
	}

	// Initialize visited array
	visited = make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}

	// Initialize maxLength to zero before running DFS
	maxLength = 0

	// Run DFS to find the longest path between start and end points
	dfs(startRow, startCol, endRow, endCol, dirs['.'], 0)

	return maxLength
}

func dfs(row, col, endR, endC int, directions [][2]int, length int) {
	// Keep track of the original state of the visited array at the current cell
	originalVisited := visited[row][col]
	visited[row][col] = true

	for _, dir := range directions {
		newR, newC := row+dir[0], col+dir[1]

		// Check if we have reached the end
		if newR == endR && newC == endC {
			if length+1 > maxLength {
				maxLength = length + 1
			}
		}

		if isInBoundary(newC, newR) && !visited[newR][newC] && grid[newR][newC] != '#' {
			dfs(newR, newC, endR, endC, dirs[grid[newR][newC]], length+1)
		}
	}

	// Backtrack: Restore the original state of the visited array at the current cell
	visited[row][col] = originalVisited
}

func isInBoundary(x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}
