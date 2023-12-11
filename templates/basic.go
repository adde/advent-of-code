package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	fmt.Printf("Elapsed time: %s\n\n", time.Since(startTime))
}
