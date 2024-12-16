package main

import (
	"container/heap"
	"fmt"
	"math"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	START   = 'S'
	END     = 'E'
	PEN_90  = 1000
	PEN_180 = 2000
)

type Cell struct {
	Row, Col, Score int
	Dir             [2]int
	PrevCell        [4]int
}

type PriorityQueue []Cell

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Score < pq[j].Score
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

var directions = [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)
	startR, startC := findGridPos(grid, START)

	fmt.Println("\nPart one:", partOne(grid, startR, startC))
	fmt.Println("Part two:", partTwo(grid, startR, startC))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Dijkstra's algorithm to find the path with the best score(low cost) to the target
func partOne(grid g.Grid, startR, startC int) int {
	pq := &PriorityQueue{{
		Row:      startR,
		Col:      startC,
		Score:    0,
		Dir:      [2]int{0, 1},
		PrevCell: [4]int{-1, -1, 0, 0},
	}}
	heap.Init(pq)
	seen := set.New[[4]int]()

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)

		// Add the current cell to the seen set
		key := [4]int{cell.Row, cell.Col, cell.Dir[0], cell.Dir[1]}
		seen.Add(key)

		// If we have reached the target, return the score
		if grid.Get(cell.Row, cell.Col) == END {
			return cell.Score
		}

		// Turn in all four directions
		for _, dir := range directions {
			newRow, newCol, dirRow := cell.Row+dir[0], cell.Col+dir[1], cell.Dir[0]
			newDirRow, newDirCol := dir[0], dir[1]

			// If the next cell is not a valid move, skip it
			if grid.Get(newRow, newCol) == '#' {
				continue
			}

			// If we have already visited the next cell, skip it
			newKey := [4]int{newRow, newCol, newDirRow, newDirCol}
			if seen.Contains(newKey) {
				continue
			}

			// If we turn 90 degrees, add 1000 to the score
			score := cell.Score + 1
			if dir != cell.Dir && dirRow != newDirRow {
				score += PEN_90
			}

			heap.Push(pq, Cell{newRow, newCol, score, dir, cell.PrevCell})
		}
	}

	return 0
}

// Dijkstra's algorithm to find all paths with the best score(low cost) to the target
// and then backtrack from the target to find all unique tiles visited by the best paths
func partTwo(grid g.Grid, startR, startC int) int {
	pq := &PriorityQueue{{
		Row:      startR,
		Col:      startC,
		Score:    0,
		Dir:      [2]int{0, 1},
		PrevCell: [4]int{-1, -1, 0, 0},
	}}
	heap.Init(pq)
	lowestScore := make(map[[4]int]int)           // Store the lowest score for each cell state
	bestScore := math.MaxInt64                    // Store the best score for the target
	backtrack := make(map[[4]int]map[[4]int]bool) // Store the previous cell for each cell state
	targetStates := set.New[[4]int]()             // Store the cell states that reach the target

	for pq.Len() > 0 {
		cell := heap.Pop(pq).(Cell)
		key := [4]int{cell.Row, cell.Col, cell.Dir[0], cell.Dir[1]}

		// If we have already visited the next cell with a lower score, skip it
		if cell.Score > getFromMap(lowestScore, key) {
			continue
		}

		// Add the current cell state and score to lowest score map
		lowestScore[key] = cell.Score

		// If we have reached the target, add the cell state to the target states
		if grid.Get(cell.Row, cell.Col) == END {
			// If the score is higher than the best score,
			// we can break since we will not find a better path
			if cell.Score > bestScore {
				break
			}
			bestScore = cell.Score
			targetStates.Add(key)
		}

		// If cell is not in backtrack, add it
		if _, ok := backtrack[key]; !ok {
			backtrack[key] = make(map[[4]int]bool)
		}
		prevCellKey := [4]int{cell.PrevCell[0], cell.PrevCell[1], cell.PrevCell[2], cell.PrevCell[3]}
		backtrack[key][prevCellKey] = true

		// Turn in all four directions
		for _, dir := range directions {
			newRow, newCol, dirRow := cell.Row+dir[0], cell.Col+dir[1], cell.Dir[0]
			newDirRow, newDirCol := dir[0], dir[1]

			// If the next cell is not a valid move, skip it
			if grid.Get(newRow, newCol) == '#' {
				continue
			}

			// If we already have a lower score for the next cell, skip it
			newKey := [4]int{newRow, newCol, newDirRow, newDirCol}
			if cell.Score > getFromMap(lowestScore, newKey) {
				continue
			}

			// If we turn 90 degrees, add 1000 to the score
			score := cell.Score + 1
			if dir != cell.Dir && dirRow != newDirRow {
				score += PEN_90
			}

			// Add the new cell to the priority queue, with the current cell as the previous cell
			heap.Push(pq, Cell{newRow, newCol, score, dir, key})
		}
	}

	// Find all tiles visited by the best paths by backtracking from the target states using BFS
	// Start by adding the target states to the queue
	tiles := make([][4]int, 0)
	seen := set.New[[4]int]()
	for key := range targetStates {
		seen.Add(key)
		tiles = append(tiles, [4]int{key[0], key[1], key[2], key[3]})
	}

	for len(tiles) > 0 {
		tile := tiles[0]
		tiles = tiles[1:]

		// Add all the tiles that can reach the current tile
		for key := range backtrack[tile] {
			// Skip the start cell state
			if key == [4]int{-1, -1, 0, 0} {
				continue
			}
			if seen.Contains(key) {
				continue
			}

			seen.Add(key)
			tiles = append(tiles, [4]int{key[0], key[1], key[2], key[3]})
		}
	}

	// Finally, return the number of unique tiles visited for the best paths
	uniqueTiles := set.New[[2]int]()
	for key := range seen {
		uniqueTiles.Add([2]int{key[0], key[1]})
	}

	return len(uniqueTiles)
}

// Get int from map, if not found return max int
func getFromMap(m map[[4]int]int, key [4]int) int {
	if val, ok := m[key]; ok {
		return val
	}

	return math.MaxInt64
}

// Find the position of a character in a grid
func findGridPos(grid [][]rune, char rune) (int, int) {
	for r, row := range grid {
		for c, val := range row {
			if val == char {
				return r, c
			}
		}
	}

	return -1, -1
}
