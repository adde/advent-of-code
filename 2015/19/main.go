package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Replacement struct {
	from string
	to   string
}

func main() {
	startTime := time.Now()
	input := utils.ReadAll("input.txt")
	replacements, molecule := parseInput(input)

	fmt.Println("\nPart one:", partOne(molecule, replacements))
	fmt.Println("Part two:", partTwo(molecule))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(molecule string, replacements []Replacement) int {
	uniqueMolecules := make(map[string]bool)

	// Try each replacement at each position
	for _, rep := range replacements {
		// Find all occurrences of the "from" pattern
		fromLen := len(rep.from)
		for i := 0; i <= len(molecule)-fromLen; i++ {
			if molecule[i:i+fromLen] == rep.from {
				// Create new molecule with this replacement
				newMolecule := molecule[:i] + rep.to + molecule[i+fromLen:]
				uniqueMolecules[newMolecule] = true
			}
		}
	}

	return len(uniqueMolecules)
}

func partTwo(molecule string) int {
	// Count Rn, Ar, Y and total uppercase letters
	rn := strings.Count(molecule, "Rn")
	ar := strings.Count(molecule, "Ar")
	y := strings.Count(molecule, "Y")

	// Count total elements (uppercase letters followed by optional lowercase)
	upperCount := 0
	for i := 0; i < len(molecule); i++ {
		if i == 0 || (molecule[i] >= 'A' && molecule[i] <= 'Z') {
			upperCount++
		}
	}

	// The formula is based on pattern analysis of the replacements:
	// - Each step reduces the count by 1 (except for special elements)
	// - Rn and Ar appear in pairs and can be removed together
	// - Y appears with additional elements that also need to be removed
	return upperCount - rn - ar - 2*y - 1
}

func parseInput(input string) ([]Replacement, string) {
	parts := strings.Split(input, "\n\n")
	rulesStr := parts[0]
	molecule := strings.TrimSpace(parts[1])

	var replacements []Replacement
	for _, line := range strings.Split(rulesStr, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.Split(line, " => ")
		replacements = append(replacements, Replacement{
			from: parts[0],
			to:   parts[1],
		})
	}

	return replacements, molecule
}
