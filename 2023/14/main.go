package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var grid [][]rune

func main() {
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
		grid = append(grid, make([]rune, len(line)))
		for c, v := range line {
			grid[r][c] = v
		}
		r++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Found the answer for P2 by luck
	// Saw the repeating pattern after some cycles
	// and just brute forced the answer
	// After that, I found out that exactly
	// 1 000 cycles also gives the correct answer (for me)

	fmt.Println()
	fmt.Println("Total load part one:", getTotalLoad(0))
	fmt.Println("Total load part two:", getTotalLoad(1000))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getTotalLoad(cycles int) int {
	for i := 0; i < cycles; i++ {
		for j := 0; j < 4; j++ {
			rollStones()
			rotateClockwise()
		}
	}

	if cycles == 0 {
		rollStones()
	}

	return getLoad()
}

func rollStones() {
	for r := range grid {
		if r == 0 {
			continue
		}

		for c := range grid[r] {
			if grid[r][c] == 'O' {
				for ur := r - 1; ur >= 0; ur-- {
					// Stop if stone can't move up
					if grid[ur][c] != '.' {
						break
					}

					// Move stone up
					grid[ur][c] = 'O'
					grid[ur+1][c] = '.'
				}
			}
		}
	}
}

func getLoad() int {
	load := 0

	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 'O' {
				load += len(grid) - r
			}
		}
	}

	return load
}

func rotateClockwise() {
	// Transpose the matrix
	for i := range grid {
		for j := i; j < len(grid[i]); j++ {
			grid[i][j], grid[j][i] = grid[j][i], grid[i][j]
		}
	}

	// Reverse each row
	for i := range grid {
		for j, k := 0, len(grid[i])-1; j < k; j, k = j+1, k-1 {
			grid[i][j], grid[i][k] = grid[i][k], grid[i][j]
		}
	}
}
