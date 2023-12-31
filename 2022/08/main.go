package main

import (
	"fmt"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Tree struct {
	Row, Col  int
	Direction [][2]int
	Score     int
}

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	grid := parseInput(lines)

	fmt.Println("\nVisible trees:", getVisibleTrees(grid))
	fmt.Println("Highest scenic score:", getHighestScenicScore(grid))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) [][]rune {
	grid := [][]rune{}

	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	return grid
}

func getVisibleTrees(grid [][]rune) int {
	visible := 0

	for r, line := range grid {
		for c := range line {
			if isVisible(grid, r, c) {
				visible++
			}
		}
	}

	return visible
}

func isVisible(grid [][]rune, r, c int) bool {
	// If the tree is on the boundary, is is visible
	if isInBoundary(r, c, grid) {
		return true
	}

	treeHeight := utils.ToInt(string(grid[r][c]))

	// Run BFS to check if this tree is visible from the boundary
	queue := []Tree{{r, c, [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}, 0}}

	for len(queue) > 0 {
		tree := queue[0]
		queue = queue[1:]

		// If we are on the boundary, we can see the tree
		if isInBoundary(tree.Row, tree.Col, grid) {
			return true
		}

		for _, d := range tree.Direction {
			newRow, newCol := tree.Row+d[0], tree.Col+d[1]
			nextHeight := utils.ToInt(string(grid[newRow][newCol]))

			// If the next tree is lower than the current tree, continue in the same direction
			if nextHeight < treeHeight {
				queue = append(queue, Tree{newRow, newCol, [][2]int{d}, 0})
			}
		}
	}

	return false
}

func getHighestScenicScore(grid [][]rune) int {
	score := 0

	for r, line := range grid {
		for c := range line {
			score = max(score, calculateScenicScore(grid, r, c))
		}
	}

	return score
}

func calculateScenicScore(grid [][]rune, r, c int) int {
	scoreAxis := map[[2]int]int{}
	treeHeight := utils.ToInt(string(grid[r][c]))

	// Run BFS to check how many trees this tree can see
	queue := []Tree{{r, c, [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}, 0}}

	for len(queue) > 0 {
		tree := queue[0]
		queue = queue[1:]

		// If we are on the boundary, collect the score for this axis
		if isInBoundary(tree.Row, tree.Col, grid) {
			scoreAxis[tree.Direction[0]] = tree.Score
			continue
		}

		for _, d := range tree.Direction {
			newRow, newCol := tree.Row+d[0], tree.Col+d[1]
			nextHeight := utils.ToInt(string(grid[newRow][newCol]))

			// If the next tree is lower than the current tree, continue in the same direction
			// Otherwise, this axis is dead and we can add the score
			if nextHeight < treeHeight {
				queue = append(queue, Tree{newRow, newCol, [][2]int{d}, tree.Score + 1})
			} else {
				scoreAxis[d] = tree.Score + 1
			}
		}
	}

	// Calculate the final score for this tree
	score := 1
	for _, s := range scoreAxis {
		score *= s
	}

	return score
}

func isInBoundary(r, c int, grid [][]rune) bool {
	return r == 0 || r == len(grid)-1 || c == 0 || c == len(grid[r])-1
}
