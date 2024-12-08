package main

import (
	"fmt"
	"github.com/aoc2024/helper"
)

type point struct {
	x, y int
}

type antenna struct {
	pos       point
	frequency rune
}

func findAntennas(input []string) []antenna {
	var antennas []antenna
	for y, line := range input {
		for x, char := range line {
			if char != '.' {
				antennas = append(antennas, antenna{
					pos:       point{x, y},
					frequency: rune(char),
				})
			}
		}
	}
	return antennas
}

func findAntinodes(antennas []antenna, maxX, maxY int) map[point]bool {
	antinodes := make(map[point]bool)
	freqGroups := make(map[rune][]antenna)
	for _, ant := range antennas {
		freqGroups[ant.frequency] = append(freqGroups[ant.frequency], ant)
	}

	for _, group := range freqGroups {
		for i := 0; i < len(group); i++ {
			for j := i + 1; j < len(group); j++ {
				ant1, ant2 := group[i], group[j]

				linePoints := getExtraPoints(ant1.pos, ant2.pos, maxX, maxY)
				for _, p := range linePoints {
					d1 := distance(p, ant1.pos)
					d2 := distance(p, ant2.pos)
					if (d1 == 4*d2) || (d2 == 4*d1) {
						antinodes[p] = true
					}
				}
			}
		}
	}
	return antinodes
}

func isCollinear(p1, p2, p3 point) bool {
	x1, y1 := float64(p1.x), float64(p1.y)
	x2, y2 := float64(p2.x), float64(p2.y)
	x3, y3 := float64(p3.x), float64(p3.y)
	area := x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)
	return area == 0
}

func distance(p1, p2 point) float64 {
	dx := float64(p2.x - p1.x)
	dy := float64(p2.y - p1.y)
	return dx*dx + dy*dy
}

func getExtraPoints(p1, p2 point, maxX, maxY int) []point {
	points := make([]point, 0)
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	if dx == 0 {
		for y := 0; y < maxY; y++ {
			points = append(points, point{p1.x, y})
		}
		return points
	}

	slope := float64(dy) / float64(dx)
	b := float64(p1.y) - slope*float64(p1.x)

	for x := 0; x < maxX; x++ {
		y := int(slope*float64(x) + b + 0.5)
		if y >= 0 && y < maxY {
			points = append(points, point{x, y})
		}
	}
	return points
}

func findAntinodesForPart2(group []antenna, maxX, maxY int) map[point]bool {
	antinodes := make(map[point]bool)

	if len(group) < 2 {
		return antinodes
	}

	for _, ant := range group {
		antinodes[ant.pos] = true
	}

	for i := 0; i < len(group); i++ {
		for j := i + 1; j < len(group); j++ {
			ant1, ant2 := group[i], group[j]

			for y := 0; y < maxY; y++ {
				for x := 0; x < maxX; x++ {
					p := point{x, y}
					if isCollinear(ant1.pos, ant2.pos, p) {
						antinodes[p] = true
					}
				}
			}
		}
	}
	return antinodes
}

func findAntinodesP2(antennas []antenna, maxX, maxY int) map[point]bool {
	antinodes := make(map[point]bool)
	freqGroups := make(map[rune][]antenna)

	for _, ant := range antennas {
		freqGroups[ant.frequency] = append(freqGroups[ant.frequency], ant)
	}

	for _, group := range freqGroups {
		groupAntinodes := findAntinodesForPart2(group, maxX, maxY)
		for p := range groupAntinodes {
			antinodes[p] = true
		}
	}
	return antinodes
}

func solve(input []string) (int, int) {
	maxY := len(input)
	maxX := len(input[0])

	antennas := findAntennas(input)
	antinodesPart1 := findAntinodes(antennas, maxX, maxY)
	antinodesPart2 := findAntinodesP2(antennas, maxX, maxY)

	return len(antinodesPart1), len(antinodesPart2)
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
