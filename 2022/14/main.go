package main

import (
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"github.com/adde/advent-of-code/utils"
	"github.com/nsf/termbox-go"
)

type Point struct {
	X, Y int
}

type Path struct {
	Points []Point
}

const (
	VIS_UPDATE_FREQ = 30
)

var (
	isDrawing  = true
	sands      map[Point]bool
	sandsMutex = sync.RWMutex{}
)

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	paths := parseInput(lines)

	// visualizeCave(paths)

	fmt.Println("\nSand come to rest(part one):", getUnitsOfSand(paths, false))
	fmt.Println("Sand come to rest(part two):", getUnitsOfSandWithFloor(paths))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Path {
	var paths []Path

	for _, line := range lines {
		path := Path{}
		lineParts := strings.Split(line, " -> ")

		for _, part := range lineParts {
			posParts := strings.Split(part, ",")
			path.Points = append(path.Points, Point{
				X: utils.ToInt(strings.TrimSpace(posParts[0])),
				Y: utils.ToInt(strings.TrimSpace(posParts[1])),
			})
		}

		paths = append(paths, path)
	}

	return paths
}

func getUnitsOfSand(paths []Path, visualize bool) int {
	sandResting := 0
	start := Point{X: 500, Y: 0}
	sand := start
	endY := getEndY(paths)
	rocks := getRockPositions(paths)

	for sand.Y < endY {
		sand = start
		falling := true

	outer:
		for falling && sand.Y < endY {
			// Check how long the sand can fall before it hits a rock
			// The sand can fall down, left or right one step at a time
			for _, dir := range []Point{{0, 1}, {-1, 1}, {1, 1}} {
				nx, ny := sand.X+dir.X, sand.Y+dir.Y

				if _, ok := rocks[Point{X: nx, Y: ny}]; !ok {
					sand.X, sand.Y = nx, ny
					continue outer
				}
			}

			// If the sand cannot fall down, left or right, it will stop falling
			sandResting++
			falling = false
			rocks[sand] = true

			if visualize {
				sandsMutex.Lock()
				sands[sand] = true
				sandsMutex.Unlock()
				time.Sleep(time.Millisecond * VIS_UPDATE_FREQ * 2)
			}
		}
	}

	return sandResting
}

func getUnitsOfSandWithFloor(paths []Path) int {
	sandResting := 0
	start := Point{X: 500, Y: 0}
	current := Point{X: -1, Y: -1}
	sand := start
	endY := getEndY(paths)
	rocks := getRockPositions(paths)

	for current.X != start.X || current.Y != start.Y {
		sand = start

	outer:
		for _, ok := rocks[start]; !ok && current != sand; {
			for _, dir := range []Point{{0, 1}, {-1, 1}, {1, 1}} {
				nx, ny := sand.X+dir.X, sand.Y+dir.Y

				if _, ok := rocks[Point{X: nx, Y: ny}]; !ok {
					sand.X, sand.Y = nx, ny

					if sand.Y < endY {
						continue outer
					} else {
						break
					}
				}
			}

			sandResting++
			current = sand
			rocks[sand] = true
		}
	}

	return sandResting
}

func getRockPositions(paths []Path) map[Point]bool {
	rocks := make(map[Point]bool)

	for _, path := range paths {
		points := path.Points

		for i := range points[:len(points)-1] {
			a, b := points[i], points[i+1]

			// Get all points between a and b
			for x := min(a.X, b.X); x <= max(a.X, b.X); x++ {
				for y := min(a.Y, b.Y); y <= max(a.Y, b.Y); y++ {
					rocks[Point{X: x, Y: y}] = true
				}
			}
		}
	}

	return rocks
}

func getEndY(paths []Path) int {
	maxY := 0

	for _, path := range paths {
		for _, point := range path.Points {
			maxY = max(maxY, point.Y)
		}
	}

	return maxY + 1
}

func printCave(paths []Path) {
	rocks := getRockPositions(paths)

	maxX, minX := 0, math.MaxInt32
	maxY, minY := 0, math.MaxInt32
	for r := range rocks {
		maxX = max(maxX, r.X)
		minX = min(minX, r.X)
		maxY = max(maxY, r.Y)
		minY = min(minY, r.Y)
	}

	// Scroll viewport when more sand is added
	diffY := 0
	if maxY > 50 && len(sands) > 125 {
		diffY = len(sands) / 6
	}

	// Print grid
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	sandsMutex.RLock()
	for y := minY - 3; y < maxY+3; y++ {
		for x := minX - 5; x < maxX+5; x++ {
			if _, ok := rocks[Point{X: x, Y: y}]; ok {
				termbox.SetCell(x-minX+5, y-minY+3-diffY, '#', termbox.ColorDarkGray, termbox.ColorDefault)
			} else if _, ok := sands[Point{X: x, Y: y}]; ok {
				termbox.SetCell(x-minX+5, y-minY+3-diffY, 'O', termbox.ColorLightYellow, termbox.ColorDefault)
			} else {
				termbox.SetCell(x-minX+5, y-minY+3-diffY, '.', termbox.ColorLightBlue, termbox.ColorDefault)
			}
		}
	}
	sandsMutex.RUnlock()
	termbox.Flush()
}

func visualizeCave(paths []Path) {
	termbox.Init()
	defer termbox.Close()

	go func() {
		sands = make(map[Point]bool)
		getUnitsOfSand(paths, true)
		time.Sleep(time.Second * 2)
		isDrawing = false
	}()

	for isDrawing {
		printCave(paths)
		time.Sleep(time.Millisecond * VIS_UPDATE_FREQ)
	}
}
