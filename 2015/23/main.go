package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Instruction struct {
	operation string
	register  string
	argument  int
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	program := make([]Instruction, len(lines))

	for i, line := range lines {
		numbers := u.GetIntsFromString(line, true)

		arg := 0
		if len(numbers) > 0 {
			arg = numbers[0]
		}

		target := line[4:5]
		if line[4:5] != "a" && line[4:5] != "b" {
			target = ""
		}

		program[i] = Instruction{line[:3], target, arg}
	}

	fmt.Println("\nPart one:", partOne(program))
	fmt.Println("Part two:", partTwo(program))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(program []Instruction) int {
	return runProgram(program, map[string]int{"a": 0, "b": 0})
}

func partTwo(program []Instruction) int {
	return runProgram(program, map[string]int{"a": 1, "b": 0})
}

func runProgram(program []Instruction, register map[string]int) int {
	counter := 0

	for {
		if counter >= len(program) {
			break
		}

		instruction := program[counter]

		switch instruction.operation {
		case "hlf":
			register[instruction.register] /= 2
		case "tpl":
			register[instruction.register] *= 3

		case "inc":
			register[instruction.register]++
		case "jmp":
			counter += instruction.argument - 1
		case "jie":
			if register[instruction.register]%2 == 0 {
				counter += instruction.argument - 1
			}
		case "jio":
			if register[instruction.register] == 1 {
				counter += instruction.argument - 1
			}
		}

		counter++
	}

	return register["b"]
}
