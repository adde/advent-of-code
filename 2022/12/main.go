package main

import (
	"container/heap"
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Cell struct {
	Row, Col, Dist int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Dist < pq[j].Dist
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Cell)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func main() {
	startTime := time.Now()

	// Populate the grid from input
	grid := parseInput(utils.ReadLines("input.txt"))
	start, end := getPos('S', grid), getPos('E', grid)

	fmt.Println("\nFewest steps to target(p1):", getShortestDistance(grid, start.Row, start.Col, false))
	fmt.Println("Fewest steps to target(p2):", getShortestDistance(grid, end.Row, end.Col, true))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) [][]rune {
	var grid [][]rune
	row := 0

	for _, line := range lines {
		grid = append(grid, make([]rune, len(line)))

		for col, val := range line {
			grid[row][col] = val
		}

		row++
	}

	return grid
}

func getShortestDistance(grid [][]rune, startR, startC int, reverse bool) int {
	seen := make(map[string]int)
	rowLen, colLen := len(grid), len(grid[0])

	target := getPos('E', grid)

	pq := &PriorityQueue{{startR, startC, 0}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// If we reached the target, just return the distance
		if (!reverse && cell.Row == target.Row && cell.Col == target.Col) ||
			(reverse && grid[cell.Row][cell.Col] == 'a') {
			return cell.Dist
		}

		// Check if we already visited this cell
		key := fmt.Sprintf("%d,%d", cell.Row, cell.Col)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = cell.Dist

		// Check all four directions
		for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
			r, c := cell.Row+dir[0], cell.Col+dir[1]

			// If the next cell is a valid move, add it to the priority queue
			if r >= 0 && r < rowLen && c >= 0 && c < colLen &&
				isValidMove(grid, cell.Row, cell.Col, r, c, reverse) {
				heap.Push(pq, Cell{r, c, cell.Dist + 1})
			}
		}
	}

	return 0
}

func getPos(char rune, grid [][]rune) Cell {
	pos := Cell{-1, -1, 0}

	for r, row := range grid {
		for c := range row {
			if grid[r][c] == char {
				pos.Row, pos.Col = r, c
				break
			}
		}
	}

	return pos
}

func isValidMove(grid [][]rune, row, col, newRow, newCol int, reverse bool) bool {
	a := grid[row][col]
	b := grid[newRow][newCol]

	if a == 'S' {
		a = 'a'
	}
	if a == 'E' {
		a = 'z'
	}
	if b == 'S' {
		b = 'a'
	}
	if b == 'E' {
		b = 'z'
	}

	return (!reverse && b-a <= 1) || (reverse && a-b <= 1)
}
