package main

import (
	"fmt"
	"utils"
)

func p1(lines []string) {
	res := 0

	for _, line := range lines {
		first := true
		last := 0

		for _, char := range line {
			if utils.IsDigit(char) {
				val := int(char - 48)
				if first {
					res += val * 10
					first = false
				}
				last = val
			}
		}
		res += last
	}
	fmt.Printf("res: %v\n", res)
}

func checkDigits(index int, line string) int {
	digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for i, digit := range digits {
		if len(line) >= index+len(digit) && line[index:index+len(digit)] == digit {
			return i + 1
		}
	}
	return 0
}

func p2(lines []string) {
	res := 0

	for _, line := range lines {
		first := true
		last := 0

		for i, char := range line {
			if utils.IsDigit(char) {
				val := int(char - 48)
				if first {
					res += val * 10
					first = false
				}
				last = val
			} else {
				val := checkDigits(i, line)
				if val != 0 {
					if first {
						res += val * 10
						first = false
					}
					last = val
				}
			}
		}
		res += last
	}
	fmt.Printf("res: %v\n", res)
}

func main() {
	lines := utils.Read()
	p1(lines)
	p2(lines)
}
