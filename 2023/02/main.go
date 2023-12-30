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
	maxCubeCount := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		bag := strings.Split(line, ": ")
		sets := strings.Split(bag[1], "; ")
		possible := true
		fewestPossibleCubes := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		for _, set := range sets {
			cubes := strings.Split(set, ", ")

			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				cubeColor := cubeParts[1]
				cubeCount, err := strconv.Atoi(cubeParts[0])
				if err != nil {
					log.Fatal(err)
				}

				if cubeCount > maxCubeCount[cubeColor] {
					possible = false
				}

				if cubeCount > fewestPossibleCubes[cubeColor] {
					fewestPossibleCubes[cubeColor] = cubeCount
				}
			}
		}

		if possible {
			game := strings.Split(bag[0], " ")
			gameId, err := strconv.Atoi(game[1])
			if err != nil {
				log.Fatal(err)
			}

			sumP1 += gameId
		}

		sumP2 += (fewestPossibleCubes["red"] * fewestPossibleCubes["green"] * fewestPossibleCubes["blue"])
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}
