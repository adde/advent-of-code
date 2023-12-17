package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
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
	var grid [][]int

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
		grid = append(grid, make([]int, len(line)))

		for c, v := range line {
			grid[r][c] = utils.ToInt(string(v))
		}

		r++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nThe shortest distance to target:", getShortestDistance(grid, len(grid[0])-1, len(grid)-1))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getShortestDistance(grid [][]int, targetX, targetY int) int {
	seen := make(map[string]int)
	rowLen, colLen := len(grid), len(grid[0])

	// Start at the top left corner
	pq := &PriorityQueue{{0, 0, grid[0][0]}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// If we reached the target, just return the distance
		if cell.Row == targetY && cell.Col == targetX {
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
			if r >= 0 && r < rowLen && c >= 0 && c < colLen {
				heap.Push(pq, Cell{r, c, cell.Dist + grid[r][c]})
			}
		}
	}

	return 0
}
