package main

import (
	"fmt"

	"github.com/aoc2024/helper"
)

type Data struct {
	a, b, c, d, x, y int64
}

func parseInput(lines []string) []Data {
	var result []Data
	for i := 0; i < len(lines); i += 4 {
		if i+2 >= len(lines) {
			break
		}

		var data Data
		var tmp1, tmp2 int64

		fmt.Sscanf(lines[i], "Button A: X+%d, Y+%d", &data.a, &data.b)
		fmt.Sscanf(lines[i+1], "Button B: X+%d, Y+%d", &data.c, &data.d)
		fmt.Sscanf(lines[i+2], "Prize: X=%d, Y=%d", &tmp1, &tmp2)
		data.x = tmp1
		data.y = tmp2

		result = append(result, data)
	}
	return result
}

type Fraction struct {
	num, den int64
}

func newFraction(num, den int64) Fraction {
	if den < 0 {
		num, den = -num, -den
	}
	g := helper.Gcd64(num, den)
	return Fraction{num / g, den / g}
}

func (f Fraction) mul(x int64) Fraction {
	return newFraction(f.num*x, f.den)
}

func (f Fraction) isInt() bool {
	return f.den == 1
}

func (f Fraction) toInt() int64 {
	return f.num / f.den
}

func solve(a, b, c, d, x, y int64) (int64, int64) {
	det := a*d - b*c
	if det == 0 {
		return 0, 0
	}

	detF := newFraction(1, det)

	aF := detF.mul(d)
	bF := detF.mul(-b)
	cF := detF.mul(-c)
	dF := detF.mul(a)

	ra := newFraction(aF.num*x*cF.den+cF.num*y*aF.den, aF.den*cF.den)
	rb := newFraction(bF.num*x*dF.den+dF.num*y*bF.den, bF.den*dF.den)

	if ra.isInt() && rb.isInt() {
		return ra.toInt(), rb.toInt()
	}

	return 0, 0
}

func solve_puzzle(input []string) (int, int) {
	part1, part2 := 0, 0

	data := parseInput(input)
	for _, d := range data {
		a, b := solve(d.a, d.b, d.c, d.d, d.x, d.y)
		if a != -1 {
			part1 += int(a*3 + b)
		}

		offset := int64(10000000000000)
		a, b = solve(d.a, d.b, d.c, d.d, d.x+offset, d.y+offset)
		if a != -1 {
			part2 += int(a*3 + b)
		}
	}

	return part1, part2
}

func main() {
	input, err := helper.ReadFileLineByLine("input.txt")
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}
	part1, part2 := solve_puzzle(input)
	fmt.Printf("Part 1: %d\n", part1)
	fmt.Printf("Part 2: %d\n", part2)
}
