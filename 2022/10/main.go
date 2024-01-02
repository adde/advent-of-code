package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Instruction struct {
	Operation string
	Argument  int
}

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	program := parseInput(lines)

	fmt.Println("\nSum of signal strengths:", getSignalStrengths(program))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Instruction {
	program := make([]Instruction, 0)

	for _, line := range lines {
		lineParts := strings.Split(line, " ")
		arg := 0

		if len(lineParts) > 1 {
			arg = utils.ToInt(lineParts[1])
		}

		program = append(program, Instruction{lineParts[0], arg})
	}

	return program
}

func getSignalStrengths(program []Instruction) int {
	x := 1
	signals := []int{}

	for _, in := range program {
		if in.Operation == "noop" {
			signals = append(signals, x)
		} else {
			signals = append(signals, x)
			signals = append(signals, x)
			x += in.Argument
		}
	}
	signals = append(signals, x)

	// Sum up the signal strengths
	sum := 0
	for i := 20; i <= 220; i += 40 {
		sum += signals[i-1] * i
	}

	// Print the image rendered by the CRT
	printCrt(signals)

	return sum
}

func printCrt(signals []int) {
	fmt.Println()
	for r := 0; r <= 220; r += 40 {
		for c := 0; c <= 40; c++ {
			if math.Abs(float64(signals[r+c]-c)) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
