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

	totalPoints := 0
	cardCopies := []int{}
	cardCount := 0

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cardCount++

		// Get card ID
		cardLine := strings.Split(line, ": ")
		cardId := toInt(strings.Replace(cardLine[0], "Card ", "", -1))

		// Get card numbers
		card := strings.Split(cardLine[1], " | ")
		winningNumbers, userNumbers := strings.Split(card[0], " "), strings.Split(card[1], " ")
		winNumMap, userNumMap := getNumberMap(winningNumbers), getNumberMap(userNumbers)

		// Get winning points (part one)
		totalPoints += getWinningPoints(cardId, winNumMap, userNumMap, &cardCopies)

		// Get card copies (part two)
		getCardCopies(cardId, winNumMap, userNumMap, &cardCopies)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum part one:", totalPoints)
	fmt.Println("Sum part two:", len(cardCopies)+cardCount)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func getWinningPoints(cardId int, winningNumMap map[string]bool, userNumMap map[string]bool, cardCopies *[]int) int {
	count := 1
	cardSum := 0

	for wn := range winningNumMap {
		if userNumMap[wn] {
			if cardSum == 0 {
				cardSum = 1
			} else {
				cardSum *= 2
			}

			*cardCopies = append(*cardCopies, cardId+count)
			count++
		}
	}

	return cardSum
}

func getCardCopies(cardId int, winningNumMap map[string]bool, userNumMap map[string]bool, cardCopies *[]int) {
	for _, c := range *cardCopies {
		if c != cardId {
			continue
		}

		count := 1
		for wn := range winningNumMap {
			if userNumMap[wn] {
				*cardCopies = append(*cardCopies, cardId+count)
				count++
			}
		}
	}
}

func getNumberMap(numbers []string) map[string]bool {
	numMap := make(map[string]bool)

	for _, n := range numbers {
		if n != "" {
			numMap[n] = true
		}
	}

	return numMap
}

func toInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
