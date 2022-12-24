package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Coord struct {
	x, y int
}

var Zero = Coord{}

func (coord Coord) Move(dir int) Coord {
	return coord.Add([]Coord{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}[dir])
}

func (coord Coord) Add(add Coord) Coord {
	return Coord{coord.x + add.x, coord.y + add.y}
}

type Blizzard struct {
	coord Coord
	dir   int
}

type Valley struct {
	width, height    int
	walls            map[Coord]bool
	blizzards        []Blizzard
	startPos, endPos Coord
}

func (valley *Valley) Occupied(coord Coord) bool {
	if coord.x < 0 || coord.x >= valley.width ||
		coord.y < 0 || coord.y >= valley.height {
		return true
	}
	if v, ok := valley.walls[coord]; ok && v {
		return true
	}
	for _, blizzard := range valley.blizzards {
		if blizzard.coord == coord {
			return true
		}
	}
	return false
}

func (valley *Valley) MoveBlizzards() {
	for i, blizzard := range valley.blizzards {
		pos := blizzard.coord
		for {
			pos = pos.Move(blizzard.dir)
			for pos.x < 0 {
				pos.x += valley.width
			}
			for pos.x >= valley.width {
				pos.x -= valley.width
			}
			for pos.y < 0 {
				pos.y += valley.height
			}
			for pos.y >= valley.height {
				pos.y -= valley.height
			}
			if v, ok := valley.walls[pos]; !(ok && v) {
				break
			}
		}
		valley.blizzards[i].coord = pos
	}
}

type CacheKey struct {
	minute int
	pos    Coord
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var valley Valley
	valley.walls = map[Coord]bool{}

	scanner := bufio.NewScanner(file)
	for y := 0; scanner.Scan(); y++ {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		valley.height++
		if len(line) > valley.width {
			valley.width = len(line)
		}

		for x := 0; x < len(line); x++ {
			pos := Coord{x, y}
			switch ch := line[x]; ch {
			case '.':
				if valley.startPos == Zero {
					valley.startPos = pos
				}
				valley.endPos = pos
			case '#':
				valley.walls[pos] = true
			default:
				var dir int
				switch ch {
				case '^':
					dir = 3
				case 'v':
					dir = 1
				case '<':
					dir = 2
				case '>':
					dir = 0
				default:
					log.Fatalf("invalid command: %v", line)
				}
				valley.blizzards = append(valley.blizzards, Blizzard{pos, dir})
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minute := 0
	for leg := 0; leg < 3; leg++ {
		var start, finish Coord
		if leg%2 == 0 {
			start, finish = valley.startPos, valley.endPos
		} else {
			start, finish = valley.endPos, valley.startPos
		}

		positions := []Coord{start}
		cache := map[CacheKey]bool{CacheKey{minute, start}: true}

	outer:
		for len(positions) > 0 {
			valley.MoveBlizzards()

			var nextPositions []Coord
			next := func(coord Coord) bool {
				if valley.Occupied(coord) {
					return false
				}
				key := CacheKey{minute + 1, coord}
				if v, ok := cache[key]; ok && v {
					return false
				}
				nextPositions = append(nextPositions, coord)
				cache[key] = true
				return true
			}
			for _, pos := range positions {
				if pos == finish {
					if leg == 0 || leg == 2 {
						fmt.Println(minute)
					}
					break outer
				}

				next(pos)
				for dir := 0; dir < 4; dir++ {
					next(pos.Move(dir))
				}
			}
			positions = nextPositions
			minute++
		}
		minute++
	}
}
