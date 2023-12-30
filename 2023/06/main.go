package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	times := []int{}
	distances := []int{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		// If first line in file, get the times (ms)
		// If second line, get the distances (mm)
		if count == 0 {
			getLineNumbers(line, &times)

		} else if count == 1 {
			getLineNumbers(line, &distances)
		}

		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Multiple races(part one):", calcMultiRace(times, distances))
	fmt.Println("One long race(part two):", calcSingleRace(times, distances))
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func calcMultiRace(times, distances []int) int {
	multipliedRecords := 1

	for i, time := range times {
		timesBeatenRecord := 0

		// Loop through all different speeds(s)
		// that we have in a race with the available time:
		// Speed 1 = 1 mm/ms, Speed 2 = 2 mm/ms, etc.
		for s := 1; s < time; s++ {
			// Calculate how far we can travel with the current speed
			// and the remaining time
			remainingTime := time - s
			distanceTraveled := s * remainingTime

			// If we traveled further than the record,
			// Count that as a win
			if distanceTraveled > distances[i] {
				timesBeatenRecord++
			}
		}

		multipliedRecords *= timesBeatenRecord
	}

	return multipliedRecords
}

func calcSingleRace(times, distances []int) int {
	concatTime := ""
	concatDistance := ""
	timesBeatenRecord := 0

	// Concatenate all times and distances into one number
	for i, v := range times {
		concatTime += strconv.Itoa(v)
		concatDistance += strconv.Itoa(distances[i])
	}

	concatTimeInt := toInt(concatTime)
	concatDistanceInt := toInt(concatDistance)

	for s := 1; s < concatTimeInt; s++ {
		remainingTime := concatTimeInt - s
		distanceTraveled := s * remainingTime

		if distanceTraveled > concatDistanceInt {
			timesBeatenRecord++
		}
	}

	return timesBeatenRecord
}

func getLineNumbers(line string, numbers *[]int) {
	reN := regexp.MustCompile(`\d+`)
	matches := reN.FindAllStringSubmatch(line, -1)

	for _, m := range matches {
		*numbers = append(*numbers, toInt(m[0]))
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
