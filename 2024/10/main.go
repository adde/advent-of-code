package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	TRAIL_HEAD = '0'
	TRAIL_END  = '9'
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)
	p1, p2 := getScoresAndRatings(grid)

	fmt.Println("\nPart one:", p1)
	fmt.Println("Part two:", p2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func getScoresAndRatings(grid g.Grid) (int, int) {
	scores := 0
	ratings := 0

	directions := [][2]int{
		{0, 1},  // Right
		{1, 0},  // Down
		{0, -1}, // Left
		{-1, 0}, // Up
	}
	trailHeads := findTrailHeads(grid)

	// Run BFS from each trail head to find all possible paths
	for _, trailHead := range trailHeads {
		queue := []g.Point{trailHead}
		visited := map[g.Point]bool{}
		canReach := set.New[g.Point]()

		rating := 0

		for len(queue) > 0 {
			current := queue[0]

			if grid[current.Row][current.Col] == TRAIL_END {
				// Keep track of which distinct trail heads we can reach
				canReach.Add(current)
				// Keep track of all possible paths for this trail head
				rating++
			}

			visited[current] = true
			queue = queue[1:]

			for _, dir := range directions {
				newRow := current.Row + dir[0]
				newCol := current.Col + dir[1]
				newPoint := g.Point{Row: newRow, Col: newCol}

				if grid.IsInsideBounds(newRow, newCol) &&
					!visited[newPoint] &&
					checkIfPointIsValid(grid[newRow][newCol], grid[current.Row][current.Col]) {
					queue = append(queue, newPoint)
				}
			}
		}

		scores += len(canReach)
		ratings += rating
	}

	return scores, ratings
}

func checkIfPointIsValid(newPoint rune, currentPoint rune) bool {
	return u.ToInt(string(newPoint)) == u.ToInt(string(currentPoint))+1
}

func findTrailHeads(grid [][]rune) []g.Point {
	trailHeads := []g.Point{}

	for row := range grid {
		for col := range grid[row] {
			if grid[row][col] == TRAIL_HEAD {
				trailHeads = append(trailHeads, g.Point{Row: row, Col: col})
			}
		}
	}

	return trailHeads
}
