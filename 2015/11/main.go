package main

import (
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	initialPassword := u.ReadAll("input.txt")

	passwordP1 := initialPassword
	for {
		passwordP1 = incrementPassword(passwordP1)
		if isValid(passwordP1) {
			break
		}
	}

	passwordP2 := incrementPassword(passwordP1)
	for {
		passwordP2 = incrementPassword(passwordP2)
		if isValid(passwordP2) {
			break
		}
	}

	fmt.Println("\nPart one:", passwordP1)
	fmt.Println("Part two:", passwordP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func incrementPassword(password string) string {
	runes := []rune(password)

	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == 'z' {
			runes[i] = 'a'
		} else {
			runes[i]++
			break
		}
	}

	return string(runes)
}

func isValid(password string) bool {
	return hasIncreasingStraight(password) && !hasForbiddenLetters(password) && hasTwoPairs(password)
}

func hasIncreasingStraight(password string) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i]+1 == password[i+1] && password[i]+2 == password[i+2] {
			return true
		}
	}

	return false
}

func hasForbiddenLetters(password string) bool {
	for _, r := range password {
		if r == 'i' || r == 'o' || r == 'l' {
			return true
		}
	}

	return false
}

func hasTwoPairs(password string) bool {
	pairs := 0

	for i := 0; i < len(password)-1; i++ {
		if password[i] == password[i+1] {
			pairs++
			i++
		}
	}

	return pairs >= 2
}
