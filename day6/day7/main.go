package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aoc2024/helper"
)

type cache struct {
	index, currentValue int
}

func isCorrectCalibrationPart2(values []int, memo map[cache]bool, index, currentValue, target int) bool {
	if index == len(values) {
		return currentValue == target
	}
	key := cache{
		index:        index,
		currentValue: currentValue,
	}
	if result, exists := memo[key]; exists {
		return result
	}
	multiplied := currentValue * values[index]
	added := currentValue + values[index]
	concatStr := fmt.Sprintf("%d%d", currentValue, values[index])
	concatenated, _ := strconv.Atoi(concatStr)
	result := isCorrectCalibrationPart2(values, memo, index+1, multiplied, target) ||
		isCorrectCalibrationPart2(values, memo, index+1, added, target) ||
		isCorrectCalibrationPart2(values, memo, index+1, concatenated, target)
	memo[key] = result
	return memo[key]
}

func isCorrectCalibration(values []int, memo map[cache]bool, index, currentValue, target int) bool {
	if index == len(values) {
		return currentValue == target
	}
	key := cache{
		index:        index,
		currentValue: currentValue,
	}
	if result, exists := memo[key]; exists {
		return result
	}
	multiplied := currentValue * values[index]
	added := currentValue + values[index]

	result := isCorrectCalibration(values, memo, index+1, multiplied, target) ||
		isCorrectCalibration(values, memo, index+1, added, target)
	memo[key] = result
	return memo[key]
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0

	for _, line := range input {
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numStrs := strings.Fields(strings.TrimSpace(parts[1]))

		values := make([]int, len(numStrs))
		for i, numStr := range numStrs {
			values[i], _ = strconv.Atoi(numStr)
		}
		memo1 := make(map[cache]bool)

		if isCorrectCalibration(values, memo1, 1, values[0], target) {
			part1 += target
		}
		memo2 := make(map[cache]bool)

		if isCorrectCalibrationPart2(values, memo2, 1, values[0], target) {
			part2 += target
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
