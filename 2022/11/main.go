package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type Monkey struct {
	Id          int
	Items       []int
	Operation   string
	Amount      int
	Divisor     int
	Receiver    map[bool]int
	Inspections int
}

func (m Monkey) GetAmount(item int) int {
	if m.Amount == -1 {
		return item
	}

	return m.Amount
}

func (m Monkey) CalcNewItem(item, amount int) int {
	newItem := item

	if m.Operation == "+" {
		newItem += amount
	} else {
		newItem *= amount
	}

	return newItem
}

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	monkeysP1 := parseInput(lines)
	monkeysP2 := parseInput(lines)

	fmt.Println("\nLevel of monkey business(part one):", getLevelOfMonkeyBusiness(monkeysP1, 20))
	fmt.Println("Level of monkey business(part two):", getLevelOfMonkeyBusiness(monkeysP2, 10000))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) []Monkey {
	monkeys := make([]Monkey, 0)
	monkey := Monkey{}
	receiver := map[bool]int{}

	for _, line := range lines {
		if line == "" {
			monkey.Receiver = receiver
			monkeys = append(monkeys, monkey)
			monkey = Monkey{}
			receiver = map[bool]int{}
			continue
		}
		if strings.Contains(line, "Monkey") {
			monkey.Id = utils.ToInt(line[len(line)-2 : len(line)-1])
		}
		if strings.Contains(line, "Starting items") {
			items := []int{}

			for _, item := range strings.Split(strings.Split(line, ": ")[1], ", ") {
				items = append(items, utils.ToInt(item))
			}

			monkey.Items = items
		}
		if strings.Contains(line, "Operation") {
			operationParts := strings.Split(strings.Split(line, "old ")[1], " ")
			monkey.Operation = operationParts[0]
			amount := -1

			if operationParts[1] != "old" {
				amount = utils.ToInt(operationParts[1])
			}

			monkey.Amount = amount
		}
		if strings.Contains(line, "Test") {
			monkey.Divisor = utils.ToInt(strings.Split(line, "by ")[1])
		}
		if strings.Contains(line, "If true") {
			receiver[true] = utils.ToInt(strings.Split(line, "monkey ")[1])
		}
		if strings.Contains(line, "If false") {
			receiver[false] = utils.ToInt(strings.Split(line, "monkey ")[1])
		}
	}

	// Add last monkey
	monkey.Receiver = receiver
	monkeys = append(monkeys, monkey)

	return monkeys
}

func getLevelOfMonkeyBusiness(monkeys []Monkey, rounds int) int {
	divisorsMultiple := getMonkeyDivisorsMultiple(monkeys)

	for i := 0; i < rounds; i++ {
		for m, monkey := range monkeys {
			for _, item := range monkey.Items {
				amount := monkey.GetAmount(item)
				newItem := monkey.CalcNewItem(item, amount)

				// If rounds is less than 21, divide the item by 3,
				// else modulo the multiple of all the divisors for the item.
				if rounds <= 20 {
					newItem = newItem / 3
				} else {
					newItem %= divisorsMultiple
				}

				// Send updated item to receiver monkey
				receiverId := monkey.Receiver[newItem%monkey.Divisor == 0]
				monkeys[receiverId].Items = append(monkeys[receiverId].Items, newItem)

				monkeys[m].Inspections++
			}

			// Remove items from sender monkey
			monkeys[m].Items = []int{}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].Inspections > monkeys[j].Inspections
	})

	return monkeys[0].Inspections * monkeys[1].Inspections
}

func getMonkeyDivisorsMultiple(monkeys []Monkey) int {
	mod := 1

	for _, monkey := range monkeys {
		mod *= monkey.Divisor
	}

	return mod
}
