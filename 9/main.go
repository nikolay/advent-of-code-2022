package main

import "fmt"

type Coord struct {
	x, y int
}

func (coord Coord) Move(dx, dy int) Coord {
	return Coord{coord.x + dx, coord.y + dy}
}

func (coord Coord) Add(add Coord) Coord {
	return Coord{coord.x + add.x, coord.y + add.y}
}

func (coord Coord) Subtract(subtract Coord) Coord {
	return Coord{coord.x - subtract.x, coord.y - subtract.y}
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return +1
	}
	return 0
}

func ParseDirection(direction string) Coord {
	switch direction {
	case "L":
		return Coord{-1, 0}
	case "R":
		return Coord{+1, 0}
	case "U":
		return Coord{0, -1}
	case "D":
		return Coord{0, +1}
	}
	return Coord{0, 0}
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
