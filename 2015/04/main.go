package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	key := u.ReadAll("input.txt")
	ansP1, ansP2, count := math.MaxInt32, 0, 0

	for {
		newKey := Md5Hash(key + strconv.Itoa(count))

		if strings.HasPrefix(newKey, "000000") {
			ansP2 = count
			break
		} else if strings.HasPrefix(newKey, "00000") {
			ansP1 = min(ansP1, count)
		}

		count++
	}

	fmt.Println("\nPart one:", ansP1)
	fmt.Println("Part two:", ansP2)
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func Md5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
