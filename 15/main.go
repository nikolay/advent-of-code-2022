package main

import "fmt"

type Coord struct {
	x, y int
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func (coord Coord) Add(add Coord) Coord {
	return Coord{coord.x + add.x, coord.y + add.y}
}

func (coord Coord) Distance(dest Coord) int {
	return Abs(coord.x-dest.x) + Abs(coord.y-dest.y)
}

type Pair struct {
	sensor, beacon Coord
	distance       int
}

func Validate(candidate Coord, pairs *[]Pair) bool {
	for _, pair := range *pairs {
		if candidate == pair.beacon {
			break
		}
		if candidate.Distance(pair.sensor) <= pair.distance {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Part1())
	fmt.Println(Part2())
}
