package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Reindeer struct {
	name     string
	speed    int
	duration int
	rest     int
	total    int
}

var reindeers []Reindeer

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		parts := strings.Split(line, " ")
		numbers := u.GetIntsFromString(line, true)

		reindeers = append(reindeers, Reindeer{
			name:     parts[0],
			speed:    numbers[0],
			duration: numbers[1],
			rest:     numbers[2],
			total:    numbers[1] + numbers[2],
		})
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	var distances = make(map[string]int)

	for _, r := range reindeers {
		second := 0
		for second < 2503 {
			duration := 0

			if second+r.duration > 2503 {
				duration = 2503 - second
			} else {
				duration = r.duration
			}

			distances[r.name] += r.speed * duration
			second += r.duration
			second += r.rest
		}
	}

	maxDistance := 0
	for _, v := range distances {
		if v > maxDistance {
			maxDistance = v
		}
	}

	return maxDistance
}

func partTwo() int {
	points := make(map[string]int)

	// Calculate points for each second
	for i := 1; i <= 2503; i++ {
		distances := make(map[string]int)
		maxDistance := 0

		// Calculate distances for all deers at current time
		for _, d := range reindeers {
			fullCycles := math.Floor(float64(i) / float64(d.total))
			remainingTime := math.Min(float64(i%d.total), float64(d.duration))

			distance := d.speed * (int(fullCycles)*d.duration + int(remainingTime))
			distances[d.name] = distance

			if distance > maxDistance {
				maxDistance = distance
			}
		}

		// Award points to deer with max distance
		for name, dist := range distances {
			if dist == maxDistance {
				points[name]++
			}
		}
	}

	maxPoints := 0
	for _, p := range points {
		if p > maxPoints {
			maxPoints = p
		}
	}

	return maxPoints
}
