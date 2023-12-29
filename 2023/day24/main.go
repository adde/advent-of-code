package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

const (
	INTER_MIN = float64(200000000000000)
	INTER_MAX = float64(400000000000000)
	VEL_RANGE = 1000
)

type Vector struct {
	X, Y, Z int
}

type Hailstone struct {
	Position Vector
	Velocity Vector
}

type Set map[int]bool

func (s Set) Add(k int) {
	s[k] = true
}

func (s Set) Pop() int {
	for k := range s {
		delete(s, k)
		return k
	}

	return 0
}

func (s Set) GetIntersectingValues(t Set) Set {
	result := make(Set)

	for k := range s {
		if _, found := t[k]; found {
			result[k] = true
		}
	}

	if len(result) == 0 {
		return t
	}

	return result
}

func main() {
	startTime := time.Now()

	lines := utils.ReadInput("input.txt")
	hailstones := parseInput(lines)

	fmt.Println("\nSum of intersections(part one):", getIntersections(hailstones))
	fmt.Println("Single throw position(part two):", getSingleThrowPosition(hailstones))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Hailstone {
	hailstones := make([]Hailstone, 0)

	for _, line := range lines {

		lineParts := strings.Split(line, " @ ")
		position := strings.Split(lineParts[0], ", ")
		velocity := strings.Split(lineParts[1], ", ")

		hailstones = append(hailstones, Hailstone{
			Position: Vector{
				X: utils.ToInt(strings.TrimSpace(position[0])),
				Y: utils.ToInt(strings.TrimSpace(position[1])),
				Z: utils.ToInt(strings.TrimSpace(position[2])),
			},
			Velocity: Vector{
				X: utils.ToInt(strings.TrimSpace(velocity[0])),
				Y: utils.ToInt(strings.TrimSpace(velocity[1])),
				Z: utils.ToInt(strings.TrimSpace(velocity[2])),
			},
		})
	}

	return hailstones
}

func getIntersections(hailstones []Hailstone) int {
	intersections := 0

	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[i+1:] {
			// Calculate the line equation for the two hailstones
			ma := float64(hs1.Velocity.Y) / float64(hs1.Velocity.X)
			mb := float64(hs2.Velocity.Y) / float64(hs2.Velocity.X)
			ca := hs1.Position.Y - int(ma*float64(hs1.Position.X))
			cb := hs2.Position.Y - int(mb*float64(hs2.Position.X))

			// If the two hailstones have the same velocity, they will never intersect
			if ma == mb {
				continue
			}

			// Calculate the intersection point
			xPos := float64(cb-ca) / (ma - mb)
			yPos := ma*xPos + float64(ca)

			// If the intersection point is not between the two hailstones, skip
			if (xPos < float64(hs1.Position.X) && hs1.Velocity.X > 0) ||
				(xPos > float64(hs1.Position.X) && hs1.Velocity.X < 0) ||
				(xPos < float64(hs2.Position.X) && hs2.Velocity.X > 0) ||
				(xPos > float64(hs2.Position.X) && hs2.Velocity.X < 0) {
				continue
			}

			// If the intersection point is between MIN and MAX, add to intersections
			if INTER_MIN <= xPos && xPos <= INTER_MAX &&
				INTER_MIN <= yPos && yPos <= INTER_MAX {
				intersections++
			}
		}
	}

	return intersections
}

func getSingleThrowPosition(hailstones []Hailstone) int {
	// Sort Hailstones in ascending order by X position
	sort.Slice(hailstones, func(i, j int) bool {
		return hailstones[i].Position.X < hailstones[j].Position.X
	})

	// Find potential velocities for each axis
	potenialVelX, potenialVelY, potenialVelZ := Set{}, Set{}, Set{}
	for i, hs1 := range hailstones {
		for _, hs2 := range hailstones[i+1:] {
			getPotentialVelocities(hs1.Position.X, hs2.Position.X, hs1.Velocity.X, hs2.Velocity.X, &potenialVelX)
			getPotentialVelocities(hs1.Position.Y, hs2.Position.Y, hs1.Velocity.Y, hs2.Velocity.Y, &potenialVelY)
			getPotentialVelocities(hs1.Position.Z, hs2.Position.Z, hs1.Velocity.Z, hs2.Velocity.Z, &potenialVelZ)
		}
	}

	// Calculate the line equation for the single throw using the found velocities
	rVelX, rVelY, rVelZ := potenialVelX.Pop(), potenialVelY.Pop(), potenialVelZ.Pop()
	ma := float64(hailstones[0].Velocity.Y-rVelY) / float64(hailstones[0].Velocity.X-rVelX)
	mb := float64(hailstones[1].Velocity.Y-rVelY) / float64(hailstones[1].Velocity.X-rVelX)
	ca := hailstones[0].Position.Y - int(ma*float64(hailstones[0].Position.X))
	cb := hailstones[1].Position.Y - int(mb*float64(hailstones[1].Position.X))

	// Calculate the start position for the single throw
	xPos := int(float64(cb-ca) / (ma - mb))
	yPos := int(ma*float64(xPos) + float64(ca))
	time := (xPos - hailstones[0].Position.X) / (hailstones[0].Velocity.X - rVelX)
	zPos := hailstones[0].Position.Z + (hailstones[0].Velocity.Z-rVelZ)*time

	return xPos + yPos + zPos
}

func getPotentialVelocities(p1, p2, v1, v2 int, potentialVel *Set) {
	if v1 == v2 && math.Abs(float64(v1)) > 100 {
		var newSet = make(Set)
		diff := p2 - p1

		for v := -VEL_RANGE; v < VEL_RANGE; v++ {
			if v == v1 {
				continue
			}

			if diff%(v-v1) == 0 {
				newSet.Add(v)
			}
		}

		*potentialVel = potentialVel.GetIntersectingValues(newSet)
	}
}
