package main

import (
	"fmt"
	"math"
	"sort"
	"time"

	u "github.com/adde/advent-of-code/utils"
	"github.com/adde/advent-of-code/utils/unionfind"
)

type Point struct {
	X, Y, Z float64
}

type Distance struct {
	Dist float64
	BoxA int
	BoxB int
}

const (
	NO_OF_PAIRS = 1000
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	boxes := make([]Point, 0)

	for _, line := range lines {
		numbers := u.GetIntsFromString(line, true)
		boxes = append(boxes, Point{
			X: float64(numbers[0]),
			Y: float64(numbers[1]),
			Z: float64(numbers[2]),
		})
	}

	distances := getDistances(boxes)

	fmt.Println("\nPart one:", partOne(boxes, distances))
	fmt.Println("Part two:", partTwo(boxes, distances))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(boxes []Point, distances []Distance) int {
	// Use Union-Find to group boxes based on closest distances
	boxesCount := len(boxes)
	uf := unionfind.New(boxesCount)

	for i := 0; i < NO_OF_PAIRS; i++ {
		uf.Union(distances[i].BoxA, distances[i].BoxB)
	}

	// Get circuit sizes
	circuits := uf.GetComponents()
	circuitSizes := make([]int, 0)
	for _, boxes := range circuits {
		circuitSizes = append(circuitSizes, len(boxes))
	}

	// Find the three largest circuits and multiply their sizes
	sort.Sort(sort.Reverse(sort.IntSlice(circuitSizes)))
	ans := 1
	for i := 0; i < 3; i++ {
		ans *= circuitSizes[i]
	}

	return ans
}

func partTwo(boxes []Point, distances []Distance) int {
	boxesCount := len(boxes)
	uf := unionfind.New(boxesCount)

	for i := 0; i < len(distances); i++ {
		uf.Union(distances[i].BoxA, distances[i].BoxB)

		if len(uf.GetComponents()) == 1 {
			return int(boxes[distances[i].BoxA].X) * int(boxes[distances[i].BoxB].X)
		}
	}

	return 0
}

func getDistances(boxes []Point) []Distance {
	distances := make([]Distance, 0)

	for i, boxA := range boxes {
		for j, boxB := range boxes {
			if i <= j {
				continue
			}

			dist := GetEuclideanDistance(boxA, boxB)
			distances = append(distances, Distance{
				Dist: dist,
				BoxA: i,
				BoxB: j,
			})
		}
	}

	// Sort distances in ascending order(closest first)
	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Dist < distances[j].Dist
	})

	return distances
}

func GetEuclideanDistance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	dz := p2.Z - p1.Z

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
