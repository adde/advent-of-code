package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Beam struct {
	X, Y int
	D    [2]int
}

var grid [][]rune
var vGrid [][]rune
var isDrawing = true
var dirs = map[string][2]int{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
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

	fmt.Println()
	fmt.Println("Tiles energized:", getEnergizedTiles(0, 0))
	fmt.Println("Max tiles energized:", getMaxEnergizedTiles())
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))

	// Uncomment to visualize the beam path(approx 1 min):
	// visualizeBeamPath()
}

func getMaxEnergizedTiles() int {
	maxEnergizedTiles := 0

	for y, row := range grid {
		for x := range row {
			if !isOuterEdge(x, y) {
				continue
			}

			energizedTiles := getEnergizedTiles(x, y)

			if energizedTiles > maxEnergizedTiles {
				maxEnergizedTiles = energizedTiles
			}
		}
	}

	return maxEnergizedTiles
}

func getEnergizedTiles(startX, startY int) int {
	rows := len(grid)
	cols := len(grid[0])

	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// BFS to find the number of energized tiles
	bfsTiles(startX, startY, visited, getStartingDirection(startX, startY), false)

	energizedTiles := 0
	for y, row := range visited {
		for x := range row {
			if visited[y][x] {
				energizedTiles++
			}
		}
	}

	return energizedTiles
}

func bfsTiles(x, y int, visited [][]bool, directions [2]int, visualize bool) {
	queue := []Beam{{x, y, directions}}
	visited[y][x] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr.Y][curr.X] = true

		// Get positon of the next tile based on the current direction
		newY, newX := curr.Y+curr.D[0], curr.X+curr.D[1]

		// If we hit a wall, this branch is dead
		if !isInBoundaries(newX, newY) {
			continue
		}

		if (isRight(curr.D) || isLeft(curr.D)) && grid[newY][newX] == '|' && !visited[newY][newX] {
			// If we hit a vertical splitter from right or left, we branch out to up and down
			queue = append(queue, Beam{newX, newY, dirs["up"]})
			queue = append(queue, Beam{newX, newY, dirs["down"]})
		} else if (isUp(curr.D) || isDown(curr.D)) && grid[newY][newX] == '-' && !visited[newY][newX] {
			// If we hit a horizontal splitter from up or down, we branch out to right and left
			queue = append(queue, Beam{newX, newY, dirs["right"]})
			queue = append(queue, Beam{newX, newY, dirs["left"]})
		} else if isRight(curr.D) && grid[newY][newX] == '/' || isLeft(curr.D) && grid[newY][newX] == '\\' {
			// If we hit a mirror from right or left, we change direction to up
			queue = append(queue, Beam{newX, newY, dirs["up"]})
		} else if isRight(curr.D) && grid[newY][newX] == '\\' || isLeft(curr.D) && grid[newY][newX] == '/' {
			// If we hit a mirror from right or left, we change direction to down
			queue = append(queue, Beam{newX, newY, dirs["down"]})
		} else if isUp(curr.D) && grid[newY][newX] == '/' || isDown(curr.D) && grid[newY][newX] == '\\' {
			// If we hit a mirror from up or down, we change direction to right
			queue = append(queue, Beam{newX, newY, dirs["right"]})
		} else if isUp(curr.D) && grid[newY][newX] == '\\' || isDown(curr.D) && grid[newY][newX] == '/' {
			// If we hit a mirror from up or down, we change direction to left
			queue = append(queue, Beam{newX, newY, dirs["left"]})
		} else if grid[newY][newX] == '.' ||
			(grid[newY][newX] == '|' && (isUp(curr.D) || isDown(curr.D))) ||
			(grid[newY][newX] == '-' && (isRight(curr.D) || isLeft(curr.D))) {
			// If we hit an empty tile or a splitter "on the tip",
			// we keep going in the same direction
			queue = append(queue, Beam{newX, newY, curr.D})
		}

		// If visualizing, update the grid and wait a bit to let drawing happen
		if visualize {
			updateGrid(visited)

			// Speed up the updates if the queue has more items
			queueLen := len(queue)
			if queueLen <= 0 {
				queueLen = 1
			}
			durationMultiplier := int(100 / queueLen)

			time.Sleep(time.Millisecond * time.Duration(durationMultiplier))
		}
	}
}

func getStartingDirection(x int, y int) [2]int {
	if !isInBoundaries(x, y) {
		return [2]int{0, 0}
	}

	switch {
	case x == 0 && y == 0:
		if grid[y][x] == '.' || grid[y][x] == '-' {
			return [2]int{0, 1}
		} else if grid[y][x] == '\\' {
			return [2]int{1, 0}
		} else if grid[y][x] == '/' {
			return [2]int{-1, 0}
		}
	case x > 0 && y == 0:
		if grid[y][x] == '.' || grid[y][x] == '|' || grid[y][x] == '-' {
			return [2]int{1, 0}
		} else if grid[y][x] == '\\' {
			return [2]int{0, 1}
		} else if grid[y][x] == '/' {
			return [2]int{0, -1}
		}
	case x >= 0 && y == len(grid)-1:
		if grid[y][x] == '.' || grid[y][x] == '|' || grid[y][x] == '-' {
			return [2]int{-1, 0}
		} else if grid[y][x] == '\\' {
			return [2]int{0, -1}
		} else if grid[y][x] == '/' {
			return [2]int{0, 1}
		}
	case x == 0 && y > 0:
		if grid[y][x] == '.' || grid[y][x] == '-' || grid[y][x] == '|' {
			return [2]int{0, 1}
		} else if grid[y][x] == '\\' {
			return [2]int{1, 0}
		} else if grid[y][x] == '/' {
			return [2]int{-1, 0}
		}
	case x == len(grid[0])-1 && y > 0:
		if grid[y][x] == '.' || grid[y][x] == '-' || grid[y][x] == '|' {
			return [2]int{0, -1}
		} else if grid[y][x] == '\\' {
			return [2]int{-1, 0}
		} else if grid[y][x] == '/' {
			return [2]int{1, 0}
		}
	}

	return [2]int{0, 0}
}

func isInBoundaries(x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func isOuterEdge(x, y int) bool {
	return x == 0 || x == len(grid[0])-1 || y == 0 || y == len(grid)-1
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

func printGrid() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for y, row := range vGrid {
		for x, val := range row {
			if val == '#' {
				termbox.SetCell(x, y, val, termbox.ColorLightYellow, termbox.ColorDefault)
			} else {
				termbox.SetCell(x, y, val, termbox.ColorDarkGray, termbox.ColorDefault)
			}
		}
	}

	termbox.Flush()
}

func updateGrid(visited [][]bool) {
	for r, row := range visited {
		for c, v := range row {
			if v {
				vGrid[r][c] = '#'
			}
		}
	}
}

func visualizeBeamPath() {
	// Initialize a copy of the grid to visualize the beam path
	vGrid = make([][]rune, len(grid))
	for i := range vGrid {
		vGrid[i] = append(vGrid[i], grid[i]...)
	}

	termbox.Init()
	defer termbox.Close()

	// Run the BFS async so we can draw the grid while it's running
	go func() {
		rows := len(grid)
		cols := len(grid[0])

		visited := make([][]bool, rows)
		for i := range visited {
			visited[i] = make([]bool, cols)
		}

		bfsTiles(0, 0, visited, getStartingDirection(0, 0), true)
		isDrawing = false
	}()

	// Draw the grid until the BFS is done
	for isDrawing {
		printGrid()
		time.Sleep(time.Millisecond * 100)
	}
}
