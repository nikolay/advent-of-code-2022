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

func (coord Coord) GetKey() string {
	return fmt.Sprintf("%v,%v", coord.x, coord.y)
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

type Space map[string]byte

func (space Space) Get(coord Coord) byte {
	if v, ok := space[coord.GetKey()]; ok {
		return v
	}
	return '.'
}

func (space Space) Put(coord Coord, ch byte) {
	space[coord.GetKey()] = ch
}

func Drop(space *Space, drop Coord, abyss int) bool {
	moves := []Coord{
		Coord{0, +1},
		Coord{-1, +1},
		Coord{+1, +1},
	}
	pos := drop
outer:
	for pos.y < abyss {
		if space.Get(pos) != '.' {
			break
		}

		for _, move := range moves {
			newPos := pos.Add(move)
			if space.Get(newPos) == '.' {
				pos = newPos
				continue outer
			}
		}

		space.Put(pos, 'o')
		return true
	}
	return false
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
				result.Put(pos, '#')
				pos = pos.Step(segment)
			}
			result.Put(pos, '#')
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
		path := []Coord{}
		for _, frag := range segments {
			coords := strings.Split(frag, ",")
			x, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
			y, _ := strconv.Atoi(strings.TrimSpace(coords[1]))
			if y > result.maxY {
				result.maxY = y
			}
			path = append(path, Coord{x, y})
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
