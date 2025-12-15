package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	q "github.com/adde/advent-of-code/utils/queue"
)

type Device struct {
	Name        string
	Connections []string
}

type State struct {
	Device string
	Flags  uint8
}

const (
	END = "out"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	devices := make(map[string]Device)

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		deviceName := parts[0]
		connections := strings.Split(parts[1], " ")
		device := Device{
			Name:        deviceName,
			Connections: connections,
		}
		devices[deviceName] = device
	}

	fmt.Println("\nPart one:", partOne(devices, "you"))
	fmt.Println("Part two:", partTwo(devices, "svr"))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// BFS to find the number of paths from start to end
func partOne(devices map[string]Device, startDevice string) int {
	paths := 0
	start := devices[startDevice]
	queue := q.New(start)

	for !queue.IsEmpty() {
		current := queue.Pop()

		for _, conn := range current.Connections {
			if conn == END {
				paths++
				continue
			}
			queue.Append(devices[conn])
		}
	}

	return paths
}

// DFS to find the number of paths that include both "dac" and "fft"
func partTwo(devices map[string]Device, startDevice string) int {
	memo := make(map[State]int)
	paths := dfs(startDevice, false, false, memo, devices)

	return paths
}

func dfs(device string, hasDac, hasFft bool, memo map[State]int, devices map[string]Device) int {
	flags := uint8(0)

	if hasDac {
		flags |= 1
	}
	if hasFft {
		flags |= 2
	}

	state := State{device, flags}
	if count, exists := memo[state]; exists {
		return count
	}

	paths := 0
	for _, conn := range devices[device].Connections {
		if conn == END {
			if hasDac && hasFft {
				paths++
			}
		} else {
			newDac := hasDac || conn == "dac"
			newFft := hasFft || conn == "fft"
			paths += dfs(conn, newDac, newFft, memo, devices)
		}
	}

	memo[state] = paths
	return paths
}
