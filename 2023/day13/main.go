package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	pg := []string{}
	patterns := [][]string{}
	sumP1, sumP2 := 0, 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			pg = append(pg, line)
		} else {
			patterns = append(patterns, pg)
			pg = []string{}
		}
	}
	patterns = append(patterns, pg) // add last group

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(patterns); i++ {
		sumP1 += getReflections(patterns[i], 0)
		sumP2 += getReflections(patterns[i], 1)
	}

	fmt.Println()
	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getReflections(pattern []string, maxDiff int) int {
	sum := 0
	rowLen, colLen := len(pattern), len(pattern[0])

	// Check for vertical symmetry
	for col := 0; col < colLen-1; col++ {
		diff := 0

		for dirCol := 0; dirCol < colLen; dirCol++ {
			left, right := col-dirCol, col+dirCol+1

			if left >= 0 && right < colLen {
				for row := 0; row < rowLen; row++ {
					if pattern[row][left] != pattern[row][right] {
						diff++
					}
				}
			}
		}

		if diff == maxDiff {
			sum += col + 1
		}
	}

	// Check for horizontal symmetry
	for row := 0; row < rowLen-1; row++ {
		diff := 0

		for dirRow := 0; dirRow < rowLen; dirRow++ {
			up, down := row-dirRow, row+dirRow+1

			if up >= 0 && down < rowLen {
				for col := 0; col < colLen; col++ {
					if pattern[up][col] != pattern[down][col] {
						diff++
					}
				}
			}
		}

		if diff == maxDiff {
			sum += 100 * (row + 1)
		}
	}

	return sum
}
