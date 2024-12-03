package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/helper"
)

func checkAdjacent(nums []int) bool {
	for i := 1; i < len(nums); i++ {
		diff := helper.Abs(nums[i] - nums[i-1])
		if diff < 1 || diff > 3 {
			return false
		}
	}
	return true
}

func isMonotonic(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	increasing := true
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			increasing = false
			break
		}
	}

	decreasing := true
	if !increasing {
		for i := 1; i < len(nums); i++ {
			if nums[i] >= nums[i-1] {
				decreasing = false
				break
			}
		}
	}

	return increasing || decreasing
}

func isSafe(nums []int) bool {
	return isMonotonic(nums) && checkAdjacent(nums)
}

func isSafeWithDampener(nums []int) bool {
	if isSafe(nums) {
		return true
	}

	for i := 0; i < len(nums); i++ {
		dampened := make([]int, 0, len(nums)-1)
		dampened = append(dampened, nums[:i]...)
		dampened = append(dampened, nums[i+1:]...)
		if isSafe(dampened) { // check if list is safe by removing current element
			return true
		}
	}
	return false
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0

	for _, line := range input {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		numStrs := strings.Fields(line)
		nums := make([]int, len(numStrs))

		for i, numStr := range numStrs {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			nums[i] = num
		}

		if isSafe(nums) {
			part1++
		}
		if isSafeWithDampener(nums) {
			part2++
		}
	}

	return part1, part2
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
