package main

import (
	"fmt"
	"math"
	"slices"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Point struct {
	X, Y float64
}

type Rectangle struct {
	X, Y, Width, Height float64
}

func (r Rectangle) Area() float64 {
	return (r.Width + 1) * (r.Height + 1)
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	tiles := make([]Point, 0)
	for _, line := range lines {
		numbers := u.GetIntsFromString(line, false)
		tiles = append(tiles, Point{X: float64(numbers[0]), Y: float64(numbers[1])})
	}

	fmt.Println("\nPart one:", partOne(tiles))
	fmt.Println("Part two:", partTwo(tiles))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(tiles []Point) int {
	biggestRect := 0

	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			width := (int(tiles[i].X) - int(tiles[j].X) + 1)
			height := (int(tiles[i].Y) - int(tiles[j].Y) + 1)
			rectArea := width * height

			if rectArea > biggestRect {
				biggestRect = rectArea
			}
		}
	}

	return biggestRect
}

func partTwo(tiles []Point) int {
	biggestRect := 0
	rects := make([]Rectangle, 0)

	// Generate all possible rectangles from pairs of points
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {
			width := math.Abs(tiles[i].X-tiles[j].X) + 1
			height := math.Abs(tiles[i].Y-tiles[j].Y) + 1
			rectArea := int(width * height)

			if rectArea <= biggestRect {
				continue
			}

			// Get the top left corner coordinates of the rectangle
			x := min(tiles[i].X, tiles[j].X)
			y := min(tiles[i].Y, tiles[j].Y)

			rect := Rectangle{X: float64(x), Y: float64(y), Width: float64(width - 1), Height: float64(height - 1)}
			rects = append(rects, rect)
		}
	}

	// Sort rectangles by area descending
	slices.SortFunc(rects, func(a, b Rectangle) int {
		return int(b.Area() - a.Area())
	})

	// Check if the rectangle is fully contained within the polygon
	for _, rect := range rects {
		area := int(rect.Area())

		if IsRectInPolygon(rect, tiles) {
			biggestRect = area
			break
		}
	}

	return biggestRect
}

// Checks if a rectangle is fully contained(including boundary) in a polygon
func IsRectInPolygon(rect Rectangle, polygon []Point) bool {
	// Check all four corners
	corners := []Point{
		{rect.X, rect.Y},
		{rect.X + rect.Width, rect.Y},
		{rect.X + rect.Width, rect.Y + rect.Height},
		{rect.X, rect.Y + rect.Height},
	}

	// All corners must be inside or on boundary
	for _, corner := range corners {
		if !IsPointInPolygon(corner.X, corner.Y, polygon) {
			return false
		}
	}

	// Rectangle edges
	rectEdges := [][2]Point{
		{corners[0], corners[1]},
		{corners[1], corners[2]},
		{corners[2], corners[3]},
		{corners[3], corners[0]},
	}

	// Check if any polygon edge properly intersects with rectangle edges
	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		polyEdge := [2]Point{polygon[i], polygon[j]}

		for _, rectEdge := range rectEdges {
			if EdgesIntersect(polyEdge[0], polyEdge[1], rectEdge[0], rectEdge[1]) {
				return false
			}
		}
	}

	return true
}

// Checks if a point is inside or on the boundary of the polygon
func IsPointInPolygon(x, y float64, polygon []Point) bool {
	// First check if point is on any edge
	for i := 0; i < len(polygon); i++ {
		j := (i + 1) % len(polygon)
		if IsPointOnSegment(Point{x, y}, polygon[i], polygon[j]) {
			return true // Point is on the boundary
		}
	}

	// Ray casting for interior points
	inside := false
	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {
		xi, yi := polygon[i].X, polygon[i].Y
		xj, yj := polygon[j].X, polygon[j].Y

		intersect := ((yi > y) != (yj > y)) &&
			(x < (xj-xi)*(y-yi)/(yj-yi)+xi)

		if intersect {
			inside = !inside
		}
		j = i
	}

	return inside
}

// Checks if a point lies on a line segment
func IsPointOnSegment(p, a, b Point) bool {
	const epsilon = 1e-10

	// Check if point is collinear with segment endpoints
	crossProduct := (p.Y-a.Y)*(b.X-a.X) - (p.X-a.X)*(b.Y-a.Y)
	if crossProduct > epsilon || crossProduct < -epsilon {
		return false // Not collinear
	}

	// Check if point is within the bounding box of the segment
	if p.X < min(a.X, b.X)-epsilon || p.X > max(a.X, b.X)+epsilon {
		return false
	}
	if p.Y < min(a.Y, b.Y)-epsilon || p.Y > max(a.Y, b.Y)+epsilon {
		return false
	}

	return true
}

// Checks if two line segments properly intersect (cross through each other)
func EdgesIntersect(p1, p2, p3, p4 Point) bool {
	ccw := func(A, B, C Point) float64 {
		return (C.Y-A.Y)*(B.X-A.X) - (B.Y-A.Y)*(C.X-A.X)
	}

	d1 := ccw(p3, p4, p1)
	d2 := ccw(p3, p4, p2)
	d3 := ccw(p1, p2, p3)
	d4 := ccw(p1, p2, p4)

	// Proper intersection only if signs differ (not touching/collinear)
	if ((d1 > 0 && d2 < 0) || (d1 < 0 && d2 > 0)) &&
		((d3 > 0 && d4 < 0) || (d3 < 0 && d4 > 0)) {
		return true
	}

	return false
}
