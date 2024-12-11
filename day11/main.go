package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/aoc2024/helper"
)

const WORKERS = 8

type Stone struct {
	value uint64
	count uint64
}

type StoneMap struct {
	mu   sync.RWMutex
	data map[uint64]uint64
}

func NewStoneMap() *StoneMap {
	return &StoneMap{
		data: make(map[uint64]uint64),
	}
}

func (sm *StoneMap) Add(value uint64, count uint64) {
	sm.mu.Lock()
	sm.data[value] += count
	sm.mu.Unlock()
}

func (sm *StoneMap) GetAndClear() map[uint64]uint64 {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	result := sm.data
	sm.data = make(map[uint64]uint64)
	return result
}

func processStone(stone uint64) []Stone {
	if stone == 0 {
		return []Stone{{value: 1, count: 1}}
	}

	strStone := strconv.FormatUint(stone, 10)
	if len(strStone)%2 == 0 {
		mid := len(strStone) / 2
		leftStr := strStone[:mid]
		rightStr := strStone[mid:]

		left, _ := strconv.ParseUint(leftStr, 10, 64)
		right, _ := strconv.ParseUint(rightStr, 10, 64)
		return []Stone{{value: left, count: 1}, {value: right, count: 1}}
	}

	return []Stone{{value: stone * 2024, count: 1}}
}

func worker(jobs <-chan Stone, results *StoneMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for stone := range jobs {
		newStones := processStone(stone.value)
		for _, ns := range newStones {
			results.Add(ns.value, ns.count*stone.count)
		}
	}
}

func blink(stones map[uint64]uint64) map[uint64]uint64 {
	jobs := make(chan Stone, len(stones))
	results := NewStoneMap()
	var wg sync.WaitGroup

	for i := 0; i < WORKERS; i++ {
		wg.Add(1)
		go worker(jobs, results, &wg)
	}

	for value, count := range stones {
		jobs <- Stone{value: value, count: count}
	}
	close(jobs)

	wg.Wait()

	return results.GetAndClear()
}

func solve(input []string) (uint64, uint64) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	currentStones := make(map[uint64]uint64)
	for _, numStr := range strings.Fields(input[0]) {
		num, _ := strconv.ParseUint(numStr, 10, 64)
		currentStones[num]++
	}

	for i := 0; i < 75; i++ {
		currentStones = blink(currentStones)
		if (i+1)%10 == 0 {
			var total uint64
			for _, count := range currentStones {
				total += count
			}
			fmt.Printf("After %d blinks: %d stones (unique: %d)\n",
				i+1, total, len(currentStones))
		}
	}

	var total uint64
	for _, count := range currentStones {
		total += count
	}
	return total, 0
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
