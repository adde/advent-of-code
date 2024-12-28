package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Seating struct {
	person1   string
	person2   string
	happiness int
}

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	seatings := []Seating{}
	persons := make(map[string]bool)
	personList := []string{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		var happiness int
		fmt.Sscanf(parts[3], "%d", &happiness)
		if parts[2] == "lose" {
			happiness = -happiness
		}
		seatings = append(seatings, Seating{
			person1:   parts[0],
			person2:   strings.Trim(parts[len(parts)-1], "."),
			happiness: happiness,
		})
	}

	// Build unique persons list
	for _, seat := range seatings {
		persons[seat.person1] = true
		persons[seat.person2] = true
	}

	// Convert persons map to slice
	for person := range persons {
		personList = append(personList, person)
	}

	// Part one
	perms := generatePermutations(personList)
	ansP1 := calculateHappiness(seatings, perms)

	// Part two
	personList = append(personList, "Me")
	perms = generatePermutations(personList)
	ansP2 := calculateHappiness(seatings, perms)

	fmt.Println("\nPart one:", ansP1)
	fmt.Println("Part two:", ansP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Calculates the happiness of a seating arrangement
func calculateHappiness(seatings []Seating, perms [][]string) int {
	ans := 0

	for _, perm := range perms {
		happiness := 0

		for i := 0; i < len(perm); i++ {
			person1 := perm[i]
			person2 := perm[(i+1)%len(perm)]

			for _, seat := range seatings {
				if seat.person1 == person1 && seat.person2 == person2 {
					happiness += seat.happiness
				} else if seat.person1 == person2 && seat.person2 == person1 {
					happiness += seat.happiness
				}
			}
		}

		if happiness > ans {
			ans = happiness
		}
	}

	return ans
}

// Generates all possible permutations of persons
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
