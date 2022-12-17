package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coord struct {
	x, y int
}

func (coord Coord) Add(add Coord) Coord {
	return Coord{coord.x + add.x, coord.y + add.y}
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

func (coord Coord) Step(dest Coord) Coord {
	return Coord{coord.x + Sign(dest.x-coord.x), coord.y + Sign(dest.y-coord.y)}
}

type Content byte

const (
	EMPTY Content = '.'
	ROCK          = '#'
	SAND          = 'o'
)

type Space map[Coord]Content

func (space Space) Get(coord Coord) Content {
	if v, ok := space[coord]; ok {
		return v
	}
	return EMPTY
}

func (space Space) Set(coord Coord, content Content) {
	space[coord] = content
}

type DropResult int

const (
	BLOCKED DropResult = iota
	ABYSS
	LANDED
)

func Drop(space *Space, drop Coord, abyss, floor int) DropResult {
	moves := []Coord{{0, +1}, {-1, +1}, {+1, +1}}
	pos := drop
outer:
	for {
		for _, move := range moves {
			newPos := pos.Add(move)
			if (floor == 0 || newPos.y < floor) && space.Get(newPos) == EMPTY {
				if abyss != 0 && newPos.y >= abyss {
					return ABYSS
				}
				pos = newPos
				continue outer
			}
		}

		if space.Get(pos) != EMPTY {
			return BLOCKED
		}

		space.Set(pos, SAND)
		return LANDED
	}
}

type Input struct {
	paths [][]Coord
	maxY  int
	drop  Coord
}

func (input *Input) GetSpace() (result Space) {
	result = Space{}
	for _, path := range input.paths {
		var pos Coord
		for i, segment := range path {
			if i == 0 {
				pos = segment
				continue
			}
			for pos != segment {
				result.Set(pos, ROCK)
				pos = pos.Step(segment)
			}
			result.Set(pos, ROCK)
		}
	}
	return
}

func GetInput(filename string) (result Input) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		segments := strings.Split(line, " -> ")
		var path []Coord
		for _, frag := range segments {
			coords := strings.Split(frag, ",")
			x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))
			coord := Coord{x, y}
			if coord.y > result.maxY {
				result.maxY = coord.y
			}
			path = append(path, coord)
		}
		result.paths = append(result.paths, path)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	result.drop = Coord{500, 0}
	return
}

func main() {
	input := GetInput("input.txt")
	fmt.Println(Part1(&input))
	fmt.Println(Part2(&input))
}
