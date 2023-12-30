package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	steps := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		steps = strings.Split(line, ",")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Println("Sum part one:", getStepsSum(steps))
	fmt.Println("Sum part two:", getFocusingPower(getLensConfigurations(steps)))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func getLensConfigurations(steps []string) map[int]*utils.OrderedMap {
	boxes := make(map[int]*utils.OrderedMap)

	// Going through all steps
	for _, step := range steps {
		lensLabel := ""
		boxNumber := 0

		// Going through all characters in a step
		for _, char := range step {
			asciiCode := int(char)

			// If we have not hit equals or dash,
			// we are still building the label and
			// calculating the box number
			if char != '=' && char != '-' {
				lensLabel += string(char)
				boxNumber = (boxNumber + asciiCode) * 17 % 256
			} else {
				// If we hit equals, add the lens
				if char == '=' {
					// Get lens focal length
					focalLength := utils.ToInt(string(step[len(step)-1]))

					// Check if we already have a box,
					// if not, create it and add the lens to it
					_, ok := boxes[boxNumber]
					if ok {
						boxes[boxNumber].Set(lensLabel, focalLength)
					} else {
						boxes[boxNumber] = &utils.OrderedMap{Values: make(map[string]int)}
						boxes[boxNumber].Set(lensLabel, focalLength)
					}
				} else {
					// If we hit dash, remove the lens
					_, ok := boxes[boxNumber]
					if ok {
						boxes[boxNumber].Delete(lensLabel)
					}
				}

				break
			}
		}
	}

	return boxes
}

func getFocusingPower(boxes map[int]*utils.OrderedMap) int {
	sum := 0

	for bn, box := range boxes {
		for slot, label := range box.Keys {
			fl, ok := box.Get(label)
			if ok {
				sum += (bn + 1) * (slot + 1) * fl
			}
		}
	}

	return sum
}

func getStepsSum(steps []string) int {
	sum := 0

	for _, step := range steps {
		currentValue := 0

		for _, char := range step {
			asciiCode := int(char)
			currentValue = (currentValue + asciiCode) * 17 % 256
		}

		sum += currentValue
	}

	return sum
}
