package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()
	totalScoreP1 := 0
	totalScoreP2 := 0

	selfMap := map[string]int{
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	winMap := map[string]string{
		"A": "Y",
		"B": "Z",
		"C": "X",
	}

	drawMap := map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
	}

	loseMap := map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")
		self := selfMap[lineParts[1]]
		totalScoreP1 += self

		if (lineParts[0] == "A" && lineParts[1] == "X") ||
			(lineParts[0] == "B" && lineParts[1] == "Y") ||
			(lineParts[0] == "C" && lineParts[1] == "Z") {
			totalScoreP1 += 3
		} else if (lineParts[0] == "A" && lineParts[1] == "Y") ||
			(lineParts[0] == "B" && lineParts[1] == "Z") ||
			(lineParts[0] == "C" && lineParts[1] == "X") {
			totalScoreP1 += 6
		}

		if lineParts[1] == "X" {
			totalScoreP2 += selfMap[loseMap[lineParts[0]]]
		} else if lineParts[1] == "Y" {
			totalScoreP2 += selfMap[drawMap[lineParts[0]]]
			totalScoreP2 += 3
		} else if lineParts[1] == "Z" {
			totalScoreP2 += selfMap[winMap[lineParts[0]]]
			totalScoreP2 += 6
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total score part one:", totalScoreP1)
	fmt.Println("Total score part two:", totalScoreP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}
