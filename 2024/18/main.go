package main

import (
	"container/heap"
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	pq "github.com/adde/advent-of-code/utils/priorityqueue"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	GRID_SIZE     = 71
	BYTES_TO_FALL = 1024
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	bytes := []g.Point{}

	for _, line := range lines {
		coords := u.GetIntsFromString(line, false)
		bytes = append(bytes, g.Point{Row: coords[1], Col: coords[0]})
	}

	fmt.Println("\nPart one:", partOne(bytes, GRID_SIZE, GRID_SIZE, BYTES_TO_FALL))
	fmt.Println("Part two:", partTwo(bytes, GRID_SIZE, GRID_SIZE))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(bytes []g.Point, rows, cols, bytesToFall int) int {
	grid := g.Create(rows, cols, '.')

	for i := 0; i < bytesToFall; i++ {
		grid.Set(bytes[i].Row, bytes[i].Col, '#')
	}

	return getMinSteps(grid, 0, 0, rows-1, cols-1)
}

func partTwo(bytes []g.Point, rows, cols int) string {
	blockingByte := ""

	// Iterate through the bytes backwards and
	// find the first byte that blocks the path to the target
	for b := len(bytes) - 1; b >= 0; b-- {
		grid := g.Create(rows, cols, '.')

		// Fill the grid with bytes up to the current byte
		for i := 0; i < b; i++ {
			grid.Set(bytes[i].Row, bytes[i].Col, '#')
		}

		blockingByte = fmt.Sprintf("%d,%d", bytes[b].Col, bytes[b].Row)

		// If the path is no longer blocked, we found the blocking byte
		if getMinSteps(grid, 0, 0, rows-1, cols-1) != 0 {
			break
		}
	}

	return blockingByte
}

// Get the shortest distance from the start cell to the target cell in a 2D grid
func getMinSteps(grid g.Grid, startRow, startCol, targetRow, targetCol int) int {
	seen := set.New[[2]int]()

	prioQueue := &pq.PriorityQueue{pq.Cell{Row: startRow, Col: startCol, Cost: 0}}
	heap.Init(prioQueue)

	for prioQueue.Len() > 0 {
		cell := heap.Pop(prioQueue).(pq.Cell)

		if cell.Row == targetRow && cell.Col == targetCol {
			return cell.Cost
		}

		key := [2]int{cell.Row, cell.Col}
		if seen.Contains(key) {
			continue
		}
		seen.Add(key)

		for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			newRow, newCol := cell.Row+dir[0], cell.Col+dir[1]

			if !grid.IsInsideBounds(newRow, newCol) || grid[newRow][newCol] == '#' {
				continue
			}

			heap.Push(prioQueue, pq.Cell{Row: newRow, Col: newCol, Cost: cell.Cost + 1})
		}
	}

	return 0
}
