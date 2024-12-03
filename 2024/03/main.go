package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := utils.ReadLines("input.txt")

	mulRegexp := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)

	sumP1 := 0
	sumP2 := 0
	enabled := true

	for _, line := range lines {
		matches := mulRegexp.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			if match[0] == "do()" {
				enabled = true
			} else if match[0] == "don't()" {
				enabled = false
			} else {
				if enabled {
					sumP2 += utils.ToInt(match[1]) * utils.ToInt(match[2])
				}
				sumP1 += utils.ToInt(match[1]) * utils.ToInt(match[2])
			}
		}
	}

	fmt.Println("\nPart one:", sumP1)
	fmt.Println("Part two:", sumP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}
