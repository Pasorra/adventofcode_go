package main

import (
	"fmt"
	"strconv"
	"strings"
	"utils"
)

type Part1 struct{}

func GatherFromString(line string) []string {
	arr := make([]string, 0)
	var sb strings.Builder
	for _, ch := range line {
		if ch == ' ' {
			if sb.Len() > 0 {
				arr = append(arr, sb.String())
				sb.Reset()
			}
		} else {
			sb.WriteRune(ch)
		}
	}
	if sb.Len() > 0 {
		arr = append(arr, sb.String())
	}
	return arr
}

func (p Part1) Solve(lines []string) int {
	symbols := GatherFromString(lines[len(lines)-1])
	res := make([]int, len(symbols))

	for _, l := range lines[:len(lines)-1] {
		gathered := GatherFromString(l)
		for i, numStr := range gathered {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}
			switch symbols[i] {
			case "+":
				res[i] += num
			case "*":
				if res[i] == 0 {
					res[i] = 1
				}
				res[i] *= num
			default:
				panic("The symbol has to be either + or *")
			}
		}
	}

	sum := 0
	for _, n := range res {
		sum += n
	}
	return sum
}

type Part2 struct{}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
func (p Part2) Solve(lines []string) int {
	s := 0
	var symbol byte = '*'

	problem := 0
	for x := range lines[0] {
		num := 0
		// gather column
		for y := range lines {
			b := lines[y][x]
			if b == ' ' {
				continue
			} else if isDigit(b) {
				num = num*10 + int(b-'0')
			} else {
				symbol = b
			}
		}

		// gathered an empty column, a problem is solved
		if num == 0 {
			s += problem
			problem = 0
			continue
		}

		if symbol == '*' {
			if problem == 0 {
				problem = 1
			}
			problem *= num
		} else {
			problem += num
		}
	}
	// no empty column for the last problem so manually add it
	s += problem
	return s
}

func main() {
	lines := utils.Read()

	p1 := Part1{}
	p2 := Part2{}

	fmt.Println("Part 1:", p1.Solve(lines))
	fmt.Println("Part 2:", p2.Solve(lines))
}
