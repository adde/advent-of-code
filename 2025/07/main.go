package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	START    = 'S'
	SPLITTER = '^'
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	fmt.Println("\nPart one:", partOne(lines))
	fmt.Println("Part two:", partTwo(lines))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(lines []string) int {
	grid := g.CreateFromLines(lines)
	startR, startC := grid.Find(START)
	seen := set.New[[2]int]()

	return getTachyonBeamsSplits(grid, seen, startR, startC)
}

func partTwo(lines []string) int {
	grid := g.CreateFromLines(lines)
	startR, startC := grid.Find(START)
	cache := make(map[[2]int]int)

	return getTachyonParticleTimelines(grid, cache, startR, startC)
}

func getTachyonBeamsSplits(grid g.Grid, seen set.Set[[2]int], r, c int) int {
	if !grid.IsInsideBounds(r, c) || seen.Contains([2]int{r, c}) {
		return 0
	}

	seen.Add([2]int{r, c})

	if grid.Get(r, c) == SPLITTER {
		// Split the beam
		return getTachyonBeamsSplits(grid, seen, r+1, c-1) +
			getTachyonBeamsSplits(grid, seen, r+1, c+1) + 1
	} else {
		// Continue downwards
		return getTachyonBeamsSplits(grid, seen, r+1, c)
	}
}

func getTachyonParticleTimelines(grid g.Grid, cache map[[2]int]int, r, c int) int {
	if val, found := cache[[2]int{r, c}]; found {
		return val
	}

	if !grid.IsInsideBounds(r, c) {
		return 0
	}

	if r == grid.RowLen()-1 {
		return 1
	}

	if grid.Get(r, c) == SPLITTER {
		val := getTachyonParticleTimelines(grid, cache, r+1, c-1) +
			getTachyonParticleTimelines(grid, cache, r+1, c+1)
		cache[[2]int{r, c}] = val
		return val
	} else {
		val := getTachyonParticleTimelines(grid, cache, r+1, c)
		cache[[2]int{r, c}] = val
		return val
	}
}
