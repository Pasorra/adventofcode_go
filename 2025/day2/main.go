package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Part1 struct{}

func (p Part1) IsRepeatedTwice(num int) bool {
	s := strconv.Itoa(num)

	if len(s)%2 != 0 {
		return false
	}

	l := 0
	r := len(s) / 2

	for r < len(s) {
		if s[l] != s[r] {
			return false
		}
		r++
		l++
	}

	return true
}

func (p Part1) Solve(lines []string) int {
	ranges := strings.Split(lines[0], ",")

	res := 0

	for _, r := range ranges {
		split := strings.Split(r, "-")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		for i := left; i <= right; i++ {
			if p.IsRepeatedTwice(i) {
				res += i
			}
		}
	}
	return res
}

type Part2 struct{}

func (p Part2) RecursivelyCheck(s string, end int) bool {
	if end > len(s)/2 {
		return false
	}
	match := s[:end]
	for i := end; i < len(s); i += end {
		if i+end > len(s) || match != s[i:i+end] {
			return p.RecursivelyCheck(s, end+1)
		}
	}
	return true
}

func (p Part2) Solve(lines []string) int {
	ranges := strings.Split(lines[0], ",")

	res := 0

	for _, r := range ranges {
		split := strings.Split(r, "-")
		left, _ := strconv.Atoi(split[0])
		right, _ := strconv.Atoi(split[1])

		for i := left; i <= right; i++ {
			if p.RecursivelyCheck(strconv.Itoa(i), 1) {
				res += i
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
