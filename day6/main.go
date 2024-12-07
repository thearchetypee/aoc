package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

const (
	up = iota
	right
	down
	left
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

type point struct {
	x, y int
}

type state struct {
	pos point
	dir int
}

func calculateDistinctPositions(lab [][]rune, startX, startY, startDir int) (int, map[point]bool) {
	rows, cols := len(lab), len(lab[0])
	visited := make(map[point]bool)
	x, y, guardDirection := startX, startY, startDir
	visited[point{x, y}] = true

	for {
		nextX := x + dx[guardDirection]
		nextY := y + dy[guardDirection]

		if nextX < 0 || nextX >= cols || nextY < 0 || nextY >= rows {
			break
		}

		if lab[nextY][nextX] == '#' {
			guardDirection = (guardDirection + 1) % 4
		} else {
			x, y = nextX, nextY
			visited[point{nextX, nextY}] = true
		}
	}

	return len(visited), visited
}

func findGuardInitialPosition(lab [][]rune) (int, int, int) {
	rows, cols := len(lab), len(lab[0])
	i, j, direc := 0, 0, up
	found := false
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			switch lab[x][y] {
			case '^':
				i = y
				j = x
				direc = up
				found = true
			case '>':
				i = y
				j = x
				direc = right
				found = true
			case 'v':
				i = y
				j = x
				direc = down
				found = true
			case '<':
				i = y
				j = x
				direc = left
				found = true
			}
			if found {
				break
			}
		}
	}

	return i, j, direc
}

func tryObstaclePosition(lab [][]rune, startX, startY, startDir int, obstaclePos point) bool {
	rows, cols := len(lab), len(lab[0])

	visited := make(map[state]bool)

	x, y := startX, startY
	dir := startDir

	for {
		state := state{point{x, y}, dir}
		if visited[state] {
			return true
		}
		visited[state] = true

		nextX := x + dx[dir]
		nextY := y + dy[dir]

		if nextX < 0 || nextX >= cols || nextY < 0 || nextY >= rows {
			return false
		}

		isObstacle := lab[nextY][nextX] == '#' ||
			(nextX == obstaclePos.x && nextY == obstaclePos.y)

		if isObstacle {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}
}

func calculateLoopPositions(lab [][]rune, startX, startY, startDir int, path map[point]bool) int {
	checkedPositions := make(map[point]bool)
	loopCount := 0

	// Try putting obstacle only on the path visited.
	// The worst case time is still same where guard have to visit every cell
	for pos := range path {
		if pos.x == startX && pos.y == startY {
			continue
		}
		if lab[pos.y][pos.x] != '.' {
			continue
		}
		if _, ok := checkedPositions[pos]; !ok {
			if tryObstaclePosition(lab, startX, startY, startDir, pos) {
				loopCount++
				checkedPositions[pos] = true
			}
		}
	}

	return loopCount
}

func buildMatrix(input []string) [][]rune {
	matrix := make([][]rune, len(input))
	for i, line := range input {
		matrix[i] = []rune(line)
	}
	return matrix
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0
	lab := buildMatrix(input)
	if len(lab) == 0 {
		return 0, 0
	}
	startX, startY, startDir := findGuardInitialPosition(lab)
	part1, path := calculateDistinctPositions(lab, startX, startY, startDir)
	part2 = calculateLoopPositions(lab, startX, startY, startDir, path)
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
