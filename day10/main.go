package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func calculateTotalPaths(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	totalScore := 0
	visited := make([][]bool, rows)
	for k := range visited {
		visited[k] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				peakMap := make(map[string]bool)
				visited[i][j] = true
				exploreTrail(visited, matrix, i, j, rows, cols, 0, peakMap)
				totalScore += len(peakMap)
			}
		}
	}
	return totalScore
}

func exploreTrail(visited [][]bool, matrix [][]int, i, j, rows, cols, currentHeight int, peakMap map[string]bool) {
	if currentHeight == 9 {
		key := fmt.Sprintf("%d,%d", i, j)
		peakMap[key] = true
		return
	}

	for _, dir := range directions {
		newRow, newCol := i+dir[0], j+dir[1]
		if newRow < 0 || newRow >= rows || newCol < 0 || newCol >= cols || visited[newRow][newCol] {
			continue
		}

		nextHeight := matrix[newRow][newCol]
		if nextHeight != currentHeight+1 {
			continue
		}

		visited[newRow][newCol] = true
		exploreTrail(visited, matrix, newRow, newCol, rows, cols, nextHeight, peakMap)
		visited[newRow][newCol] = false
	}
}

func buildMatrix(input []string) [][]int {
	matrix := make([][]int, len(input))
	for i, line := range input {
		matrix[i] = make([]int, len(line))
		for j, char := range line {
			matrix[i][j] = int(char - '0')
		}
	}
	return matrix
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0
	matrix := buildMatrix(input)
	part1 = calculateTotalPaths(matrix)
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
