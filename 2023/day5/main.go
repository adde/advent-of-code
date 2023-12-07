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

type Seed struct {
	start int
	end   int
}

type Converter struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func (c Converter) Convert(value int) (int, bool) {
	if value >= c.sourceRangeStart && value < c.sourceRangeStart+c.rangeLength {
		pos := value - c.sourceRangeStart
		newPos := c.destinationRangeStart + pos

		return newPos, true
	}
	return value, false
}

func (c Converter) ConvertReverse(value int) (int, bool) {
	if value >= c.destinationRangeStart && value < c.destinationRangeStart+c.rangeLength {
		pos := value - c.destinationRangeStart
		newPos := c.sourceRangeStart + pos

		return newPos, true
	}
	return value, false
}

func main() {
	startTime := time.Now()

	seeds := []int{}
	mappers := [][]Converter{}
	mapSection := []Converter{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Skip map title lines
		linePartsMap := strings.Split(line, " ")
		if len(linePartsMap) > 1 && linePartsMap[1] == "map:" {
			continue
		}

		lineParts := strings.Split(line, ": ")

		if len(lineParts) > 1 {
			if lineParts[0] == "seeds" {
				for _, seed := range strings.Split(lineParts[1], " ") {
					seeds = append(seeds, toInt(seed))
				}
			}
		} else if len(lineParts) == 1 && lineParts[0] != "" {
			dest, source, length := 0, 0, 0
			fmt.Sscanf(lineParts[0], "%d %d %d", &dest, &source, &length)
			mapSection = append(mapSection, Converter{dest, source, length})
		} else {
			mappers = append(mappers, mapSection)
			mapSection = []Converter{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Add last "missing" map section after the loop
	mappers = append(mappers, mapSection)

	fmt.Println("Lowest part one:", getLowestLocationFromSeeds(seeds, mappers))
	fmt.Println("Lowest part two:", getLowestLocationFromRange(getSeedRanges(seeds), mappers))
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func getLowestLocationFromSeeds(seeds []int, mappers [][]Converter) int {
	lowest := []int{}

	// Convert seeds to locations
	for _, seed := range seeds {
		lowest = append(lowest, convertSeedToLocation(seed, mappers))
	}

	// Sort locations ascending to get lowest location first
	for i := 0; i < len(lowest); i++ {
		for j := i + 1; j < len(lowest); j++ {
			if lowest[i] > lowest[j] {
				lowest[i], lowest[j] = lowest[j], lowest[i]
			}
		}
	}

	if len(lowest) > 0 {
		return lowest[0]
	}

	return 0
}

func getLowestLocationFromRange(seedRanges []Seed, mappers [][]Converter) int {
	lowest := 0

	// Loop potential locations from 0 to max int
	// When lowest location that matches a seed range is found,
	// break all loops and return the lowest location
outer:
	for i := 0; i < 9223372036854775807; i++ {
		// Convert the location(i) back to a seed
		seed := convertLocationToSeed(i, mappers)

		// Check if seed exists in seeds range
		for _, s := range seedRanges {
			if seed >= s.start && seed < s.end {
				lowest = i // i represents the location
				break outer
			}
		}
	}

	return lowest
}

func convertSeedToLocation(val int, mappers [][]Converter) int {
	converted := false

	for _, ms := range mappers {
		for _, mapper := range ms {
			val, converted = mapper.Convert(val)

			if converted {
				break
			}
		}
	}

	return val
}

func convertLocationToSeed(val int, mappers [][]Converter) int {
	converted := false

	for j := len(mappers) - 1; j >= 0; j-- {
		for k := len(mappers[j]) - 1; k >= 0; k-- {
			val, converted = mappers[j][k].ConvertReverse(val)

			if converted {
				break
			}
		}
	}

	return val
}

func getSeedRanges(seeds []int) []Seed {
	seedRange := []Seed{}

	for i, seed := range seeds {
		if i%2 == 0 {
			start, end := seed, seed+seeds[i+1]
			seedRange = append(seedRange, Seed{start, end})
		}
	}

	return seedRange
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
