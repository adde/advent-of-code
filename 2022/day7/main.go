package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type File struct {
	parentPath string
	size       int
}

type Directory struct {
	path       string
	parentPath string
	totalSize  int
	files      []File
}

func main() {
	startTime := time.Now()

	sumP1 := 0
	sumP2 := 0

	files := []File{}
	dirs := []Directory{{path: "root"}}
	dirsCreated := map[string]bool{}
	parent := "root"
	totalDiskSpace := 70000000
	requiredDiskSpace := 30000000

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read input file and create files and directories
	for scanner.Scan() {
		line := scanner.Text()
		cmdParts := strings.Split(line, " ")

		if line[0] == '$' {
			if cmdParts[1] == "cd" && cmdParts[2] != "/" && cmdParts[2] != ".." {
				parent += "/" + cmdParts[2]
			} else if cmdParts[1] == "cd" && cmdParts[2] == ".." {
				parent = parent[:strings.LastIndex(parent, "/")]
			}
		} else if line[0] == 'd' && dirsCreated[parent+"/"+cmdParts[1]] == false {
			dirs = append(dirs, Directory{
				parentPath: parent,
				path:       parent + "/" + cmdParts[1]})
			dirsCreated[parent+"/"+cmdParts[1]] = true
		} else {
			files = append(files, File{
				parentPath: parent,
				size:       toInt(cmdParts[0])})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Add files to directories
	for i, dir := range dirs {
		for _, file := range files {
			if file.parentPath == dir.path {
				dirs[i].files = append(dirs[i].files, file)
			}
		}
	}

	// Calculate total size for each directory
	// Sum up all directories with total size <= 100000
	calculateTotalSize(dirs[0], dirs)
	for _, dir := range dirs {
		if dir.totalSize <= 100000 {
			sumP1 += dir.totalSize
		}
	}

	freeDiskSpace := totalDiskSpace - dirs[0].totalSize

	// Order dirs by total size
	for i := 0; i < len(dirs); i++ {
		for j := i + 1; j < len(dirs); j++ {
			if dirs[i].totalSize > dirs[j].totalSize {
				dirs[i], dirs[j] = dirs[j], dirs[i]
			}
		}
	}

	// Get the smallest directory that can be deleted
	for _, dir := range dirs {
		if dir.totalSize > requiredDiskSpace-freeDiskSpace {
			sumP2 = dir.totalSize
			break
		}
	}

	fmt.Println("Sum P1:", sumP1)
	fmt.Println("Sum P2:", sumP2)
	fmt.Println("Elapsed time", time.Since(startTime))
}

func calculateTotalSize(dir Directory, dirs []Directory) int {
	totalSize := 0

	for _, file := range dir.files {
		totalSize += file.size
	}

	for _, d := range dirs {
		if dir.path == d.parentPath && dir.path != d.path {
			totalSize += calculateTotalSize(d, dirs)
		}
	}

	for i, d := range dirs {
		if d.path == dir.path {
			(dirs)[i].totalSize = totalSize
			(dirs)[i].files = []File{}
		}
	}

	return totalSize
}

func toInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}
