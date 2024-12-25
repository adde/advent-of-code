package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
)

const (
	COLUMN = '#'
)

var (
	locks []g.Grid
	keys  []g.Grid
)

func main() {
	startTime := time.Now()
	data := u.ReadAll("input.txt")
	parseInput(data)

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two: -")
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	sum := 0

	for _, lock := range locks {
		for _, key := range keys {
			if keyFitsLock(key, lock) {
				sum++
			}
		}
	}

	return sum
}

func keyFitsLock(key, lock g.Grid) bool {
	keyColumns := getKeyColumns(key)
	lockColumns := getLockColumns(lock)

	for i := range keyColumns {
		if keyColumns[i]+lockColumns[i] >= key.RowLen()-1 {
			return false
		}
	}

	return true
}

func getLockColumns(lock g.Grid) []int {
	columns := []int{}

	for c := range lock[0] {
		columnHeight := 0

		for r := 1; r < lock.RowLen(); r++ {
			if lock.Get(r, c) == COLUMN {
				columnHeight++
			} else {
				break
			}
		}

		columns = append(columns, columnHeight)
	}

	return columns
}

func getKeyColumns(key g.Grid) []int {
	columns := []int{}

	for c := range key[0] {
		columnHeight := 0

		for r := key.RowLen() - 2; r >= 0; r-- {
			if key.Get(r, c) == COLUMN {
				columnHeight++
			} else {
				break
			}
		}

		columns = append(columns, columnHeight)
	}

	return columns
}

func isLock(grid g.Grid) bool {
	for c := range grid[0] {
		if grid.Get(0, c) != COLUMN {
			return false
		}
	}

	return true
}

func parseInput(data string) {
	sections := strings.Split(data, "\n\n")

	for _, section := range sections {
		lines := strings.Split(section, "\n")
		grid := g.CreateFromLines(lines)

		if isLock(grid) {
			locks = append(locks, grid)
		} else {
			keys = append(keys, grid)
		}
	}
}
