package main

import (
	"fmt"
	"slices"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/set"
)

const (
	GRID_ROWS      = 103
	GRID_COLS      = 101
	SECONDS_TO_RUN = 100
)

type Point struct {
	x, y int
}

type Robot struct {
	position, velocity Point
}

func (r *Robot) Move(gridRows, gridCols int) {
	newX := r.position.x + r.velocity.x
	newY := r.position.y + r.velocity.y

	// if the new position is outside the grid, teleport the robot to the opposite side
	// Use the velocity to determine which side to teleport to and how much to move
	if newY < 0 {
		newY = gridRows + newY
	} else if newY >= gridRows {
		newY = newY - gridRows
	}

	if newX < 0 {
		newX = gridCols + newX
	} else if newX >= gridCols {
		newX = newX - gridCols
	}

	r.position.x = newX
	r.position.y = newY
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	robots := make([]Robot, 0)

	for _, line := range lines {
		numbers := u.GetIntsFromString(line, true)
		robots = append(robots, Robot{
			Point{numbers[0], numbers[1]},
			Point{numbers[2], numbers[3]},
		})
	}

	partOne := partOne(robots, GRID_ROWS, GRID_COLS, SECONDS_TO_RUN)
	partTwo := partTwo(robots, GRID_ROWS, GRID_COLS, SECONDS_TO_RUN*100, false) // Switch false to true to print the easter egg :)

	fmt.Println("\nPart one:", partOne)
	fmt.Println("Part two:", partTwo)

	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(robots []Robot, rows, cols, secondsToRun int) int {
	robots = slices.Clone(robots)

	// Move the robots for the specified number of seconds
	for s := 0; s < secondsToRun; s++ {
		for i := range robots {
			robots[i].Move(rows, cols)
		}
	}

	// Count the number of robots in each of the four quadrants
	// not including the center lines
	tl, tr, bl, br := 0, 0, 0, 0
	for _, robot := range robots {
		// Check if the robot is in the top left quadrant
		if robot.position.x < cols/2 && robot.position.y < rows/2 {
			tl++
		}

		// Check if the robot is in the top right quadrant
		if robot.position.x > cols/2 && robot.position.y < rows/2 {
			tr++
		}

		// Check if the robot is in the bottom left quadrant
		if robot.position.x < cols/2 && robot.position.y > rows/2 {
			bl++
		}

		// Check if the robot is in the bottom right quadrant
		if robot.position.x > cols/2 && robot.position.y > rows/2 {
			br++
		}
	}

	return tl * tr * bl * br
}

func partTwo(robots []Robot, rows, cols, secondsToRun int, isPrint bool) int {
	robots = slices.Clone(robots)
	grid := g.Create(rows, cols, '.')

	for s := 0; s < secondsToRun; s++ {
		seen := set.New[Point]()

		if isPrint {
			grid.Clear('.')
		}

		// Move the robots and update the grid
		for i := range robots {
			robots[i].Move(rows, cols)
			x, y := robots[i].position.x, robots[i].position.y

			seen.Add(Point{x, y})

			if isPrint {
				grid.Set(y, x, '#')
			}
		}

		// Check if all robots have unique positions
		// If they do, we have (most likely) found the easter egg
		if len(seen) == len(robots) {
			if isPrint {
				fmt.Print("\nFound the easter egg after ", s+1, " seconds!\n\n")
				grid.Print()
			}

			return s + 1
		}
	}

	return 0
}
