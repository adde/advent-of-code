package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

var dirs = map[rune][2]int{
	'^': {-1, 0},
	'>': {0, 1},
	'v': {1, 0},
	'<': {0, -1},
}

func main() {
	startTime := time.Now()
	data := u.ReadAll("input.txt")

	fmt.Println("\nPart one:", partOne(data))
	fmt.Println("Part two:", partTwo(data))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(santaDirections string) int {
	row, col := 0, 0
	visited := make(map[[2]int]bool)
	visited[[2]int{row, col}] = true

	for _, dirChar := range santaDirections {
		row += dirs[dirChar][0]
		col += dirs[dirChar][1]
		visited[[2]int{row, col}] = true
	}

	return len(visited)
}

func partTwo(santaDirections string) int {
	santaRow, santaCol, roboRow, roboCol := 0, 0, 0, 0
	visited := make(map[[2]int]bool)
	visited[[2]int{santaRow, santaCol}] = true

	for i, dirChar := range santaDirections {
		if i%2 == 0 {
			santaRow += dirs[dirChar][0]
			santaCol += dirs[dirChar][1]
			visited[[2]int{santaRow, santaCol}] = true
		} else {
			roboRow += dirs[dirChar][0]
			roboCol += dirs[dirChar][1]
			visited[[2]int{roboRow, roboCol}] = true
		}
	}

	return len(visited)
}
