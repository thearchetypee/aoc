package main

import (
	"fmt"
	"github.com/aoc2024/helper"
)

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

func findXMAS(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	isValid := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 'X' {
				continue
			}
			for _, dir := range directions {
				ni, nj := i+dir[0], j+dir[1]
				if !isValid(ni, nj) || grid[ni][nj] != 'M' {
					continue
				}
				ni, nj = ni+dir[0], nj+dir[1]
				if !isValid(ni, nj) || grid[ni][nj] != 'A' {
					continue
				}
				ni, nj = ni+dir[0], nj+dir[1]
				if !isValid(ni, nj) || grid[ni][nj] != 'S' {
					continue
				}
				count++
			}
		}
	}
	return count
}

func findXMASPart2(grid [][]rune) int {
	rows, cols := len(grid), len(grid[0])
	count := 0

	isValid := func(i, j int) bool {
		return i >= 0 && i < rows && j >= 0 && j < cols
	}

	checkMS := func(i1, j1, i2, j2 int) bool {
		return isValid(i1, j1) && isValid(i2, j2) &&
			((grid[i1][j1] == 'M' && grid[i2][j2] == 'S') ||
				(grid[i1][j1] == 'S' && grid[i2][j2] == 'M'))
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] != 'A' {
				continue
			}

			if checkMS(i-1, j-1, i+1, j+1) &&
				checkMS(i+1, j-1, i-1, j+1) {
				count++
			}
		}
	}
	return count
}

func buildMatrix(input []string) [][]rune {
	matrix := make([][]rune, len(input))
	for i, line := range input {
		matrix[i] = []rune(line)
	}
	return matrix
}

func solve(input []string) (int, int) {
	matrix := buildMatrix(input)
	part1 := findXMAS(matrix)
	part2 := findXMASPart2(matrix)
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
