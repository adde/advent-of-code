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
