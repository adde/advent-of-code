package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Sequence struct {
	Values    []int
	NextValue int
	PrevValue int
}

func main() {
	startTime := time.Now()

	sumP1, sumP2 := 0, 0
	report := make([]Sequence, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		values := make([]int, 0)

		for _, part := range lineParts {
			values = append(values, toInt(part))
		}

		report = append(report, Sequence{values, 0, 0})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Extrapolate next values for all sequences (part one)
	for _, sequence := range report {
		sequence = extrapolateValue(sequence, true)
		sumP1 += sequence.NextValue
	}

	// Extrapolate previous values for all sequences (part two)
	for _, sequence := range report {
		sequence = extrapolateValue(sequence, false)
		sumP2 += sequence.PrevValue
	}

	fmt.Println("Sum P1:", sumP1)
	fmt.Println("Sum P2:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func extrapolateValue(sequence Sequence, next bool) Sequence {
	sets, inc, idx := make([][]int, 0), 0, 0
	currentValues := sequence.Values
	sets = append(sets, currentValues)

	// Create new sequences from the differences of the values
	// until all values in a new sequence are 0
	for !isAllValuesZeroes(currentValues) {
		newValues := make([]int, 0)

		for i, v := range currentValues {
			if i != len(currentValues)-1 {
				newValues = append(newValues, currentValues[i+1]-v)
			}
		}

		currentValues = newValues
		sets = append(sets, newValues)
	}

	// Extrapolate new values in the sequences
	// Reverse the sequences to go from bottom to top
	// Handles both next and previous values
	sets = reverse(sets)
	for i := 1; i < len(sets); i++ {
		if next {
			idx = len(sets[i]) - 1
			inc = sets[i][idx] + inc
		} else {
			inc = sets[i][idx] - inc
		}

		if next {
			sets[i] = append(sets[i], inc)
		} else {
			sets[i] = append([]int{inc}, sets[i]...)
		}
	}

	if next {
		sequence.NextValue = sets[len(sets)-1][len(sets[len(sets)-1])-1]
	} else {
		sequence.PrevValue = sets[len(sets)-1][0]
	}

	return sequence
}

func isAllValuesZeroes(values []int) bool {
	for _, value := range values {
		if value != 0 {
			return false
		}
	}

	return true
}

func reverse(arr [][]int) [][]int {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}

	return arr
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
