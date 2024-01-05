package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/adde/advent-of-code/utils"
)

type PacketPair struct {
	Left, Right any
}

func main() {
	startTime := time.Now()

	lines := utils.ReadLines("input.txt")
	pairs, packets := parseInput(lines)

	fmt.Println("\nPackets in the right order:", getPairsInRightOrder(pairs))
	fmt.Println("Decoder key:", getDecoderKey(packets))
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}

func parseInput(lines []string) ([]PacketPair, []any) {
	packets := []any{}
	pairs := []PacketPair{}
	pair := PacketPair{}
	count := 0

	for _, line := range lines {
		if line == "" {
			pairs = append(pairs, pair)
			pair = PacketPair{}
			continue
		}
		packets = append(packets, toJson(line))
		if count%2 == 0 {
			pair.Left = toJson(line)
		} else {
			pair.Right = toJson(line)
		}
		count++
	}
	pairs = append(pairs, pair)

	return pairs, packets
}

func getPairsInRightOrder(pairs []PacketPair) int {
	sum := 0

	for i, pair := range pairs {
		if compare(pair.Left, pair.Right) <= 0 {
			sum += i + 1
		}
	}

	return sum
}

func getDecoderKey(packets []any) int {
	key := 1

	packets = append(packets, []any{[]any{2.0}}, []any{[]any{6.0}})

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) < 0
	})

	for i, pair := range packets {
		if fmt.Sprint(pair) == "[[2]]" || fmt.Sprint(pair) == "[[6]]" {
			key *= i + 1
		}
	}

	return key
}

func compare(left, right any) int {
	a, okA := left.([]any)
	b, okB := right.([]any)

	switch {
	case !okA && !okB:
		return int(left.(float64) - right.(float64))
	case !okA:
		a = []any{left}
	case !okB:
		b = []any{right}
	}

	for i := 0; i < len(a) && i < len(b); i++ {
		if res := compare(a[i], b[i]); res != 0 {
			return res
		}
	}

	return len(a) - len(b)
}

func toJson(str string) any {
	var jsonObj any
	json.Unmarshal([]byte(str), &jsonObj)
	return jsonObj
}
