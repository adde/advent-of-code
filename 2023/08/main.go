package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const (
	LEFT        = "L"
	START_NODE  = "AAA"
	END_NODE    = "ZZZ"
	START_NODES = "A"
	END_NODES   = "Z"
)

type Node struct {
	Value string
	Left  string
	Right string
}

func main() {
	startTime := time.Now()

	count := 0
	instructions := make([]string, 0)
	nodes := make([]Node, 0)
	startingNodes := make([]string, 0)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// If first line, get instructions,
		// otherwise get nodes
		if count == 0 {
			for _, v := range line {
				instructions = append(instructions, string(v))
			}
		} else if line != "" {
			lineParts := strings.Split(line, " = ")
			insParts := strings.Split(lineParts[1], ", ")
			nodes = append(nodes, Node{lineParts[0], insParts[0][1:], insParts[1][:len(insParts[1])-1]})

			if lineParts[0][2:] == START_NODES {
				startingNodes = append(startingNodes, lineParts[0])
			}
		}

		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Steps taken part one:", getSteps(nodes, instructions, START_NODE, false))
	fmt.Println("Steps taken part two:", getStepsMulti(nodes, instructions, startingNodes))
	fmt.Println("Elapsed time:", time.Since(startTime))
}

// Get number of steps required to reach all end nodes simultaneously
func getStepsMulti(nodes []Node, instructions []string, startingNodes []string) int {
	steps := make([]int, 0)

	for _, startNode := range startingNodes {
		steps = append(steps, getSteps(nodes, instructions, startNode, true))
	}

	return lcmOfArray(steps)
}

// Get number of steps required to reach the end node
func getSteps(nodes []Node, instructions []string, startNode string, last bool) int {
	step := 0
	currentNode := startNode

	for !isEndNode(currentNode, last) {
		instruction := instructions[step%len(instructions)]

		for _, s := range nodes {
			if currentNode == s.Value {
				if instruction == LEFT {
					currentNode = s.Left
				} else {
					currentNode = s.Right
				}
				break
			}
		}

		step++
	}

	return step
}

// Check if the current node is the end node
// If last = false, check the full node name (part one)
// If last = true, check only the last character (part two)
func isEndNode(node string, last bool) bool {
	if (last && node[2:] != END_NODES) || (!last && node != END_NODE) {
		return false
	}

	return true
}

// Function to calculate the greatest common divisor (GCD) using Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Function to calculate the least common multiple (LCM) of two numbers
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

// Function to calculate the least common multiple (LCM) for an array of numbers
func lcmOfArray(nums []int) int {
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		result = lcm(result, nums[i])
	}

	return result
}
