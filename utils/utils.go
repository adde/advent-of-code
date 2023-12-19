package utils

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		log.Fatal(err)
	}

	return i
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
