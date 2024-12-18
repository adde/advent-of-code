package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Register struct {
	A, B, C int
}

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	fileParts := strings.Split(file, "\n\n")
	lines := strings.Split(fileParts[0], "\n")
	program := u.GetIntsFromString(fileParts[1], false)
	register := Register{
		A: u.GetIntsFromString(lines[0], false)[0],
		B: u.GetIntsFromString(lines[1], false)[0],
		C: u.GetIntsFromString(lines[2], false)[0],
	}

	fmt.Println("\nPart one:", partOne(register, program))
	fmt.Println("Part two:", partTwo(register, program))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(register Register, program []int) string {
	return strings.Join(u.IntsToStrings(runProgram(register, program)), ",")
}

func partTwo(register Register, program []int) int {
	register.A = 0

	// Loop through the program backwards
	for pos := len(program) - 1; pos >= 0; pos-- {
		// Bitwise shift left by 3 (same as multiplying by 8)
		register.A <<= 3

		// If the output of the program is not equal to the program itself, increment A
		for !reflect.DeepEqual(runProgram(register, program), program[pos:]) {
			register.A++
		}
	}

	return register.A
}

func runProgram(register Register, program []int) []int {
	instructionPointer := 0
	output := []int{}

	for instructionPointer < len(program) {
		opcode := program[instructionPointer]
		operand := program[instructionPointer+1]
		instructionPointer += 2

		switch opcode {
		case 0:
			register.A = register.A / int(math.Pow(2, float64(getComboOp(operand, register))))
		case 1:
			register.B = register.B ^ operand
		case 2:
			register.B = getComboOp(operand, register) % 8
		case 3:
			if register.A == 0 {
				continue
			}
			instructionPointer = operand
		case 4:
			register.B = register.B ^ register.C
		case 5:
			output = append(output, getComboOp(operand, register)%8)
		case 6:
			register.B = register.A / int(math.Pow(2, float64(getComboOp(operand, register))))
		case 7:
			register.C = register.A / int(math.Pow(2, float64(getComboOp(operand, register))))
		}
	}

	return output
}

func getComboOp(literalOperand int, register Register) int {
	if literalOperand == 4 {
		return register.A
	}
	if literalOperand == 5 {
		return register.B
	}
	if literalOperand == 6 {
		return register.C
	}

	return literalOperand
}
