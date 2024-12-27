package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Route struct {
	from     string
	to       string
	distance int
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	routes := []Route{}
	cities := make(map[string]bool)

	// Parse input
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		cities := strings.Split(parts[0], " to ")
		var distance int
		fmt.Sscanf(parts[1], "%d", &distance)
		routes = append(routes, Route{
			from:     cities[0],
			to:       cities[1],
			distance: distance,
		})
	}

	// Build unique cities list
	for _, route := range routes {
		cities[route.from] = true
		cities[route.to] = true
	}

	// Convert cities map to slice
	var cityList []string
	for city := range cities {
		cityList = append(cityList, city)
	}

	// Generate all possible routes
	permutations := generatePermutations(cityList)

	// Record shortest and longest route
	shortestDistance := math.MaxInt32
	longestDistance := 0

	for _, perm := range permutations {
		totalDistance := 0
		valid := true

		for i := 0; i < len(perm)-1; i++ {
			distance, err := findDistance(perm[i], perm[i+1], routes)
			if err != nil {
				valid = false
				break
			}
			totalDistance += distance
		}

		if valid && totalDistance < shortestDistance {
			shortestDistance = totalDistance
		}

		if valid && totalDistance > longestDistance {
			longestDistance = totalDistance
		}
	}

	fmt.Println("\nPart one:", shortestDistance)
	fmt.Println("Part two:", longestDistance)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Generates all possible permutations of cities
func generatePermutations(arr []string) [][]string {
	var result [][]string
	if len(arr) <= 1 {
		return [][]string{arr}
	}

	for i := 0; i < len(arr); i++ {
		curr := arr[i]
		remaining := make([]string, 0)
		remaining = append(remaining, arr[:i]...)
		remaining = append(remaining, arr[i+1:]...)

		subPerms := generatePermutations(remaining)
		for _, perm := range subPerms {
			newPerm := make([]string, 0)
			newPerm = append(newPerm, curr)
			newPerm = append(newPerm, perm...)
			result = append(result, newPerm)
		}
	}

	return result
}

// Looks up the distance between two cities
func findDistance(from, to string, routes []Route) (int, error) {
	for _, route := range routes {
		if (route.from == from && route.to == to) || (route.from == to && route.to == from) {
			return route.distance, nil
		}
	}

	return 0, fmt.Errorf("no route found between %s and %s", from, to)
}
