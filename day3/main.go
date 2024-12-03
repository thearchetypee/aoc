package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"

	"github.com/aoc2024/helper"
)

type Instruction struct {
	typ      string
	position int
	x, y     int
}

func processCorruptedMemory(input string, enabled bool) (int, bool) {
	mulPattern := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	doPattern := regexp.MustCompile(`do\(\)`)
	dontPattern := regexp.MustCompile(`don't\(\)`)

	var instructions []Instruction

	mulMatches := mulPattern.FindAllStringSubmatchIndex(input, -1)
	for _, match := range mulMatches {
		x, _ := strconv.Atoi(input[match[2]:match[3]])
		y, _ := strconv.Atoi(input[match[4]:match[5]])
		instructions = append(instructions, Instruction{
			typ:      "mul",
			position: match[0],
			x:        x,
			y:        y,
		})
	}

	doMatches := doPattern.FindAllStringIndex(input, -1)
	for _, match := range doMatches {
		instructions = append(instructions, Instruction{
			typ:      "do",
			position: match[0],
		})
	}

	dontMatches := dontPattern.FindAllStringIndex(input, -1)
	for _, match := range dontMatches {
		instructions = append(instructions, Instruction{
			typ:      "dont",
			position: match[0],
		})
	}

	sort.Slice(instructions, func(i, j int) bool {
		return instructions[i].position < instructions[j].position
	})

	sum := 0
	for _, inst := range instructions {
		switch inst.typ {
		case "mul":
			if enabled {
				result := inst.x * inst.y
				sum += result
			}
		case "do":
			enabled = true
		case "dont":
			enabled = false
		}
	}

	return sum, enabled
}

func extractMultiplicationPairs(input string) ([][2]int, error) {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(input, -1)
	pairs := make([][2]int, 0, len(matches))

	for _, match := range matches {
		x, err := strconv.Atoi(match[1])
		if err != nil {
			return nil, fmt.Errorf("error converting first number: %v", err)
		}

		y, err := strconv.Atoi(match[2])
		if err != nil {
			return nil, fmt.Errorf("error converting second number: %v", err)
		}

		pairs = append(pairs, [2]int{x, y})
	}

	return pairs, nil
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0
	// Keep track of enable and disable from previous line
	enabled := true
	for _, line := range input {
		pairs, err := extractMultiplicationPairs(line)
		if err != nil {
			fmt.Printf("Error extracting input: %v\n", err)
		}
		for _, pair := range pairs {
			part1 += pair[0] * pair[1]
		}
		lineSum := 0
		lineSum, enabled = processCorruptedMemory(line, enabled)
		part2 += lineSum
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
