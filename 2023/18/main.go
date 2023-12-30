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

type Instruction struct {
	Direction [2]int
	Amount    int
}

func main() {
	instructionsP1 := []Instruction{}
	instructionsP2 := []Instruction{}

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
		hexCode := lineParts[2][2 : len(lineParts[2])-1]

		// Set intructions for part one
		instructionsP1 = append(instructionsP1, Instruction{
			Direction: getDirectionByString(lineParts[0]),
			Amount:    utils.ToInt(lineParts[1]),
		})

		// Set intructions for part two
		instructionsP2 = append(instructionsP2, Instruction{
			Direction: getDirectionByString(hexCode[len(hexCode)-1:]),
			Amount:    utils.HexToDec(hexCode[:len(hexCode)-1]),
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nLagoon capacity part one:", getLagoonCapacity(instructionsP1))
	fmt.Println("Lagoon capacity part two:", getLagoonCapacity(instructionsP2))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

// Use Shoelace algorithm to calculate area of polygon
func getLagoonCapacity(instructions []Instruction) int {
	startX, startY, boundary := 0, 0, 0
	vertices := [][2]int{}
	vertices = append(vertices, [2]int{startX, startY})

	// Create vertices for the polygon
	for _, instruction := range instructions {
		startX, startY, boundary = getVertexPos(startX, startY, boundary, instruction)
		vertices = append(vertices, [2]int{startX, startY})
	}

	// Calculate area of polygon
	numberOfVertices := len(vertices)
	sum1, sum2 := 0, 0

	for i := 0; i < numberOfVertices-1; i++ {
		sum1 = sum1 + vertices[i][0]*vertices[i+1][1]
		sum2 = sum2 + vertices[i][1]*vertices[i+1][0]
	}

	sum1 = sum1 + vertices[numberOfVertices-1][0]*vertices[0][1]
	sum2 = sum2 + vertices[0][0]*vertices[numberOfVertices-1][1]

	area := utils.Abs(sum1-sum2) / 2

	// Add boundary to the total area
	area += boundary/2 + 1

	return area
}

func getVertexPos(startX, startY, boundary int, instruction Instruction) (int, int, int) {
	if instruction.Direction[0] == 1 { // Down
		startY += instruction.Amount
	} else if instruction.Direction[0] == -1 { // Up
		startY -= instruction.Amount
	} else if instruction.Direction[1] == 1 { // Right
		startX += instruction.Amount
	} else if instruction.Direction[1] == -1 { // Left
		startX -= instruction.Amount
	}

	boundary += instruction.Amount

	return startX, startY, boundary
}

func getDirectionByString(direction string) [2]int {
	switch direction {
	case "3":
		return [2]int{-1, 0}
	case "U":
		return [2]int{-1, 0}
	case "1":
		return [2]int{1, 0}
	case "D":
		return [2]int{1, 0}
	case "2":
		return [2]int{0, -1}
	case "L":
		return [2]int{0, -1}
	case "0":
		return [2]int{0, 1}
	case "R":
		return [2]int{0, 1}
	}

	return [2]int{0, 0}
}
