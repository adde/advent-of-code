package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	JOKER = "J"
)

type Hand struct {
	cards       []string
	bid         int
	occurrences map[string]int
	joker       bool
}

type ByTypeAndValues []Hand

func (h ByTypeAndValues) Len() int      { return len(h) }
func (h ByTypeAndValues) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h ByTypeAndValues) Less(i, j int) bool {
	if h[i].GetType() < h[j].GetType() {
		return false
	} else if h[i].GetType() == h[j].GetType() {
		for k := 0; k < 5; k++ {
			if h[i].GetCardValue(k) == h[j].GetCardValue(k) {
				continue
			} else if h[i].GetCardValue(k) > h[j].GetCardValue(k) {
				return false
			} else if h[i].GetCardValue(k) < h[j].GetCardValue(k) {
				return true
			}
		}
	}
	return true
}

func (h *Hand) CountCards() {
	if len(h.occurrences) > 0 {
		return
	}

	occurrences := make(map[string]int)

	for _, card := range h.cards {
		occurrences[card]++
	}

	h.occurrences = occurrences
}

func (h Hand) GetType() int {
	h.CountCards()

	if h.IsFiveOfAKind() {
		return 1
	} else if h.IsFourOfAKind() {
		return 2
	} else if h.IsFullHouse() {
		return 3
	} else if h.IsThreeOfAKind() {
		return 4
	} else if h.IsTwoPair() {
		return 5
	} else if h.IsOnePair() {
		return 6
	}

	return 7
}

func (h Hand) GetCardValue(i int) int {
	cardValue := 0

	switch h.cards[i] {
	case "T":
		cardValue = 10
	case "J":
		cardValue = 11
	case "Q":
		cardValue = 12
	case "K":
		cardValue = 13
	case "A":
		cardValue = 14
	default:
		cardValue = toInt(h.cards[i])
	}

	// If jokers are enabled, lower its card value to 1
	if h.joker && cardValue == 11 {
		cardValue = 1
	}

	return cardValue
}

func (h Hand) GetJokers() int {
	jokers, ok := h.occurrences[JOKER]

	if !ok {
		jokers = 0
	}

	return jokers
}

func (h Hand) IsFiveOfAKind() bool {
	jokers := h.GetJokers()

	for card, count := range h.occurrences {
		if count == 5 {
			return true
		} else if h.joker && card != JOKER && count+jokers == 5 {
			return true
		}
	}

	return false
}

func (h Hand) IsFourOfAKind() bool {
	jokers := h.GetJokers()

	for card, count := range h.occurrences {
		if count == 4 {
			return true
		} else if h.joker && card != JOKER && count+jokers == 4 {
			return true
		}
	}

	return false
}

func (h Hand) IsFullHouse() bool {
	jokers := h.GetJokers()

	hasThree := false
	hasTwo := false
	for card, count := range h.occurrences {
		if count == 3 && card != JOKER {
			hasThree = true
		} else if h.joker && card != JOKER && count+jokers == 3 {
			hasThree = true
			jokers = 0
		} else if count == 2 && card != JOKER {
			hasTwo = true
		} else if h.joker && card != JOKER && count+jokers == 2 {
			hasTwo = true
			jokers = 0
		}
	}

	if !hasThree && hasTwo && jokers == 3 ||
		hasThree && !hasTwo && jokers == 2 ||
		hasThree && hasTwo && jokers == 1 {
		return true
	}

	return hasThree && hasTwo
}

func (h Hand) IsThreeOfAKind() bool {
	jokers := h.GetJokers()

	for card, count := range h.occurrences {
		if count == 3 {
			return true
		} else if h.joker && card != JOKER && count+jokers == 3 {
			return true
		}
	}

	return false
}

func (h Hand) IsTwoPair() bool {
	jokers := h.GetJokers()

	pairCount := 0
	for card, count := range h.occurrences {
		if count == 2 && card != JOKER {
			pairCount++
		} else if h.joker && card != JOKER && count+jokers == 2 {
			pairCount++
			jokers = 0
		}
	}

	if pairCount == 1 && jokers == 2 || pairCount == 2 && jokers == 1 {
		return true
	}

	return pairCount == 2
}

func (h Hand) IsOnePair() bool {
	jokers := h.GetJokers()

	for key, count := range h.occurrences {
		if count == 2 {
			return true
		} else if h.joker && key != JOKER && count+jokers == 2 {
			return true
		}
	}

	return false
}

func main() {
	startTime := time.Now()

	sumP1, sumP2 := 0, 0
	handsP1, handsP2 := []Hand{}, []Hand{}

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineParts := strings.Split(line, " ")

		cards := make([]string, len(lineParts[0]))
		for i, char := range lineParts[0] {
			cards[i] = string(char)
		}

		handsP1 = append(handsP1, Hand{cards, toInt(lineParts[1]), make(map[string]int), false})
		handsP2 = append(handsP2, Hand{cards, toInt(lineParts[1]), make(map[string]int), true})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sumP1 = getSum(sortHandsByCards(handsP1))
	sumP2 = getSum(sortHandsByCards(handsP2))

	fmt.Println("Sum part one:", sumP1)
	fmt.Println("Sum part two:", sumP2)
	fmt.Println("Elapsed time:", time.Since(startTime))
}

func getSum(hands []Hand) int {
	sum := 0

	for i, hand := range hands {
		rank := i + 1
		sum += rank * hand.bid
	}

	return sum
}

func sortHandsByCards(hands []Hand) []Hand {
	sort.Sort(ByTypeAndValues(hands))
	return hands
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}
