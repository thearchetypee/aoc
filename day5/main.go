package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

func solve(input []string) (int, int) {
	part1, part2 := 0, 0
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
