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

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)

	distance := getShortestDistance(grid, 0, 0, len(grid)-1, len(grid[0])-1)

	fmt.Println("\nPart one:", distance)
	fmt.Println("Part two:")
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Get the shortest distance from the start cell to the target cell in 2D grid
func getShortestDistance(grid g.Grid, startRow, startCol, targetRow, targetCol int) int {
	seen := set.New[[2]int]()

	// Create a priority queue and add the starting cell
	prioQueue := &pq.PriorityQueue{pq.Cell{Row: startRow, Col: startCol, Cost: 0}}
	heap.Init(prioQueue)

	for prioQueue.Len() > 0 {
		cell := heap.Pop(prioQueue).(pq.Cell)

		// Check if we reached the target cell, return the cost
		if cell.Row == targetRow && cell.Col == targetCol {
			return cell.Cost
		}

		// Check if we already visited this cell
		key := [2]int{cell.Row, cell.Col}
		if seen.Contains(key) {
			continue
		}
		seen.Add(key)

		// Check all four directions(Up, Right, Down, Left)
		for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			newRow, newCol := cell.Row+dir[0], cell.Col+dir[1]

			// Check if the next cell is outside the grid
			if !grid.IsInsideBounds(newRow, newCol) {
				continue
			}

			// If the next cell is a valid move, add it to the priority queue
			heap.Push(prioQueue, pq.Cell{Row: newRow, Col: newCol, Cost: cell.Cost + 1})
		}
	}

	return 0
}
