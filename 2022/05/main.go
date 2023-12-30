package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Crate struct {
	start int
	char  string
}

type Instruction struct {
	move int
	from int
	to   int
}

func main() {
	startTime := time.Now()

	messageP1 := ""
	messageP2 := ""
	isInstructions := false

	crates := []Crate{}
	stackNumbers := map[int]int{}
	stacks := map[int][]Crate{}
	instructions := []Instruction{}

	re := regexp.MustCompile(`\w`)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			isInstructions = true
			continue
		}

		if isInstructions {
			move, from, to := 0, 0, 0
			fmt.Sscanf(line, "move %d from %d to %d", &move, &from, &to)
			instructions = append(instructions, Instruction{move, from, to})
		} else {
			matches := re.FindAllStringIndex(line, -1)

			for _, match := range matches {
				start, end := match[0], match[1]
				crate := line[start:end]

				stackNumber, err := strconv.Atoi(crate)
				if err != nil {
					crates = append(crates, Crate{start, crate})
				} else {
					stackNumbers[start] = stackNumber
				}
			}
		}
	}

	// Create stacks
	for _, crate := range crates {
		stacks[stackNumbers[crate.start]] = append(stacks[stackNumbers[crate.start]], crate)
	}
	stacksP2 := make(map[int][]Crate)
	for k, v := range stacks {
		stacksP2[k] = append([]Crate(nil), v...)
	}

	// Move crates
	for _, instruction := range instructions {
		stack := stacks[instruction.from]

		for i := 0; i < instruction.move; i++ {
			crate := stack[0]
			// Move crate to another stack
			stacks[instruction.to] = append([]Crate{crate}, stacks[instruction.to]...)
			// Remove crate from previous stack
			stacks[instruction.from] = append(stacks[instruction.from][:0], stacks[instruction.from][1:]...)
		}
	}

	// Move multiple crates at once
	for _, instruction := range instructions {
		crates := append([]Crate(nil), stacksP2[instruction.from][:instruction.move]...)
		// Move crate to another stack
		stacksP2[instruction.to] = append(crates, stacksP2[instruction.to]...)
		// Remove crates slice from previous stack
		stacksP2[instruction.from] = append(stacksP2[instruction.from][:0], stacksP2[instruction.from][instruction.move:]...)
	}

	messageP1 = getMessage(stacks)
	messageP2 = getMessage(stacksP2)

	fmt.Println("Message part one:", messageP1)
	fmt.Println("Message part two:", messageP2)
	fmt.Println("Elapsed time", time.Since(startTime))
}

func getMessage(stacks map[int][]Crate) string {
	message := ""

	keys := make([]int, 0, len(stacks))
	for k := range stacks {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		message += stacks[k][0].char
	}

	return message
}
