package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Point struct {
	x, y int
}

type Machine struct {
	A     Point
	B     Point
	Prize Point
}

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	sections := strings.Split(file, "\n\n")
	machines := make([]Machine, 0)

	for _, section := range sections {
		numbers := u.GetIntsFromString(section, false)

		machines = append(machines, Machine{
			A:     Point{x: numbers[0], y: numbers[1]},
			B:     Point{x: numbers[2], y: numbers[3]},
			Prize: Point{x: numbers[4], y: numbers[5]},
		})
	}

	fmt.Println("\nPart one:", getTokensSpent(machines, 0))
	fmt.Println("Part two:", getTokensSpent(machines, 1e13))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func getTokensSpent(machines []Machine, prizeIncrease int) int {
	sum := 0

	for _, m := range machines {
		a, b, prize := m.A, m.B, m.Prize
		ap, bp := float64(0), float64(0)

		// Increase prize for part two
		if prizeIncrease > 0 {
			prize.x += prizeIncrease
			prize.y += prizeIncrease
		}

		ap = float64(prize.x*b.y-prize.y*b.x) / float64(a.x*b.y-b.x*a.y)
		bp = float64(prize.y*a.x-prize.x*a.y) / float64(a.x*b.y-b.x*a.y)

		// Button presses must be whole numbers
		if math.Round(ap) != ap || math.Round(bp) != bp {
			continue
		}

		// Limit to 100 button presses for part one
		if (ap > 100 || bp > 100) && prizeIncrease == 0 {
			continue
		}

		sum += 3*int(ap) + int(bp)
	}

	return sum
}
