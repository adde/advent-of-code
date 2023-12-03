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
	sumP1 := 0
	sumP2 := 0

	aToZMap := map[string]int{
		"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6,
		"g": 7, "h": 8, "i": 9, "j": 10, "k": 11, "l": 12,
		"m": 13, "n": 14, "o": 15, "p": 16, "q": 17, "r": 18,
		"s": 19, "t": 20, "u": 21, "v": 22, "w": 23, "x": 24,
		"y": 25, "z": 26, "A": 27, "B": 28, "C": 29, "D": 30,
		"E": 31, "F": 32, "G": 33, "H": 34, "I": 35, "J": 36,
		"K": 37, "L": 38, "M": 39, "N": 40, "O": 41, "P": 42,
		"Q": 43, "R": 44, "S": 45, "T": 46, "U": 47, "V": 48,
		"W": 49, "X": 50, "Y": 51, "Z": 52,
	}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rows := 0
	group := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		mid := len(line) / 2
		firstHalf := line[:mid]
		secondHalf := line[mid:]

		for _, char := range firstHalf {
			if strings.Contains(secondHalf, string(char)) {
				sumP1 += aToZMap[string(char)]
				break
			}
		}

		group = append(group, line)
		rows++
		if rows%3 == 0 {
			for _, char := range group[0] {
				if strings.Contains(group[1], string(char)) && strings.Contains(group[2], string(char)) {
					sumP2 += aToZMap[string(char)]
					break
				}
			}
			group = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part one:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}
