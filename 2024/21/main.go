package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/queue"
)

type Cell struct {
	Pos   g.Point
	Moves []rune
}

var numKeypad = [][]rune{
	[]rune("789"),
	[]rune("456"),
	[]rune("123"),
	[]rune(" 0A"),
}

var dirKeypad = [][]rune{
	[]rune(" ^A"),
	[]rune("<v>"),
}

var directions = map[g.Point]rune{
	{Row: -1, Col: 0}: '^',
	{Row: 1, Col: 0}:  'v',
	{Row: 0, Col: -1}: '<',
	{Row: 0, Col: 1}:  '>',
}

var codes [][]rune

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")

	for _, line := range lines {
		codes = append(codes, []rune(line))
	}

	fmt.Println("\nPart one:", partOne())
	fmt.Println("Part two:", partTwo())
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne() int {
	return getComplexitySum(2)
}

func partTwo() int {
	return getComplexitySum(25)
}

func getComplexitySum(depth int) int {
	sum := 0
	cache := map[string]int{}
	numSequences, dirSequences := getSequences(numKeypad), getSequences(dirKeypad)

	for _, code := range codes {
		nextRobot := getDirSequencesFromNumKeypad(code, numSequences)
		optimal := math.MaxInt64

		for _, seq := range nextRobot {
			length := 0

			nextSeq := u.Zip(append([]rune{'A'}, []rune(seq)...), []rune(seq))
			for _, next := range nextSeq {
				length += getSequenceLen(next.Left, next.Right, depth, dirSequences, cache)
			}

			optimal = u.Min(optimal, length)
		}

		sum += optimal * u.ToInt(string(code[:len(code)-1]))
	}

	return sum
}

func getDirSequencesFromNumKeypad(code []rune, sequences map[[2]rune][][]rune) []string {
	optionKeys := u.Zip(append([]rune{'A'}, code...), code)
	options := [][]string{}

	for _, option := range optionKeys {
		seq := sequences[[2]rune{option.Left, option.Right}]
		opt := []string{}
		for _, s := range seq {
			opt = append(opt, string(s))
		}
		options = append(options, opt)
	}

	nextSequences := u.CartesianProduct(options...)
	nextSequencesJoined := []string{}
	for _, c := range nextSequences {
		nextSequencesJoined = append(nextSequencesJoined, strings.Join(c, ""))
	}

	return nextSequencesJoined
}

func getSequenceLen(key1, key2 rune, depth int, sequences map[[2]rune][][]rune, cache map[string]int) int {
	key := fmt.Sprintf("%s%s%d", string(key1), string(key2), depth)
	if val, ok := cache[key]; ok {
		return val
	}

	if depth == 1 {
		return len(sequences[[2]rune{key1, key2}][0])
	}

	optimal := math.MaxInt64

	for _, seq := range sequences[[2]rune{key1, key2}] {
		length := 0

		nextSeq := u.Zip(append([]rune{'A'}, seq...), seq)
		for _, next := range nextSeq {
			length += getSequenceLen(next.Left, next.Right, depth-1, sequences, cache)
		}

		optimal = u.Min(optimal, length)
	}

	cache[key] = optimal
	return optimal
}

func getSequences(keypad [][]rune) map[[2]rune][][]rune {
	pos := make(map[rune]g.Point)

	for r, row := range keypad {
		for c, v := range row {
			if v == ' ' {
				continue
			}
			pos[v] = g.Point{Row: r, Col: c}
		}
	}

	sequences := map[[2]rune][][]rune{}

	for x := range pos {
		for y := range pos {
			if x == y {
				sequences[[2]rune{x, y}] = [][]rune{{'A'}}
				continue
			}

			possibilities := [][]rune{}
			optimal := math.MaxInt64
			q := queue.New(Cell{Pos: pos[x], Moves: []rune{}})

		outer:
			for !q.IsEmpty() {
				entry := q.Pop()

				for _, dir := range [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
					newPos := g.Point{Row: entry.Pos.Row + dir[0], Col: entry.Pos.Col + dir[1]}
					dirChar := directions[g.Point{Row: dir[0], Col: dir[1]}]

					if newPos.Row < 0 || newPos.Row >= len(keypad) || newPos.Col < 0 || newPos.Col >= len(keypad[0]) {
						continue
					}

					if keypad[newPos.Row][newPos.Col] == ' ' {
						continue
					}

					if keypad[newPos.Row][newPos.Col] == y {
						if optimal < len(entry.Moves)+1 {
							break outer
						}

						optimal = len(entry.Moves) + 1
						poss := append([]rune{}, entry.Moves...)
						poss = append(poss, dirChar, 'A')
						possibilities = append(possibilities, poss)
					} else {
						moves := append([]rune{}, entry.Moves...)
						moves = append(moves, dirChar)
						q.Append(Cell{Pos: newPos, Moves: append([]rune{}, moves...)})
					}
				}
			}

			if len(possibilities) > 0 {
				sequences[[2]rune{x, y}] = possibilities
			}
		}
	}

	return sequences
}
