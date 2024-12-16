package main

import (
	"fmt"
	"github.com/aoc2024/helper"
)

type Point struct {
	x, y int
}

type State struct {
	pos   Point
	dir   int
	score int
}

func findStartEnd(grid []string) (Point, Point) {
	var start, end Point
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == 'S' {
				start = Point{x, y}
			} else if grid[y][x] == 'E' {
				end = Point{x, y}
			}
		}
	}
	return start, end
}

func solve(input []string) (int, int) {
	start, end := findStartEnd(input)

	dx := []int{1, 0, -1, 0} // East, South, West, North
	dy := []int{0, 1, 0, -1}

	// Track minimum scores to reach each point from each direction
	minScores := make(map[string]int)
	queue := []State{{pos: start, dir: 0, score: 0}}
	minEndScore := -1

	// Phase 1: Find minimum end score
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.pos == end {
			if minEndScore == -1 || curr.score < minEndScore {
				minEndScore = curr.score
			}
			continue
		}

		key := fmt.Sprintf("%d,%d,%d", curr.pos.x, curr.pos.y, curr.dir)
		if score, exists := minScores[key]; exists && score <= curr.score {
			continue
		}
		minScores[key] = curr.score

		for turn := -1; turn <= 1; turn++ {
			newDir := (curr.dir + turn + 4) % 4
			turnCost := 0
			if turn != 0 {
				turnCost = 1000
			}

			newX := curr.pos.x + dx[newDir]
			newY := curr.pos.y + dy[newDir]

			if newX >= 0 && newX < len(input[0]) && newY >= 0 && newY < len(input) && input[newY][newX] != '#' {
				queue = append(queue, State{
					pos:   Point{newX, newY},
					dir:   newDir,
					score: curr.score + turnCost + 1,
				})
			}
		}
	}

	// Phase 2: Backtrack from end to find all optimal paths
	bestPaths := make(map[Point]bool)
	queue = []State{{pos: end, dir: 0, score: minEndScore}}
	visited := make(map[string]bool)
	bestPaths[end] = true
	bestPaths[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for turn := -1; turn <= 1; turn++ {
			for prevDir := 0; prevDir < 4; prevDir++ {
				newDir := (prevDir + turn + 4) % 4
				turnCost := 0
				if turn != 0 {
					turnCost = 1000
				}

				prevX := curr.pos.x - dx[newDir]
				prevY := curr.pos.y - dy[newDir]
				prevPos := Point{prevX, prevY}

				if prevX >= 0 && prevX < len(input[0]) && prevY >= 0 && prevY < len(input) && input[prevY][prevX] != '#' {
					key := fmt.Sprintf("%d,%d,%d", prevX, prevY, prevDir)
					if score, exists := minScores[key]; exists {
						if score+turnCost+1 == curr.score {
							bestPaths[prevPos] = true
							nextKey := fmt.Sprintf("%d,%d,%d,%d", prevX, prevY, prevDir, score)
							if !visited[nextKey] {
								visited[nextKey] = true
								queue = append(queue, State{
									pos:   prevPos,
									dir:   prevDir,
									score: score,
								})
							}
						}
					}
				}
			}
		}
	}
	return minEndScore, len(bestPaths)
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
