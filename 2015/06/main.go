package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Light struct {
	x, y int
}

type Instruction struct {
	action string
	start  Light
	end    Light
}

var instructions = []Instruction{}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		coords := u.GetIntsFromString(line, false)

		start, end := Light{x: coords[0], y: coords[1]}, Light{x: coords[2], y: coords[3]}
		action := ""

		if strings.Contains(line, "turn on") {
			action = "on"
		} else if strings.Contains(line, "turn off") {
			action = "off"
		} else if strings.Contains(line, "toggle") {
			action = "toggle"
		}

		instructions = append(instructions, Instruction{
			action: action,
			start:  start,
			end:    end,
		})
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	grid := make([][]bool, 1000)
	for i := range grid {
		grid[i] = make([]bool, 1000)
	}

	for _, instr := range instructions {
		for x := instr.start.x; x <= instr.end.x; x++ {
			for y := instr.start.y; y <= instr.end.y; y++ {
				if instr.action == "on" {
					grid[x][y] = true
				} else if instr.action == "off" {
					grid[x][y] = false
				} else if instr.action == "toggle" {
					grid[x][y] = !grid[x][y]
				}
			}
		}
	}

	ans := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				ans++
			}
		}
	}

	return ans
}

func partTwo() int {
	grid := make([][]int, 1000)
	for i := range grid {
		grid[i] = make([]int, 1000)
	}

	for _, instr := range instructions {
		for x := instr.start.x; x <= instr.end.x; x++ {
			for y := instr.start.y; y <= instr.end.y; y++ {
				if instr.action == "on" {
					grid[x][y] += 1
				} else if instr.action == "off" {
					grid[x][y] -= 1
					if grid[x][y] < 0 {
						grid[x][y] = 0
					}
				} else if instr.action == "toggle" {
					grid[x][y] += 2
				}
			}
		}
	}

	ans := 0
	for _, row := range grid {
		for _, cell := range row {
			ans += cell
		}
	}

	return ans
}
