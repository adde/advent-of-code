package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	SCALE_FACTOR   = 3
	STARTING_CHAR  = 'S'
	LOOP_PIPE_CHAR = 'X'
	GROUND_CHAR    = '.'
)

var grid []string
var expandedGrid [][]rune
var pipeDirs = map[byte][][2]int{
	'|': {{1, 0}, {-1, 0}},
	'-': {{0, 1}, {0, -1}},
	'7': {{1, 0}, {0, -1}},
	'J': {{-1, 0}, {0, -1}},
	'L': {{-1, 0}, {0, 1}},
	'F': {{1, 0}, {0, 1}},
}

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	defer file.Close()

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Original grid
	steps, visited := stepsToFarthestPoint()
	// printOriginalGrid(visited)

	// Expanded grid
	expandedGrid = expandGrid(visited, SCALE_FACTOR)
	tilesEnclosed, visited := getEnclosedTiles()
	// printExpandedGrid()

	// Expanded visited grid (contrast between visited and unvisited tiles)
	// printExpandedVisitedGrid(visited)

	fmt.Println("\nSteps to farthest point(part one):", steps)
	fmt.Println("Tiles enclosed by loop(part two):", tilesEnclosed)
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func stepsToFarthestPoint() (int, [][]bool) {
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // Right, Left, Down, Up
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Finding the starting position 'S'
	startY, startX := getStarterPipePos()

	// DFS to find the number of steps to the farthest point in loop
	dfsSteps(startY, startX, visited, directions)

	steps := 0
	for y, row := range visited {
		for x, col := range row {
			_ = col
			if visited[y][x] {
				steps++
			}
		}
	}

	return steps / 2, visited
}

func getEnclosedTiles() (int, [][]bool) {
	rows := len(expandedGrid)
	cols := len(expandedGrid[0])
	enclosedTilesCount := 0

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Mark all tiles from the outer boundaries to the boundary of the loop as visited
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if (y == 0 || y == rows-1 || x == 0 || x == cols-1) &&
				!visited[y][x] && expandedGrid[y][x] == GROUND_CHAR {
				dfsEnclosed(y, x, visited)
			}
		}
	}

	// Count the number of unreachable tiles
	scaleFactor := SCALE_FACTOR
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			isVisited := true

			for i := 0; i < scaleFactor; i++ {
				for j := 0; j < scaleFactor; j++ {
					y := r*scaleFactor + i
					x := c*scaleFactor + j

					// Avoid hitting the outer boundaries
					if y < 1 || y > rows-1 || x < 1 || x > cols-1 {
						continue
					}

					// Ugly check! Could probably be improved, but my patience ran out...
					// Check around the current tile part (1x1 of the 3x3 expanded tile)
					// If all the sub tiles around the current tile are not '.', then the tile is not enclosed
					if !visited[y][x] && expandedGrid[y][x] == GROUND_CHAR && expandedGrid[y-1][x-1] != LOOP_PIPE_CHAR &&
						expandedGrid[y-1][x] != LOOP_PIPE_CHAR && expandedGrid[y-1][x+1] != LOOP_PIPE_CHAR &&
						expandedGrid[y][x-1] != LOOP_PIPE_CHAR && expandedGrid[y][x+1] != LOOP_PIPE_CHAR &&
						expandedGrid[y+1][x-1] != LOOP_PIPE_CHAR && expandedGrid[y+1][x] != LOOP_PIPE_CHAR &&
						expandedGrid[y+1][x+1] != LOOP_PIPE_CHAR {
						isVisited = false
					}
				}
			}

			if !isVisited {
				enclosedTilesCount++
			}
		}
	}

	return enclosedTilesCount, visited
}

func dfsSteps(y, x int, visited [][]bool, directions [][2]int) {
	visited[y][x] = true

	for _, dir := range directions {
		newY, newX := y+dir[0], x+dir[1]

		if isValid(newX, newY) &&
			!visited[newY][newX] &&
			grid[newY][newX] != GROUND_CHAR &&
			(canGoUp(newX, newY, dir) || canGoDown(newX, newY, dir) ||
				canGoLeft(newX, newY, dir) || canGoRight(newX, newY, dir)) {
			dfsSteps(newY, newX, visited, pipeDirs[grid[newY][newX]])
		}
	}
}

func dfsEnclosed(y, x int, visited [][]bool) {
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} // Right, Left, Down, Up
	visited[y][x] = true

	for _, dir := range directions {
		newX, newY := x+dir[1], y+dir[0]

		if isValidExpanded(newX, newY) && !visited[newY][newX] && expandedGrid[newY][newX] != LOOP_PIPE_CHAR {
			dfsEnclosed(newY, newX, visited)
		}
	}
}

