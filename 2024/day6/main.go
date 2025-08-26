package main

import (
	"fmt"
	"utils"
)

type Vector struct {
	X int
	Y int
}

func (v Vector) Add(other Vector) Vector {
	return Vector{X: v.X + other.X, Y: v.Y + other.Y}
}

func (v Vector) Equals(x, y int) bool {
	return v.X == x && v.Y == y
}

type Part1 struct{}

func isInBounds(lines []string, x, y int) bool {
	return y >= 0 && y < len(lines) && x >= 0 && x < len(lines[y])
}

func hasVisited(visited []Vector, x, y int) bool {
	for _, v := range visited {
		if v.Equals(x, y) {
			return true
		}
	}
	return false
}

func hasVisitedState(visited []State, x, y, dir int) bool {
	for _, v := range visited {
		if v.dir == dir && v.pos.Equals(x, y) {
			return true
		}
	}
	return false
}

func isObstacle(obstacles []Vector, x, y int) bool {
	for _, obstacle := range obstacles {
		if obstacle.Equals(x, y) {
			return true
		}
	}
	return false
}

func (p Part1) Solve(lines []string) (res int, visited []Vector) {
	var obstacles []Vector
	guard := Vector{X: 0, Y: 0}

	// obstacles and guard position detection
	for y := range lines {
		for x := 0; x < len(lines[y]); x++ {
			switch lines[y][x] {
			case '#':
				obstacles = append(obstacles, Vector{X: x, Y: y})
			case '^':
				guard.X = x
				guard.Y = y
			}
		}
	}

	visited = []Vector{{X: guard.X, Y: guard.Y}}
	directions := []Vector{
		{X: 0, Y: -1}, // up
		{X: 1, Y: 0},  // right
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
	}
	directionIndex := 0

	for isInBounds(lines, guard.X, guard.Y) {
		nextX := guard.X + directions[directionIndex].X
		nextY := guard.Y + directions[directionIndex].Y

		// path blocked?
		if isObstacle(obstacles, nextX, nextY) {
			directionIndex = (directionIndex + 1) % 4
			continue
		}

		guard.X = nextX
		guard.Y = nextY

		// in bounds after moving?
		if !isInBounds(lines, guard.X, guard.Y) {
			break
		}

		if !hasVisited(visited, guard.X, guard.Y) {
			visited = append(visited, Vector{X: guard.X, Y: guard.Y})
		}
	}

	return len(visited), visited
}

type Part2 struct{}

type State struct {
	pos Vector
	dir int
}

func (p Part2) Solve(lines []string, guardVisited []Vector) int {
	var originalObstacles []Vector
	guard := Vector{X: 0, Y: 0}
	res := 0

	// obstacles and guard position detection
	for y := range lines {
		for x := 0; x < len(lines[y]); x++ {
			switch lines[y][x] {
			case '#':
				originalObstacles = append(originalObstacles, Vector{X: x, Y: y})
			case '^':
				guard.X = x
				guard.Y = y
			}
		}
	}
	guardStartPos := Vector{X: guard.X, Y: guard.Y}

	for _, v := range guardVisited {
		if v.X == guardStartPos.X && v.Y == guardStartPos.Y {
			continue
		}
		guard = guardStartPos
		obstacles := append([]Vector(nil), originalObstacles...)
		obstacles = append(obstacles, Vector{X: v.X, Y: v.Y})

		directions := []Vector{
			{X: 0, Y: -1}, // up
			{X: 1, Y: 0},  // right
			{X: 0, Y: 1},  // down
			{X: -1, Y: 0}, // left
		}
		directionIndex := 0
		visitedStates := make(map[State]bool)

		for isInBounds(lines, guard.X, guard.Y) {
			currentState := State{pos: Vector{X: guard.X, Y: guard.Y}, dir: directionIndex}

			if visitedStates[currentState] {
				res += 1
				break
			}
			visitedStates[currentState] = true

			nextX := guard.X + directions[directionIndex].X
			nextY := guard.Y + directions[directionIndex].Y

			// path blocked?
			if isObstacle(obstacles, nextX, nextY) {
				directionIndex = (directionIndex + 1) % 4
				continue
			}

			guard.X = nextX
			guard.Y = nextY

			// in bounds after moving?
			if !isInBounds(lines, guard.X, guard.Y) {
				break
			}
		}

	}

	return res
}

func main() {
	lines := utils.Read()

	p1 := Part1{}
	p2 := Part2{}

	res, visited := p1.Solve(lines)
	fmt.Println("Part 1:", res)
	fmt.Println("Part 2:", p2.Solve(lines, visited))
}
