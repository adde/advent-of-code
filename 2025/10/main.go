package main

import (
	"fmt"
	"math"
	"slices"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	"github.com/adde/advent-of-code/utils/queue"
)

type Machine struct {
	IndicatorLights        []bool
	ButtonWiringSchematics [][]int
	ButtonsBitmask         []int
	JoltageRequirements    []int
}

const (
	N   = 13
	EPS = 1e-8
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	machines := []Machine{}
	for _, line := range lines {
		machines = append(machines, parseMachine(line))
	}

	fmt.Println("\nPart one:", partOne(machines))
	fmt.Println("Part two:", partTwo(machines))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(machines []Machine) int {
	total := 0

	for _, machine := range machines {
		presses := getButtonPresses(machine)
		total += presses
	}

	return total
}

func partTwo(machines []Machine) int {
	var total int

	for _, machine := range machines {
		solver := JoltReqSolver{buttons: machine.ButtonsBitmask}
		presses := solver.Solve(machine.JoltageRequirements)
		total += presses
	}

	return total
}

type State struct {
	Lights  []bool
	Presses int
}

// BFS to find minimum button presses to reach desired indicator lights
func getButtonPresses(machine Machine) int {
	numLights := len(machine.IndicatorLights)
	numButtons := len(machine.ButtonWiringSchematics)
	visited := make(map[string]bool)

	// Helper to serialize state
	serialize := func(lights []bool) string {
		b := make([]byte, numLights)
		for i, v := range lights {
			if v {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
		}
		return string(b)
	}

	// Start from all lights off
	initial := make([]bool, numLights)
	start := State{Lights: initial, Presses: 0}
	q := queue.New(start)
	visited[serialize(initial)] = true

	for !q.IsEmpty() {
		curr := q.Pop()

		if slices.Equal(curr.Lights, machine.IndicatorLights) {
			return curr.Presses
		}

		// Try pressing each button
		for btnIdx := 0; btnIdx < numButtons; btnIdx++ {
			nextLights := make([]bool, numLights)
			copy(nextLights, curr.Lights)

			for _, lightIdx := range machine.ButtonWiringSchematics[btnIdx] {
				nextLights[lightIdx] = !nextLights[lightIdx]
			}

			key := serialize(nextLights)
			if !visited[key] {
				visited[key] = true
				q.Append(State{Lights: nextLights, Presses: curr.Presses + 1})
			}
		}
	}

	// If unreachable
	return -1
}

type Linear struct {
	A [N]float64
	B float64
}

type Variable struct {
	Expr Linear
	Free bool
	Val  int
	Max  int
}

type JoltReqSolver struct {
	buttons []int
}

// Solve finds the minimum button presses to meet desired joltage requirements
func (s *JoltReqSolver) Solve(desired []int) int {
	// Set up variables
	vars := make([]Variable, len(s.buttons))
	for i := range vars {
		vars[i].Max = math.MaxInt
	}

	// Set up equations
	eqs := make([]Linear, len(desired))
	for i, jolt := range desired {
		eq := Linear{B: float64(-jolt)}
		for j, b := range s.buttons {
			if b&(1<<i) != 0 {
				eq.A[j] = 1
				vars[j].Max = min(vars[j].Max, jolt)
			}
		}
		eqs[i] = eq
	}

	// Solve equations
	for i := range vars {
		vars[i].Free = true

		for _, eq := range eqs {
			if expr, ok := extract(eq, i); ok {
				vars[i].Free = false
				vars[i].Expr = expr

				for j := range eqs {
					eqs[j] = substitute(eqs[j], i, expr)
				}

				break
			}
		}
	}

	// Collect free variables
	free := []int(nil)
	for i, v := range vars {
		if v.Free {
			free = append(free, i)
		}
	}

	// Search for best solution
	best, _ := evalRecursive(vars, free, 0)
	return best
}

// evalRecursive tries all combinations of free variables to find minimum total
func evalRecursive(vars []Variable, free []int, index int) (int, bool) {
	if index == len(free) {
		vals := [N]int{}
		total := 0

		for i := len(vars) - 1; i >= 0; i-- {
			x := eval(vars[i], vals)
			if x < -EPS || math.Abs(x-math.Round(x)) > EPS {
				return 0, false
			}
			vals[i] = int(math.Round(x))
			total += vals[i]
		}

		return total, true
	}

	best, found := math.MaxInt, false
	for x := 0; x <= vars[free[index]].Max; x++ {
		vars[free[index]].Val = x
		total, ok := evalRecursive(vars, free, index+1)

		if ok {
			found = true
			best = min(best, total)
		}
	}

	return best, found
}

// extract isolates variable at index in the linear equation
func extract(lin Linear, index int) (Linear, bool) {
	a := -lin.A[index]
	if math.Abs(a) < EPS {
		return Linear{}, false
	}

	r := Linear{B: lin.B / a}
	for i := 0; i < N; i++ {
		if i != index {
			r.A[i] = lin.A[i] / a
		}
	}
	return r, true
}

// substitute replaces variable at index in lin with expr
func substitute(lin Linear, index int, expr Linear) Linear {
	r := Linear{}

	a := lin.A[index]
	lin.A[index] = 0
	for i := 0; i < N; i++ {
		r.A[i] = lin.A[i] + a*expr.A[i]
	}
	r.B = lin.B + a*expr.B
	return r
}

// eval computes the value of variable v given vals for other variables
func eval(v Variable, vals [N]int) float64 {
	if v.Free {
		return float64(v.Val)
	}

	x := v.Expr.B
	for i := 0; i < N; i++ {
		x += v.Expr.A[i] * float64(vals[i])
	}
	return x
}

func parseMachine(line string) Machine {
	// Parse indicator lights
	il := strings.Split(line, "[")[1]
	il = strings.Split(il, "]")[0]
	indicatorLights := []bool{}
	for _, ch := range il {
		if ch == '#' {
			indicatorLights = append(indicatorLights, true)
		} else {
			indicatorLights = append(indicatorLights, false)
		}
	}

	// Parse button wiring schematics
	r := strings.Split(line, "] ")[1]
	b := strings.Split(r, " {")[0]
	buttons := strings.Split(b, " ")
	buttonWiringSchematics := [][]int{}
	buttonBitmask := []int{}
	for _, button := range buttons {
		bw := strings.Trim(button, "()")
		buttonParts := strings.Split(bw, ",")
		buttonWiring := []int{}
		bitmask := 0
		for _, part := range buttonParts {
			buttonWiring = append(buttonWiring, u.ToInt(strings.TrimSpace(part)))
			bitmask |= 1 << u.ToInt(strings.TrimSpace(part))
		}
		buttonWiringSchematics = append(buttonWiringSchematics, buttonWiring)
		buttonBitmask = append(buttonBitmask, bitmask)
	}

	// Parse joltage requirements
	j := strings.Split(line, "{")[1]
	joltageStr := strings.TrimRight(j, "}")
	joltageParts := strings.Split(joltageStr, ",")
	joltageRequirements := []int{}
	for _, part := range joltageParts {
		joltageRequirements = append(joltageRequirements, u.ToInt(strings.TrimSpace(part)))
	}

	return Machine{
		IndicatorLights:        indicatorLights,
		ButtonWiringSchematics: buttonWiringSchematics,
		ButtonsBitmask:         buttonBitmask,
		JoltageRequirements:    joltageRequirements,
	}
}
