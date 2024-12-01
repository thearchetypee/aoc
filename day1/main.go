package main

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/aoc2024/helper"
)

func main() {
	if err := part1(); err != nil {
		log.Fatalf("Error in part1: %v", err)
	}

	if err := part2(); err != nil {
		log.Fatalf("Error in part2: %v", err)
	}
}

// part1 processes the first part of the problem
func part1() error {
	lines, err := helper.ReadFileLineByLine("input1.txt")
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	list1, list2, err := processLines(lines)
	if err != nil {
		return fmt.Errorf("processing lines: %w", err)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	distance := calculateDistance(list1, list2)
	fmt.Printf("Total Distance: %d\n", distance)
	return nil
}

// part2 processes the second part of the problem using maps for both lists
func part2() error {
	lines, err := helper.ReadFileLineByLine("input2.txt")
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	// Maps to store counts for both lists
	list1Counts := make(map[int]int)
	list2Counts := make(map[int]int)

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return fmt.Errorf("invalid line format: %s", line)
		}

		// Parse and count first number
		val1, err := parseLocation(fields[0])
		if err != nil {
			return fmt.Errorf("parsing location 1: %w", err)
		}
		list1Counts[val1]++

		// Parse and count second number
		val2, err := parseLocation(fields[1])
		if err != nil {
			return fmt.Errorf("parsing location 2: %w", err)
		}
		list2Counts[val2]++
	}

	total := 0
	// Calculate total by multiplying matching counts
	for val, count1 := range list1Counts {
		if count2, exists := list2Counts[val]; exists {
			total += val * count1 * count2
		}
	}

	fmt.Printf("Total Distance part2: %d\n", total)
	return nil
}

// parseLocation converts a string to int with error handling
func parseLocation(s string) (int, error) {
	if s == "" {
		return 0, nil
	}
	return strconv.Atoi(s)
}

// processLines parses input lines and returns two lists of locations
func processLines(lines []string) ([]int, []int, error) {
	// Pre-allocate slices for better performance
	list1 := make([]int, 0, len(lines))
	list2 := make([]int, 0, len(lines))

	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		loc1, err := parseLocation(fields[0])
		if err != nil {
			return nil, nil, fmt.Errorf("parsing location 1: %w", err)
		}

		loc2, err := parseLocation(fields[1])
		if err != nil {
			return nil, nil, fmt.Errorf("parsing location 2: %w", err)
		}

		list1 = append(list1, loc1)
		list2 = append(list2, loc2)
	}

	return list1, list2, nil
}

// calculateDistance computes total distance between sorted lists
func calculateDistance(list1, list2 []int) int {
	minLen := len(list1)
	if len(list2) < minLen {
		minLen = len(list2)
	}

	total := 0
	for i := 0; i < minLen; i++ {
		total += helper.Abs(list1[i] - list2[i])
	}
	return total
}
