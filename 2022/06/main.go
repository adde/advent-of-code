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

	startOfPacket := 0
	startOfMessage := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		processedChars := []string{}

		for _, char := range line {
			processedChars = append(processedChars, string(char))

			// Check if the last 4 characters are not the same
			if len(processedChars) >= 4 && isMarker(processedChars, 4) && startOfPacket == 0 {
				startOfPacket = len(processedChars)
			}

			// Check if the last 14 characters are not the same
			if len(processedChars) >= 14 && isMarker(processedChars, 14) && startOfMessage == 0 {
				startOfMessage = len(processedChars)
			}
		}

		fmt.Println("Start of packet:", startOfPacket)
		fmt.Println("Start of message:", startOfMessage)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Elapsed time:", time.Since(startTime))
}

func isMarker(processedChars []string, numberOfChars int) bool {
	isMarker := true

outer:
	for i := len(processedChars) - 1; i >= len(processedChars)-numberOfChars; i-- {
		for j := len(processedChars) - 1; j >= len(processedChars)-numberOfChars; j-- {
			if processedChars[i] == processedChars[j] && i != j {
				isMarker = false
				break outer
			}
		}
	}

	return isMarker
}
