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

				linePoints := findPointsOnLine(ant1.pos, ant2.pos, maxX, maxY)
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

func distance(p1, p2 point) float64 {
	dx := float64(p2.x - p1.x)
	dy := float64(p2.y - p1.y)
	return dx*dx + dy*dy
}

/**
 Bresenham's Line-Drawing Algorithm:
This is tricky part I used to optimise my algorithm. Instead of checking every
position, I used the Bresenham algorithm tofind all the points that lie on the
line between two given points. Let's understand this algorithm - Imagine you're
drawing a line between two points on a pixel grid. At each step, you need to
decide which pixel to color next. The key insight of Bresenham's algorithm is
that it makes this decision using only integer arithmetic, avoiding the
computationally expensive floating-point calculations typically required to
determine the exact path of the line.
**/

func findPointsOnLine(p1, p2 point, maxX, maxY int) []point {
	points := []point{}
	/**
		Vector Calculation: we calculate the direction vector between two points.
		I can't add diagram here so think of it as the total horizontal and vertical
		distances we need to cover.
		  /|
	     / | dy = y2-y1
		/__|
		dx = x2-x1
		**/
	dx := p2.x - p1.x
	dy := p2.y - p1.y

	if dx == 0 && dy == 0 {
		return []point{p1}
	}

	/**
		The key mathematical insight is using GCD to find the smallest
		possible step size. Why GCD? Consider a line from (0,0) to (6,4):
			- Raw vector is (6,4)
			- GCD(6,4) = 2
			- Therefore minimal step vector is (3,2)
			- This ensures we hit every possible integer point
			  on the line(because we are dealing with matrix indexes)
	**/
	g := helper.Gcd(helper.Abs(dx), helper.Abs(dy))
	stepX := dx / g
	stepY := dy / g

	// Now I traverse in both direction to get points on both side on the line.
	// I don't want points outside matrix so I handled that too.
	x, y := p1.x, p1.y
	for x >= 0 && x < maxX && y >= 0 && y < maxY {
		points = append(points, point{x, y})
		x -= stepX
		y -= stepY
	}

	x, y = p1.x+stepX, p1.y+stepY
	for x >= 0 && x < maxX && y >= 0 && y < maxY {
		points = append(points, point{x, y})
		x += stepX
		y += stepY
	}

	return points
}

func findAntinodesForPart2(group []antenna, maxX, maxY int) map[point]bool {
	antinodes := make(map[point]bool)
	if len(group) < 2 {
		return antinodes
	}

	for i := 0; i < len(group)-1; i++ {
		for j := i + 1; j < len(group); j++ {
			points := findPointsOnLine(group[i].pos, group[j].pos, maxX, maxY)
			for _, p := range points {
				antinodes[p] = true
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
