package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

const (
	UP = iota
	RIGHT
	DOWN
	LEFT
)

var dx = []int{0, 1, 0, -1}
var dy = []int{-1, 0, 1, 0}

type Point struct {
	x, y int
}

type State struct {
	pos Point
	dir int
}

func calculateDistinctPositions(lab [][]rune) int {
	if len(lab) == 0 {
		return 0
	}
	rows, cols := len(lab), len(lab[0])

	visited := make(map[Point]bool)

	x, y, guardDirection := findGuardInitialPosition(lab, rows, cols)

	visited[Point{x, y}] = true

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
			visited[Point{nextX, nextY}] = true
		}
	}

	return len(visited)
}

func findGuardInitialPosition(lab [][]rune, rows, cols int) (int, int, int) {
	i, j, direc := 0, 0, UP
	found := false
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			switch lab[x][y] {
			case '^':
				i = y
				j = x
				direc = UP
				found = true
			case '>':
				i = y
				j = x
				direc = RIGHT
				found = true
			case 'v':
				i = y
				j = x
				direc = DOWN
				found = true
			case '<':
				i = y
				j = x
				direc = LEFT
				found = true
			}
			if found {
				break
			}
		}
	}

	return i, j, direc
}

func tryObstaclePosition(lab [][]rune, startX, startY, startDir int, obstaclePos Point) bool {
	rows, cols := len(lab), len(lab[0])

	originalValue := lab[obstaclePos.y][obstaclePos.x]
	lab[obstaclePos.y][obstaclePos.x] = '#'

	visited := make(map[State]bool)

	x, y := startX, startY
	dir := startDir

	for {
		state := State{Point{x, y}, dir}
		if visited[state] {
			lab[obstaclePos.y][obstaclePos.x] = originalValue
			return true
		}
		visited[state] = true

		nextX := x + dx[dir]
		nextY := y + dy[dir]

		if nextX < 0 || nextX >= cols || nextY < 0 || nextY >= rows {
			lab[obstaclePos.y][obstaclePos.x] = originalValue
			return false
		}

		if lab[nextY][nextX] == '#' {
			dir = (dir + 1) % 4
		} else {
			x, y = nextX, nextY
		}
	}
}

func calculateLoopPositions(lab [][]rune) int {
	if len(lab) == 0 {
		return 0
	}

	rows, cols := len(lab), len(lab[0])
	startX, startY, startDir := findGuardInitialPosition(lab, rows, cols)
	loopCount := 0

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if lab[y][x] != '.' {
				continue
			}
			if x == startX && y == startY {
				continue
			}

			if tryObstaclePosition(lab, startX, startY, startDir, Point{x, y}) {
				loopCount++
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
	part1 = calculateDistinctPositions(lab)
	part2 = calculateLoopPositions(lab)
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
