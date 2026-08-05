package main

import (
	"fmt"
	"utils"
)

type Part1 struct{}

func (p Part1) Solve(lines []string) int {
	res := 0
	beamPositions := make([]bool, len(lines[0]))
	beamPositions[(len(lines[0])-1)/2] = true

	for _, line := range lines {
		for pos, ch := range line {
			if ch == '^' {
				if beamPositions[pos] {
					res += 1
					beamPositions[pos-1] = true
					beamPositions[pos+1] = true
					beamPositions[pos] = false
				}
			}
		}
	}

	return res
}

type Part2 struct{}

func (p Part2) Solve(lines []string) int {
	beamPositions := make([]int, len(lines[0]))
	beamPositions[(len(lines[0])-1)/2] = 1

	for _, line := range lines {
		for pos, ch := range line {
			if ch == '^' {
				if beamPositions[pos] != 0 {
					beamPositions[pos-1] += beamPositions[pos]
					beamPositions[pos+1] += beamPositions[pos]
					beamPositions[pos] = 0
				}
			}
		}
	}
	res := 0
	for _, n := range beamPositions {
		res += n
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
