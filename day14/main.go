package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aoc2024/helper"
)

type Point struct {
	x, y int
}

type Robot struct {
	pos Point
	vel Point
}

func parseRobot(line string) Robot {
	parts := strings.Split(line, " ")
	pos := strings.TrimPrefix(parts[0], "p=")
	vel := strings.TrimPrefix(parts[1], "v=")

	posCoords := strings.Split(pos, ",")
	velCoords := strings.Split(vel, ",")

	x, _ := strconv.Atoi(posCoords[0])
	y, _ := strconv.Atoi(posCoords[1])
	vx, _ := strconv.Atoi(velCoords[0])
	vy, _ := strconv.Atoi(velCoords[1])

	return Robot{Point{x, y}, Point{vx, vy}}
}

func simulate(robots []Robot, width, height int) []Robot {
	result := make([]Robot, len(robots))
	for i, robot := range robots {
		newX := helper.Mod(robot.pos.x+robot.vel.x, width)
		newY := helper.Mod(robot.pos.y+robot.vel.y, height)
		result[i] = Robot{Point{newX, newY}, robot.vel}
	}
	return result
}

func robotsInQuads(robots []Robot, width, height int) [4]int {
	midX := width / 2
	midY := height / 2
	quads := [4]int{}

	for _, robot := range robots {
		if robot.pos.x == midX || robot.pos.y == midY {
			continue
		}

		if robot.pos.x < midX {
			if robot.pos.y < midY {
				quads[0]++
			} else {
				quads[2]++
			}
		} else {
			if robot.pos.y < midY {
				quads[1]++
			} else {
				quads[3]++
			}
		}
	}

	return quads
}

func robotDensity(robots []Robot) float64 {
	sum := 0.0
	count := 0

	for i := 0; i < len(robots); i++ {
		for j := i + 1; j < len(robots); j++ {
			dx := float64(robots[i].pos.x - robots[j].pos.x)
			dy := float64(robots[i].pos.y - robots[j].pos.y)
			sum += math.Sqrt(dx*dx + dy*dy)
			count++
		}
	}

	if count == 0 {
		return math.MaxFloat64
	}
	return sum / float64(count)
}

func solve(input []string) (int, int) {
	width, height := 101, 103
	robots := make([]Robot, 0)

	for _, line := range input {
		robots = append(robots, parseRobot(line))
	}

	// Part 1
	currentState := robots
	for i := 0; i < 100; i++ {
		currentState = simulate(currentState, width, height)
	}
	quads := robotsInQuads(currentState, width, height)
	part1 := quads[0] * quads[1] * quads[2] * quads[3]

	// Part 2
	minDensity := math.MaxFloat64
	minTime := 0
	currentState = robots

	for t := 0; t < 20000; t++ {
		density := robotDensity(currentState)
		if density < minDensity {
			minDensity = density
			minTime = t
		}
		currentState = simulate(currentState, width, height)
	}

	return part1, minTime
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
