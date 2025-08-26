package main

import (
	"fmt"
	"utils"
)

type Part1 struct{}

const xmas = "XMAS"

func outOfBounds(lines []string, x int, y int) bool {
	return (y < 0 || y >= len(lines)) || (x < 0 || x >= len(lines[y]))
}

func (p Part1) checkDirection(lines []string, x int, y int, offsetX int, offsetY int) int {
	for i := range 4 {
		if outOfBounds(lines, x, y) || lines[y][x] != xmas[i] {
			return 0
		}
		if i == 3 { // last character found
			return 1
		}
		x += offsetX
		y += offsetY
	}
	return 0
}

func (p Part1) checkDirections(lines []string, x int, y int) int {
	dirs := []int{-1, 0, 1}
	sum := 0
	for _, offsetX := range dirs {
		for _, offsetY := range dirs {
			if offsetX == 0 && offsetY == 0 {
				continue
			}
			sum += p.checkDirection(lines, x, y, offsetX, offsetY)
		}
	}
	return sum
}

func (p Part1) Solve(lines []string) int {
	res := 0
	for y, line := range lines {
		for x, ch := range line {
			if byte(ch) == 'X' {
				res += p.checkDirections(lines, x, y)
			}
		}
	}
	return res
}

type Part2 struct{}

func (p Part2) checkCorners(lines []string, x int, y int) int {
	boundCheckOffsets := []int{-1, 1}
	for _, yOffset := range boundCheckOffsets {
		for _, xOffset := range boundCheckOffsets {
			if outOfBounds(lines, x+xOffset, y+yOffset) {
				return 0
			}
		}
	}

	tl := lines[y-1][x-1]
	tr := lines[y-1][x+1]
	bl := lines[y+1][x-1]
	br := lines[y+1][x+1]

	res := 0
	if (tl == 'S' && br == 'M') || (tl == 'M' && br == 'S') {
		res += 1
	}
	if (tr == 'S' && bl == 'M') || (tr == 'M' && bl == 'S') {
		res += 1
	}

	if res == 2 {
		return 1
	}
	return 0
}

func (p Part2) Solve(lines []string) int {
	res := 0
	for y, line := range lines {
		for x, ch := range line {
			if byte(ch) == 'A' {
				res += p.checkCorners(lines, x, y)
			}
		}
	}
	return res
}

func main() {
	lines := utils.Read()

	p1 := Part1{}
	p2 := Part2{}

	fmt.Println("Part 1:", p1.Solve(lines))
	fmt.Println("Part 2:", p2.Solve(lines))
}
