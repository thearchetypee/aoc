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

// This is similar to part1 with one item in recurrence relation i.e. concatination
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
	if currentValue > target {
		memo[key] = false
		return memo[key]
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

/**
*Intution:
* The problem requires using two operators: addition (+) and multiplication (×).
* At each index, we must decide whether to add or multiply the current number with
* our running total. To find a valid solution, we explore every possible combination
* of these operations through recursion.
* The recurrence relation is simple: at each step, we try both adding and multiplying
* the current number, checking if either path reaches our target. The base case occurs
* when we reach the end of our input (index out of bounds) - we return true if the
* current value equals the target, false otherwise.
*
* Approach:
* 1. Top-Down Dynamic Programming with Memoization
*    - Use a recursive function that tries all possible operations at each index
*    - Cache results to avoid recomputing same states
*
* 2. State Parameters:
*    - index: current position in values array
*    - currentValue: accumulated value so far
*    - target: target value to achieve
*    - memo: map to store computed results
*
* 3. Recurrence Relation:
*    dp(index, currentValue) =
*        dp(index+1, currentValue + values[index]) OR
*        dp(index+1, currentValue * values[index])
*
* 4. Base Cases:
*    - If index == len(values): return currentValue == target
*    - If currentValue > target: return false (optimization)
*
* 5. Optimizations:
*    - Early pruning: If currentValue > target, no need to explore further
*    as both + and * will only increase the value further
*    - Memoization using a cache struct with {index, currentValue} as key
*    - This makes it faster than bottom-up as we avoid exploring impossible paths
*
* 6. Time Complexity:
*    - O(N * V) where N is length of values array and V is range of possible values
*    - With early pruning, actual time is much less as we skip impossible paths
*
* 7. Space Complexity:
*    - O(N * V) for memoization cache
*    - O(N) recursion stack depth
**/
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
	// Because we are doing + and x operation
	// that means if current value increase target then current
	// value will never be equal to target for next indexes
	// and because of that we can early return (this conditions
	// makes this algorithm faster then bottom up approach)
	if currentValue > target {
		memo[key] = false
		return memo[key]
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
