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

func main() {
	startTime := time.Now()

	sumP1 := 0
	sumP2 := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		pairs := strings.Split(line, ",")

		if isRangeContained(getInt(strings.Split(pairs[0], "-")[0]), getInt(strings.Split(pairs[0], "-")[1]),
			getInt(strings.Split(pairs[1], "-")[0]), getInt(strings.Split(pairs[1], "-")[1])) == true {
			sumP1++
		}

		if isRangeOverlap(getInt(strings.Split(pairs[0], "-")[0]), getInt(strings.Split(pairs[0], "-")[1]),
			getInt(strings.Split(pairs[1], "-")[0]), getInt(strings.Split(pairs[1], "-")[1])) == true {
			sumP2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func isRangeContained(start1, end1, start2, end2 int) bool {
	return start1 >= start2 && end1 <= end2 || start2 >= start1 && end2 <= end1
}

func isRangeOverlap(start1, end1, start2, end2 int) bool {
	return start1 <= end2 && end1 >= start2 || start2 <= end1 && end2 >= start1
}
