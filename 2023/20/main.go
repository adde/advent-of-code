package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

const (
	MODULE_RX = "rx"
)

type Module struct {
	Type    string
	Targets []string
	State   bool
	Memory  map[string]int
}

type Pulse struct {
	Source      string
	Destination string
	Strength    int
}

func main() {
	modules := map[string]Module{}

	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		moduleParts := strings.Split(line, " -> ")
		targets := strings.Split(moduleParts[1], ", ")

		moduleName := ""
		if moduleParts[0] == "broadcaster" {
			moduleName = "broadcaster"
		} else {
			moduleName = moduleParts[0][1:]
		}

		modules[moduleName] = Module{
			Type:    string(moduleParts[0][0]),
			Targets: targets,
			State:   false,
			Memory:  map[string]int{},
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nPulses:", getPulses(modules))
	fmt.Println("Fewest button presses:", getFewestButtonPresses(modules))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getPulses(modules map[string]Module) int {
	low, high := 0, 0

	// Init memory and state for modules
	modules = initModuleMemory(modules)

	// Push the button 1000 times,
	// sending a low pulse to the broadcaster module each time.
	for i := 0; i < 1000; i++ {
		low++

		queue := []Pulse{}

		for _, target := range modules["broadcaster"].Targets {
			queue = append(queue, Pulse{Source: "broadcaster", Destination: target, Strength: 0})
		}

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			if pulse.Strength == 0 {
				low++
			} else {
				high++
			}

			module, ok := modules[pulse.Destination]
			if !ok {
				continue
			}

			if module.Type == "%" {
				// If low pulse, switch on or off
				if pulse.Strength == 0 {
					newState := !module.State

					strength := 0
					if newState {
						strength = 1
					}

					modules[pulse.Destination] = Module{
						Type:    module.Type,
						Targets: module.Targets,
						State:   newState,
					}

					for _, dest := range module.Targets {
						queue = append(queue, Pulse{
							Source:      pulse.Destination,
							Destination: dest,
							Strength:    strength,
						})
					}
				}
			} else if module.Type == "&" {
				// Update memory
				memory := module.Memory
				memory[pulse.Source] = pulse.Strength
				modules[pulse.Destination] = Module{
					Type:    module.Type,
					Targets: module.Targets,
					Memory:  memory,
				}

				// If all inputs are high, send a low pulse
				// otherwise send a high pulse
				strength := 0
				for _, mem := range modules[pulse.Destination].Memory {
					if mem == 0 {
						strength = 1
						break
					}
				}

				for _, dest := range module.Targets {
					queue = append(queue, Pulse{
						Source:      pulse.Destination,
						Destination: dest,
						Strength:    strength,
					})
				}
			}
		}
	}

	return low * high
}

func getFewestButtonPresses(modules map[string]Module) int {
	presses := 0
	rxSource := ""
	cycles := map[string]int{}
	seen := map[string]int{}

	// Init memory and state for modules
	modules = initModuleMemory(modules)

	// Find the source module for the "rx" module
	for m := range modules {
		for _, t := range modules[m].Targets {
			if t == MODULE_RX {
				rxSource = m
			}
		}
	}

	// Find all modules that can send a pulse to the rxSource module
	// and add them to the seen map
	for m := range modules {
		for _, t := range modules[m].Targets {
			if t == rxSource {
				seen[m] = 0
			}
		}
	}

	for true {
		presses++

		queue := []Pulse{}

		for _, target := range modules["broadcaster"].Targets {
			queue = append(queue, Pulse{Source: "broadcaster", Destination: target, Strength: 0})
		}

		for len(queue) > 0 {
			pulse := queue[0]
			queue = queue[1:]

			module, ok := modules[pulse.Destination]
			if !ok {
				continue
			}

			// Check if we have seen a pulse from a source module of rxSource
			// and if it was a high pulse
			if pulse.Destination == rxSource && pulse.Strength == 1 {
				seen[pulse.Source]++

				// Add cyles of presses for the source module of rxSource
				_, ok := cycles[pulse.Source]
				if !ok {
					cycles[pulse.Source] = presses
				}

				// If we have seen a high pulse from all sources of rxSource,
				// calculate the LCM of cycles
				allSeen := true
				for m := range seen {
					if seen[m] < 1 {
						allSeen = false
					}
				}
				if allSeen {
					fewestPresses := 1

					for _, presses := range cycles {
						fewestPresses = utils.Lcm(fewestPresses, presses)
					}

					return fewestPresses
				}
			}

			if module.Type == "%" {
				if pulse.Strength == 0 {
					newState := !module.State

					strength := 0
					if newState {
						strength = 1
					}

					modules[pulse.Destination] = Module{
						Type:    module.Type,
						Targets: module.Targets,
						State:   newState,
					}

					for _, dest := range module.Targets {
						queue = append(queue, Pulse{
							Source:      pulse.Destination,
							Destination: dest,
							Strength:    strength,
						})
					}
				}
			} else if module.Type == "&" {
				// Update memory
				memory := module.Memory
				memory[pulse.Source] = pulse.Strength
				modules[pulse.Destination] = Module{
					Type:    module.Type,
					Targets: module.Targets,
					Memory:  memory,
				}

				// If all inputs are high, send a low pulse
				// otherwise send a high pulse
				strength := 0
				for _, mem := range modules[pulse.Destination].Memory {
					if mem == 0 {
						strength = 1
						break
					}
				}

				for _, dest := range module.Targets {
					queue = append(queue, Pulse{
						Source:      pulse.Destination,
						Destination: dest,
						Strength:    strength,
					})
				}
			}
		}
	}

	return presses
}

func initModuleMemory(modules map[string]Module) map[string]Module {
	for m := range modules {
		for _, t := range modules[m].Targets {
			if modules[t].Type == "&" {
				modules[t].Memory[m] = 0
				modules[t] = Module{
					Type:    modules[t].Type,
					Targets: modules[t].Targets,
					State:   false,
					Memory:  modules[t].Memory,
				}
			} else if modules[t].Type == "%" {
				modules[t] = Module{
					Type:    modules[t].Type,
					Targets: modules[t].Targets,
					State:   false,
					Memory:  map[string]int{},
				}
			}
		}
	}

	return modules
}
