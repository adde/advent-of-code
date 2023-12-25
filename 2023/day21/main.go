package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type Step struct {
	X, Y, Count int
}

var grid [][]rune

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

	fmt.Println("\nGarden plots reached:", gardenPlotsReached(64))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func gardenPlotsReached(steps int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Find the starting position marked 'S' in the grid
	startX, startY := getStartPos()

	// BFS to find the number of garden plots that can be reached
	return bfs(startX, startY, steps, visited)
}

func bfs(x, y, steps int, visited [][]bool) int {
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	reached := make(map[string]bool)
	queue := []Step{{x, y, 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.Count <= steps && curr.Count%2 == 0 {
			reached[fmt.Sprintf("%d,%d", curr.X, curr.Y)] = true
		}

		if visited[curr.Y][curr.X] {
			continue
		}
		visited[curr.Y][curr.X] = true

		for _, d := range directions {
			newX, newY := curr.X+d[1], curr.Y+d[0]

			// If we hit a wall or a stone, this branch is dead
			if !isInBoundary(newX, newY) || grid[newY][newX] == '#' {
				continue
			}

			// Otherwise, add the new tile to the queue
			queue = append(queue, Step{newX, newY, curr.Count + 1})
		}
	}

	return len(reached)
}

func getStartPos() (int, int) {
	startX, startY := 0, 0

	for y, row := range grid {
		for x := range row {
			if grid[y][x] == 'S' {
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
