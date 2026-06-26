package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
	"utils"
)

type Part1 struct{}

type FreshRange struct {
	lower int
	upper int
}

func (p Part1) Solve(lines []string) int {
	res := 0
	checkingFresh := true
	ranges := make([]FreshRange, 0)
	for _, line := range lines {
		if line == "" {
			checkingFresh = false
			continue
		}

		if checkingFresh {
			split := strings.Split(line, "-")
			lower, _ := strconv.Atoi(split[0])
			upper, _ := strconv.Atoi(split[1])
			ranges = append(ranges, FreshRange{lower: lower, upper: upper})
		} else {
			val, _ := strconv.Atoi(line)
			for _, r := range ranges {
				if r.lower <= val && val <= r.upper {
					res += 1
					break
				}
			}
		}

	}
	return res
}

type Part2 struct{}

func (p Part2) CollapseRanges(ranges []FreshRange) []FreshRange {
	sort.Slice(ranges, func(i, j int) bool {
		if ranges[i].lower == ranges[j].lower {
			return ranges[i].upper > ranges[j].upper
		}
		return ranges[i].lower < ranges[j].lower
	})

	// 1-10 20-30 25-50 40-70
	//             ^
	// new becomes 25-70 and 40-70 gets deleted
	// 1-10 20-30 25-70
	//       ^

	for i := len(ranges) - 2; i >= 0; i-- {
		curr := &ranges[i]
		next := &ranges[i+1]
		if curr.upper >= next.lower {
			if next.upper > curr.upper {
				curr.upper = next.upper
			}
			ranges = slices.Delete(ranges, i+1, i+2)
			if i < len(ranges)-1 {
				i++ // we might have new contenders to check on the right after delete, so check on this again
			}
		}
	}

	return ranges
}

func (p Part2) Solve(lines []string) int {
	res := 0

	ranges := make([]FreshRange, 0)
	for _, line := range lines {
		if line == "" {
			break
		}

		split := strings.Split(line, "-")
		lower, _ := strconv.Atoi(split[0])
		upper, _ := strconv.Atoi(split[1])
		ranges = append(ranges, FreshRange{lower: lower, upper: upper})
	}

	ranges = p.CollapseRanges(ranges)
	for _, v := range ranges {
		res += v.upper - v.lower + 1
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
