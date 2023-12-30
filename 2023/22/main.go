package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Vector struct {
	X, Y, Z int
}

type Brick struct {
	ID        string
	Start     Vector
	End       Vector
	Supports  map[string]bool
	Supported map[string]bool
}

func main() {
	startTime := time.Now()
	bricks := parseInput(utils.ReadLines("input.txt"))

	fmt.Println("\nDisintegrated bricks:", getDisintegrated(getSettled(bricks)))
	fmt.Println("Sum of fallen bricks:", getFallenBricks(getSettled(bricks)))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Brick {
	bricks := make([]Brick, 0)

	for _, line := range lines {
		lineParts := strings.Split(line, "~")
		startCoords := strings.Split(lineParts[0], ",")
		endCoords := strings.Split(lineParts[1], ",")

		bricks = append(bricks, Brick{
			ID: line,
			Start: Vector{
				X: utils.ToInt(startCoords[0]),
				Y: utils.ToInt(startCoords[1]),
				Z: utils.ToInt(startCoords[2]),
			},
			End: Vector{
				X: utils.ToInt(endCoords[0]),
				Y: utils.ToInt(endCoords[1]),
				Z: utils.ToInt(endCoords[2]),
			},
			Supports:  make(map[string]bool),
			Supported: make(map[string]bool),
		})
	}

	return bricks
}

func getSettled(bricks []Brick) []Brick {
	// Sort the bricks by Z coordinate ascending
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})

	// Let all bricks fall down until they settle,
	// the brick is settled when it hits another brick or the floor
	for i, brick := range bricks {
		// Check if the brick can fall down without hitting another brick
		newZ := 1
		for _, brick2 := range bricks[:i] {
			if isBrickOverlap(brick, brick2) {
				newZ = max(newZ, brick2.End.Z+1)
			}
		}

		bricks[i].Start.Z = newZ
		bricks[i].End.Z = newZ + brick.End.Z - brick.Start.Z
	}

	// Sort bricks again since they might have changed Z coordinate
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].Start.Z < bricks[j].Start.Z
	})

	return bricks
}

func getDisintegrated(settled []Brick) int {
	disintegrated := 0

	// Check if bricks are supporting each other
	for i, brick := range settled {
		for j, brick2 := range settled[:i] {
			if isBrickOverlap(brick, brick2) && brick.Start.Z == brick2.End.Z+1 {
				settled[j].Supports[brick.ID] = true
				settled[i].Supported[brick2.ID] = true
			}
		}
	}

	for _, brick := range settled {
		canDisintegrate := true

		// Check which bricks this brick is supporting
		for brickId := range brick.Supports {
			// If the supported brick is not supported by any other brick
			// then this brick can't be disintegrated
			if len(getBrick(brickId, settled).Supported) < 2 {
				canDisintegrate = false
				break
			}
		}

		// A brick can only be disintegrated
		// if it's not supporting a brick alone
		if canDisintegrate {
			disintegrated++
		}
	}

	return disintegrated
}

func getFallenBricks(settled []Brick) int {
	fallen := 0
	falling := make(map[string]bool)

	// Check if bricks are supporting each other
	for i, brick := range settled {
		for j, brick2 := range settled[:i] {
			if isBrickOverlap(brick, brick2) && brick.Start.Z == brick2.End.Z+1 {
				settled[j].Supports[brick.ID] = true
				settled[i].Supported[brick2.ID] = true
			}
		}
	}

	// Check how many bricks that would fall if we remove a brick
	for _, brick := range settled {
		q := make([]string, 0)

		// Add all bricks that are supported by only this brick to the queue
		for j := range brick.Supports {
			if len(getBrick(j, settled).Supported) == 1 {
				q = append(q, j)
			}
		}

		// Add the current brick to the queue and mark it as falling
		q = append(q, brick.ID)
		falling = make(map[string]bool)
		falling[brick.ID] = true

		// BFS to find all bricks that would fall
		for len(q) > 0 {
			brickId := q[0]
			q = q[1:]

			// Add all bricks that are supported by only falling bricks to the queue
			for b1 := range getBrick(brickId, settled).Supports {
				if !falling[b1] {
					intersection := true

					for b2 := range getBrick(b1, settled).Supported {
						if !falling[b2] {
							intersection = false
							break
						}
					}

					if intersection {
						q = append(q, b1)
						falling[b1] = true
					}
				}
			}
		}

		// Subtract 1 since the current brick is included in the count
		fallen += len(falling) - 1
	}

	return fallen
}

func getBrick(ID string, bricks []Brick) *Brick {
	for i := range bricks {
		if bricks[i].ID == ID {
			return &bricks[i]
		}
	}

	return nil
}

func isBrickOverlap(b1, b2 Brick) bool {
	return max(b1.Start.X, b2.Start.X) <= min(b1.End.X, b2.End.X) &&
		max(b1.Start.Y, b2.Start.Y) <= min(b1.End.Y, b2.End.Y)
}
