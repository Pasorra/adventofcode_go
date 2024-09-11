package main

import (
	"fmt"
	"log"
	"strconv"
	"utils"
)

type Pair struct {
	x int
	y int
}

func extractNumber(line string, startingX int, checked *[]Pair, y int) int {
	x := startingX
	beginning := startingX
	ending := startingX
	for x >= 0 && utils.IsDigit(line[x]) {
		beginning = x
		x--
	}
	x = startingX + 1
	for x < len(line) && utils.IsDigit(line[x]) {
		ending = x
		x++
	}

	if utils.Contains(Pair{beginning, y}, *checked) {
		return 0
	}
	*checked = append(*checked, Pair{beginning, y})

	res, err := strconv.Atoi(line[beginning : ending+1])
	if err != nil {
		log.Fatalf("err: %v\n", err)
	}
	return res
}

func p1(lines []string) {
	res := 0
	checked := []Pair{}
	for y := range lines {
		for x := range lines[y] {
			char := rune(lines[y][x])

			if char == '.' || utils.IsDigit(char) {
				continue
			}

			for dy := -1; dy <= 1; dy++ {
				for dx := -1; dx <= 1; dx++ {
					if utils.IsDigit(lines[y+dy][x+dx]) {
						res += extractNumber(lines[y+dy], x+dx, &checked, y+dy)
					}
				}
			}
		}
	}
	fmt.Printf("res: %v\n", res)
}

func p2(lines []string) {
	res := 0
	checked := []Pair{}
	for y := range lines {
		for x := range lines[y] {
			char := rune(lines[y][x])

			if char != '*' {
				continue
			}

			gears := []int{0, 0}
			i := 0
			shouldBreak := false
			for dy := -1; dy <= 1; dy++ {
				if shouldBreak {
					break
				}
				for dx := -1; dx <= 1; dx++ {
					if utils.IsDigit(lines[y+dy][x+dx]) {
						gear := extractNumber(lines[y+dy], x+dx, &checked, y+dy)
						if gear == 0 {
							continue
						}
						if i > 1 {
							gears[0] = 0
							shouldBreak = true
							break
						}
						gears[i] = gear
						i++
					}
				}
			}
			res += gears[0] * gears[1]
		}
	}
	fmt.Printf("res: %v\n", res)
}

func main() {
	lines := utils.Read()

	p1(lines)
	p2(lines)
}
