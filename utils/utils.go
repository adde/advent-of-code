package utils

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Read a file and return its content as a string
func ReadAll(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

// Read a file and return its content as a slice of strings
func ReadLines(filename string) []string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(data), "\n")
}

// Get all integers from a string
func GetIntsFromString(s string, includeNegatives bool) []int {
	re := regexp.MustCompile(`\d+`)
	if includeNegatives {
		re = regexp.MustCompile(`-?\d+`)
	}

	matches := re.FindAllString(s, -1)

	return ToIntSlice(matches)
}

// Convert a string to an integer
func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
}

// Convert a slice of strings to a slice of integers
func ToIntSlice(s []string) []int {
	ints := make([]int, len(s))

	for i, v := range s {
		ints[i] = ToInt(v)
	}

	return ints
}

// Convert a boolean to an integer
func BoolToInt(b bool) int {
	if b {
		return 1
	}

	return 0
}

// Convert a hexadecimal string to an integer
func HexToDec(hex string) int {
	i, err := strconv.ParseInt(hex, 16, 64)

	if err != nil {
		log.Fatal(err)
	}

	return int(i)
}

// Get the absolute value of an integer
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// Calculate the greatest common divisor (GCD) of two numbers
func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

// Calculate the least common multiple (LCM) of two numbers
func Lcm(a, b int) int {
	return (a * b) / Gcd(a, b)
}

// Calculate the sign of an integer
func Sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}

	return 0
}
