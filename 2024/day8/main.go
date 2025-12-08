package main

import (
	"fmt"
	"utils"
)

type Vector struct {
	X int
	Y int
}

func sign(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	}
	panic("Should not happen")
}

type Part1 struct{}

func (p Part1) Solve(lines []string) int {
	var antennas [256][]Vector
	antinodeSet := make(map[Vector]bool)

	for y, line := range lines {
		for x, ch := range line {
			if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				antennas[ch] = append(antennas[ch], Vector{X: x, Y: y})
			}
		}
	}

	for _, antenna := range antennas {
		for _, position := range antenna {
			for _, other := range antenna {
				if position == other {
					continue
				}
				antinode := Vector{X: 2*position.X - other.X, Y: 2*position.Y - other.Y}

				if antinode.Y >= 0 && antinode.Y < len(lines) && antinode.X >= 0 && antinode.X < len(lines[antinode.Y]) {
					antinodeSet[antinode] = true
				}
			}
		}
	}

	return len(antinodeSet)
}

type Part2 struct{}

func (p Part2) Solve(lines []string) int {
	var antennas [256][]Vector
	antinodeSet := make(map[Vector]bool)

	for y, line := range lines {
		for x, ch := range line {
			if (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
				antennas[ch] = append(antennas[ch], Vector{X: x, Y: y})
			}
		}
	}

	for _, antenna := range antennas {
		for _, position := range antenna {
			for _, other := range antenna {
				if position == other {
					continue
				}
				dx := other.X - position.X
				dy := other.Y - position.Y
				antinode := Vector{X: position.X - dx, Y: position.Y - dy}

				for antinode.Y >= 0 && antinode.Y < len(lines) && antinode.X >= 0 && antinode.X < len(lines[antinode.Y]) {
					antinodeSet[antinode] = true
					antinode.X -= dx
					antinode.Y -= dy
				}
			}
		}
	}

	return len(antinodeSet)
}
func main() {
	lines := utils.Read()

	p1 := Part1{}
	p2 := Part2{}

	fmt.Println("Part 1:", p1.Solve(lines))
	fmt.Println("Part 2:", p2.Solve(lines))
}
