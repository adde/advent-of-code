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
	Row, Col, HeatLoss, Steps int
	CurrDir                   []int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].HeatLoss < pq[j].HeatLoss
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

	fmt.Println("\nLeast heat loss(part one):", getLeastHeatLoss(grid, false))
	fmt.Println("Least heat loss(part two):", getLeastHeatLoss(grid, true))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

// Use Dijkstra's algorithm to find the path with the least heat loss
// Starts from the top left corner and finishes in the bottom right corner
func getLeastHeatLoss(grid [][]int, useUltraCrucibles bool) int {
	seen := make(map[string]int)
	rowLen, colLen := len(grid), len(grid[0])

	// Initialize the priority queue with the top left corner
	pq := &PriorityQueue{{0, 0, 0, 0, []int{0, 1}}}
	heap.Init(pq)

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// If we have reached the bottom right corner, return the heat loss
		// Also check if we are using ultra crucibles (part two),
		// if we are, we can only end if we have moved atleast 4 steps
		if cell.Row == rowLen-1 && cell.Col == colLen-1 &&
			(useUltraCrucibles && cell.Steps >= 4 || !useUltraCrucibles) {
			return cell.HeatLoss
		}

		// Check if we have already visited this cell
		// with the same direction and steps
		key := fmt.Sprintf(
			"%d,%d,%d,%d,%d",
			cell.Row,
			cell.Col,
			cell.Steps,
			cell.CurrDir[0],
			cell.CurrDir[1],
		)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = cell.HeatLoss

		for _, dir := range getPossibleDirections(cell.CurrDir) {
			r, c := cell.Row+dir[0], cell.Col+dir[1]
			isNewDir := cell.CurrDir[0] != dir[0] || cell.CurrDir[1] != dir[1]

			// Reset the steps if we are moving in a new direction,
			// otherwise increase the current steps by one
			var steps int
			if isNewDir {
				steps = 1
			} else {
				steps = cell.Steps + 1
			}

			// Check if we are using ultra crucibles (part two)
			// If we are, we can move max 10 steps in the same direction
			// and we can only change direction if we have moved atleast 4 steps
			// Otherwise we can only move max 3 steps in the same direction (part one)
			var isValid bool
			if useUltraCrucibles {
				isValid = steps <= 10 && (cell.Steps >= 4 || cell.Steps == 0 || !isNewDir)
			} else {
				isValid = steps <= 3
			}

			// If the next cell is a valid move, then add it to the priority queue
			if r >= 0 && r < rowLen && c >= 0 && c < colLen && isValid {
				heap.Push(pq, Cell{r, c, cell.HeatLoss + grid[r][c], steps, dir})
			}
		}
	}

	return 0
}

// Get possible directions that we can move in,
// based on the current direction
func getPossibleDirections(dir []int) [][]int {
	var dirs [][]int

	switch {
	case dir[0] == 0 && dir[1] == 1: // right
		dirs = [][]int{{0, 1}, {-1, 0}, {1, 0}} // right, up, down
	case dir[0] == 1 && dir[1] == 0: // down
		dirs = [][]int{{1, 0}, {0, 1}, {0, -1}} // down, right, left
	case dir[0] == 0 && dir[1] == -1: // left
		dirs = [][]int{{0, -1}, {-1, 0}, {1, 0}} // left, up, down
	case dir[0] == -1 && dir[1] == 0: // up
		dirs = [][]int{{-1, 0}, {0, 1}, {0, -1}} // up, right, left
	}

	return dirs
}
