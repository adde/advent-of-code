package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

type Region struct {
	Character rune
	Area      int
	Perimeter int
	Sides     int
	Cells     map[[2]int]bool
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)
	regions := findRegions(grid)

	sumP1 := 0
	for _, region := range regions {
		sumP1 += region.Area * region.Perimeter
	}

	sumP2 := 0
	for _, region := range regions {
		sumP2 += region.Area * region.Sides
	}

	fmt.Println("\nPart one:", sumP1)
	fmt.Println("Part two:", sumP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Identifies and analyzes all unique regions in the grid
func findRegions(grid g.Grid) []Region {
	var regions []Region
	rows, cols := len(grid), len(grid[0])
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Explore each cell in the grid
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if !visited[r][c] {
				char := grid[r][c]
				region := Region{
					Character: char,
					Cells:     make(map[[2]int]bool),
				}

				dfs(grid, visited, r, c, char, &region)

				// Only add non-empty regions
				if region.Area > 0 {
					// Get the number of sides of the region for part two
					region.Sides = countRegionSides(region)
					regions = append(regions, region)
				}
			}
		}
	}

	return regions
}

// Performs depth-first search to find all connected cells of the same character
func dfs(grid g.Grid, visited [][]bool, r, c int, target rune, region *Region) {
	// Check bounds and matching character
	if !grid.IsInsideBounds(r, c) ||
		visited[r][c] || grid[r][c] != target {
		return
	}

	// Mark as visited and add to region
	visited[r][c] = true
	region.Cells[[2]int{r, c}] = true
	region.Area++

	// Calculate perimeter by checking neighbors
	for _, dir := range [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		nr, nc := r+dir[0], c+dir[1]

		// Out of bounds or different character means boundary
		if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) ||
			grid[nr][nc] != target {
			region.Perimeter++
			continue
		}

		// Explore neighboring cells
		if !visited[nr][nc] {
			dfs(grid, visited, nr, nc, target, region)
		}
	}
}

// Counts the number of sides of a region
func countRegionSides(region Region) int {
	sides := 0
	possibleCorners := set.New[[2]float64]()

	// Get all possible corners for the cells in the region
	for cell := range region.Cells {
		r, c := cell[0], cell[1]

		for _, corner := range getCornerPositions(float64(r), float64(c)) {
			possibleCorners.Add([2]float64{corner[0], corner[1]})
		}
	}

	for corner := range possibleCorners {
		cr, cc := corner[0], corner[1]
		cellsInRegion := []bool{}
		nrOfCellsInRegion := 0

		// Check if surrounding cells are in the region
		for _, cell := range getCornerPositions(cr, cc) {
			sr, sc := int(cell[0]), int(cell[1])
			_, ok := region.Cells[[2]int{sr, sc}]

			cellsInRegion = append(cellsInRegion, ok)
			if ok {
				nrOfCellsInRegion++
			}
		}

		// Check how many cells in the region are adjacent to the corner
		// If 1 or 3, we have one side
		// If 2, we have two sides if the cells are in opposite corners
		if nrOfCellsInRegion == 1 || nrOfCellsInRegion == 3 {
			sides++
		} else if nrOfCellsInRegion == 2 {
			if cellsInRegion[0] && !cellsInRegion[1] &&
				!cellsInRegion[2] && cellsInRegion[3] ||
				!cellsInRegion[0] && cellsInRegion[1] &&
					cellsInRegion[2] && !cellsInRegion[3] {
				sides += 2
			}
		}
	}

	return sides
}

// Get the corner positions of a cell
func getCornerPositions(r, c float64) [][2]float64 {
	return [][2]float64{
		{r - 0.5, c - 0.5},
		{r + 0.5, c - 0.5},
		{r - 0.5, c + 0.5},
		{r + 0.5, c + 0.5},
	}
}
