package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

type Point struct {
	x, y int
}

type Region struct {
	char   byte
	points map[Point]bool
}

func newRegion(char byte) *Region {
	return &Region{
		char:   char,
		points: make(map[Point]bool),
	}
}

func (r *Region) addPoint(p Point) {
	r.points[p] = true
}

func (r *Region) size() int {
	return len(r.points)
}

func doRegion(x, y int, grid []string, seen map[Point]bool) *Region {
	c := grid[y][x]
	region := newRegion(c)
	todo := []Point{{x, y}}
	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(todo) > 0 {
		curr := todo[0]
		todo = todo[1:]

		if seen[curr] {
			continue
		}

		seen[curr] = true
		region.addPoint(curr)

		for _, dir := range dirs {
			nx, ny := curr.x+dir.x, curr.y+dir.y
			next := Point{nx, ny}
			if nx >= 0 && nx < len(grid[0]) && ny >= 0 && ny < len(grid) &&
				grid[ny][nx] == c && !seen[next] {
				todo = append(todo, next)
			}
		}
	}
	return region
}

func getRegions(grid []string) []*Region {
	regions := []*Region{}
	seen := make(map[Point]bool)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			p := Point{x, y}
			if !seen[p] {
				region := doRegion(x, y, grid, seen)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func perimeter(region *Region) int {
	n := 0
	dirs := []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for p := range region.points {
		for _, dir := range dirs {
			next := Point{p.x + dir.x, p.y + dir.y}
			if !region.points[next] {
				n++
			}
		}
	}
	return n
}

func price(region *Region) int {
	return perimeter(region) * region.size()
}

func perimeter2(region *Region) int {
	n := 0
	checks := []struct {
		next, p1, p2 Point
	}{
		{Point{1, 0}, Point{0, -1}, Point{1, -1}},   // right
		{Point{-1, 0}, Point{0, -1}, Point{-1, -1}}, // left
		{Point{0, 1}, Point{-1, 0}, Point{-1, 1}},   // down
		{Point{0, -1}, Point{-1, 0}, Point{-1, -1}}, // up
	}

	for p := range region.points {
		for _, check := range checks {
			next := Point{p.x + check.next.x, p.y + check.next.y}
			p1 := Point{p.x + check.p1.x, p.y + check.p1.y}
			p2 := Point{p.x + check.p2.x, p.y + check.p2.y}

			if !region.points[next] && !(region.points[p1] && !region.points[p2]) {
				n++
			}
		}
	}
	return n
}

func price2(region *Region) int {
	return perimeter2(region) * region.size()
}

func solve(input []string) (int, int) {
	regions := getRegions(input)
	part1, part2 := 0, 0

	for _, region := range regions {
		part1 += price(region)
		part2 += price2(region)
	}

	return part1, part2
}

func main() {
	input, err := helper.ReadFileLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}
	part1, part2 := solve(input)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
