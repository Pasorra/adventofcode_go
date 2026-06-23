package main

import (
	"fmt"
	"utils"
)

type Part1 struct{}

func (p Part1) CheckDirections(lines []string, x, y int) bool {
	found := 0
	for yD := -1; yD <= 1; yD++ {
		for xD := -1; xD <= 1; xD++ {
			newX := x + xD
			newY := y + yD

			if xD == 0 && yD == 0 {
				continue
			}

			if newY < 0 || newY >= len(lines) {
				continue
			}
			if newX < 0 || newX >= len(lines[newY]) {
				continue
			}

			if lines[newY][newX] == '@' {
				found += 1
			}
		}
	}

	return found < 4
}

func (p Part1) Solve(lines []string) int {
	res := 0
	for y, l := range lines {
		for x, c := range l {
			if c == '@' && p.CheckDirections(lines, x, y) {
				res += 1
			}
		}
	}
	return res
}

type Part2 struct{}

func (p Part2) CheckDirections(grid [][]byte, x, y int) bool {
	found := 0
	for yD := -1; yD <= 1; yD++ {
		for xD := -1; xD <= 1; xD++ {
			newX := x + xD
			newY := y + yD

			if xD == 0 && yD == 0 {
				continue
			}

			if newY < 0 || newY >= len(grid) {
				continue
			}
			if newX < 0 || newX >= len(grid[newY]) {
				continue
			}

			if grid[newY][newX] == '@' {
				found += 1
			}
		}
	}

	return found < 4
}

func (p Part2) Solve(lines []string) int {
	res := 0
	grid := make([][]byte, len(lines))
	for i, line := range lines {
		grid[i] = []byte(line)
	}
	changed := true
	for changed {
		changed = false
		for y, l := range grid {
			for x, c := range l {
				if c == '@' && p.CheckDirections(grid, x, y) {
					res += 1
					changed = true
					grid[y][x] = '.'
				}
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
