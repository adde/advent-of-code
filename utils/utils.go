package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}

func ToIntSlice(s []string) []int {
	ints := make([]int, len(s))

	for i, v := range s {
		ints[i] = ToInt(v)
	}

	return ints
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}

func HexToDec(hex string) int {
	i, err := strconv.ParseInt(hex, 16, 64)

	if err != nil {
		log.Fatal(err)
	}

	return int(i)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Function to calculate the least common multiple (LCM) of two numbers
func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}

func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}

	return 0
}
