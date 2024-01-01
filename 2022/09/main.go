package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Point struct {
	X, Y int
}

type Move struct {
	Direction Point
	Steps     int
}

var directions = map[string]Point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	moves := parseInput(lines)

	fmt.Println("\nPositions visited by the tail(p1):", getTailPositions(moves, 2))
	fmt.Println("Positions visited by the tail(p2):", getTailPositions(moves, 10))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Move {
	moves := make([]Move, 0)

	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		moves = append(moves, Move{directions[lineParts[0]], utils.ToInt(lineParts[1])})
	}

	return moves
}

func getTailPositions(moves []Move, ropeLength int) int {
	knots := []Point{}
	for i := 0; i < ropeLength; i++ {
		knots = append(knots, Point{X: 0, Y: 0})
	}

	visited := make(map[Point]bool)
	visited[knots[len(knots)-1]] = true

	for _, move := range moves {
		for i := 0; i < move.Steps; i++ {
			// Move the head
			knots[0].Y += move.Direction.Y
			knots[0].X += move.Direction.X

			// Check if knots are more than one step away from the previous knot
			// If so, move it one step closer
			for k := 1; k < len(knots); k++ {
				distanceX := knots[k-1].X - knots[k].X
				distanceY := knots[k-1].Y - knots[k].Y

				if math.Abs(float64(distanceX)) > 1 || math.Abs(float64(distanceY)) > 1 {
					knots[k].X += utils.Sign(distanceX)
					knots[k].Y += utils.Sign(distanceY)
				}
			}

			// Record tail positions
			visited[knots[len(knots)-1]] = true
		}
	}

	return len(visited)
}
