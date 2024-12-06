package main

import (
	"fmt"
	"github.com/aoc2024/helper"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

func parseRules(lines []string) ([]rule, int) {
	var rules []rule
	rulesEndIndex := 0

	for i, line := range lines {
		if line == "" {
			rulesEndIndex = i
			break
		}
		parts := strings.Split(line, "|")
		before, _ := strconv.Atoi(parts[0])
		after, _ := strconv.Atoi(parts[1])
		rules = append(rules, rule{before: before, after: after})
	}
	return rules, rulesEndIndex
}

func parseUpdate(line string) []int {
	var update []int
	numStrs := strings.Split(line, ",")
	for _, numStr := range numStrs {
		num, _ := strconv.Atoi(numStr)
		update = append(update, num)
	}
	return update
}

func isValidOrder(update []int, rules []rule) bool {
	positions := make(map[int]int)
	for i, page := range update {
		positions[page] = i
	}

	for _, rule := range rules {
		beforePos, beforeExists := positions[rule.before]
		afterPos, afterExists := positions[rule.after]

		if beforeExists && afterExists && beforePos > afterPos {
			return false
		}
	}
	return true
}

func getMiddlePage(update []int) int {
	return update[len(update)/2]
}

func buildDependencyGraph(rules []rule, pages []int) map[int][]int {
	graph := make(map[int][]int)
	for _, page := range pages {
		if _, exists := graph[page]; !exists {
			graph[page] = []int{}
		}
	}
	for _, rule := range rules {
		if containsPage(pages, rule.before) && containsPage(pages, rule.after) {
			graph[rule.after] = append(graph[rule.after], rule.before)
		}
	}
	return graph
}

func containsPage(pages []int, page int) bool {
	for _, p := range pages {
		if p == page {
			return true
		}
	}
	return false
}

func topologicalSort(graph map[int][]int) []int {
	visited := make(map[int]bool)
	temp := make(map[int]bool)
	order := make([]int, 0)

	var visit func(int)
	visit = func(node int) {
		if temp[node] {
			return
		}
		if !visited[node] {
			temp[node] = true
			for _, dep := range graph[node] {
				visit(dep)
			}
			visited[node] = true
			temp[node] = false
			order = append([]int{node}, order...)
		}
	}

	for node := range graph {
		if !visited[node] {
			visit(node)
		}
	}
	return order
}

func correctOrder(update []int, rules []rule) []int {
	graph := buildDependencyGraph(rules, update)
	return topologicalSort(graph)
}

func solve(input []string) (int, int) {
	part1, part2 := 0, 0

	rules, rulesEndIndex := parseRules(input)

	for i := rulesEndIndex + 1; i < len(input); i++ {
		if input[i] == "" {
			continue
		}
		update := parseUpdate(input[i])

		if isValidOrder(update, rules) {
			part1 += getMiddlePage(update)
		} else {
			correctedUpdate := correctOrder(update, rules)
			part2 += getMiddlePage(correctedUpdate)
		}
	}

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
