package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	fileParts := strings.Split(file, "\n\n")
	patterns, designs := make([]string, 0), make([]string, 0)
	patterns = append(patterns, strings.Split(fileParts[0], ", ")...)
	designs = append(designs, strings.Split(fileParts[1], "\n")...)

	fmt.Println("\nPart one:", partOne(patterns, designs))
	fmt.Println("Part two:", partTwo(patterns, designs))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(patterns []string, designs []string) int {
	sum := 0

	for _, design := range designs {
		if containsCombination(patterns, design) > 0 {
			sum++
		}
	}

	return sum
}

func partTwo(patterns []string, designs []string) int {
	sum := 0

	for _, design := range designs {
		sum += containsCombination(patterns, design)
	}

	return sum
}

// Use dynamic programming to find the number of combinations of words in a line
func containsCombination(words []string, line string) int {
	lineLen := len(line)
	dp := make([]int, lineLen+1)
	dp[0] = 1 // Base case for empty string

	for i := 1; i <= lineLen; i++ {
		for _, word := range words {
			wordLen := len(word)

			// Check if the word fits in the line and if the substring is equal to the word
			if i >= wordLen && line[i-wordLen:i] == word {
				// Add the number of ways to form the string up to the position before the word
				dp[i] += dp[i-wordLen]
			}
		}
	}

	return dp[lineLen]
}
