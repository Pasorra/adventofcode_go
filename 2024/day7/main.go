package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

func stringSliceToInt(strings []string) []int {
	ints := make([]int, len(strings))
	for i, str := range strings {
		ints[i], _ = strconv.Atoi(str)
	}
	return ints
}

type Part1 struct{}

func (p Part1) recurse(target int, soFar int, nums []int, i int) bool {
	if soFar > target {
		return false
	}

	if i >= len(nums) {
		return soFar == target
	}

	return p.recurse(target, soFar+nums[i], nums, i+1) || p.recurse(target, soFar*nums[i], nums, i+1)

}

func (p Part1) Solve(lines []string) int {
	res := 0
	for _, line := range lines {
		splitRes := strings.Split(line, ": ")
		target, _ := strconv.Atoi(splitRes[0])
		numbers := stringSliceToInt(strings.Split(splitRes[1], " "))
		if p.recurse(target, 0, numbers, 0) {
			res += target
		}
	}
	return res
}

type Part2 struct{}

func (p Part2) concatNumbersFast(n1 int, n2 int) int {
	multiplicator := 1
	n := n2
	for n != 0 {
		n /= 10
		multiplicator *= 10
	}
	n1 *= multiplicator
	n1 += n2
	return n1
}

func (p Part2) recurse(target int, soFar int, nums []int, i int) bool {
	if soFar > target {
		return false
	}

	if i >= len(nums) {
		return soFar == target
	}

	return p.recurse(target, soFar+nums[i], nums, i+1) || p.recurse(target, soFar*nums[i], nums, i+1) || p.recurse(target, p.concatNumbersFast(soFar, nums[i]), nums, i+1)

}

func (p Part2) Solve(lines []string) int {
	res := 0
	for _, line := range lines {
		splitRes := strings.Split(line, ": ")
		target, _ := strconv.Atoi(splitRes[0])
		numbers := stringSliceToInt(strings.Split(splitRes[1], " "))
		if p.recurse(target, 0, numbers, 0) {
			res += target
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
