package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	EMPTY = '.'
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	grid := g.CreateFromLines(lines)
	antennas := getAntennas(grid)

	fmt.Println("\nPart one:", partOne(grid, antennas))
	fmt.Println("Part two:", partTwo(grid, antennas))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(grid [][]rune, antennas map[rune][][2]int) int {
	sum := 0

	antinodes := getAntinodes(grid, antennas, false)

	for an := range antinodes {
		if g.IsInsideBounds(grid, an[0], an[1]) {
			sum++
		}
	}

	return sum
}

func partTwo(grid [][]rune, antennas map[rune][][2]int) int {
	return len(getAntinodes(grid, antennas, true))
}

func getAntennas(grid [][]rune) map[rune][][2]int {
	antennas := make(map[rune][][2]int)

	// Find all antennas and add them to a map grouped by frequency
	for r, row := range grid {
		for c, char := range row {
			if char != EMPTY {
				if _, ok := antennas[char]; !ok {
					antennas[char] = make([][2]int, 0)
				}

				antennas[char] = append(antennas[char], [2]int{r, c})
			}
		}
	}

	return antennas
}

func getAntinodes(grid [][]rune, antennas map[rune][][2]int, useResonantHarmonics bool) set.Set[[2]int] {
	antinodes := set.New[[2]int]()

	// Find all antinodes and add them ta a set
	for _, antenna := range antennas {
		for i := 0; i < len(antenna); i++ {
			for j := 0; j < len(antenna); j++ {
				if i == j {
					continue
				}

				// Get the coordinates for the two antennas
				r1, c1 := antenna[i][0], antenna[i][1]
				r2, c2 := antenna[j][0], antenna[j][1]

				// Add all possible antinodes along the direction vector(dr, dc) (part two)
				if useResonantHarmonics {
					// Calculate the distance between the two antennas
					dr := r2 - r1
					dc := c2 - c1

					r := r1
					c := c1

					for g.IsInsideBounds(grid, r, c) {
						antinodes.Add([2]int{r, c})
						r += dr
						c += dc
					}
				} else {
					// Add the antinode by calculating a symmetric point (part one)
					antinodes.Add([2]int{2*r1 - r2, 2*c1 - c2})
				}
			}
		}
	}

	return antinodes
}
