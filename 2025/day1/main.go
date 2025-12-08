package main

import (
	"fmt"
	"strconv"
	"utils"
)

type Part1 struct{}

func (p Part1) Solve(lines []string) int {
	dial := 50
	res := 0

	for _, line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println(err)
			return 0
		}

		switch direction {
		case 'R':
			dial = (dial + amount) % 100
		case 'L':
			dial = (dial - amount) % 100
		}

		if dial == 0 {
			res++
		}
	}
	return res
}

type Part2 struct{}

func (p Part2) Solve(lines []string) int {
	dial := 50
	res := 0

	for _, line := range lines {
		direction := line[0]
		amount, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println(err)
			return 0
		}

		switch direction {
		case 'R':
			res += (dial + amount) / 100
			dial = (dial + amount) % 100
		case 'L':
			distToZero := dial
			if dial == 0 {
				distToZero = 100
			}

			if amount >= distToZero {
				res += 1 + (amount-distToZero)/100
			}

			dial = (dial - amount) % 100
			if dial < 0 {
				dial += 100
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
