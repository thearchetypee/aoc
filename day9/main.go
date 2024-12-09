package main

import (
	"fmt"
	"github.com/aoc2024/helper"
	"strconv"
)

func parseInput(input string) []int {
	var nums []int
	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		nums = append(nums, num)
	}
	return nums
}

type FileInfo struct {
	id     int
	start  int
	length int
}

func createInitialDisk(nums []int) []int {
	var disk []int
	fileID := 0

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			for j := 0; j < nums[i]; j++ {
				disk = append(disk, fileID)
			}
			fileID++
		} else {
			for j := 0; j < nums[i]; j++ {
				disk = append(disk, -1)
			}
		}
	}
	return disk
}

func findFileInfo(disk []int) []FileInfo {
	var files []FileInfo
	var currentFile FileInfo
	currentFile.id = -1

	for i := 0; i < len(disk); i++ {
		if disk[i] != -1 {
			if currentFile.id != disk[i] {
				if currentFile.id != -1 {
					files = append(files, currentFile)
				}
				currentFile = FileInfo{id: disk[i], start: i, length: 1}
			} else {
				currentFile.length++
			}
		}
	}
	if currentFile.id != -1 {
		files = append(files, currentFile)
	}
	return files
}

func findFreeSpace(disk []int, start int, length int) int {
	for i := 0; i < start; i++ {
		if disk[i] == -1 {
			canFit := true
			for j := 0; j < length && i+j < len(disk); j++ {
				if disk[i+j] != -1 {
					canFit = false
					break
				}
			}
			if canFit && i+length <= start {
				return i
			}
		}
	}
	return -1
}

func moveFilePart1(disk []int) bool {
	firstEmpty := -1
	lastFile := -1
	lastFileStart := -1

	for i := 0; i < len(disk); i++ {
		if disk[i] == -1 && firstEmpty == -1 {
			firstEmpty = i
		}
		if disk[i] != -1 {
			lastFile = disk[i]
			lastFileStart = i
		}
	}

	if firstEmpty == -1 || lastFileStart <= firstEmpty {
		return false
	}

	fileLength := 0
	for i := lastFileStart; i < len(disk) && disk[i] == lastFile; i++ {
		fileLength++
	}

	for i := 0; i < fileLength; i++ {
		disk[firstEmpty+i] = lastFile
		disk[lastFileStart+i] = -1
	}

	return true
}

func moveFile(disk []int, file FileInfo) bool {
	newStart := findFreeSpace(disk, file.start, file.length)
	if newStart == -1 {
		return false
	}

	for i := 0; i < file.length; i++ {
		disk[newStart+i] = file.id
		disk[file.start+i] = -1
	}
	return true
}

func compactDiskPart2(disk []int) {
	files := findFileInfo(disk)

	// Process files in descending order of ID
	for i := len(files) - 1; i >= 0; i-- {
		moveFile(disk, files[i])
	}
}

func calculateChecksum(disk []int) int {
	checksum := 0
	for pos, fileID := range disk {
		if fileID != -1 {
			checksum += pos * fileID
		}
	}
	return checksum
}

func solve(input []string) (int, int) {
	nums := parseInput(input[0])

	// Part 1
	disk1 := createInitialDisk(nums)
	disk2 := make([]int, len(disk1))
	copy(disk2, disk1)
	for moveFilePart1(disk1) {
	}
	compactDiskPart2(disk2)
	checksum1 := calculateChecksum(disk1)
	checksum2 := calculateChecksum(disk2)

	return checksum1, checksum2
}

func main() {
	input, err := helper.ReadFileLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
