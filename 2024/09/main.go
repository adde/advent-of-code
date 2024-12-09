package main

import (
	"fmt"
	"slices"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

func main() {
	startTime := time.Now()
	lines := u.ReadLines("input.txt")
	disk := make([]int, 0)
	fileId := 0

	for i, char := range lines[0] {
		size := u.ToInt(string(char))

		if size == 0 {
			continue
		}

		for j := 0; j < size; j++ {
			if i%2 == 0 {
				// Add even characters to disk as files
				disk = append(disk, fileId)
			} else {
				// Add odd characters to disk as empty space
				disk = append(disk, -1)
			}
		}

		if i%2 == 0 {
			fileId++
		}
	}

	fmt.Println("\nPart one:", partOne(disk))
	fmt.Println("Part two:", partTwo(disk))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(disk []int) int {
	disk = slices.Clone[[]int](disk)

	// Go through the disk and find empty space(-)
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			for j := len(disk) - 1; j > i; j-- {
				// Move file blocks from the end of the disk to the empty space
				if disk[j] != -1 {
					disk[i] = disk[j]
					disk[j] = -1
					break
				}
			}
		}
	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			break
		}
		sum += disk[i] * i
	}

	return sum
}

func partTwo(disk []int) int {
	disk = slices.Clone[[]int](disk)

	// Find all empty space(-1) between files
	emptySpace := getEmptySpace(disk)
	moved := make(map[int]bool)

	for i := len(disk) - 1; i >= 0; i-- {
		// Skip files that have already been moved
		if moved[i] {
			continue
		}
		// Skip free space
		if disk[i] == -1 {
			continue
		}

		// Find the start and size of the file
		fileStart := i
		for fileStart >= 0 && disk[fileStart] != -1 && disk[fileStart] == disk[i] {
			fileStart--
		}
		fileStart++
		fileSize := i - fileStart + 1

		// Continue to the next file in the disk
		i = fileStart

		for k, space := range emptySpace {
			// If the empty space can fit the file and the space is before the file
			if space[1]-space[0]+1 >= fileSize && space[0] < fileStart {
				for j := 0; j < fileSize; j++ {
					// Move file to empty space
					disk[space[0]+j] = disk[fileStart+j]
					disk[fileStart+j] = -1
					moved[space[0]+j] = true
				}
				emptySpace[k][0] += fileSize

				// Remove empty space if it's full
				if emptySpace[k][0] > emptySpace[k][1] {
					emptySpace = append(emptySpace[:k], emptySpace[k+1:]...)
				}

				break
			}
		}

	}

	sum := 0
	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			sum += (disk[i] * i)
		}
	}

	return sum
}

func getEmptySpace(disk []int) [][2]int {
	freeSpace := [][2]int{}

	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 {
			for j := i + 1; j < len(disk); j++ {
				if disk[j] != -1 {
					freeSpace = append(freeSpace, [2]int{i, j - 1})
					i = j
					break
				}
			}
		}
	}

	return freeSpace
}
