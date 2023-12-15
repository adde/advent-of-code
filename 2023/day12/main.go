package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

const (
	REPEAT = 5
)

type Record struct {
	Springs string
	Sizes   []int
}

var records []Record
var cache map[string]int

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		sizes := []int{}

		for _, size := range strings.Split(lineParts[1], ",") {
			sizes = append(sizes, utils.ToInt(size))
		}

		records = append(records, Record{lineParts[0], sizes})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumP1 := 0
	for _, record := range records {
		cache = make(map[string]int)
		sumP1 += numArrangements(record.Springs+".", record.Sizes, 0)
	}

	sumP2 := 0
	for _, record := range records {
		cache = make(map[string]int)
		sumP2 += numArrangements(
			unFoldSprings(record.Springs, REPEAT)+".",
			unFoldSizes(record.Sizes, REPEAT), 0)
	}

	fmt.Println()
	fmt.Println("Sum of arrangements part one:", sumP1)
	fmt.Println("Sum of arrangements part two:", sumP2)
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func numArrangements(springs string, sizes []int, numDoneInGroup int) int {
	// Create a key for this combination of parameters
	key := springs + strconv.Itoa(len(sizes)) + strconv.Itoa(numDoneInGroup)

	// If we already computed this, return the result
	if value, ok := cache[key]; ok {
		return value
	}

	if len(springs) == 0 {
		// Is this an arrangement? Did we handle and close all groups?
		return utils.BoolToInt(len(sizes) == 0 && numDoneInGroup == 0)
	}

	num := 0

	// If next letter is a "?", we branch
	var possible []rune
	if springs[0] == '?' {
		possible = []rune{'.', '#'}
	} else {
		possible = []rune{rune(springs[0])}
	}

	for _, c := range possible {
		if c == '#' {
			// Extend current group
			num += numArrangements(springs[1:], sizes, numDoneInGroup+1)
		} else {
			if numDoneInGroup > 0 {
				// If we were in a group that can be closed, close it
				if len(sizes) > 0 && sizes[0] == numDoneInGroup {
					num += numArrangements(springs[1:], sizes[1:], 0)
				}
			} else {
				// If we are not in a group, move on to next symbol
				num += numArrangements(springs[1:], sizes, 0)
			}
		}
	}

	// Cache the result
	cache[key] = num

	return num
}

func unFoldSprings(springs string, repeat int) string {
	repeated := ""

	for i := 0; i < repeat; i++ {
		repeated += springs + "?"
	}
	repeated = repeated[:len(repeated)-1] // Remove last "?"

	return repeated
}

func unFoldSizes(sizes []int, repeat int) []int {
	uSizes := []int{}

	for i := 0; i < repeat; i++ {
		uSizes = append(uSizes, sizes...)
	}

	return uSizes
}
