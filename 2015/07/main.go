package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Gate struct {
	op   string
	x, y string
}

var wiresP1 = make(map[string]uint16)
var wiresP2 = make(map[string]uint16)
var gates = make(map[string]Gate)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	parseInput(lines)

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() uint16 {
	for name := range gates {
		wiresP1[name] = evalGate(name, wiresP1)
	}

	return wiresP1["a"]
}

func partTwo() uint16 {
	wiresP2["b"] = partOne()

	for name := range gates {
		wiresP2[name] = evalGate(name, wiresP2)
	}

	return wiresP2["a"]
}

func evalGate(gateName string, wires map[string]uint16) uint16 {
	gate := gates[gateName]
	for _, wire := range []string{gate.x, gate.y} {
		if _, exists := wires[wire]; !exists {
			if wire != "" && !IsInt(wire) && wires[wire] == 0 {
				wires[wire] = evalGate(wire, wires)
			}
		}
	}

	if gate.op == "ASSIGN" {
		if IsInt(gate.x) {
			return uint16(u.ToInt(gate.x))
		}
		return wires[gate.x]
	}

	if gate.op == "NOT" {
		if IsInt(gate.x) {
			return ^uint16(u.ToInt(gate.x))
		}
		return ^wires[gate.x]
	}

	x, y := uint16(0), uint16(0)
	if IsInt(gate.x) {
		x = uint16(u.ToInt(gate.x))
	} else {
		x = wires[gate.x]
	}

	if IsInt(gate.y) {
		y = uint16(u.ToInt(gate.y))
	} else {
		y = wires[gate.y]
	}

	if gate.op == "AND" {
		return x & y
	}
	if gate.op == "OR" {
		return x | y
	}
	if gate.op == "LSHIFT" {
		return x << y
	}
	if gate.op == "RSHIFT" {
		return x >> y
	}

	return 0
}

func parseInput(lines []string) {
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		inputs := strings.Split(parts[0], " ")

		if len(inputs) == 1 {
			if IsInt(inputs[0]) {
				wiresP1[parts[1]] = uint16(u.ToInt(inputs[0]))
			} else {
				gates[parts[1]] = Gate{"ASSIGN", inputs[0], ""}
			}
			continue
		}

		if inputs[0] == "NOT" {
			gates[parts[1]] = Gate{"NOT", inputs[1], ""}
		} else {
			gates[parts[1]] = Gate{inputs[1], inputs[0], inputs[2]}
		}
	}
}

func IsInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
