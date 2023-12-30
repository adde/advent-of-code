package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"time"
)

const (
	GALAXY_MULTIPLIER = 1000000 - 1
)

type Galaxy struct {
	y int
	x int
}

var grid []string
var galaxies []Galaxy

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Get shortest paths between all galaxies pairs and sum them up
	galaxies = getGalaxies()
	sumP1 := getGalaxyPathsSum(1) // For empty rows and cols, expand grid by 1 in all directions
	galaxies = getGalaxies()
	sumP2 := getGalaxyPathsSum(GALAXY_MULTIPLIER) // For empty rows and cols, expand grid by 1,000,000 in all directions

	fmt.Println("\nSum of path lengths(part one):", sumP1)
	fmt.Println("Sum of path lengths(part two):", sumP2)
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getGalaxyPathsSum(multiplier int) int {
	dist := make([][]int, len(galaxies))
	sum := 0

	// Set correct coordinates for galaxies based on provided multiplier
	for i := 0; i < len(galaxies); i++ {
		galaxies[i].x += multiplier * getEmptyCols(galaxies[i].x)
		galaxies[i].y += multiplier * getEmptyRows(galaxies[i].y)
	}

	// Initialize the adjacency matrix
	for i := range dist {
		dist[i] = make([]int, len(galaxies))
		for j := range dist[i] {
			if i != j {
				dist[i][j] = math.MaxInt32
			}
		}
	}

	// Calculate the distances between all pairs of galaxies
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i != j {
				dist[i][j] = abs(g1.x-g2.x) + abs(g1.y-g2.y)
			}
		}
	}

	// Use the Floyd-Warshall algorithm to find the shortest paths
	for k := range galaxies {
		for i := range galaxies {
			for j := range galaxies {
				if dist[i][j] > dist[i][k]+dist[k][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Sum up the distances for each galaxy pair
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += dist[i][j]
		}
	}

	return sum
}

func getGalaxies() []Galaxy {
	g := []Galaxy{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				g = append(g, Galaxy{y: i, x: j})
			}
		}
	}

	return g
}

func getEmptyRows(row int) int {
	offset := 0

	for i := 0; i < len(grid); i++ {
		hasGalaxy := false

		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			if i < row {
				offset++
			}
		}
	}

	return offset
}

func getEmptyCols(col int) int {
	offset := 0

	for j := 0; j < len(grid[0]); j++ {
		hasGalaxy := false

		for i := 0; i < len(grid); i++ {
			if grid[i][j] == '#' {
				hasGalaxy = true
			}
		}

		if !hasGalaxy {
			if j < col {
				offset++
			}
		}
	}

	return offset
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