// Check if a position is within grid boundaries
func isValid(x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func isValidExpanded(x, y int) bool {
	return x >= 0 && x < len(expandedGrid[0]) && y >= 0 && y < len(expandedGrid)
}

func canGoUp(x, y int, dir [2]int) bool {
	if isUp(dir) && grid[y][x] == '|' {
		return true
	} else if isUp(dir) && grid[y][x] == 'F' {
		return true
	} else if isUp(dir) && grid[y][x] == '7' {
		return true
	}
	return false
}

func canGoDown(x, y int, dir [2]int) bool {
	if isDown(dir) && grid[y][x] == '|' {
		return true
	} else if isDown(dir) && grid[y][x] == 'L' {
		return true
	} else if isDown(dir) && grid[y][x] == 'J' {
		return true
	}
	return false
}

func canGoLeft(x, y int, dir [2]int) bool {
	if isLeft(dir) && grid[y][x] == '-' {
		return true
	} else if isLeft(dir) && grid[y][x] == 'F' {
		return true
	} else if isLeft(dir) && grid[y][x] == 'L' {
		return true
	}
	return false
}

func canGoRight(x, y int, dir [2]int) bool {
	if isRight(dir) && grid[y][x] == '-' {
		return true
	} else if isRight(dir) && grid[y][x] == '7' {
		return true
	} else if isRight(dir) && grid[y][x] == 'J' {
		return true
	}
	return false
}

func isRight(dir [2]int) bool {
	return dir[0] == 0 && dir[1] == 1
}

func isLeft(dir [2]int) bool {
	return dir[0] == 0 && dir[1] == -1
}

func isUp(dir [2]int) bool {
	return dir[0] == -1 && dir[1] == 0
}

func isDown(dir [2]int) bool {
	return dir[0] == 1 && dir[1] == 0
}

func getStarterPipePos() (int, int) {
	var startY, startX int

	for y, row := range grid {
		for x, char := range row {
			if char == STARTING_CHAR {
				startY = y
				startX = x
				break
			}
		}
	}

	return startY, startX
}

func getStarterPipe() byte {
	y, x := getStarterPipePos()
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	connections := [][2]int{}

	// Check if the starting pipe is connected to two other pipes
	for _, dir := range directions {
		newY, newX := y+dir[0], x+dir[1]

		if isValid(newX, newY) &&
			grid[newY][newX] != GROUND_CHAR &&
			(canGoUp(newX, newY, dir) || canGoDown(newX, newY, dir) ||
				canGoLeft(newX, newY, dir) || canGoRight(newX, newY, dir)) {
			connections = append(connections, dir)
		}
	}

	// Return the correct pipe character based on the connections
	if len(connections) == 2 {
		if isDown(connections[0]) && isUp(connections[1]) {
			return '|'
		} else if isRight(connections[0]) && isLeft(connections[1]) {
			return '-'
		} else if isLeft(connections[0]) && isDown(connections[1]) {
			return '7'
		} else if isLeft(connections[0]) && isUp(connections[1]) {
			return 'J'
		} else if isRight(connections[0]) && isUp(connections[1]) {
			return 'L'
		} else if isRight(connections[0]) && isDown(connections[1]) {
			return 'F'
		}
	}

	return ' '
}

func expandGrid(visited [][]bool, scaleFactor int) [][]rune {
	rows, cols := len(grid), len(grid[0])
	expandedRows, expandedCols := rows*scaleFactor, cols*scaleFactor
	expandedGrid = make([][]rune, expandedRows)
	for i := range expandedGrid {
		expandedGrid[i] = make([]rune, expandedCols)
	}

	// Create a new grid with the expanded size
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			// Fill the expanded grid with '.'
			for i := 0; i < scaleFactor; i++ {
				for j := 0; j < scaleFactor; j++ {
					expandedGrid[r*scaleFactor+i][c*scaleFactor+j] = GROUND_CHAR
				}
			}

			// Leave unvisited tiles as '.'
			if !visited[r][c] {
				continue
			}

			// Print pipes as 'X' in the expanded grid
			if grid[r][c] == '|' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == '|') {
				expandedGrid[r*scaleFactor+0][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+2][c*scaleFactor+1] = LOOP_PIPE_CHAR
			} else if grid[r][c] == '-' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == '-') {
				expandedGrid[r*scaleFactor+1][c*scaleFactor+0] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+2] = LOOP_PIPE_CHAR
			} else if grid[r][c] == '7' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == '7') {
				expandedGrid[r*scaleFactor+1][c*scaleFactor+0] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+2][c*scaleFactor+1] = LOOP_PIPE_CHAR
			} else if grid[r][c] == 'F' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == 'F') {
				expandedGrid[r*scaleFactor+2][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+2] = LOOP_PIPE_CHAR
			} else if grid[r][c] == 'L' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == 'L') {
				expandedGrid[r*scaleFactor+0][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+2] = LOOP_PIPE_CHAR
			} else if grid[r][c] == 'J' || (grid[r][c] == STARTING_CHAR && getStarterPipe() == 'J') {
				expandedGrid[r*scaleFactor+1][c*scaleFactor+0] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+1][c*scaleFactor+1] = LOOP_PIPE_CHAR
				expandedGrid[r*scaleFactor+0][c*scaleFactor+1] = LOOP_PIPE_CHAR
			}

		}
	}

	return expandedGrid
}

func printOriginalGrid(visited [][]bool) {
	fmt.Println("\nOriginal Grid:")

	for y, row := range visited {
		for x, col := range row {
			if col {
				fmt.Print(string(grid[y][x]))
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func printExpandedGrid() {
	fmt.Println("\nExpanded Grid:")

	for _, row := range expandedGrid {
		fmt.Println(string(row))
	}

	fmt.Println()
}

func printExpandedVisitedGrid(visited [][]bool) {
	fmt.Println("\nExpanded visited Grid:")

	for y, row := range visited {
		for x, col := range row {
			if col {
				fmt.Print("O")
			} else {
				fmt.Print(string(expandedGrid[y][x]))
			}
		}

		fmt.Println()
	}

	fmt.Println()
}
