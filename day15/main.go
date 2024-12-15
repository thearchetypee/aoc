package main

import (
	"fmt"
	"strings"

	"github.com/aoc2024/helper"
)

type Movement struct {
	x, y int
}

var movements = map[rune]Movement{
	'<': {-1, 0},
	'>': {1, 0},
	'^': {0, -1},
	'v': {0, 1},
}

type Location struct {
	x, y int
}

func processWarehouseInput(input []string) ([][]rune, Location, string) {
	warehouse := make([][]rune, len(input))
	var robotStart Location
	var instructions string

	parsingMap := true
	for i, line := range input {
		if line == "" {
			parsingMap = false
			continue
		}

		if parsingMap {
			warehouse[i] = []rune(line)
			for j, ch := range warehouse[i] {
				if ch == '@' {
					robotStart = Location{j, i}
					warehouse[i][j] = '.'
				}
			}
		} else {
			instructions += strings.TrimSpace(line)
		}
	}

	return warehouse, robotStart, instructions
}

func cloneWarehouse(warehouse [][]rune) [][]rune {
	newWarehouse := make([][]rune, len(warehouse))
	for i := range warehouse {
		newWarehouse[i] = make([]rune, len(warehouse[i]))
		copy(newWarehouse[i], warehouse[i])
	}
	return newWarehouse
}

func moveRobot(warehouse [][]rune, robotPos Location, instruction rune) ([][]rune, Location) {
	move := movements[instruction]
	pathLength := 1

	for warehouse[robotPos.y+pathLength*move.y][robotPos.x+pathLength*move.x] == 'O' {
		pathLength++
	}

	if warehouse[robotPos.y+pathLength*move.y][robotPos.x+pathLength*move.x] == '#' {
		return warehouse, robotPos
	}

	updatedWarehouse := cloneWarehouse(warehouse)
	if pathLength > 1 {
		updatedWarehouse[robotPos.y+pathLength*move.y][robotPos.x+pathLength*move.x] = 'O'
		updatedWarehouse[robotPos.y+move.y][robotPos.x+move.x] = '.'
	}
	return updatedWarehouse, Location{robotPos.x + move.x, robotPos.y + move.y}
}

func moveWideRobot(warehouse [][]rune, robotPos Location, instruction rune) ([][]rune, Location) {
	move := movements[instruction]
	updatedWarehouse := cloneWarehouse(warehouse)

	if warehouse[robotPos.y+move.y][robotPos.x+move.x] == '#' {
		return warehouse, robotPos
	}

	containers := scanForContainers(warehouse, robotPos, move)
	if len(containers) == 0 {
		return updatedWarehouse, Location{robotPos.x + move.x, robotPos.y + move.y}
	}

	for _, container := range containers {
		if warehouse[container.y+move.y][container.x+move.x] == '#' {
			return warehouse, robotPos
		}
		if warehouse[container.y+move.y][container.x+1+move.x] == '#' {
			return warehouse, robotPos
		}
	}

	for _, container := range containers {
		updatedWarehouse[container.y][container.x] = '.'
		updatedWarehouse[container.y][container.x+1] = '.'
	}

	for _, container := range containers {
		updatedWarehouse[container.y+move.y][container.x+move.x] = '['
		updatedWarehouse[container.y+move.y][container.x+1+move.x] = ']'
	}

	return updatedWarehouse, Location{robotPos.x + move.x, robotPos.y + move.y}
}

func scanForContainers(warehouse [][]rune, pos Location, move Movement) []Location {
	containers := make([]Location, 0)
	x, y := pos.x+move.x, pos.y+move.y

	var scanContainer func(int, int)
	scanContainer = func(bx, by int) {
		if warehouse[by][bx] == '[' {
			containers = append(containers, Location{bx, by})
			scanContainer(bx+move.x, by+move.y)
			if move.x != -1 {
				scanContainer(bx+1+move.x, by+move.y)
			}
		} else if warehouse[by][bx] == ']' {
			containers = append(containers, Location{bx - 1, by})
			if move.x != 1 {
				scanContainer(bx-1+move.x, by+move.y)
			}
			scanContainer(bx+move.x, by+move.y)
		}
	}

	scanContainer(x, y)
	return containers
}

func expandWarehouse(warehouse [][]rune) [][]rune {
	wideWarehouse := make([][]rune, len(warehouse))
	for i, row := range warehouse {
		expandedRow := make([]rune, len(row)*2)
		for j, cell := range row {
			switch cell {
			case '#':
				expandedRow[j*2] = '#'
				expandedRow[j*2+1] = '#'
			case 'O':
				expandedRow[j*2] = '['
				expandedRow[j*2+1] = ']'
			case '.':
				expandedRow[j*2] = '.'
				expandedRow[j*2+1] = '.'
			case '@':
				expandedRow[j*2] = '@'
				expandedRow[j*2+1] = '.'
			}
		}
		wideWarehouse[i] = expandedRow
	}
	return wideWarehouse
}

func executeInstructions(warehouse [][]rune, start Location, instructions string) ([][]rune, Location) {
	currentWarehouse := cloneWarehouse(warehouse)
	currentPos := start

	for _, instruction := range instructions {
		currentWarehouse, currentPos = moveRobot(currentWarehouse, currentPos, instruction)
	}

	return currentWarehouse, currentPos
}

func executeWideInstructions(warehouse [][]rune, start Location, instructions string) ([][]rune, Location) {
	wideWarehouse := expandWarehouse(warehouse)
	currentPos := Location{start.x * 2, start.y}

	for _, instruction := range instructions {
		wideWarehouse, currentPos = moveWideRobot(wideWarehouse, currentPos, instruction)
	}

	return wideWarehouse, currentPos
}

func calculateWarehouseScore(warehouse [][]rune) int {
	score := 0
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == 'O' {
				score += y*100 + x
			}
		}
	}
	return score
}

func calculateWideWarehouseScore(warehouse [][]rune) int {
	score := 0
	for y, row := range warehouse {
		for x, cell := range row {
			if cell == '[' {
				score += y*100 + x
			}
		}
	}
	return score
}

func solve(input []string) (int, int) {
	warehouse, robotStart, instructions := processWarehouseInput(input)

	finalWarehouse, _ := executeInstructions(warehouse, robotStart, instructions)
	normalScore := calculateWarehouseScore(finalWarehouse)

	wideWarehouse, _ := executeWideInstructions(warehouse, robotStart, instructions)
	wideScore := calculateWideWarehouseScore(wideWarehouse)

	return normalScore, wideScore
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
