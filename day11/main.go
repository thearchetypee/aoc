package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/helper"
)

func processStone(stone int64) []int64 {
	strStone := strconv.FormatInt(stone, 10)

	if stone == 0 {
		return []int64{1}
	}

	if len(strStone)%2 == 0 {
		mid := len(strStone) / 2
		leftStr := strStone[:mid]
		rightStr := strStone[mid:]

		left, _ := strconv.ParseInt(leftStr, 10, 64)
		right, _ := strconv.ParseInt(rightStr, 10, 64)
		return []int64{left, right}
	}

	return []int64{stone * 2024}
}

func blink(stones []int64) []int64 {
	var result []int64

	for _, stone := range stones {
		result = append(result, processStone(stone)...)
	}

	return result
}

func solve(input []string) (int, int) {
	var stones []int64

	for _, numStr := range strings.Fields(input[0]) {
		num, _ := strconv.ParseInt(numStr, 10, 64)
		stones = append(stones, num)
	}

	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return len(stones), 0
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
