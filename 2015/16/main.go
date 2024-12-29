package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Sue struct {
	number int
	things map[string]int
}

var sues []Sue
var things = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		var thing1, thing2, thing3, sueNumber string
		var value1, value2, value3 int
		fmt.Sscanf(line, "Sue %s %s %d, %s %d, %s %d", &sueNumber, &thing1, &value1, &thing2, &value2, &thing3, &value3)

		sue := Sue{
			number: u.ToInt(strings.TrimSuffix(sueNumber, ":")),
			things: make(map[string]int),
		}
		sue.things[strings.TrimSuffix(thing1, ":")] = value1
		sue.things[strings.TrimSuffix(thing2, ":")] = value2
		sue.things[strings.TrimSuffix(thing3, ":")] = value3

		sues = append(sues, sue)
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	for _, sue := range sues {
		match := true

		for thing, value := range sue.things {
			if things[thing] != value {
				match = false
				break
			}
		}

		if match {
			return sue.number
		}
	}

	return 0
}

func partTwo() int {
	for _, sue := range sues {
		match := true

		for thing, value := range sue.things {
			if thing == "cats" || thing == "trees" {
				if things[thing] >= value {
					match = false
					break
				}
			} else if thing == "pomeranians" || thing == "goldfish" {
				if things[thing] <= value {
					match = false
					break
				}
			} else {
				if things[thing] != value {
					match = false
					break
				}
			}
		}

		if match {
			return sue.number
		}
	}

	return 0
}
