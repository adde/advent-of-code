package main

import (
	"fmt"
	"sort"
	"time"

	u "github.com/adde/advent-of-code/utils"
	"github.com/adde/advent-of-code/utils/set"
)

type Sensor struct {
	sx, sy, bx, by int
}

const (
	TARGET_Y        = 2000000
	SEARCH_BOUNDARY = 4000000
	MULTIPLIER      = 4000000
)

var sensors []Sensor

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		numbers := u.GetIntsFromString(line, true)
		sensors = append(sensors, Sensor{numbers[0], numbers[1], numbers[2], numbers[3]})
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	known := set.New[int]()
	intervals := [][2]int{}

	// Get intervals
	for _, sensor := range sensors {
		dist := u.Abs(sensor.sx-sensor.bx) + u.Abs(sensor.sy-sensor.by)
		offset := dist - u.Abs(sensor.sy-TARGET_Y)

		if offset < 0 {
			continue
		}

		loX := sensor.sx - offset
		hiX := sensor.sx + offset

		intervals = append(intervals, [2]int{loX, hiX})

		if sensor.by == TARGET_Y {
			known.Add(sensor.bx)
		}
	}

	// Sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})

	// Merge intervals
	merged := [][2]int{}
	for _, interval := range intervals {
		lo, hi := interval[0], interval[1]

		if len(merged) == 0 {
			merged = append(merged, interval)
			continue
		}

		mhi := merged[len(merged)-1][1]

		if lo > mhi+1 {
			merged = append(merged, interval)
			continue
		}

		merged[len(merged)-1][1] = u.Max(mhi, hi)
	}

	// Find spots that cannot contain a beacon
	cannot := set.New[int]()
	for _, interval := range merged {
		lo, hi := interval[0], interval[1]

		for x := lo; x <= hi; x++ {
			cannot.Add(x)
		}
	}

	return len(cannot.Difference(known))
}

func partTwo() int {
	for y := 0; y <= SEARCH_BOUNDARY; y++ {
		intervals := [][2]int{}

		// Get intervals
		for _, sensor := range sensors {
			dist := u.Abs(sensor.sx-sensor.bx) + u.Abs(sensor.sy-sensor.by)
			offset := dist - u.Abs(sensor.sy-y)

			if offset < 0 {
				continue
			}

			loX := sensor.sx - offset
			hiX := sensor.sx + offset

			intervals = append(intervals, [2]int{loX, hiX})
		}

		// Sort intervals
		sort.Slice(intervals, func(i, j int) bool {
			if intervals[i][0] == intervals[j][0] {
				return intervals[i][1] < intervals[j][1]
			}
			return intervals[i][0] < intervals[j][0]
		})

		// Merge intervals
		merged := [][2]int{}
		for _, interval := range intervals {
			lo, hi := interval[0], interval[1]

			if len(merged) == 0 {
				merged = append(merged, interval)
				continue
			}

			mhi := merged[len(merged)-1][1]

			if lo > mhi+1 {
				merged = append(merged, interval)
				continue
			}

			merged[len(merged)-1][1] = u.Max(mhi, hi)
		}

		// Find free spot
		x := 0
		for _, interval := range merged {
			lo, hi := interval[0], interval[1]

			if x < lo {
				return x*MULTIPLIER + y
			}

			x = max(x, hi+1)

			if x > SEARCH_BOUNDARY {
				break
			}
		}
	}

	return 0
}
