package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	totalKcal := 0
	mostKcals := 0
	allKcals := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if totalKcal > mostKcals {
				mostKcals = totalKcal
			}
			allKcals = append(allKcals, totalKcal)
			totalKcal = 0
			continue
		}
		kcal, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		totalKcal += kcal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(allKcals))) // sort the slice in descending order

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Most kcals", mostKcals)
	fmt.Println("Top three kcals", (allKcals[0] + allKcals[1] + allKcals[2]))
	fmt.Println("Elapsed time", time.Since(startTime))
}
