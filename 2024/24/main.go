package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Gate struct {
	a  string
	op string
	b  string
}

var (
	wires = make(map[string]int)
	gates = make(map[string]Gate)
)

func main() {
	startTime := time.Now()
	data := u.ReadAll("input.txt")
	parseInput(data)

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	// Find all z gates and sort them in reverse order
	zGates := []string{}
	for gate := range gates {
		if strings.HasPrefix(gate, "z") {
			zGates = append(zGates, gate)
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(zGates)))

	// Evaluate bit values for all z gates
	output := ""
	for _, gate := range zGates {
		output += fmt.Sprintf("%d", evalGate(gate))
	}

	// Convert binary string to int
	val, _ := strconv.ParseInt(output, 2, 64)

	return int(val)
}

func partTwo() string {
	swapped := []string{}

	// Find the four pairs of wires that have been swapped
	for wire, gate := range gates {
		if isSwapped(wire, gate) {
			swapped = append(swapped, wire)
		}
	}
	sort.Strings(swapped)

	return strings.Join(swapped, ",")
}

// Check if the wire has been swapped
func isSwapped(wire string, gate Gate) bool {
	// Check if the wire starts with 'z', the operation is not 'XOR', and the wire is not 'z45'
	if strings.HasPrefix(wire, "z") && gate.op != "XOR" && wire != "z45" {
		return true
	}

	// Check if the operation is 'XOR' and none of the wires involved start with 'x', 'y', or 'z'
	if gate.op == "XOR" &&
		!strings.HasPrefix(gate.a, "x") && !strings.HasPrefix(gate.a, "y") && !strings.HasPrefix(gate.a, "z") &&
		!strings.HasPrefix(gate.b, "x") && !strings.HasPrefix(gate.b, "y") && !strings.HasPrefix(gate.b, "z") &&
		!strings.HasPrefix(wire, "x") && !strings.HasPrefix(wire, "y") && !strings.HasPrefix(wire, "z") {
		return true
	}

	// Check if the operation is 'AND' and neither input wire is 'x00'
	if gate.op == "AND" && gate.a != "x00" && gate.b != "x00" {
		// Check if the wire is used in any gate with an operation other than 'OR'
		for _, g := range gates {
			if (wire == g.a || wire == g.b) && g.op != "OR" {
				return true
			}
		}
	}

	// Check if the operation is 'XOR' and the wire is used in any gate with an 'OR' operation
	if gate.op == "XOR" {
		for _, g := range gates {
			if (wire == g.a || wire == g.b) && g.op == "OR" {
				return true
			}
		}
	}

	// If none of the conditions are met, the wire has not been swapped
	return false
}

// Evaluate the gate, recursively evaluating the input wires
func evalGate(gateName string) int {
	gate := gates[gateName]

	for _, wire := range []string{gate.a, gate.b} {
		if _, exists := wires[wire]; !exists {
			wires[wire] = evalGate(wire)
		}
	}

	return applyOp(gate.op, wires[gate.a], wires[gate.b])
}

// Apply bitwise operation
func applyOp(op string, a, b int) int {
	switch op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "XOR":
		return a ^ b
	default:
		return 0
	}
}

func parseInput(data string) {
	sections := strings.Split(string(data), "\n\n")

	// Parse wires
	wireLines := strings.Split(sections[0], "\n")
	for _, line := range wireLines {
		wireParts := strings.Split(line, ": ")
		wires[wireParts[0]] = u.ToInt(wireParts[1])
	}

	// Parse gates
	gateLines := strings.Split(sections[1], "\n")
	for _, line := range gateLines {
		gateParts := strings.Split(strings.Replace(line, "-> ", "", 1), " ")
		gates[gateParts[3]] = Gate{
			a:  gateParts[0],
			op: gateParts[1],
			b:  gateParts[2],
		}
	}
}
