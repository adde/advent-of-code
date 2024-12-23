package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	computers := make(map[string]map[string]struct{})

	for _, line := range lines {
		link := strings.Split(line, "-")

		// Add the computers to the map, in both directions
		if _, ok := computers[link[0]]; !ok {
			computers[link[0]] = make(map[string]struct{})
		}
		computers[link[0]][link[1]] = struct{}{}

		if _, ok := computers[link[1]]; !ok {
			computers[link[1]] = make(map[string]struct{})
		}
		computers[link[1]][link[0]] = struct{}{}
	}

	fmt.Println("\nPart one:", partOne(computers))
	fmt.Println("Part two:", partTwo(computers))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(computers map[string]map[string]struct{}) int {
	connected := map[string]struct{}{}

	for name, computer := range computers {
		if !strings.HasPrefix(name, "t") {
			continue
		}

		// Check if three computers are connected to each other
		for comp := range computer {
			for comp2 := range computer {
				if comp == comp2 {
					continue
				}

				// Skip if the connection is already added
				if connectionExists(connected, name, comp, comp2) {
					continue
				}

				_, conn1 := computers[comp][comp2]
				_, conn2 := computers[comp2][comp]
				if conn1 || conn2 {
					key := name + "," + comp + "," + comp2
					connected[key] = struct{}{}
				}
			}
		}
	}

	return len(connected)
}

func partTwo(computers map[string]map[string]struct{}) string {
	largestSet := map[string]struct{}{}

	for name, computer := range computers {
		connected := map[string]struct{}{}

		// Find the largest set of computers that are connected to each other
		for comp := range computer {
			for comp2 := range computer {
				if comp == comp2 {
					continue
				}

				// Skip if the connection is already added
				if connectionExists(connected, name, comp, comp2) {
					continue
				}

				_, conn1 := computers[comp][comp2]
				_, conn2 := computers[comp2][comp]
				if conn1 || conn2 {
					key := name + "," + comp + "," + comp2
					connected[key] = struct{}{}
				}
			}
		}

		if len(connected) > len(largestSet) {
			largestSet = connected
		}
	}

	// Extract the password from the largest set of connected computers
	password := []string{}
	for key := range largestSet {
		parts := strings.Split(key, ",")

		for _, part := range parts {
			if !u.SliceContainsString(password, part) {
				password = append(password, part)
			}
		}
	}
	sort.Strings(password)

	return strings.Join(password, ",")
}

// Check if a connection between three computers already exists
func connectionExists(connected map[string]struct{}, computer, comp, comp2 string) bool {
	connExists := false

	for key := range connected {
		if strings.Contains(key, comp) && strings.Contains(key, comp2) && strings.Contains(key, computer) {
			connExists = true
			break
		}
	}

	return connExists
}
