package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

func parseRound(round string) (int, int, int) {
	vals := strings.Split(round, ", ")
	red, green, blue := 0, 0, 0
	for _, val := range vals {
		pair := strings.Split(val, " ")
		v, err := strconv.Atoi(string(pair[0]))
		if err != nil {
			log.Fatalf("err: %v\n", err)
		}

		switch pair[1] {
		case "red":
			red += v
		case "green":
			green += v
		case "blue":
			blue += v
		}
	}
	return red, green, blue
}

func p1(lines []string) {
	impossibles := 0
	for gameNo, line := range lines {
		start := strings.Index(line, ":") + 2
		rounds := strings.Split(line[start:], "; ")
		for _, round := range rounds {
			red, green, blue := parseRound(round)
			if red > 12 || green > 13 || blue > 14 {
				impossibles -= gameNo + 1
				break
			}
		}
	}
	noGames := len(lines)
	res := (noGames * (noGames + 1)) / 2
	res += impossibles
	fmt.Printf("res: %v\n", res)
}
func p2(lines []string) {
	res := 0
	for _, line := range lines {
		start := strings.Index(line, ":") + 2
		rounds := strings.Split(line[start:], "; ")
		minRed, minGreen, minBlue := 0, 0, 0
		for _, round := range rounds {
			red, green, blue := parseRound(round)
			minRed = utils.Max(red, minRed)
			minGreen = utils.Max(green, minGreen)
			minBlue = utils.Max(blue, minBlue)
		}
		res += minRed * minGreen * minBlue
	}
	fmt.Printf("res: %v\n", res)
}

func main() {
	lines := utils.Read()

	p1(lines)
	p2(lines)
}
