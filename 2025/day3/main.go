package main

import (
	"fmt"
	"math"
	"utils"
)

type Part1 struct{}

func (p Part1) Solve(lines []string) int {
	res := 0

	for _, l := range lines {
		largest := 0
		largest_i := -1
		for i, c := range l[:len(l)-1] {
			num := int(c - '0')
			if num > largest {
				largest = num
				largest_i = i
			}
		}
		second_largest := 0
		for _, c := range l[largest_i+1:] {
			num := int(c - '0')
			if num > second_largest {
				second_largest = num
			}
		}
		res += largest*10 + second_largest
	}

	return res
}

type Part2 struct{}

func (p Part2) Solve(lines []string) int {
	res := 0

	for _, l := range lines {
		arr := [12]int{}
		largest_i := 0
		for x := 11; x >= 0; x-- {
			for i := largest_i; i < len(l)-x; i++ {
				c := l[i]
				num := int(c - '0')
				if num > arr[x] {
					arr[x] = num
					largest_i = i + 1
				}
			}
		}

		for i := 11; i >= 0; i-- {
			res += int(math.Pow10(i)) * arr[i]
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
