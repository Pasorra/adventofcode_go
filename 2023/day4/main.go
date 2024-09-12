package main

import (
	"fmt"
	"strings"
	"utils"
)

func parseRow(line string) []int {
	numofnums := (len(line) + 1) / 3
	res := make([]int, numofnums)
	for i := range numofnums {
		res[i] += int(line[i*3]-'0')*10 + int(line[(i*3)+1])
	}
	return res
}

func p1(lines []string) {
	res := 0
	for _, line := range lines {
		start := strings.Index(line, ": ") + 2
		results := strings.Split(line[start:], " | ")
		winners := parseRow(results[0])
		numbers := parseRow(results[1])
		i := 0
		for _, number := range numbers {
			if utils.Contains(number, winners) {
				i++
			}
		}
		if i > 0 {
			res += 1 << (i - 1)
		}
	}
	fmt.Printf("res: %v\n", res)
}

func p2(lines []string) {
	res := 0
	scratchies := make([]int, len(lines))
	for i := range scratchies {
		scratchies[i] = 1
	}
	for i, line := range lines {
		start := strings.Index(line, ": ") + 2
		results := strings.Split(line[start:], " | ")
		winners := parseRow(results[0])
		numbers := parseRow(results[1])
		matches := 0
		for _, number := range numbers {
			if utils.Contains(number, winners) {
				matches++
			}
		}
		for scratchIndex := i + 1; scratchIndex < i+1+matches; scratchIndex++ {
			scratchies[scratchIndex] += scratchies[i]
		}
	}
	for _, scratch := range scratchies {
		res += scratch
	}
	fmt.Printf("res: %v\n", res)
}

func main() {
	lines := utils.Read()

	p1(lines)
	p2(lines)
}
